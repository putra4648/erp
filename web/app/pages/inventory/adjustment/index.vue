<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Stock Adjustment</h1>
            <UButton label="Add Adjustment" @click="addAdjustment" />
        </div>

        <UModal v-model:open="open" :title="isEdit ? 'Stock Adjustment Details' : 'Add Stock Adjustment'"
            :description="isEdit ? 'View adjustment details' : 'Create a new stock adjustment'"
            :ui="{ content: 'sm:max-w-4xl' }">

            <template #body>
                <div v-if="isEdit" class="space-y-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <p class="text-sm text-gray-500">Adjustment No</p>
                            <p class="font-medium">{{ state.adjustment_no }}</p>
                        </div>
                        <div>
                            <p class="text-sm text-gray-500">Warehouse</p>
                            <p class="font-medium">{{ state.warehouse_id }}</p>
                        </div>
                        <div>
                            <p class="text-sm text-gray-500">Transaction Date</p>
                            <p class="font-medium">{{ state.transaction_date }}</p>
                        </div>
                        <div>
                            <p class="text-sm text-gray-500">Status</p>
                            <p class="font-medium">{{ state.status }}</p>
                        </div>
                    </div>
                    <div>
                        <p class="text-sm text-gray-500">Note</p>
                        <p>{{ state.note }}</p>
                    </div>

                    <UTable :data="state.items" :columns="itemDisplayColumns" />

                    <div class="flex justify-end mt-4">
                        <UButton label="Close" color="neutral" variant="ghost" @click="open = false" />
                    </div>
                </div>

                <UForm v-else :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <div class="grid grid-cols-2 gap-4">
                        <UFormField label="Warehouse" name="warehouse_id">
                            <USelectMenu class="w-full" v-model="(state.warehouse_id)" :items="allWarehouses"
                                value-key="id" label-key="name" placeholder="Select warehouse" />
                        </UFormField>
                        <UFormField label="Transaction Date" name="transaction_date">
                            <UInput type="date" class="w-full" v-model="state.transaction_date" />
                        </UFormField>
                        <UFormField label="Status" name="status">
                            <USelectMenu class="w-full" v-model="(state.status)" :items="allStatus" value-key="id"
                                label-key="name" placeholder="Select status" />
                        </UFormField>
                    </div>

                    <UFormField label="Note" name="note">
                        <UTextarea class="w-full" v-model="state.note" />
                    </UFormField>

                    <div class="space-y-2">
                        <div class="flex justify-between items-center border-b pb-2">
                            <h3 class="font-bold">Adjusted Items</h3>
                            <UButton icon="i-lucide-plus" size="xs" label="Add Item" @click="addItem" />
                        </div>

                        <div v-for="(item, index) in state.items" :key="index"
                            class="p-4 border rounded-lg bg-gray-50 dark:bg-gray-800 space-y-4 relative">
                            <UButton icon="i-lucide-x" color="error" variant="ghost" size="xs"
                                class="absolute top-2 right-2" @click="removeItem(index)" />

                            <div class="grid grid-cols-2 gap-4">
                                <UFormField label="Product" :name="`items.${index}.product_id`">
                                    <USelectMenu class="w-full" v-model="item.product_id" :items="allProducts"
                                        value-key="id" label-key="name" placeholder="Select product" />
                                </UFormField>
                                <UFormField label="Reason" :name="`items.${index}.reason_id`">
                                    <USelectMenu class="w-full" v-model="item.reason_id" :items="allReasons"
                                        value-key="id" label-key="name" placeholder="Select reason" />
                                </UFormField>
                            </div>

                            <div class="grid grid-cols-2 gap-4">
                                <UFormField label="System Qty" :name="`items.${index}.system_qty`">
                                    <UInput type="number" class="w-full" v-model.number="item.system_qty" />
                                </UFormField>
                                <UFormField label="Actual Qty" :name="`items.${index}.actual_qty`">
                                    <UInput type="number" class="w-full" v-model.number="item.actual_qty" />
                                </UFormField>
                            </div>

                            <div class="grid grid-cols-2 gap-4">
                                <UFormField label="Adjustment Qty" :name="`items.${index}.adjustment_qty`">
                                    <UInput type="number" class="w-full" v-model.number="item.adjustment_qty" />
                                </UFormField>
                            </div>
                        </div>
                    </div>

                    <div class="flex justify-end gap-2 pt-4">
                        <UButton label="Cancel" color="neutral" variant="ghost" @click="open = false" />
                        <UButton type="submit">Save Adjustment</UButton>
                    </div>
                </UForm>
            </template>
        </UModal>

        <UTable :loading="status === 'pending'" :data="adjustments" :columns="adjustmentColumns" />

        <div class="flex justify-end mt-4">
            <UPagination v-model:page="page" :total="total" :items-per-page="size" />
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    layout: 'default',
    label: "Adjustment"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { StockAdjustment, StockAdjustmentItem, AdjustmentReason } from '~/types/models/stock_adjustment';
import type { Warehouse } from '~/types/models/warehouse';
import type { Product } from '~/types/models/product';
import { StockAdjustmentSchema } from '~/validations/schemas/stock_adjustment_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';
import { Status } from '~/types/enums/status_enum';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = StockAdjustmentSchema
const toast = useToast()

const state = reactive<StockAdjustment>({
    id: "",
    adjustment_no: "",
    warehouse_id: "",
    transaction_date: new Date().toISOString().split('T')[0] as string,
    status: "DRAFT",
    note: "",
    created_by: "",
    items: []
})

const page = ref(1)
const size = ref(10)
const open = ref(false)
const isEdit = ref(false)

const { data, status, refresh } = await useFetch<PaginationResponse<StockAdjustment>>('/api/stock-adjustments', {
    query: {
        page,
        size
    },
    watch: [page, size]
})

const { data: warehouseData } = await useFetch<PaginationResponse<Warehouse>>('/api/warehouses', {
    query: { size: 100 }
})
const { data: productData } = await useFetch<PaginationResponse<Product>>('/api/products', {
    query: { size: 100 }
})

const { data: reasonData } = await useFetch<AdjustmentReason[]>('/api/adjustment-reasons')

const adjustments = computed(() => (data.value?.items || []) as StockAdjustment[])
const total = computed(() => data.value?.total || 0)

const allStatus = ref<{ id: string; name: string }[]>(Object.values(Status).map((status) => ({ id: status, name: status })))
const allWarehouses = computed(() => warehouseData.value?.items || [])
const allProducts = computed(() => productData.value?.items || [])
const allReasons = computed(() => reasonData.value || [])

const adjustmentColumns = ref<TableColumn<StockAdjustment>[]>([
    {
        accessorKey: "adjustment_no",
        header: "No",
    },
    {
        accessorKey: "transaction_date",
        header: "Date",
        cell: ({ row }) => new Date(row.original.transaction_date).toLocaleDateString()
    },
    {
        accessorKey: "status",
        header: "Status",
    },
    {
        accessorKey: "note",
        header: "Note",
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

const itemDisplayColumns = ref<TableColumn<StockAdjustmentItem>[]>([
    {
        accessorKey: "product_name",
        header: "Product",
    },
    {
        accessorKey: "reason_name",
        header: "Reason",
    },
    {
        accessorKey: "system_qty",
        header: "System Qty",
    },
    {
        accessorKey: "actual_qty",
        header: "Actual Qty",
    },
    {
        accessorKey: "adjustment_qty",
        header: "Diff",
        cell: ({ row }) => {
            const val = row.original.actual_qty - row.original.system_qty
            return h('span', { class: val >= 0 ? 'text-green-600' : 'text-red-600' }, val > 0 ? `+${val}` : val)
        }
    }
])

function getRowActions(row: TableRow<StockAdjustment>): DropdownMenuItem[] {
    return [
        {
            type: 'label',
            label: 'Actions',
        },
        {
            type: 'separator'
        },
        {
            label: 'View',
            icon: 'i-lucide-eye',
            onSelect: async () => {
                try {
                    const data = await $fetch<StockAdjustment>(`/api/stock-adjustments/${row.original.id}`)
                    isEdit.value = true
                    Object.assign(state, data)
                    open.value = true
                } catch (error) {
                    toast.add({ title: 'Error', description: 'Failed to fetch details', color: 'error' })
                }
            }
        }
    ]
}

function addAdjustment() {
    isEdit.value = false
    Object.assign(state, {
        id: "",
        adjustment_no: "",
        warehouse_id: "",
        transaction_date: new Date().toISOString().split('T')[0],
        status: "DRAFT",
        note: "",
        created_by: "",
        items: []
    })
    addItem()
    open.value = true
}

function addItem() {
    state.items.push({
        id: "",
        product_id: "",
        reason_id: "",
        actual_qty: 0,
        system_qty: 0,
        adjustment_qty: 0
    })
}

function removeItem(index: number) {
    state.items.splice(index, 1)
}

async function onSubmit(event: FormSubmitEvent<StockAdjustment>) {
    try {
        await $fetch('/api/stock-adjustments', {
            method: 'POST',
            body: event.data
        })
        toast.add({ title: 'Success', description: 'Stock adjustment has been created.' })
        open.value = false
        refresh()
    } catch (error: any) {
        toast.add({
            title: 'Error',
            description: error.data?.error || 'Failed to save adjustment',
            color: 'error'
        })
    }
}

function onError(event: { errors: FormError[] }) {
    console.log(event.errors)
    toast.add({ title: 'Validation Error', description: `Please fill in all required fields accurately.`, color: 'error' });
}

</script>