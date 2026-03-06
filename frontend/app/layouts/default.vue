<template>
    <div class="min-h-screen flex bg-gray-50 dark:bg-gray-900">
        <!-- Sidebar -->
        <aside
            class="w-64 border-r border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-950 flex-col fixed h-full z-10 hidden md:flex">
            <div class="h-16 flex items-center px-6 border-b border-gray-200 dark:border-gray-800">
                <span class="ml-3 font-bold text-xl text-gray-900 dark:text-white">My ERP</span>
            </div>

            <div class="flex-1 overflow-y-auto py-4 px-3">
                <UNavigationMenu :items="links" orientation="vertical" />
            </div>


            <!-- User Profile -->
            <div class="p-4 border-t border-gray-200 dark:border-gray-800 flex flex-row justify-between">
                <UUser :name="user?.name ?? ''" :avatar="{
                    src: user?.picture,
                    loading: 'lazy',
                    icon: 'i-lucide-image'
                }" :chip="{
                    color: 'primary',
                    position: 'top-right'
                }" :description="user?.email ?? ''" />
                <a href="/auth/logout" variant="subtle" icon="i-lucide-square-arrow-right-exit">
                </a>
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
                <span class="ml-2 font-bold text-lg">My ERP</span>
            </div>
            <UButton icon="i-heroicons-bars-3" variant="subtle" color="secondary" @click="isOpen = true" />
        </div>

        <!-- Mobile Sidebar Drawer -->
        <USlideover v-model:open="isOpen" side="left" class="md:hidden">
            <template #content>
                <div class="p-4 flex-1 flex flex-col h-full bg-white dark:bg-gray-950">
                    <div class="flex items-center justify-between mb-6">
                        <div class="flex items-center">
                            <span class="ml-2 font-bold text-xl">My ERP</span>
                        </div>
                        <UButton icon="i-heroicons-x-mark" variant="subtle" color="secondary" @click="isOpen = false" />
                    </div>
                    <UNavigationMenu :items="links" orientation="vertical" @click="isOpen = false" />
                    <!-- User Profile -->
                    <div class="mt-auto  flex flex-row justify-between">
                        <UUser :name="user?.name ?? ''" :avatar="{
                            src: user?.picture,
                            loading: 'lazy',
                            icon: 'i-lucide-image'
                        }" :chip="{
                            color: 'primary',
                            position: 'top-right'
                        }" :description="user?.email ?? ''" />
                        <a href="/auth/logout" variant="subtle" icon="i-lucide-square-arrow-right-exit">
                        </a>
                    </div>
                    <div class="pt-4 border-t border-gray-200 dark:border-gray-800 flex justify-between items-center">
                        <span class="text-sm text-gray-500">Theme</span>
                        <UColorModeButton />
                    </div>

                </div>
            </template>
        </USlideover>

        <!-- Main Content -->
        <main class="flex-1 md:ml-64 p-4 md:p-8 pt-20 md:pt-8">
            <UBreadcrumb :items="items" />
            <slot />
        </main>
    </div>
</template>

<script setup lang="ts">
import type { NavigationMenuItem, BreadcrumbItem } from '@nuxt/ui'

const route = useRoute()
const isOpen = ref(false);
const user = useUser()


const items = computed<BreadcrumbItem[]>(() => {
    const crumbs: BreadcrumbItem[] = [
        {
            label: 'Inventory',
        }
    ]

    const pathParts = route.path.split('/').filter(Boolean)
    // If path is /inventory/master/product, parts are [inventory, master, product]
    if (pathParts.includes('master')) {
        crumbs.push({ label: 'Master' })
    }

    crumbs.push({
        label: (route.meta.label as string) || 'Page',
        to: route.fullPath
    })

    return crumbs
})

const links = computed<NavigationMenuItem[]>(() => {
    const baseLinks: NavigationMenuItem[] = [
        {
            label: 'Dashboard',
            icon: 'i-heroicons-home',
            to: '/',
        },
        {
            label: 'Inventory',
            icon: 'i-heroicons-archive-box',
            children: [
                {
                    label: 'Movement',
                    icon: "i-lucide-truck",
                    to: "/inventory/movement",
                },
                {
                    label: 'Stock Level',
                    icon: "i-lucide-truck",
                    to: "/inventory/stock-level",
                },
                {
                    label: 'Adjustment',
                    icon: "i-lucide-edit",
                    to: "/inventory/adjustment",
                },
                {
                    label: 'Stock Transaction',
                    icon: "i-lucide-history",
                    to: "/inventory/stock-transaction",
                }
            ],
        }
    ]

    const inventoryLink = baseLinks.find(link => link.label === 'Inventory')
    if (inventoryLink && inventoryLink.children) {
        inventoryLink.children.unshift({
            label: "Master",
            icon: "i-lucide-database",
            children: [
                {
                    label: 'Product',
                    to: '/inventory/master/product',
                },
                {
                    label: 'Category',
                    to: '/inventory/master/category',
                },
                {
                    label: 'UOM',
                    to: '/inventory/master/uom',
                },
                {
                    label: 'Supplier',
                    to: '/inventory/master/supplier',
                },
                {
                    label: 'Warehouse',
                    to: '/inventory/master/warehouse',
                }
            ]
        })
    }

    return baseLinks
})

</script>