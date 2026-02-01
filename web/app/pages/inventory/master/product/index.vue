<template>
    <div>
        <UModal title="Stock" description="Add Stock">
            <div class="flex justify-between items-center mb-4">
                <h1 class="text-2xl font-bold">Stock Management</h1>
                <UButton label="Add Stock" />
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

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import Joi from 'joi';
import type { FormError } from '#ui/types';
import type { Product, UOM } from '~/types/models/product';
import type { Supplier } from '~/types/models/supplier';
import type { Warehouse } from '~/types/models/warehouse';
import { TransactionType } from '~/types/enums/transaction_enum';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')


const schema = Joi.object<Product>({
    id: Joi.string().allow(''),
    sku: Joi.string().required(),
    name: Joi.string().required(),
    min_stock: Joi.number().required(),
    category: Joi.object<Product['category']>({
        id: Joi.string().required(),
        name: Joi.string().required()
    }).required(),
    uom: Joi.object<UOM>({
        id: Joi.string().required(),
        name: Joi.string().required()
    }).required()

})

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
            onSelect: (event: Event) => {
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
