<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Category Management</h1>
            <UButton label="Add Category" @click="addCategory" />
        </div>

        <UModal v-model:open="open" :title="isEdit ? 'Edit Category' : 'Add Category'"
            :description="isEdit ? 'Update category details' : 'Create a new category'">

            <template #body>
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="Name" name="name">
                        <UInput class="w-full" v-model="state.name" />
                    </UFormField>
                    <UButton type="submit">{{ isEdit ? 'Update' : 'Save' }}</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :loading="status === 'pending'" :data="categories" :columns="columns">
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
    label: "Category"
})

import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Category } from '~/types/models/product';
import { CategorySchema } from '~/validations/schemas/category_schema';
import type PaginationResponse from '~/../server/utils/pagination_response';


const schema = CategorySchema
const toast = useToast()
const open = ref(false)
const isEdit = ref(false)
const state = reactive<Category>({
    id: "",
    name: '',
})

const page = ref(1)
const size = ref(10)

const { data, status, refresh } = await useFetch<PaginationResponse<Category>>('/api/categories', {
    query: {
        page,
        size
    },
    watch: [page, size]
})

const categories = computed(() => (data.value?.items || []) as Category[])
const total = computed(() => data.value?.total || 0)

const columns = ref<TableColumn<Category>[]>([
    {
        accessorKey: "name",
        header: "Name",
    },
    {
        id: 'actions',
    }
])

function getRowActions(row: TableRow<Category>): DropdownMenuItem[] {
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
                open.value = true
            }
        },
        {
            label: 'Remove',
            icon: 'i-lucide-trash',
            color: 'error',
            onSelect: async () => {
                try {
                    await $fetch(`/api/categories/${row.original.id}`, {
                        method: 'DELETE'
                    })
                    toast.add({ title: 'Success', description: 'Category has been removed.' })
                    refresh()
                } catch (error: any) {
                    toast.add({ title: 'Error', description: error.data?.error || 'Failed to remove category', color: 'error' })
                }
            }
        },
    ]
}

function addCategory() {
    isEdit.value = false
    state.id = ""
    state.name = ""
    open.value = true
}

async function onSubmit(event: FormSubmitEvent<Category>) {
    try {
        if (isEdit.value) {
            await $fetch(`/api/categories/${state.id}`, {
                method: 'PUT',
                body: event.data
            })
            toast.add({ title: 'Success', description: 'Category updated successfully.' })
        } else {
            await $fetch('/api/categories', {
                method: 'POST',
                body: event.data
            })
            toast.add({ title: 'Success', description: 'Category created successfully.' })
        }
        open.value = false
        refresh()
    } catch (error: any) {
        toast.add({
            title: 'Error',
            description: error.data?.error || 'Failed to save category',
            color: 'error'
        })
    }
}

function onError(event: { errors: FormError[] }) {
    toast.add({ title: 'Validation Error', description: `Please fill in the required fields ${event.errors.map((e) => e.name).join(", ")}.`, color: 'error' });
}

</script>
