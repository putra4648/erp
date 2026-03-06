<template>
    <div class="space-y-8">
        <!-- Header with Welcome Message -->
        <div class="flex items-center justify-between">
            <div>
                <h1 class="text-3xl font-bold text-gray-900 dark:text-white">ERP Dashboard</h1>
                <p class="text-gray-500 dark:text-gray-400 mt-1">Welcome back, {{ userName }}! Here's what's happening
                    today.</p>
            </div>
            <div class="flex gap-2">
                <UButton icon="i-heroicons-arrow-path" color="secondary" variant="outline" @click="refreshAll"
                    :loading="pending">Refresh
                </UButton>
            </div>
        </div>

        <!-- Quick Stats Overview -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
            <UCard v-for="stat in quickStats" :key="stat.label" class="overflow-hidden">
                <div class="flex items-center gap-4">
                    <div :class="[stat.bgClass, 'p-3 rounded-lg text-white']">
                        <UIcon :name="stat.icon" class="w-6 h-6" />
                    </div>
                    <div>
                        <p class="text-sm text-gray-500 dark:text-gray-400 font-medium">{{ stat.label }}</p>
                        <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stat.value }}</p>
                    </div>
                </div>
            </UCard>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <!-- Low Stock Alerts -->
            <UCard class="lg:col-span-1">
                <template #header>
                    <div class="flex items-center justify-between">
                        <h3 class="font-semibold text-gray-900 dark:text-white">Low Stock Alerts</h3>
                        <UBadge v-if="lowStockItems.length" color="error" variant="subtle">{{ lowStockItems.length }}
                            items</UBadge>
                    </div>
                </template>

                <div v-if="lowStockItems.length === 0" class="py-8 text-center">
                    <UIcon name="i-heroicons-check-circle" class="w-12 h-12 text-green-500 mx-auto mb-2" />
                    <p class="text-gray-500">All inventory levels are healthy.</p>
                </div>

                <div v-else class="space-y-4">
                    <div v-for="item in lowStockItems.slice(0, 5)" :key="item.id"
                        class="flex items-center justify-between p-3 rounded-lg bg-gray-50 dark:bg-gray-800/50">
                        <div>
                            <p class="font-medium text-sm">{{ item.product_name }}</p>
                            <p class="text-xs text-gray-500">{{ item.warehouse_name }}</p>
                        </div>
                        <div class="text-right">
                            <p class="text-sm font-bold text-red-600 dark:text-red-400">{{ item.quantity }}</p>
                            <p class="text-[10px] text-gray-400">Qty Remaining</p>
                        </div>
                    </div>
                    <UButton v-if="lowStockItems.length > 5" to="/inventory/stock-level" variant="subtle" block
                        size="sm">View All
                        Alerts
                    </UButton>
                </div>
            </UCard>

            <!-- Recent Stock Movements -->
            <UCard class="lg:col-span-2">
                <template #header>
                    <div class="flex items-center justify-between">
                        <h3 class="font-semibold text-gray-900 dark:text-white">Recent Movements</h3>
                        <UButton to="/inventory/movement" variant="link" size="sm" icon-right="i-heroicons-arrow-right">
                            View History</UButton>
                    </div>
                </template>

                <UTable :data="recentMovements" :columns="movementColumns" class="w-full">
                    <template #type-cell="{ row }">
                        <StockMovementBadge :type="row.original.type" />
                    </template>
                    <template #status-cell="{ row }">
                        <StatusBadge :status="row.original.status" />
                    </template>
                </UTable>
            </UCard>
        </div>

        <!-- Quick Navigation Tiles -->
        <div>
            <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">Quick Actions</h3>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
                <UButton v-for="action in quickActions" :key="action.label" :to="action.to" variant="outline"
                    color="neutral"
                    class="flex flex-col items-center justify-center p-6 h-auto hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors gap-3">
                    <UIcon :name="action.icon" class="w-8 h-8 text-primary" />
                    <span class="font-medium">{{ action.label }}</span>
                </UButton>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { TableColumn } from '@nuxt/ui'
import type { StockLevelResponse } from '~/types/models/stock_level'
import type { StockMovement } from '~/types/models/stock_movement'
import type PaginationResponse from '~/../server/utils/pagination_response'

definePageMeta({
    label: "Dashboard"
})

const user = useUser()
const userName = computed(() => user.value?.name || 'User')

// Fetch Data
const { data: stockLevels, pending: stockPending, refresh: refreshStock } = await useFetch<PaginationResponse<StockLevelResponse>>('/api/stock-levels', {
    query: { page: 1, size: 50 }
})

const { data: movements, pending: movementPending, refresh: refreshMovements } = await useFetch<PaginationResponse<StockMovement>>('/api/stock-movements', {
    query: { page: 1, size: 5 }
})

const { data: productsData, refresh: refreshProducts } = await useFetch<PaginationResponse<any>>('/api/products', {
    query: { page: 1, size: 1 }
})

const pending = computed(() => stockPending.value || movementPending.value)

const refreshAll = () => {
    refreshStock()
    refreshMovements()
    refreshProducts()
}

// Stats Logic
const quickStats = computed(() => [
    {
        label: 'Total Products',
        value: productsData.value?.total || 0,
        icon: 'i-heroicons-cube',
        bgClass: 'bg-blue-500'
    },
    {
        label: 'Low Stock Alert',
        value: lowStockItems.value.length,
        icon: 'i-heroicons-exclamation-triangle',
        bgClass: 'bg-red-500'
    },
    {
        label: 'Recent Movements',
        value: movements.value?.total || 0,
        icon: 'i-heroicons-arrows-right-left',
        bgClass: 'bg-indigo-500'
    },
    {
        label: 'Pending Adjustments',
        value: 0,
        icon: 'i-heroicons-clipboard-document-check',
        bgClass: 'bg-orange-500'
    }
])

// Low Stock Logic
const lowStockItems = computed(() => {
    const items = stockLevels.value?.items || []
    return items.filter(item => item.quantity <= 10) // Mock threshold of 10
})

// Movement Logic
const recentMovements = computed(() => movements.value?.items || [])

const movementColumns: TableColumn<StockMovement>[] = [
    {
        accessorKey: 'movement_no',
        header: 'Order #'
    },
    {
        accessorKey: 'type',
        header: 'Type'
    },
    {
        accessorKey: 'transaction_date',
        header: 'Date',
        cell: ({ row }) => new Date(row.original.transaction_date).toLocaleDateString()
    },
    {
        accessorKey: 'status',
        header: 'Status'
    }
]

// Quick Actions Navigation
const quickActions = [
    { label: 'New Adjustment', icon: 'i-heroicons-plus-circle', to: '/inventory/adjustment' },
    { label: 'Stock Movement', icon: 'i-heroicons-truck', to: '/inventory/movement' },
    { label: 'Product List', icon: 'i-heroicons-list-bullet', to: '/inventory/master/product' },
    { label: 'Check Levels', icon: 'i-heroicons-chart-bar', to: '/inventory/stock-level' }
]


</script>