<template>
    <div class="min-h-screen flex bg-gray-50 dark:bg-gray-900">
        <!-- Sidebar -->
        <aside
            class="w-64 border-r border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-950 flex flex-col fixed h-full z-10 hidden md:flex">
            <div class="h-16 flex items-center px-6 border-b border-gray-200 dark:border-gray-800">
                <Logo class="h-8 w-auto text-primary-500" />
                <span class="ml-3 font-bold text-xl text-gray-900 dark:text-white">WMS PRO</span>
            </div>

            <div class="flex-1 overflow-y-auto py-4 px-3">
                <UNavigationMenu :items="filteredMenu" orientation="vertical" />
            </div>

            <div class="p-4 border-t border-gray-200 dark:border-gray-800">
                <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500 dark:text-gray-400">© {{ new Date().getFullYear() }}</span>
                    <UColorModeButton />
                </div>
            </div>
        </aside>

        <!-- Mobile Header -->
        <div
            class="md:hidden fixed w-full z-20 bg-white dark:bg-gray-950 border-b border-gray-200 dark:border-gray-800 flex items-center justify-between px-4 h-16">
            <div class="flex items-center">
                <Logo class="h-8 w-auto text-primary-500" />
                <span class="ml-2 font-bold text-lg">WMS PRO</span>
            </div>
            <UButton icon="i-heroicons-bars-3" variant="ghost" color="gray" @click="isOpen = true" />
        </div>

        <!-- Mobile Sidebar Drawer -->
        <USlideover v-model:open="isOpen" side="left" class="md:hidden">
            <template #content>
                <div class="p-4 flex-1 flex flex-col h-full bg-white dark:bg-gray-950">
                    <div class="flex items-center justify-between mb-6">
                        <div class="flex items-center">
                            <Logo class="h-8 w-auto text-primary-500" />
                            <span class="ml-2 font-bold text-xl">WMS PRO</span>
                        </div>
                        <UButton icon="i-heroicons-x-mark" variant="ghost" color="gray" @click="isOpen = false" />
                    </div>
                    <UNavigationMenu :items="filteredMenu" orientation="vertical" @click="isOpen = false" />
                    <div
                        class="mt-auto pt-4 border-t border-gray-200 dark:border-gray-800 flex justify-between items-center">
                        <span class="text-sm text-gray-500">Theme</span>
                        <UColorModeButton />
                    </div>
                </div>
            </template>
        </USlideover>

        <!-- Main Content -->
        <main class="flex-1 md:ml-64 p-4 md:p-8 pt-20 md:pt-8">
            <slot />
        </main>
    </div>
</template>

<script setup>
const isOpen = ref(false);
const { canAccess } = useNavAccess()

const links = [
    {
        label: 'Dashboard',
        icon: 'i-heroicons-home',
        to: '/'
    },
    {
        label: 'Jurnal Akuntansi',
        to: '/finance/journal',
        // Hanya tampil jika dia group 'finance' DAN punya role 'admin-accounting'
        access: { groups: ['finance'], roles: ['admin-accounting'] }
    },
    {
        label: 'Gudang',
        to: '/inventory',
        // Hanya butuh group 'warehouse' tanpa role spesifik
        access: { groups: ['warehouse', 'admin'] }
    }
]

const filteredMenu = computed(() => {
    return links.filter(item => canAccess(item.access || {}))
})
</script>