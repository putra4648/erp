<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Product Management</h1>
            <UButton label="Add Product" @click="addProduct" />
        </div>

        <UModal v-model:open="open" :title="isEdit ? 'Edit Product' : 'Add Product'"
            :description="isEdit ? 'Update product details' : 'Create a new product'">

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="SKU" name="sku">
                        <UInput class="w-full" v-model="state.sku" />
                    </UFormField>
                    <UFormField label="Name" name="name">
                        <UInput class="w-full" v-model="state.name" />
                    </UFormField>
                    <UFormField label="Supplier" name="supplier_id">
                        <USelectMenu class="w-full" v-model="state.supplier_id" :items="allSuppliers" value-key="id"
                            label-key="name" placeholder="Select supplier" searchable />
                    </UFormField>
                    <UFormField label="Category" name="categories">
                        <USelectMenu class="w-full" v-model="state.categories" :items="allCategories" multiple
                            label-key="name" placeholder="Select categories" />
                    </UFormField>
                    <UFormField label="UOM" name="uoms">
                        <USelectMenu class="w-full" :model-value="state.uoms[0]"
                            @update:model-value="(val) => state.uoms = (val as Category) ? [val] : []" :items="allUoms"
                            label-key="name" placeholder="Select UOM" />
                    </UFormField>
                    <UFormField label="Min Stock" name="min_stock">
                        <UInput class="w-full" v-model.number="state.min_stock" />
                    </UFormField>
                    <UButton type="submit">{{ isEdit ? 'Update' : 'Save' }}</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :loading="status === 'pending'" :data="products" :columns="productColumns" />

        <div class="flex justify-end mt-4">
            <UPagination v-model:page="page" :total="total" :items-per-page="size" />
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    label: "Product"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem, BreadcrumbItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Product, Category, UOM } from '~/types/models/product';
import type { Supplier } from '~/types/models/supplier';
import { ProductSchema } from '~/validations/schemas/product_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = ProductSchema
const toast = useToast()
const state = reactive<Product>({
    id: "",
    sku: '',
    name: '',
    supplier_id: '',
    categories: [],
    uoms: [],
    min_stock: 0
})
const page = ref(1)
const size = ref(10)
const open = ref(false)
const isEdit = ref(false)

const { data, status, refresh } = await useFetch<PaginationResponse<Product>>('/api/products', {
    query: {
        page,
        size
    },
    watch: [page, size]
})

const { data: categoriesData } = await useFetch<PaginationResponse<Category>>('/api/categories', {
    query: {
        page,
        size
    },
    watch: [page, size]
})

const { data: uomsData } = await useFetch<PaginationResponse<UOM>>('/api/uoms', {
    query: {
        page,
        size
    },
    watch: [page, size]
})

const { data: suppliersData } = await useFetch<PaginationResponse<Supplier>>('/api/suppliers', {
    query: { page: 1, size: 100 }
})

const products = computed(() => (data.value?.items || []) as Product[])
const total = computed(() => data.value?.total || 0)

const allCategories = computed(() => categoriesData.value?.items || [])
const allUoms = computed(() => uomsData.value?.items || [])
const allSuppliers = computed(() => suppliersData.value?.items || [])

const productColumns = ref<TableColumn<Product>[]>([
    {
        accessorKey: "name",
        header: "Name",
    },
    {
        accessorKey: "sku",
        header: "SKU",
    },
    {
        accessorKey: "category",
        header: "Category",
        cell: ({ row }) => {
            if (!row.original.categories) return ""
            return row.original.categories.map(c => c.name).join(", ")
        }
    },
    {
        accessorKey: "uom",
        header: "UOM",
        cell: ({ row }) => {
            if (!row.original.uoms) return ""
            return row.original.uoms.map(c => c.name).join(", ")
        }
    },
    {
        accessorKey: "supplier",
        header: "Supplier",
        cell: ({ row }) => row.original.supplier?.name || "-"
    },
    {
        accessorKey: "min_stock",
        header: "Min Stock",
    },

    {
        accessorKey: 'actions', header: 'Actions', cell: ({ row }) => {
            return h(
                UDropdownMenu,
                {
                    content: {
                        align: 'end'
                    },
                    items: getRowActions(row),
                    'aria-label': 'Actions dropdown'
                },
                () =>
                    h(UButton, {
                        icon: 'i-lucide-ellipsis-vertical',
                        color: 'neutral',
                        variant: 'ghost',
                        'aria-label': 'Actions dropdown'
                    })
            )
        }
    }
])

function getRowActions(row: TableRow<Product>): DropdownMenuItem[] {
    return [
        {
            type: 'label',
            label: 'Actions',
        },
        {
            type: 'separator'
        },
        {
            label: 'Edit',
            icon: 'i-lucide-pencil',
            onSelect: () => {
                isEdit.value = true
                state.id = row.original.id
                state.sku = row.original.sku
                state.name = row.original.name
                state.supplier_id = row.original.supplier_id
                state.categories = row.original.categories
                state.uoms = row.original.uoms
                state.min_stock = row.original.min_stock
                open.value = true
            }
        },
        {
            label: 'Remove',
            icon: 'i-lucide-trash',
            color: 'error',
            onSelect: async () => {
                try {
                    await $fetch(`/api/products/${row.original.id}`, {
                        method: 'DELETE'
                    })
                    toast.add({ title: 'Success', description: 'Product has been removed.' })
                    refresh()
                } catch (error: any) {
                    toast.add({ title: 'Error', description: error.data?.error || 'Failed to remove product', color: 'error' })
                }
            }
        },
    ]
}

function addProduct() {
    isEdit.value = false
    state.id = ""
    state.sku = ""
    state.name = ""
    state.supplier_id = ""
    state.categories = []
    state.uoms = []
    state.min_stock = 0
    open.value = true
}

async function onSubmit(event: FormSubmitEvent<Product>) {
    try {
        if (isEdit.value) {
            await $fetch(`/api/products/${state.id}`, {
                method: 'PUT',
                body: event.data
            })
            toast.add({ title: 'Success', description: 'Product updated successfully.' })
        } else {
            await $fetch('/api/products', {
                method: 'POST',
                body: event.data
            })
            toast.add({ title: 'Success', description: 'Product has been created successfully.' })
        }
        open.value = false
        refresh()
    } catch (error: any) {
        toast.add({
            title: 'Error',
            description: error.data?.error || 'Failed to save product',
            color: 'error'
        })
    }
}

function onError(event: { errors: FormError[] }) {
    toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}

</script>
