<template>
  <div>
    <UModal title="Warehouse" description="Add Warehouse">
      <div class="flex justify-between items-center mb-4">
        <h1 class="text-2xl font-bold">Warehouse Management</h1>
        <UButton label="Add Warehouse" />
      </div>

      <template #body>
        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
          <UFormField label="Name" name="name">
            <UInput class="w-full" v-model="state.name" />
          </UFormField>
          <UFormField label="Location" name="location">
            <UInput class="w-full" v-model="state.location" />
          </UFormField>
          <div class="flex flex-row justify-between items-center">
            <h2 class="font-bold ">Add Stock Level</h2>
            <UButton @click="addStock">Add Stock</UButton>
          </div>
          <UTable :columns="stockLevelColumns" :data="state.stock_levels"></UTable>
          <UButton type="submit">Save</UButton>
        </UForm>
      </template>
    </UModal>

    <UTable :data="warehouse" :columns="warehouseColumns" />
  </div>
</template>

<script setup lang="ts">
import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Warehouse, StockLevel } from '~/types/models/warehouse';
import { WarehouseSchema } from '~/validations/schemas/warehouse_schema';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = WarehouseSchema
const toast = useToast()
const state = reactive<Warehouse>({
  id: "",
  name: "",
  location: "",
  is_active: true,
  stock_levels: []

})
const warehouse = ref<Warehouse[]>([
  {
    id: "1",
    name: "Warehouse A",
    location: "Location A",
    is_active: true,
    stock_levels: []
  }])

const warehouseColumns = ref<TableColumn<Warehouse>[]>([
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "location",
    header: "Location"
  },
  {
    accessorKey: 'actions', header: 'Actions', cell: ({ row }) => {
      return h(
        UDropdownMenu,
        {
          content: {
            align: 'end'
          },
          items: warehouseActions(row),
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

const stockLevelColumns = ref<TableColumn<StockLevel>[]>([
  {
    accessorKey: "id",
    header: "ID",
  },
  {
    accessorKey: "product",
    header: "Product",
    cell: ({ row }) => {
      // FIX: Use `modelValue` and `@update:modelValue` (as `'onUpdate:modelValue'`)
      // to create a two-way binding that updates the state.
      return h(UInput, {
        modelValue: row.original.product?.name,
        'onUpdate:modelValue': (value: string) => {
          if (row.original.product) {
            row.original.product.name = value
          }
        }
      })
    }
  },
  {
    accessorKey: "quantity",
    header: "Quantity",
    cell: ({ row }) => {
      // FIX: The same two-way binding fix is applied here for quantity.
      return h(UInput, {
        type: 'number',
        modelValue: row.original.quantity,
        'onUpdate:modelValue': (value: string) => {
          // The input value is a string, so convert it back to a number.
          row.original.quantity = Number(value)
        }
      })
    }
  },
  {
    accessorKey: 'actions', header: 'Actions', cell: ({ row }) => {
      return h(
        UDropdownMenu,
        {
          content: {
            align: 'end'
          },
          items: stockLevelActions(row),
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

function stockLevelActions(row: TableRow<StockLevel>): DropdownMenuItem[] {
  return [
    {
      type: 'label',
      label: 'Actions',
    },
    {
      type: 'separator'
    },
    {
      label: 'Remove',
      onSelect: (event: Event) => {
        warehouse.value = warehouse.value.filter(s => s.id !== row.original.id)
      }
    },
  ]
}

function warehouseActions(row: TableRow<Warehouse>): DropdownMenuItem[] {
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
        warehouse.value = warehouse.value.filter(s => s.id !== row.original.id)
      }
    },
  ]
}

async function onSubmit(event: FormSubmitEvent<Warehouse>) {
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


const addStock = () => {
  // FIX: When adding a new stock level, initialize the nested `product` object.
  // This prevents errors when the cell rendering function tries to access `product.name`.
  state.stock_levels.push({
    id: "",
    product: { name: '' }, // Initialize product object
    quantity: 0
  } as StockLevel)
}

</script>
