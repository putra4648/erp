<template>
    <UBadge :color="getStatusColor(status)" variant="subtle" class="capitalize">{{
        status.replace('_', ' ') }}</UBadge>
</template>

<script setup lang="ts">
import { Status } from '~/types/enums/status_enum';

const props = withDefaults(defineProps<{
    status: Status
}>(), {
    status: Status.PENDING
})

const getStatusColor = (status: Status) => {
    switch (status) {
        case Status.PENDING:
            return "warning";
        case Status.DRAFT:
            return "warning";
        case Status.COMPLETED:
            return "primary";
        case Status.CANCELLED:
            return "error";
        case Status.VOID:
            return "error";
        case Status.CLOSED:
            return "success";
        case Status.APPROVED:
            return "success";
        case Status.REJECTED:
            return "error";
        default:
            return "neutral";
    }
};
</script>