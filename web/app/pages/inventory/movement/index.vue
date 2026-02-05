<template>
    <div>
        <UModal v-model="isModalOpen" :title="modalTitle" :description="modalDescription">
            <div class="flex justify-between items-center mb-4">
                <h1 class="text-2xl font-bold">Stock Movement</h1>
                <UButton label="Add Stock" @click="openAddModal" />
            </div>

            <template #body>
                <UForm ref="form" :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                    <UFormField label="Type" name="type">
                        <USelect class="w-full" v-model="state.type" :items="transactionTypes" />
                    </UFormField>
                    <UFormField label="Date" name="transaction_date">
                        <UInput class="w-full" v-model="state.transaction_date" type="date" />
                    </UFormField>
                    <UFormField v-if="showSupplier" label="Supplier" name="supplier">
                        <UInput class="w-full" v-model="state.supplier.name" />
                    </UFormField>
                    <UFormField v-if="showSourceWarehouse" label="Source Warehouse" name="source_warehouse">
                        <UInput class="w-full" v-model="state.source_warehouse.name" />
                    </UFormField>
                    <UFormField v-if="showTargetWarehouse" label="Target Warehouse" name="target_warehouse">
                        <UInput class="w-full" v-model="state.target_warehouse.name" />
                    </UFormField>
                    <UFormField v-if="showCustomer" label="Customer" name="customer">
                        <UInput class="w-full" v-model="state.customer.name" />
                    </UFormField>
                    <UFormField label="Notes" name="notes">
                        <UTextarea class="w-full" v-model="state.notes" />
                    </UFormField>
                    <UButton type="submit">Save</UButton>
                </UForm>
            </template>
        </UModal>

        <UTable :data="stock" :columns="stockColumns" />
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, h } from 'vue'
import type { TableRow, TableColumn, FormSubmitEvent, DropdownMenuItem, SelectItem } from '@nuxt/ui'
import type { FormError } from '#ui/types';
import type { Stock } from '~/types/models/stock';
import { TransactionType } from '~/types/enums/transaction_enum';
import { Status } from '~/types/enums/status_enum';
import { StockSchema } from '~/validations/schemas/stock_schema';

// Initialize UI components
const UInput = resolveComponent('UInput')
const UButton = resolveComponent('UButton')
const UDropdownMenu = resolveComponent('UDropdownMenu')

// State management
const isModalOpen = ref(false)
const isEditing = ref(false)
const currentStockItem = ref<Stock | null>(null)
const form = ref(null)

// Computed properties for modal
const modalTitle = computed(() => isEditing.value ? 'Edit Stock Movement' : 'Add Stock Movement');
const modalDescription = computed(() => isEditing.value ? 'Edit the selected stock movement' : 'Add a new stock movement');

// Form schema and state
const schema = StockSchema

// Transaction types for dropdown
const transactionTypes = ref<SelectItem[]>(Object.keys(TransactionType).filter((key) => isNaN(Number(key))).map((key) => ({
    label: key,
    value: TransactionType[key as keyof typeof TransactionType]
})));

// Form state with default values
const state = reactive<Stock>({
    id: "",
    transaction_no: '',
    type: TransactionType.IN, // Default to IN transaction
    transaction_date: new Date(),
    source_warehouse: {
        id: '',
        name: '',
        location: '',
        is_active: false,
        stock_levels: []
    },
    target_warehouse: {
        id: '',
        name: '',
        location: '',
        is_active: false,
        stock_levels: []
    },
    supplier: {
        id: '',
        name: '',
        contact_person: '',
        email: '',
        address: ''
    },
    customer: {
        id: '',
        name: '',
        contact_person: '',
        email: '',
        phone: '',
        address: '',
        is_active: false
    },
    status: Status.DRAFT,
    notes: '',
});

// Computed properties to show/hide fields based on transaction type
const showSupplier = computed(() => state.type === TransactionType.IN);
const showCustomer = computed(() => state.type === TransactionType.OUT);
const showSourceWarehouse = computed(() => state.type === TransactionType.OUT || state.type === TransactionType.TRANSFER);
const showTargetWarehouse = computed(() => state.type === TransactionType.IN || state.type === TransactionType.TRANSFER);

// Toast notifications
const toast = useToast()

// Stock data - this would be replaced with API calls in a real application
const stock = ref<Stock[]>([
    {
        id: "1",
        transaction_no: 'TRN-001',
        type: TransactionType.IN,
        transaction_date: new Date('2024-01-15'),
        source_warehouse: { id: '', name: '', location: '', is_active: false, stock_levels: [] },
        target_warehouse: { id: 'WH-001', name: 'Main Warehouse', location: 'Location A', is_active: true, stock_levels: [] },
        supplier: { id: 'SUP-001', name: 'Supplier A', contact_person: '', email: '', address: '' },
        customer: { id: '', name: '', contact_person: '', email: '', phone: '', address: '', is_active: false },
        status: Status.APPROVED,
        notes: 'Initial stock received',
    },
    {
        id: "2",
        transaction_no: 'TRN-002',
        type: TransactionType.OUT,
        transaction_date: new Date('2024-01-20'),
        source_warehouse: { id: 'WH-001', name: 'Main Warehouse', location: 'Location A', is_active: true, stock_levels: [] },
        target_warehouse: { id: '', name: '', location: '', is_active: false, stock_levels: [] },
        supplier: { id: '', name: '', contact_person: '', email: '', address: '' },
        customer: { id: 'CUS-001', name: 'Customer X', contact_person: '', email: '', phone: '', address: '', is_active: true },
        status: Status.APPROVED,
        notes: 'Customer order fulfillment',
    },
    {
        id: "3",
        transaction_no: 'TRN-003',
        type: TransactionType.TRANSFER,
        transaction_date: new Date('2024-01-25'),
        source_warehouse: { id: 'WH-001', name: 'Main Warehouse', location: 'Location A', is_active: true, stock_levels: [] },
        target_warehouse: { id: 'WH-002', name: 'Secondary Warehouse', location: 'Location B', is_active: true, stock_levels: [] },
        supplier: { id: '', name: '', contact_person: '', email: '', address: '' },
        customer: { id: '', name: '', contact_person: '', email: '', phone: '', address: '', is_active: false },
        status: Status.PENDING,
        notes: 'Stock transfer for regional distribution',
    },
])

// Table columns definition
const stockColumns = ref<TableColumn<Stock>[]>([
    {
        accessorKey: "transaction_no",
        header: "Transaction ID",
    },
    {
        accessorKey: "type",
        header: "Type",
    },
    {
        accessorKey: "supplier",
        header: "Supplier",
        cell: ({ row }) => {
            return row.original.supplier.name
        }
    },
    {
        accessorKey: "target_warehouse",
        header: "Warehouse ID (Destination)",
        cell: ({ row }) => {
            return row.original.target_warehouse.name
        }
    },
    {
        accessorKey: "notes",
        header: "Notes",
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

// CRUD Functions

/**
 * Opens the modal for adding a new stock movement
 */
function openAddModal() {
    isEditing.value = false
    currentStockItem.value = null
    // Reset form state with default values
    Object.assign(state, {
        id: "",
        transaction_no: '',
        type: TransactionType.IN,
        transaction_date: new Date(),
        source_warehouse: {
            id: '',
            name: '',
            location: '',
            is_active: false,
            stock_levels: []
        },
        target_warehouse: {
            id: '',
            name: '',
            location: '',
            is_active: false,
            stock_levels: []
        },
        supplier: {
            id: '',
            name: '',
            contact_person: '',
            email: '',
            address: ''
        },
        customer: {
            id: '',
            name: '',
            contact_person: '',
            email: '',
            phone: '',
            address: '',
            is_active: false
        },
        status: Status.DRAFT,
        notes: '',
    });
    isModalOpen.value = true
}

/**
 * Opens the modal for editing an existing stock movement
 * @param stockItem - The stock item to edit
 */
function openEditModal(stockItem: Stock) {
    isEditing.value = true
    currentStockItem.value = stockItem
    // Copy the stock item data to the form state
    Object.assign(state, { ...stockItem });
    isModalOpen.value = true
}

/**
 * Handles form submission for both create and update operations
 * @param event - The form submit event
 */
async function onSubmit(event: FormSubmitEvent<Stock>) {
    try {
        if (isEditing.value && currentStockItem.value) {
            // Update existing stock movement
            await updateStockMovement(state);
        } else {
            // Create new stock movement
            await createStockMovement(state);
        }
        isModalOpen.value = false;
    } catch (error) {
        console.error('Error submitting form:', error);
        toast.add({ title: 'Error', description: 'Failed to save stock movement.', color: 'error' });
    }
}

/**
 * Creates a new stock movement
 * @param stockData - The stock movement data to create
 */
async function createStockMovement(stockData: Stock) {
    // try {
    //     // In a real application, this would be an API call
    //     const response = await $fetch('/api/stock-movements', {
    //         method: 'POST',
    //         body: stockData
    //     });

    //     // Add the new item to the local state
    //     stock.value.push(response);

    //     toast.add({ title: 'Success', description: 'Stock movement created successfully.' });
    // } catch (error) {
    //     console.error('Error creating stock movement:', error);
    //     throw error;
    // }
}

/**
 * Updates an existing stock movement
 * @param stockData - The updated stock movement data
 */
async function updateStockMovement(stockData: Stock) {
    // try {
    //     // In a real application, this would be an API call
    //     const response = await $fetch(`/api/stock-movements/${stockData.id}`, {
    //         method: 'PUT',
    //         body: stockData
    //     });

    //     // Update the item in the local state
    //     const index = stock.value.findIndex(s => s.id === stockData.id);
    //     if (index !== -1) {
    //         stock.value[index] = response;
    //     }

    //     toast.add({ title: 'Success', description: 'Stock movement updated successfully.' });
    // } catch (error) {
    //     console.error('Error updating stock movement:', error);
    //     throw error;
    // }
}

/**
 * Deletes a stock movement
 * @param id - The ID of the stock movement to delete
 */
async function deleteStockMovement(id: string) {
    try {
        // In a real application, this would be an API call
        await $fetch(`/api/stock-movements/${id}`, {
            method: 'DELETE'
        });

        // Remove the item from the local state
        stock.value = stock.value.filter(s => s.id !== id);

        toast.add({ title: 'Success', description: 'Stock movement deleted successfully.' });
    } catch (error) {
        console.error('Error deleting stock movement:', error);
        throw error;
    }
}

/**
 * Gets actions for a table row
 * @param row - The table row
 * @returns Array of dropdown menu items
 */
function getRowActions(row: TableRow<Stock>): DropdownMenuItem[] {
    return [
        {
            type: 'label',
            label: 'Actions',
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
                    await deleteStockMovement(row.original.id);
                } catch (error) {
                    console.error('Error deleting stock:', error);
                }
            }
        },
    ]
}

/**
 * Handles form validation errors
 * @param event - The form error event
 */
function onError(event: { errors: FormError[] }) {
    const errorMessages = event.errors.map((e) => e.message).join(", ");
    toast.add({ title: 'Validation Error', description: `Please fix the following errors: ${errorMessages}`, color: 'error' });
}

// Lifecycle hook to load data when component mounts
onMounted(async () => {
    // try {
    //     // In a real application, this would fetch data from an API
    //     const response = await $fetch('/api/stock-movements');
    //     stock.value = response;
    // } catch (error) {
    //     console.error('Error loading stock movements:', error);
    //     toast.add({ title: 'Error', description: 'Failed to load stock movements.', color: 'error' });
    // }
});
</script>