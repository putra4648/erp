<template>
    <div>
        You are currently {{ status }}
    </div>

    <button @click="signout">Signout</button>

    <div v-if="data">
        <pre>{{ data }}</pre>
    </div>
    <div v-else>
        You are not logged in.
    </div>
</template>

<script setup lang="ts">
const { status, data, signOut, refresh } = useAuth()

const signout = async () => {
    await signOut();
    await refresh()

}

</script>