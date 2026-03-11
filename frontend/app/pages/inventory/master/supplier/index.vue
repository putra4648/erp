<template>
  <div>
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">Supplier Management</h1>
      <UButton label="Add Supplier" @click="addSupplier" />
    </div>

    <UModal v-model:open="open" :title="isEdit ? 'Edit Supplier' : 'Add Supplier'"
      :description="isEdit ? 'Update supplier details' : 'Create a new supplier'">

      <template #body>
        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
          <UFormField label="Code" name="code">
            <UInput class="w-full" v-model="state.code" />
          </UFormField>
          <UFormField label="Name" name="name">
            <UInput class="w-full" v-model="state.name" />
          </UFormField>
          <UFormField label="Email" name="email">
            <UInput class="w-full" v-model="state.email" />
          </UFormField>
          <UFormField label="Phone" name="phone">
            <UInput class="w-full" v-model="state.phone" />
          </UFormField>
          <UFormField label="Address" name="address">
            <UInput class="w-full" v-model="state.address" />
          </UFormField>
          <UButton type="submit">{{ isEdit ? 'Update' : 'Save' }}</UButton>
        </UForm>
      </template>
    </UModal>

    <UTable :loading="status === 'pending'" :data="suppliers" :columns="supplierColumns">
      <template #actions-cell="{ row }">
        <UDropdownMenu :items="getRowActions(row)" :ui="{ item: 'cursor-pointer' }">
          <UButton icon="i-lucide-ellipsis-vertical" color="neutral" variant="subtle" />
        </UDropdownMenu>
      </template>
    </UTable>

    <div class="flex justify-end mt-4">
      <UPagination v-model:page="page" :total="total" :items-per-page="size" />
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  label: "Supplier"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Supplier } from '~/types/models/supplier';
import { SupplierSchema } from '~/validations/schemas/supplier_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';

const schema = SupplierSchema
const toast = useToast()
const state = reactive<Supplier>({
  id: "",
  name: '',
  code: '',
  email: '',
  phone: '',
  address: '',
  is_active: true
})
const page = ref(1)
const size = ref(10)
const open = ref(false)
const isEdit = ref(false)

const { data, status, refresh } = await useFetch<PaginationResponse<Supplier>>('/api/suppliers', {
  query: {
    page,
    size
  },
  watch: [page, size]
})

const suppliers = computed(() => (data.value?.items || []) as Supplier[])
const total = computed(() => data.value?.total || 0)

const supplierColumns = ref<TableColumn<Supplier>[]>([
  {
    accessorKey: "code",
    header: "Code",
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "email",
    header: "Email",
  },
  {
    accessorKey: "phone",
    header: "Phone",
  },
  {
    accessorKey: "address",
    header: "Address",
  },
  {
    id: 'actions',
  }
])

function getRowActions(row: TableRow<Supplier>): DropdownMenuItem[] {
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
        state.email = row.original.email
        state.phone = row.original.phone
        state.address = row.original.address
        state.is_active = row.original.is_active ?? true
        open.value = true
      }
    },
    {
      label: 'Remove',
      color: 'error',
      onSelect: async () => {
        try {
          await $fetch(`/api/suppliers/${row.original.id}`, {
            method: 'DELETE'
          })
          toast.add({ title: 'Success', description: 'Supplier has been removed.' })
          refresh()
        } catch (error: any) {
          toast.add({ title: 'Error', description: error.data?.data.error || 'Failed to remove supplier', color: 'error' })
        }
      }
    },
  ]
}

function addSupplier() {
  isEdit.value = false
  state.id = ""
  state.name = ""
  state.code = ""
  state.email = ""
  state.phone = ""
  state.address = ""
  state.is_active = true
  open.value = true
}

async function onSubmit(event: FormSubmitEvent<Supplier>) {
  try {
    if (isEdit.value) {
      await $fetch(`/api/suppliers/${state.id}`, {
        method: 'PUT',
        body: event.data
      })
      toast.add({ title: 'Success', description: 'Supplier updated successfully.' })
    } else {
      await $fetch('/api/suppliers', {
        method: 'POST',
        body: event.data
      })
      toast.add({ title: 'Success', description: 'Supplier has been created successfully.' })
    }
    open.value = false
    refresh()
  } catch (error: any) {
    toast.add({
      title: 'Error',
      description: error.data?.data.error || 'Failed to create supplier',
      color: 'error'
    })
  }
}

function onError(event: { errors: FormError[] }) {
  toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}

</script>