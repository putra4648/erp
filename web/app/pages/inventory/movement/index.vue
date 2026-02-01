<template>
    <div>
        <UModal title="Stock" description="Add Stock">
            <div class="flex justify-between items-center mb-4">
                <h1 class="text-2xl font-bold">Stock Movement</h1>
                <UButton label="Add Stock" />
            </div>

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="Type" name="type">
                        <USelect class="w-full" v-model="state.type" :items="transactionTypes" @change="typeChange" />
                    </UFormField>
                    <UFormField label="Date" name="transaction_date">
                        <UInput class="w-full" v-model="state.transaction_date" type="date" />
                    </UFormField>
                    <UFormField v-if="showHiddenField.is_supplier" label="Supplier" name="supplier">
                        <UInput class="w-full" v-model="state.supplier.name" />
                    </UFormField>
                    <UFormField v-if="!showHiddenField.is_supplier || showHiddenField.is_customer"
                        label="Source Warehouse" name="source_warehouse">
                        <UInput class="w-full" v-model="state.source_warehouse.name" />
                    </UFormField>
                    <UFormField v-if="!showHiddenField.is_customer || showHiddenField.is_supplier"
                        label="Target Warehouse" name="target_warehouse">
                        <UInput class="w-full" v-model="state.target_warehouse.name" />
                    </UFormField>
                    <UFormField v-if="showHiddenField.is_customer" label="Customer" name="customer">
                        <UInput class="w-full" v-model="state.customer.name" />
                    </UFormField>
                    <UFormField label="Notes" name="notes">
                        <UTextarea class="w-full" v-model="state.notes" />
                    </UFormField>
                    <UButton type="submit">Save</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :data="stock" :columns="stockColumns" />
    </div>
</template>

<script setup lang="ts">

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem, SelectItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Stock } from '~/types/models/stock';
import { TransactionType } from '~/types/enums/transaction_enum';
import { Status } from '~/types/enums/status_enum';
import { StockSchema } from '~/validations/schemas/stock_schema';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = StockSchema
const showHiddenField = ref({
    is_supplier: false,
    is_customer: false,
})
const toast = useToast()
const transactionTypes = ref<SelectItem[]>(Object.keys(TransactionType).filter((key) => isNaN(Number(key))).map((key) => ({
    label: key,
    value: TransactionType[key as keyof typeof TransactionType]
})))
const state = reactive<Stock>({
    id: "",
    transaction_no: '',
    source_warehouse: {
        id: '',
        name: '',
        location: '',
        is_active: false,
        stock_levels: []
    },
    target_warehouse: {
        id: '',
        name: '',
        location: '',
        is_active: false,
        stock_levels: []
    },
    supplier: {
        id: '',
        name: '',
        contact_person: '',
        email: '',
        address: ''
    },
    customer: {
        id: '',
        name: '',
        contact_person: '',
        email: '',
        phone: '',
        address: '',
        is_active: false
    },
    status: Status.DRAFT,
    notes: '',
})

const stock = ref<Stock[]>([])
const stockColumns = ref<TableColumn<Stock>[]>([
    {
        accessorKey: "transaction_no",
        header: "Transaction ID",
    },
    {
        accessorKey: "type",
        header: "Type",
    },
    {
        accessorKey: "supplier",
        header: "Supplier",
        cell: ({ row }) => {
            return row.original.supplier.name
        }
    },
    {
        accessorKey: "target_warehouse",
        header: "Warehouse ID (Destination)",
        cell: ({ row }) => {
            return row.original.target_warehouse.name
        }
    },
    {
        accessorKey: "notes",
        header: "Notes",
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


function typeChange() {
    const type = state.type
    console.log(type == TransactionType.IN)
    if (type == TransactionType.IN) {
        showHiddenField.value.is_supplier = true
        showHiddenField.value.is_customer = false
    }

    if (type == TransactionType.OUT) {
        showHiddenField.value.is_supplier = false
        showHiddenField.value.is_customer = true
    }

    if (type == TransactionType.TRANSFER) {
        showHiddenField.value.is_supplier = false
        showHiddenField.value.is_customer = false
    }
}

function getRowActions(row: TableRow<Stock>): DropdownMenuItem[] {
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
            onSelect: (event: Event) => {
                stock.value = stock.value.filter(s => s.id !== row.original.id)
            }
        },
    ]
}

async function onSubmit(event: FormSubmitEvent<Stock>) {
    // FIX: The submit event is firing! The `state` object now contains the updated values
    // from your table. Instead of logging `event.data` (which only has schema fields),
    // log the whole `state` to see all your data.
    toast.add({ title: 'Success', description: 'The form has been submitted. Check the console for the data.' })

    console.log('Submitted State:', event.data)
}

/**
 * This function will run if the form validation fails.
 */
function onError(event: { errors: FormError[] }) {
    console.log(event.errors)
    toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}


</script>
