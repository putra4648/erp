<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Stock Levels</h1>
            <div class="flex gap-2">
                <UInput v-model="search" placeholder="Search product..." icon="i-lucide-search" />
                <USelectMenu v-model="warehouseFilter" :items="allWarehouses" value-key="id" label-key="name"
                    placeholder="Warehouse" class="w-48" clear />
                <UButton icon="i-lucide-refresh-cw" color="neutral" variant="subtle" @click="refresh" />
            </div>
        </div>

        <UTable :loading="status === 'pending'" :data="stocks" :columns="columns" />

        <div class="flex justify-end mt-4">
            <UPagination v-model:page="page" :total="total" :items-per-page="size" />
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    label: 'Stock Level'
})

import type { TableColumn } from '@nuxt/ui'
import type { Warehouse } from '~/types/models/warehouse'
import type { StockLevelResponse } from '~/types/models/stock_level'
import type PaginationResponse from '~/../server/utils/pagination_response'


const page = ref(1)
const size = ref(10)
const search = ref('')
const warehouseFilter = ref<string>('')

const refresh = () => {
    refreshStockLevels()
}

// Fetch Stock Levels
const { data, status, refresh: refreshStockLevels } = await useFetch<PaginationResponse<StockLevelResponse>>('/api/stock-levels', {
    query: {
        page,
        size,
        search,
        warehouse_id: warehouseFilter
    },
    watch: [page, size, search, warehouseFilter]
})

// Fetch Warehouses for filter
const { data: warehouseData } = await useFetch<PaginationResponse<Warehouse>>('/api/warehouses', {
    query: { page: 1, size: 100 }
})

const stocks = computed(() => data.value?.items || [])
const total = computed(() => data.value?.total || 0)
const allWarehouses = computed(() => [
    ...(warehouseData.value?.items || [])
])

const columns = ref<TableColumn<StockLevelResponse>[]>([
    {
        accessorKey: "product_name",
        header: "Product",
    },
    {
        accessorKey: "warehouse_name",
        header: "Warehouse",
    },
    {
        accessorKey: "quantity",
        header: "Quantity",
        cell: ({ row }) => {
            const qty = Number(row.original.quantity)
            return h('span', { class: qty <= 0 ? 'text-error' : '' }, qty.toLocaleString())
        }
    },
    {
        accessorKey: "last_updated",
        header: "Last Updated",
        cell: ({ row }) => new Date(row.original.last_updated).toLocaleString()
    }
])
</script>