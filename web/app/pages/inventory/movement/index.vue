<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Stock Movement</h1>
            <UButton label="Add Movement" @click="addMovement" />
        </div>

        <UModal v-model:open="open" :title="isEdit ? 'Edit Movement' : 'Add Movement'"
            :description="isEdit ? 'Update stock movement details' : 'Create a new stock movement'">

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <div class="grid grid-cols-2 gap-4">
                        <UFormField label="Type" name="type">
                            <USelectMenu class="w-full" v-model="state.type" :items="['IN', 'OUT', 'TRANSFER']" />
                        </UFormField>
                        <UFormField label="Date" name="transaction_date">
                            <UInput class="w-full" v-model="state.transaction_date" type="date" />
                        </UFormField>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <UFormField v-if="state.type === 'OUT' || state.type === 'TRANSFER'" label="Origin Warehouse"
                            name="origin_warehouse_id">
                            <USelectMenu class="w-full" v-model="state.origin_warehouse_id" :items="allWarehouses"
                                value-key="id" label-key="name" placeholder="Select origin warehouse" />
                        </UFormField>
                        <UFormField v-if="state.type === 'IN' || state.type === 'TRANSFER'"
                            label="Destination Warehouse" name="destination_warehouse_id">
                            <USelectMenu class="w-full" v-model="state.destination_warehouse_id" :items="allWarehouses"
                                value-key="id" label-key="name" placeholder="Select destination warehouse" />
                        </UFormField>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <UFormField label="Reference No" name="reference_no">
                            <UInput class="w-full" v-model="state.reference_no" />
                        </UFormField>
                        <UFormField label="Status" name="status">
                            <USelectMenu class="w-full" v-model="state.status" :items="availableStatuses" disabled />
                        </UFormField>
                    </div>

                    <UFormField label="Note" name="note">
                        <UTextarea class="w-full" v-model="state.note" />
                    </UFormField>

                    <div class="space-y-2">
                        <div class="flex justify-between items-center">
                            <h3 class="font-medium text-sm">Items</h3>
                            <UButton icon="i-lucide-plus" size="xs" color="neutral" variant="ghost" @click="addItem" />
                        </div>

                        <div v-for="(item, index) in state.items" :key="index"
                            class="grid grid-cols-12 gap-2 items-start">
                            <div class="col-span-6">
                                <USelectMenu v-model="item.product_id" :items="allProducts" value-key="id"
                                    label-key="name" placeholder="Select product" class="w-full" />
                            </div>
                            <div class="col-span-3">
                                <UInput v-model.number="item.quantity" type="number" placeholder="Qty" class="w-full" />
                            </div>
                            <div class="col-span-2">
                                <UInput v-model="item.note" placeholder="Item Note" class="w-full" />
                            </div>
                            <div class="col-span-1 flex justify-center pt-1">
                                <UButton icon="i-lucide-trash" size="xs" color="error" variant="ghost"
                                    @click="removeItem(index)" />
                            </div>
                        </div>
                    </div>

                    <div class="flex justify-end pt-4">
                        <UButton type="submit" :loading="loading">{{ isEdit ? 'Update' : 'Save' }}</UButton>
                    </div>
                </UForm>
            </template>
        </UModal>

        <UTable :loading="status === 'pending'" :data="movements" :columns="columns" />

        <div class="flex justify-end mt-4">
            <UPagination v-model:page="page" :total="total" :items-per-page="size" />
        </div>
    </div>
</template>

<script setup lang="ts">
import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import { Status } from '~/types/enums/status_enum';
import type { StockMovement, StockMovementDTO } from '~/types/models/stock_movement';
import type { Warehouse } from '~/types/models/warehouse';
import type { Product } from '~/types/models/product';
import { StockMovementSchema } from '~/validations/schemas/stock_movement_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';

const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')
const UBadge = resolveComponent('UBadge')

const schema = StockMovementSchema
const toast = useToast()
const loading = ref(false)

const initialState: StockMovementDTO = {
    type: 'IN',
    origin_warehouse_id: undefined,
    destination_warehouse_id: undefined,
    reference_no: '',
    status: Status.DRAFT,
    transaction_date: new Date().toISOString().split('T')[0],
    note: '',
    items: [{ product_id: '', quantity: 1, note: '' }]
}

const state = reactive<StockMovementDTO>({ ...initialState })
const page = ref(1)
const size = ref(10)
const open = ref(false)
const isEdit = ref(false)
const selectedId = ref("")

const availableStatuses = Object.values(Status)

// Fetch Data
const { data, status, refresh } = await useFetch<PaginationResponse<StockMovement>>('/api/stock-movement', {
    query: { page, size },
    watch: [page, size]
})

const { data: warehouseData } = await useFetch<PaginationResponse<Warehouse>>('/api/warehouses', {
    query: { page: 1, size: 100 }
})

const { data: productData } = await useFetch<PaginationResponse<Product>>('/api/products', {
    query: { page: 1, size: 100 }
})

const movements = computed(() => data.value?.items || [])
const total = computed(() => data.value?.total || 0)
const allWarehouses = computed(() => warehouseData.value?.items || [])
const allProducts = computed(() => productData.value?.items || [])

const columns = ref<TableColumn<StockMovement>[]>([
    {
        accessorKey: "movement_no",
        header: "Movement No",
    },
    {
        accessorKey: "type",
        header: "Type",
        cell: ({ row }) => {
            const colors: Record<string, any> = { IN: 'success', OUT: 'error', TRANSFER: 'primary' }
            return h(UBadge, { color: colors[row.original.type] || 'neutral', variant: 'subtle' }, () => row.original.type)
        }
    },
    {
        accessorKey: "origin_warehouse",
        header: "Origin",
        cell: ({ row }) => row.original.origin_warehouse?.name || '-'
    },
    {
        accessorKey: "destination_warehouse",
        header: "Destination",
        cell: ({ row }) => row.original.destination_warehouse?.name || '-'
    },
    {
        accessorKey: "transaction_date",
        header: "Date",
        cell: ({ row }) => new Date(row.original.transaction_date).toLocaleDateString()
    },
    {
        accessorKey: "status",
        header: "Status",
        cell: ({ row }) => {
            const colors: Record<string, any> = { DRAFT: 'neutral', APPROVED: 'success', CANCELLED: 'error', VOID: 'error' }
            return h(UBadge, { color: colors[row.original.status] || 'neutral' }, () => row.original.status)
        }
    },
    {
        accessorKey: 'actions',
        header: 'Actions',
        cell: ({ row }) => {
            return h(
                UDropdownMenu,
                {
                    content: { align: 'end' },
                    items: getRowActions(row),
                },
                () => h(UButton, {
                    icon: 'i-lucide-ellipsis-vertical',
                    color: 'neutral',
                    variant: 'ghost',
                })
            )
        }
    }
])

function getRowActions(row: TableRow<StockMovement>): DropdownMenuItem[] {
    return [
        { label: 'Edit', icon: 'i-lucide-pencil', onSelect: () => editMovement(row.original) },
        {
            label: 'Remove',
            icon: 'i-lucide-trash',
            color: 'error',
            onSelect: async () => {
                if (confirm('Are you sure you want to remove this movement?')) {
                    try {
                        await $fetch(`/api/stock-movement/${row.original.id}`, { method: 'DELETE' })
                        toast.add({ title: 'Success', description: 'Movement removed.' })
                        refresh()
                    } catch (error: any) {
                        toast.add({ title: 'Error', description: error.data?.error || 'Failed to remove', color: 'error' })
                    }
                }
            }
        }
    ]
}

function addMovement() {
    isEdit.value = false
    selectedId.value = ""
    Object.assign(state, {
        ...initialState,
        items: [{ product_id: '', quantity: 1, note: '' }]
    })
    open.value = true
}

function editMovement(movement: StockMovement) {
    isEdit.value = true
    selectedId.value = movement.id
    Object.assign(state, {
        id: movement.id,
        type: movement.type,
        origin_warehouse_id: movement.origin_warehouse_id,
        destination_warehouse_id: movement.destination_warehouse_id,
        reference_no: movement.reference_no,
        status: movement.status,
        transaction_date: movement.transaction_date.split('T')[0],
        note: movement.note,
        items: movement.items.map(i => ({
            id: i.id,
            product_id: i.product_id,
            quantity: i.quantity,
            note: i.note
        }))
    })
    open.value = true
}

function addItem() {
    state.items.push({ product_id: '', quantity: 1, note: '' })
}

function removeItem(index: number) {
    if (state.items.length > 1) {
        state.items.splice(index, 1)
    }
}

async function onSubmit(event: FormSubmitEvent<StockMovementDTO>) {
    loading.value = true
    try {
        const url = isEdit.value ? `/api/stock-movement/${selectedId.value}` : '/api/stock-movement'
        const method = isEdit.value ? 'PUT' : 'POST'

        await $fetch(url, {
            method,
            body: event.data
        })

        toast.add({ title: 'Success', description: `Movement ${isEdit.value ? 'updated' : 'created'} successfully.` })
        open.value = false
        refresh()
    } catch (error: any) {
        toast.add({
            title: 'Error',
            description: error.data?.error || 'Failed to save movement',
            color: 'error'
        })
    } finally {
        loading.value = false
    }
}

function onError(event: { errors: FormError[] }) {
    toast.add({ title: 'Validation Error', description: 'Please check your input.', color: 'error' });
}
</script>