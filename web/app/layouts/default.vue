<template>
    <div class="min-h-screen flex bg-gray-50 dark:bg-gray-900">
        <!-- Sidebar -->
        <aside
            class="w-64 border-r border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-950 flex-col fixed h-full z-10 hidden md:flex">
            <div class="h-16 flex items-center px-6 border-b border-gray-200 dark:border-gray-800">
                <span class="ml-3 font-bold text-xl text-gray-900 dark:text-white">WMS PRO</span>
            </div>

            <div class="flex-1 overflow-y-auto py-4 px-3">
                <UNavigationMenu :items="links.filter((link) => link.canAccess)" orientation="vertical" />
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
                <span class="ml-2 font-bold text-lg">WMS PRO</span>
            </div>
            <UButton icon="i-heroicons-bars-3" variant="ghost" color="secondary" @click="isOpen = true" />
        </div>

        <!-- Mobile Sidebar Drawer -->
        <USlideover v-model:open="isOpen" side="left" class="md:hidden">
            <template #content>
                <div class="p-4 flex-1 flex flex-col h-full bg-white dark:bg-gray-950">
                    <div class="flex items-center justify-between mb-6">
                        <div class="flex items-center">
                            <span class="ml-2 font-bold text-xl">WMS PRO</span>
                        </div>
                        <UButton icon="i-heroicons-x-mark" variant="ghost" color="secondary" @click="isOpen = false" />
                    </div>
                    <UNavigationMenu :items="links.filter((link) => link.canAccess)" orientation="vertical"
                        @click="isOpen = false" />
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

<script setup lang="ts">
const isOpen = ref(false);
const { canAccess } = useNavAccess()
import type { NavigationMenuItem } from '@nuxt/ui'

type CustomNavigationMenuItem = NavigationMenuItem & {
    canAccess: boolean,
    children?: (NavigationMenuItem & { canAccess: boolean })[]
}

const links = ref<CustomNavigationMenuItem[]>([
    {
        label: 'Dashboard',
        icon: 'i-heroicons-home',
        to: '/',
        canAccess: true
    },
    {
        label: 'Inventory',
        icon: 'i-heroicons-archive-box',
        canAccess: canAccess({ groups: ['inventory', 'admin'] }),
        children: [
            {
                label: 'Stock',
                to: '/inventory/stock',
                canAccess: canAccess({ groups: ['inventory', 'admin'] })
            },
            {
                label: 'Supplier',
                to: '/inventory/supplier',
                canAccess: canAccess({ groups: ['inventory', 'admin'] })
            },
            {
                label: 'Warehouse',
                to: '/inventory/warehouse',
                canAccess: canAccess({ groups: ['warehouse', 'admin'] })
            }
        ].filter(child => child.canAccess)
    },
    {
        label: 'Settings',
        to: '/settings',
        canAccess: canAccess({ groups: ['admin'] })
    }
])

</script>