<template>
    <UContainer class="py-12">
        <div class="max-w-md mx-auto text-center">
            <div class="mb-6">
                <div
                    class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-red-100 dark:bg-red-900/20">
                    <UIcon name="i-lucide-lock" class="w-10 h-10 text-red-600 dark:text-red-400" />
                </div>
            </div>

            <h1 class="text-3xl font-bold tracking-tight text-gray-900 dark:text-white mb-4">
                Access Denied
            </h1>

            <p class="text-gray-600 dark:text-gray-300 mb-8">
                You don't have permission to access this page. Please contact your administrator if you believe this is
                an error.
            </p>

            <div class="flex flex-col sm:flex-row gap-3 justify-center">
                <UButton color="neutral" variant="outline" to="/" size="lg">
                    Go to Dashboard
                </UButton>
                <a color="primary" size="lg" href="/auth/logout">
                    Logout
                </a>
            </div>
        </div>
    </UContainer>
</template>

<script setup lang="ts">

</script>