<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <h1 class="text-2xl font-bold">Stock Transactions</h1>
            <div class="flex gap-2">
                <USelectMenu v-model="productFilter" :items="allProducts" value-key="id" label-key="name"
                    placeholder="Product" class="w-48" clear />
                <USelectMenu v-model="warehouseFilter" :items="allWarehouses" value-key="id" label-key="name"
                    placeholder="Warehouse" class="w-48" clear />
                <UButton icon="i-lucide-refresh-cw" color="neutral" variant="subtle" @click="refresh" />
            </div>
        </div>

        <UTable :loading="status === 'pending'" :data="transactions" :columns="columns">
            <template #type-cell="{ row }">
                <TransactionBadge :type="row.original.type" />
            </template>
        </UTable>

        <div class="flex justify-end mt-4">
            <UPagination v-model:page="page" :total="total" :items-per-page="size" />
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    label: 'Stock Transaction'
})

import type { TableColumn } from '@nuxt/ui'
import type { Warehouse } from '~/types/models/warehouse'
import type { Product } from '~/types/models/product'
import type { StockTransactionResponse } from '~/types/models/stock_transaction'
import type PaginationResponse from '~/../server/utils/pagination_response'

const page = ref(1)
const size = ref(10)
const productFilter = ref<string>('')
const warehouseFilter = ref<string>('')

const refresh = () => {
    refreshStockTransactions()
}

// Fetch Stock Transactions
const { data, status, refresh: refreshStockTransactions } = await useFetch<PaginationResponse<StockTransactionResponse>>('/api/stock-movements/transactions', {
    query: {
        page,
        size,
        product_id: productFilter,
        warehouse_id: warehouseFilter
    },
    watch: [page, size, productFilter, warehouseFilter]
})

// Fetch Products and Warehouses for filter
const { data: warehouseData } = await useFetch<PaginationResponse<Warehouse>>('/api/warehouses', {
    query: { page: 1, size: 100 }
})

const { data: productData } = await useFetch<PaginationResponse<Product>>('/api/products', {
    query: { page: 1, size: 100 }
})

const transactions = computed(() => data.value?.items || [])
const total = computed(() => data.value?.total || 0)
const allWarehouses = computed(() => [
    ...(warehouseData.value?.items || [])
])
const allProducts = computed(() => [
    ...(productData.value?.items || [])
])

const columns = ref<TableColumn<StockTransactionResponse>[]>([
    {
        accessorKey: "created_at",
        header: "Date",
        cell: ({ row }) => new Date(row.original.created_at).toLocaleString()
    },
    {
        accessorKey: "type",
        header: "Type",

    },
    {
        accessorKey: "product_name",
        header: "Product",
    },
    {
        accessorKey: "warehouse_name",
        header: "Warehouse",
    },
    {
        accessorKey: "supplier_name",
        header: "Supplier",
        cell: ({ row }) => row.original.supplier_name || "-"
    },
    {
        accessorKey: "quantity",
        header: "Qty",
        cell: ({ row }) => {
            const qty = Number(row.original.quantity)
            const prefix = qty > 0 ? '+' : ''
            return h('span', { class: qty > 0 ? 'text-success' : 'text-error' }, `${prefix}${qty}`)
        }
    },
    {
        accessorKey: "reference_no",
        header: "Reference",
    }
])
</script>
