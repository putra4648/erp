<template>
  <div>
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">Warehouse Management</h1>
      <UButton label="Add Warehouse" @click="addWarehouse" />
    </div>

    <UModal v-model:open="open" :title="isEdit ? 'Edit Warehouse' : 'Add Warehouse'"
      :description="isEdit ? 'Update warehouse details' : 'Create a new warehouse'">

      <template #body>
        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
          <UFormField label="Code" name="code">
            <UInput class="w-full" v-model="state.code" />
          </UFormField>
          <UFormField label="Name" name="name">
            <UInput class="w-full" v-model="state.name" />
          </UFormField>
          <!-- Stock levels are currently not supported by the backend DTO -->
          <!-- <div class="flex flex-row justify-between items-center">
                        <h2 class="font-bold ">Add Stock Level</h2>
                        <UButton @click="addStock">Add Stock</UButton>
                    </div>
                    <UTable :columns="stockLevelColumns" :data="state.stock_levels"></UTable> -->
          <UButton type="submit">{{ isEdit ? 'Update' : 'Save' }}</UButton>
        </UForm>
      </template>
    </UModal>

    <UTable :loading="status === 'pending'" :data="warehouses" :columns="warehouseColumns" />

    <div class="flex justify-end mt-4">
      <UPagination v-model:page="page" :total="total" :items-per-page="size" />
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  label: "Warehouse"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Warehouse, StockLevel } from '~/types/models/warehouse';
import { WarehouseSchema } from '~/validations/schemas/warehouse_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';

const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

const schema = WarehouseSchema
const toast = useToast()
const state = reactive<Warehouse>({
  id: "",
  name: "",
  code: "",
  is_active: true,
  stock_levels: []
})
const page = ref(1)
const size = ref(10)
const open = ref(false)
const isEdit = ref(false)

const { data, status, refresh } = await useFetch<PaginationResponse<Warehouse>>('/api/warehouses', {
  query: {
    page,
    size
  },
  watch: [page, size]
})

const warehouses = computed(() => (data.value?.items || []) as Warehouse[])
const total = computed(() => data.value?.total || 0)

const warehouseColumns = ref<TableColumn<Warehouse>[]>([
  {
    accessorKey: "code",
    header: "Code",
  },
  {
    accessorKey: "name",
    header: "Name",
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
      label: 'Edit',
      icon: 'i-lucide-pencil',
      onSelect: () => {
        isEdit.value = true
        state.id = row.original.id
        state.name = row.original.name
        state.code = row.original.code
        state.is_active = row.original.is_active ?? true
        open.value = true
      }
    },
    {
      type: 'separator'
    },
    {
      label: 'Remove',
      color: 'error',
      onSelect: async () => {
        try {
          await $fetch(`/api/warehouses/${row.original.id}`, {
            method: 'DELETE'
          })
          toast.add({ title: 'Success', description: 'Warehouse has been removed.' })
          refresh()
        } catch (error: any) {
          toast.add({ title: 'Error', description: error.data?.error || 'Failed to remove warehouse', color: 'error' })
        }
      }
    },
  ]
}

function addWarehouse() {
  isEdit.value = false
  state.id = ""
  state.name = ""
  state.code = ""
  state.is_active = true
  open.value = true
}

async function onSubmit(event: FormSubmitEvent<Warehouse>) {
  try {
    if (isEdit.value) {
      await $fetch(`/api/warehouses/${state.id}`, {
        method: 'PUT',
        body: event.data
      })
      toast.add({ title: 'Success', description: 'Warehouse updated successfully.' })
    } else {
      await $fetch('/api/warehouses', {
        method: 'POST',
        body: event.data
      })
      toast.add({ title: 'Success', description: 'Warehouse has been created successfully.' })
    }
    open.value = false
    refresh()
  } catch (error: any) {
    toast.add({
      title: 'Error',
      description: error.data?.error || 'Failed to create warehouse',
      color: 'error'
    })
  }
}

function onError(event: { errors: FormError[] }) {
  toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}

// const addStock = () => {
//     state.stock_levels.push({
//         id: "",
//         product: { name: '' },
//         quantity: 0
//     } as StockLevel)
// }

</script>
