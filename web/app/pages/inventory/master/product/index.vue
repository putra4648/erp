<template>
    <div>
        <UModal title="Proudct" description="Add Product">
            <div class="flex justify-between items-center mb-4">
                <h1 class="text-2xl font-bold">Product Management</h1>
                <UButton label="Add Product" />
            </div>

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="SKU" name="sku">
                        <UInput class="w-full" v-model="state.sku" />
                    </UFormField>
                    <UFormField label="Name" name="name">
                        <UInput class="w-full" v-model="state.name" />
                    </UFormField>
                    <UFormField label="Category" name="category">
                        <UInput class="w-full" v-model="state.category.name" />
                    </UFormField>
                    <UFormField label="UOM" name="uom">
                        <UInput class="w-full" v-model="state.uom.name" />
                    </UFormField>
                    <UFormField label="Min Stock" name="min_stock">
                        <UInput class="w-full" v-model.number="state.min_stock" />
                    </UFormField>
                    <UButton type="submit">Save</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :data="product" :columns="productColumns" />
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    layout: 'master-layout',
    label: "Product"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem, BreadcrumbItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Product } from '~/types/models/product';
import { ProductSchema } from '~/validations/schemas/product_schema';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = ProductSchema
const toast = useToast()
const state = reactive<Product>({
    id: "",
    sku: '',
    name: '',
    category: {
        id: '',
        name: ''
    },
    uom: {
        id: '',
        name: ''
    },
    min_stock: 0
})
const product = ref<Product[]>([])

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
            return row.original.category.name
        }
    },
    {
        accessorKey: "uom",
        header: "UOM",
        cell: ({ row }) => {
            return row.original.uom.name
        }
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
            label: 'View customer'
        },
        {
            label: 'Remove',
            onSelect: () => {
                product.value = product.value.filter(s => s.id !== row.original.id)
            }
        },
    ]
}

async function onSubmit(event: FormSubmitEvent<Product>) {
    // FIX: The submit event is firing! The `state` object now contains the updated values
    // from your table. Instead of logging `event.data` (which only has schema fields),
    // log the whole `state` to see all your data.
    toast.add({ title: 'Success', description: 'The form has been submitted. Check the console for the data.' })

}

/**
 * This function will run if the form validation fails.
 */
function onError(event: { errors: FormError[] }) {
    toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}


</script>
