<template>
    <div>
        <UModal title="Stock" description="Add Stock">
            <div class="flex justify-between items-center mb-4">
                <h1 class="text-2xl font-bold">Stock Management</h1>
                <UButton label="Add Stock" />
            </div>

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="Name" name="name">
                        <UInput class="w-full" v-model="state.product.name" />
                    </UFormField>
                    <UButton type="submit">Save</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :data="stock" :columns="stockColumns" />
    </div>
</template>

<script setup lang="ts">

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import Joi from 'joi';
import type { FormError } from '#ui/types';
import type Product from '~/types/models/product';
import type { Stock } from '~/types/models/stock';
import type { Supplier } from '~/types/models/supplier';
import type { Warehouse } from '~/types/models/warehouse';
import { TransactionType } from '~/types/enums/transaction_enum';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')


const schema = Joi.object<Stock>({
    id: Joi.string().allow(''),
    product: Joi.object<Product>({
        id: Joi.string().allow(''),
        name: Joi.string().required()
    }).required(),
    supplier: Joi.object<Supplier>({
        id: Joi.string().allow(''),
        name: Joi.string().required(),
    }).required(),
    warehouse: Joi.object<Warehouse>({
        id: Joi.string().allow(''),
        name: Joi.string().required(),
    }).required(),
    quantity: Joi.number().min(1).required(),
    reference_no: Joi.string()
})

const toast = useToast()

const state = reactive<Stock>({
    id: "",
    product: {
        id: "",
        name: '',
        sku: '',
        min_stock: 0,
        category: {
            id: '',
            name: ''
        },
        uom: {
            id: '',
            name: ''
        }
    },
    supplier: {
        id: "",
        name: '',
        contact_person: '',
        email: '',
        address: ''
    },
    warehouse: {
        id: "",
        name: '',
        location: '',
        is_active: false,
        stock_levels: []
    },
    quantity: 0,
    reference_no: '',
    type: TransactionType.IN
})

const stock = ref<Stock[]>([])
const stockColumns = ref<TableColumn<Stock>[]>([
    {
        accessorKey: "name",
        header: "Name",
    },
    {
        accessorKey: "product",
        header: "Product",
        cell: ({ row }) => {
            return row.original.product.name
        }
    },
    {
        accessorKey: "warehouse",
        header: "Warehouse",
        cell: ({ row }) => {
            return row.original.warehouse.name
        }
    },
    {
        accessorKey: "supplier",
        header: "Supplier",
        cell: ({ row }) => {
            return row.original.supplier.name
        }
    },
    {
        accessorKey: "quantity",
        header: "Qty"
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
