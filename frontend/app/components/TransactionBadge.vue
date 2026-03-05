<template>
    <UBadge :color="getTransactionColor(type)" variant="subtle" class="capitalize">{{ type.replace('_', ' ') }}</UBadge>
</template>

<script setup lang="ts">
import { TransactionType } from '~/types/enums/transaction_enum';

const props = withDefaults(defineProps<{
    type: TransactionType;
}>(), {
    type: TransactionType.IN
})

const getTransactionColor = (type: TransactionType) => {
    switch (type) {
        case TransactionType.IN:
            return "success";
        case TransactionType.OUT:
            return "error"
        case TransactionType.TRANSFER:
            return "warning"
        default:
            return "neutral";
    }
};
</script>