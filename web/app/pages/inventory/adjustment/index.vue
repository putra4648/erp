<template>
    <div class="container mx-auto px-4 py-8">
        <!-- Header with title and add button -->
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-bold">Stock Adjustment</h1>
            <UButton label="Add Adjustment" icon="i-lucide-plus" @click="openAddModal" color="primary" />
        </div>

        <!-- Stock adjustment form modal -->
        <UModal v-model="isModalOpen" :title="modalTitle" :description="modalDescription">
            <template #body>
                <UForm ref="form" :schema="adjustmentSchema" :state="state" class="space-y-4" @submit="onSubmit"
                    @error="onError">
                    <!-- Adjustment Type -->
                    <UFormField label="Adjustment Type" name="type">
                        <USelect v-model="state.type" :items="adjustmentTypes" class="w-full" />
                    </UFormField>

                    <!-- Reason -->
                    <UFormField label="Reason" name="reason">
                        <UTextarea v-model="state.reason" placeholder="Enter reason for adjustment" class="w-full" />
                    </UFormField>

                    <!-- Date -->
                    <UFormField label="Adjustment Date" name="adjustment_date">
                        <UInput type="date" v-model="state.adjustment_date" class="w-full" />
                    </UFormField>

                    <!-- Product Selection -->
                    <UFormField label="Product" name="product">
                        <USelect v-model="selectedProductId" :items="productOptions" placeholder="Select a product"
                            @change="onProductSelected" class="w-full mb-4" />
                    </UFormField>

                    <!-- Quantity Adjustment -->
                    <div v-if="selectedProduct" class="border rounded-lg p-4 bg-gray-50">
                        <h3 class="font-semibold mb-2">Current Stock: {{ selectedProduct.current_stock }}</h3>
                        <UFormField label="Adjustment Quantity" name="quantity">
                            <UInput type="number" v-model.number="state.quantity" placeholder="Enter quantity change"
                                class="w-full" />
                        </UFormField>
                    </div>

                    <!-- Notes -->
                    <UFormField label="Notes" name="notes">
                        <UTextarea v-model="state.notes" placeholder="Additional notes" class="w-full" />
                    </UFormField>

                    <!-- Submit button -->
                    <div class="flex justify-end mt-6">
                        <UButton type="submit" color="primary">Save Adjustment</UButton>
                    </div>
                </UForm>
            </template>
        </UModal>

        <!-- Stock adjustment table -->
        <UTable :data="adjustments" :columns="adjustmentColumns" class="mt-6" />
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, h } from 'vue'
import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem, SelectItem } from '@nuxt/ui'
import type { FormError } from '#ui/types'
import type { StockAdjustment } from '~/types/models/stock_adjustment'
import { AdjustmentType } from '~/types/enums/adjustment_enum'
import { Status } from '~/types/enums/status_enum'
import { StockAdjustmentSchema } from '~/validations/schemas/stock_adjustment_schema'

// Initialize UI components
const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

// State management
const isModalOpen = ref(false)
const isEditing = ref(false)
const currentAdjustment = ref<StockAdjustment | null>(null)
const form = ref(null)

// Modal title and description based on mode
const modalTitle = computed(() => isEditing.value ? 'Edit Stock Adjustment' : 'Add Stock Adjustment')
const modalDescription = computed(() => isEditing.value ?
    'Update the details of an existing stock adjustment' :
    'Create a new stock adjustment to correct inventory discrepancies')

// Form schema and state
const adjustmentSchema = StockAdjustmentSchema

// Product selection state
const selectedProductId = ref('')
const selectedProduct = ref<any>(null)

// Adjustment types for dropdown
const adjustmentTypes = ref<SelectItem[]>([
    { label: 'Increase', value: AdjustmentType.INCREASE },
    { label: 'Decrease', value: AdjustmentType.DECREASE },
    { label: 'Damage/Loss', value: AdjustmentType.DAMAGE },
    { label: 'Return to Supplier', value: AdjustmentType.RETURN }
])

// Form state with default values
const state = reactive<StockAdjustment>({
    id: '',
    type: AdjustmentType.INCREASE,
    reason: '',
    adjustment_date: new Date().toISOString().split('T')[0],
    product_id: '',
    quantity: 0,
    notes: '',
    status: Status.DRAFT
})

// Sample products data (would come from API in real app)
const products = ref([
    { id: 'prod-1', name: 'Laptop', current_stock: 50 },
    { id: 'prod-2', name: 'Mouse', current_stock: 200 },
    { id: 'prod-3', name: 'Keyboard', current_stock: 150 },
    { id: 'prod-4', name: 'Monitor', current_stock: 80 }
])

// Product options for dropdown
const productOptions = computed(() => products.value.map(product => ({
    label: `${product.name} (${product.current_stock} in stock)`,
    value: product.id
})))

// Stock adjustments data
const adjustments = ref<StockAdjustment[]>([
    {
        id: 'adj-1',
        type: AdjustmentType.DECREASE,
        reason: 'Damaged during shipping',
        adjustment_date: '2024-01-15',
        product_id: 'prod-1',
        quantity: 5,
        notes: '5 laptops damaged in transit',
        status: Status.APPROVED
    },
    {
        id: 'adj-2',
        type: AdjustmentType.INCREASE,
        reason: 'Inventory count correction',
        adjustment_date: '2024-01-20',
        product_id: 'prod-2',
        quantity: 10,
        notes: 'Counted extra mice during audit',
        status: Status.APPROVED
    },
    {
        id: 'adj-3',
        type: AdjustmentType.DAMAGE,
        reason: 'Storage damage',
        adjustment_date: '2024-01-25',
        product_id: 'prod-3',
        quantity: 3,
        notes: '3 keyboards damaged in warehouse',
        status: Status.PENDING
    }
])

// Table columns definition
const adjustmentColumns = ref<TableColumn<StockAdjustment>[]>([
    {
        accessorKey: 'id',
        header: 'ID'
    },
    {
        accessorKey: 'type',
        header: 'Type',
        cell: ({ row }) => {
            const typeMap: Record<string, string> = {
                [AdjustmentType.INCREASE]: 'Increase',
                [AdjustmentType.DECREASE]: 'Decrease',
                [AdjustmentType.DAMAGE]: 'Damage/Loss',
                [AdjustmentType.RETURN]: 'Return'
            }
            return typeMap[row.original.type] || row.original.type
        }
    },
    {
        accessorKey: 'product_id',
        header: 'Product',
        cell: ({ row }) => {
            const product = products.value.find(p => p.id === row.original.product_id)
            return product ? product.name : row.original.product_id
        }
    },
    {
        accessorKey: 'quantity',
        header: 'Quantity'
    },
    {
        accessorKey: 'adjustment_date',
        header: 'Date'
    },
    {
        accessorKey: 'status',
        header: 'Status'
    },
    {
        accessorKey: 'actions',
        header: 'Actions',
        cell: ({ row }) => h(
            UDropdownMenu,
            {
                content: { align: 'end' },
                items: getRowActions(row),
                'aria-label': 'Actions dropdown'
            },
            () => h(UButton, {
                icon: 'i-lucide-ellipsis-vertical',
                color: 'neutral',
                variant: 'ghost',
                'aria-label': 'Actions dropdown'
            })
        )
    }
])

// Toast notifications
const toast = useToast()

// CRUD Functions

/**
 * Opens the modal for adding a new stock adjustment
 */
function openAddModal() {
    isEditing.value = false
    currentAdjustment.value = null
    selectedProductId.value = ''
    selectedProduct.value = null

    // Reset form state with default values
    Object.assign(state, {
        id: '',
        type: AdjustmentType.INCREASE,
        reason: '',
        adjustment_date: new Date().toISOString().split('T')[0],
        product_id: '',
        quantity: 0,
        notes: '',
        status: Status.DRAFT
    })

    isModalOpen.value = true
}

/**
 * Opens the modal for editing an existing stock adjustment
 * @param adjustment - The adjustment to edit
 */
function openEditModal(adjustment: StockAdjustment) {
    isEditing.value = true
    currentAdjustment.value = adjustment

    // Copy the adjustment data to form state
    Object.assign(state, { ...adjustment })
    selectedProductId.value = adjustment.product_id
    selectedProduct.value = products.value.find(p => p.id === adjustment.product_id)

    isModalOpen.value = true
}

/**
 * Handles product selection in the form
 */
function onProductSelected() {
    if (selectedProductId.value) {
        selectedProduct.value = products.value.find(p => p.id === selectedProductId.value)
        state.product_id = selectedProductId.value
    } else {
        selectedProduct.value = null
        state.product_id = ''
    }
}

/**
 * Handles form submission for both create and update operations
 * @param event - The form submit event
 */
async function onSubmit(event: FormSubmitEvent<StockAdjustment>) {
    try {
        if (isEditing.value && currentAdjustment.value) {
            // Update existing adjustment
            await updateStockAdjustment(state)
        } else {
            // Create new adjustment
            await createStockAdjustment(state)
        }

        isModalOpen.value = false
    } catch (error) {
        console.error('Error submitting form:', error)
        toast.add({ title: 'Error', description: 'Failed to save stock adjustment.', color: 'error' })
    }
}

/**
 * Creates a new stock adjustment
 * @param adjustmentData - The adjustment data to create
 */
async function createStockAdjustment(adjustmentData: StockAdjustment) {
    // try {
    //     // In a real application, this would be an API call
    //     const response = await $fetch('/api/stock-adjustments', {
    //         method: 'POST',
    //         body: adjustmentData
    //     })

    //     // Add the new item to the local state
    //     adjustments.value.push(response)

    //     toast.add({ title: 'Success', description: 'Stock adjustment created successfully.' })
    // } catch (error) {
    //     console.error('Error creating stock adjustment:', error)
    //     throw error
    // }
}

/**
 * Updates an existing stock adjustment
 * @param adjustmentData - The updated adjustment data
 */
async function updateStockAdjustment(adjustmentData: StockAdjustment) {
    // try {
    //     // In a real application, this would be an API call
    //     const response = await $fetch(`/api/stock-adjustments/${adjustmentData.id}`, {
    //         method: 'PUT',
    //         body: adjustmentData
    //     })

    //     // Update the item in the local state
    //     const index = adjustments.value.findIndex(a => a.id === adjustmentData.id)
    //     if (index !== -1) {
    //         adjustments.value[index] = response
    //     }

    //     toast.add({ title: 'Success', description: 'Stock adjustment updated successfully.' })
    // } catch (error) {
    //     console.error('Error updating stock adjustment:', error)
    //     throw error
    // }
}

/**
 * Deletes a stock adjustment
 * @param id - The ID of the adjustment to delete
 */
async function deleteStockAdjustment(id: string) {
    try {
        // In a real application, this would be an API call
        await $fetch(`/api/stock-adjustments/${id}`, {
            method: 'DELETE'
        })

        // Remove the item from the local state
        adjustments.value = adjustments.value.filter(a => a.id !== id)

        toast.add({ title: 'Success', description: 'Stock adjustment deleted successfully.' })
    } catch (error) {
        console.error('Error deleting stock adjustment:', error)
        throw error
    }
}

/**
 * Gets actions for a table row
 * @param row - The table row
 * @returns Array of dropdown menu items
 */
function getRowActions(row: TableRow<StockAdjustment>): DropdownMenuItem[] {
    return [
        {
            type: 'label',
            label: 'Actions'
        },
        {
            type: 'separator'
        },
        {
            label: 'View Details'
        },
        {
            label: 'Edit',
            onSelect: (event: Event) => {
                openEditModal(row.original)
            }
        },
        {
            label: 'Delete',
            onSelect: async (event: Event) => {
                try {
                    await deleteStockAdjustment(row.original.id)
                } catch (error) {
                    console.error('Error deleting adjustment:', error)
                }
            }
        }
    ]
}

/**
 * Handles form validation errors
 * @param event - The form error event
 */
function onError(event: { errors: FormError[] }) {
    const errorMessages = event.errors.map((e) => e.message).join(', ')
    toast.add({ title: 'Validation Error', description: `Please fix the following errors: ${errorMessages}`, color: 'error' })
}

// Lifecycle hook to load data when component mounts
onMounted(async () => {
    // try {
    //     // In a real application, this would fetch data from an API
    //     const response = await $fetch('/api/stock-adjustments')
    //     adjustments.value = response

    //     // Also load products if needed
    //     const productsResponse = await $fetch('/api/products')
    //     products.value = productsResponse
    // } catch (error) {
    //     console.error('Error loading stock adjustments:', error)
    //     toast.add({ title: 'Error', description: 'Failed to load stock adjustments.', color: 'error' })
    // }
})
</script>