<template>
    <UBadge :color="getMovementColor(type)" variant="subtle" class="capitalize">{{ type.replace('_', ' ') }}</UBadge>
</template>

<script setup lang="ts">
const props = defineProps({
    type: {
        type: String,
        required: true
    }
})

const getMovementColor = (type: string) => {
    switch (type.toLowerCase()) {
        case "in":
            return "success";
        case "out":
            return "error"
        case "transfer":
            return "warning"
        default:
            return "neutral";
    }
};
</script>