<template>
  <div>
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">Supplier Management</h1>
      <UModal title="Supplier" description="Add Supplier">
        <UButton>New Supplier</UButton>
        <template #body>
          <UForm :state="state" :schema="schema" class="space-y-4" @submit="onSubmit">
            <UFormField label="Name" name="name">
              <UInput v-model="state.name" class="w-full" />
            </UFormField>
            <UFormField label="Email" name="email">
              <UInput v-model="state.email" class="w-full" />
            </UFormField>
            <UFormField label="Contact Person" name="contact_person">
              <UInput v-model="state.contact_person" class="w-full" />
            </UFormField>
            <UFormField label="Address">
              <UInput v-model="state.address" class="w-full" />
            </UFormField>
            <UButton type="submit">Save</UButton>
          </UForm>
        </template>
      </UModal>
    </div>

    <UTable :data="suppliers" :columns="columns" />



  </div>
</template>

<script setup lang="ts">
import type { TableColumn, TableRow, DropdownMenuItem } from '@nuxt/ui'
import Joi from 'joi'
import type { Supplier } from '~/types/models/supplier'

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')

const schema = Joi.object<Supplier>({
  id: Joi.string().required(),
  name: Joi.string().required(),
  contact_person: Joi.string().required(),
  email: Joi.string().email(),
  address: Joi.string().required(),
  is_active: Joi.boolean().required().default(true),
})

const state = reactive<Supplier>({ id: "", name: '', contact_person: '', address: '', email: "" })

const suppliers = ref<Supplier[]>([
  { id: "1", name: 'Supplier A', contact_person: 'John Doe', email: "", address: '123 Main St', is_active: true },
  { id: "2", name: 'Supplier B', contact_person: 'Jane Smith', email: "", address: '456 Oak Ave', is_active: true },
  { id: "3", name: 'Supplier C', contact_person: 'Peter Jones', email: "", address: '789 Pine Ln', is_active: true },
])

const columns = ref<TableColumn<Supplier>[]>(
  [
    { accessorKey: 'id', header: 'ID' },
    { accessorKey: 'name', header: 'Name' },
    { accessorKey: 'email', header: 'Email' },
    { accessorKey: 'contact_person', header: 'Contact' },
    { accessorKey: 'address', header: 'Address' },
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
  ]
)

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
      label: 'Remove',
      onSelect: (event: Event) => {
        suppliers.value = suppliers.value.filter(s => s.id !== row.original.id)
      }
    },
  ]
}


async function onSubmit() {

}

</script>