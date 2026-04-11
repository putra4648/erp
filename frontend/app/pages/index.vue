<template>
    <div
        class="min-h-screen bg-white dark:bg-zinc-950 text-zinc-900 dark:text-zinc-100 font-sans selection:bg-indigo-100 selection:text-indigo-900">
        <!-- Navigation -->
        <nav
            class="fixed top-0 w-full z-50 border-b border-zinc-100 bg-white/80 dark:bg-zinc-950/80 backdrop-blur-md dark:border-zinc-800">
            <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
                <div class="flex items-center gap-3 group cursor-pointer">
                    <div
                        class="w-9 h-9 bg-zinc-900 dark:bg-white rounded-xl flex items-center justify-center transition-transform group-hover:rotate-12">
                        <UIcon name="i-lucide-box" class="text-white dark:text-black w-5 h-5" />
                    </div>
                    <span class="font-black text-2xl tracking-tighter uppercase italic">My<span
                            class="text-indigo-600">ERP</span></span>
                </div>

                <div class="hidden md:flex items-center gap-8 text-sm font-semibold text-zinc-500 dark:text-zinc-400">
                    <a href="#features"
                        class="hover:text-zinc-900 dark:hover:text-white transition-colors">Solutions</a>
                    <a href="#demo" class="hover:text-zinc-900 dark:hover:text-white transition-colors">Platform</a>
                    <a href="#about" class="hover:text-zinc-900 dark:hover:text-white transition-colors">Company</a>
                </div>

                <div class="flex items-center gap-4">
                    <UButton to="/dashboard" label="Go to Dashboard" variant="ghost" color="neutral"
                        class="font-bold" />
                    <!-- <UButton to="/dashboard" label="Start Free Trial" color="primary"
                        class="font-bold rounded-full px-6 shadow-lg shadow-indigo-500/20" /> -->
                </div>
            </div>
        </nav>

        <!-- Hero Section -->
        <section class="relative pt-32 pb-20 md:pt-48 md:pb-32 overflow-hidden">
            <!-- Background pattern -->
            <div
                class="absolute inset-0 -z-10 bg-[radial-gradient(circle_at_top_right,_var(--tw-gradient-stops))] from-indigo-50 via-transparent to-transparent dark:from-indigo-950/20">
            </div>

            <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">
                <div class="space-y-8">
                    <div
                        class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 text-xs font-bold uppercase tracking-widest border border-indigo-100 dark:border-indigo-800">
                        <span class="relative flex h-2 w-2">
                            <span
                                class="animate-ping absolute inline-flex h-full w-full rounded-full bg-indigo-400 opacity-75"></span>
                            <span class="relative inline-flex rounded-full h-2 w-2 bg-indigo-500"></span>
                        </span>
                        Now v2.4 is Live
                    </div>

                    <h1
                        class="text-6xl md:text-8xl font-black tracking-tight leading-[0.9] text-zinc-900 dark:text-white">
                        Enterprise <br />
                        <span class="text-indigo-600">Simplified.</span>
                    </h1>

                    <p class="text-xl text-zinc-500 dark:text-zinc-400 max-w-lg leading-relaxed font-medium">
                        MyERP provides a unified platform to manage your entire business lifecycle—from procurement
                        and inventory to human capital and financials.
                    </p>

                    <div class="flex flex-wrap gap-4 pt-4">
                        <UButton to="/dashboard" label="Request a Demo" size="xl" color="primary"
                            class="rounded-full px-10 font-bold h-14" />
                        <UButton label="View Case Studies" size="xl" variant="ghost" color="neutral"
                            icon="i-lucide-play-circle" class="rounded-full px-8 h-14 font-bold" />
                    </div>

                    <div class="flex items-center gap-6 pt-8 grayscale opacity-50 dark:invert">
                        <span class="text-xs font-bold uppercase text-zinc-400 tracking-[0.2em]">Trusted by</span>
                        <div class="flex gap-4 items-center">
                            <UIcon name="i-lucide-shield" class="w-6 h-6" />
                            <UIcon name="i-lucide-zap" class="w-6 h-6" />
                            <UIcon name="i-lucide-globe" class="w-6 h-6" />
                            <UIcon name="i-lucide-cpu" class="w-6 h-6" />
                        </div>
                    </div>
                </div>

                <!-- Hero Graphic (Mock Dashboard) -->
                <div class="relative group">
                    <div
                        class="absolute -inset-4 bg-indigo-500/20 rounded-[2.5rem] blur-3xl opacity-0 group-hover:opacity-100 transition-opacity duration-1000">
                    </div>
                    <div
                        class="relative rounded-[2rem] border border-zinc-200 dark:border-zinc-800 bg-white/50 dark:bg-zinc-900/50 backdrop-blur-xl p-4 shadow-2xl skew-y-3 md:-rotate-2 hover:rotate-0 transition-transform duration-700">
                        <div
                            class="border border-zinc-100 dark:border-zinc-800 rounded-2xl overflow-hidden shadow-inner bg-zinc-50 dark:bg-zinc-950">
                            <!-- Mock Dashboard Content -->
                            <div class="p-6 space-y-4">
                                <div class="flex items-center justify-between">
                                    <div class="flex gap-2">
                                        <div class="w-2 h-2 rounded-full bg-rose-500"></div>
                                        <div class="w-2 h-2 rounded-full bg-amber-500"></div>
                                        <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
                                    </div>
                                </div>
                                <div class="h-6 w-32 bg-zinc-200 dark:bg-zinc-800 rounded-md"></div>
                                <div class="grid grid-cols-3 gap-4">
                                    <div
                                        class="h-20 bg-white dark:bg-zinc-900 rounded-xl border border-zinc-100 dark:border-zinc-800 shadow-sm">
                                    </div>
                                    <div
                                        class="h-20 bg-white dark:bg-zinc-900 rounded-xl border border-zinc-100 dark:border-zinc-800 shadow-sm">
                                    </div>
                                    <div
                                        class="h-20 bg-white dark:bg-zinc-900 rounded-xl border border-zinc-100 dark:border-zinc-800 shadow-sm">
                                    </div>
                                </div>
                                <div
                                    class="h-40 bg-white dark:bg-zinc-900 rounded-xl border border-zinc-100 dark:border-zinc-800 shadow-sm">
                                </div>
                            </div>
                        </div>
                        <!-- Floating stats card -->
                        <div
                            class="absolute -bottom-10 -right-6 md:-right-12 p-6 rounded-3xl bg-white dark:bg-zinc-900 shadow-2xl border border-zinc-200 dark:border-zinc-800 max-w-[240px] animate-bounce-slow">
                            <div class="flex items-center gap-4 mb-4">
                                <div
                                    class="w-10 h-10 rounded-full bg-emerald-100 dark:bg-emerald-900/30 flex items-center justify-center text-emerald-600">
                                    <UIcon name="i-lucide-trending-up" class="w-5 h-5" />
                                </div>
                                <div>
                                    <p class="text-[10px] font-bold uppercase text-zinc-400 tracking-wider">Growth Rate
                                    </p>
                                    <p class="text-xl font-black">+24.8%</p>
                                </div>
                            </div>
                            <div class="h-2 w-full bg-zinc-100 dark:bg-zinc-800 rounded-full overflow-hidden">
                                <div class="h-full bg-emerald-500 w-3/4"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Features Grid -->
        <section id="features" class="py-32 bg-zinc-50 dark:bg-zinc-900/50">
            <div class="max-w-7xl mx-auto px-6">
                <div class="text-center max-w-2xl mx-auto mb-20 space-y-4">
                    <h2 class="text-4xl md:text-5xl font-black tracking-tight text-zinc-900 dark:text-white">Built for
                        High-Growth Teams</h2>
                    <p class="text-lg text-zinc-500 font-medium leading-relaxed">Everything you need to scale your
                        business operations without the complexity of traditional ERP systems.</p>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                    <div v-for="feature in features" :key="feature.title"
                        class="p-10 rounded-[2.5rem] bg-white dark:bg-zinc-950 border border-zinc-100 dark:border-zinc-800 hover:shadow-2xl hover:shadow-indigo-500/5 transition-all duration-500 group">
                        <div
                            :class="[feature.color, 'w-16 h-16 rounded-2xl flex items-center justify-center text-white mb-8 group-hover:scale-110 group-hover:rotate-3 transition-transform duration-500 shadow-lg shadow-indigo-500/10']">
                            <UIcon :name="feature.icon" class="w-7 h-7" />
                        </div>
                        <h3 class="text-2xl font-bold mb-4 text-zinc-900 dark:text-white">{{ feature.title }}</h3>
                        <p class="text-zinc-500 dark:text-zinc-400 font-medium leading-relaxed mb-6">{{
                            feature.description }}</p>
                        <ul class="space-y-3">
                            <li v-for="benefit in feature.benefits" :key="benefit"
                                class="flex items-center gap-3 text-sm font-bold text-zinc-600 dark:text-zinc-300">
                                <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-indigo-500" />
                                {{ benefit }}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>

        <!-- Call to Action -->
        <section class="py-20 md:py-40">
            <div class="max-w-5xl mx-auto px-6">
                <div class="relative rounded-[3rem] bg-zinc-900 dark:bg-white overflow-hidden p-12 md:p-24 text-center">
                    <!-- Decorative elements -->
                    <div
                        class="absolute top-0 right-0 w-64 h-64 bg-indigo-500/40 rounded-full blur-[100px] -translate-y-1/2 translate-x-1/2">
                    </div>
                    <div
                        class="absolute bottom-0 left-0 w-64 h-64 bg-indigo-500/20 rounded-full blur-[100px] translate-y-1/2 -translate-x-1/2">
                    </div>

                    <div class="relative z-10 space-y-8">
                        <h2
                            class="text-5xl md:text-7xl font-black text-white dark:text-zinc-900 tracking-tight leading-none">
                            Ready to transform your <br /> business?
                        </h2>
                        <p class="text-xl text-zinc-400 dark:text-zinc-500 max-w-xl mx-auto font-medium">
                            Join 5,000+ businesses running on MyERP. No credit card required, instant setup.
                        </p>
                        <div class="flex flex-wrap justify-center gap-4">
                            <!-- <UButton to="/dashboard" label="Start Free Trial" size="xl" color="primary"
                                class="rounded-full px-12 h-16 font-black scale-105 hover:scale-110 shadow-2xl transition-transform" /> -->
                            <UButton label="Contact Sales" size="xl" variant="outline" color="secondary"
                                class="rounded-full px-12 h-16 font-bold dark:border-zinc-300 dark:text-zinc-900" />
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Simple Footer -->
        <footer class="border-t border-zinc-100 dark:border-zinc-800 py-12 bg-white dark:bg-zinc-950">
            <div class="max-w-7xl mx-auto px-6 flex flex-col md:flex-row justify-between items-center gap-8">
                <div class="flex items-center gap-3 grayscale opacity-30">
                    <div class="w-7 h-7 bg-zinc-900 dark:bg-white rounded-lg flex items-center justify-center">
                        <UIcon name="i-lucide-box" class="text-white dark:text-black w-4 h-4" />
                    </div>
                    <span class="font-bold text-xl tracking-tight uppercase italic">MyERP</span>
                </div>

                <div class="flex flex-wrap justify-center gap-8 text-sm font-bold text-zinc-400">
                    <a href="#" class="hover:text-zinc-900 dark:hover:text-white transition-colors">Documentation</a>
                    <a href="#" class="hover:text-zinc-900 dark:hover:text-white transition-colors">API Status</a>
                    <a href="#" class="hover:text-zinc-900 dark:hover:text-white transition-colors">Security</a>
                    <a href="#" class="hover:text-zinc-900 dark:hover:text-white transition-colors">Privacy</a>
                </div>

                <!-- <p class="text-xs font-bold text-zinc-400 tracking-widest uppercase">© 2026 Ultra Systems Group</p> -->
            </div>
        </footer>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    label: 'Welcome',
    layout: 'main'
})

const features = [
    {
        title: 'Global Supply Chain',
        description: 'End-to-end visibility into your inventory, logistics, and multi-warehouse operations.',
        icon: 'i-lucide-box',
        color: 'bg-indigo-600',
        benefits: ['Real-time stock tracking', 'Multi-currency support', 'Automated restock']
    },
    {
        title: 'Human Capital',
        description: 'Manage your global workforce, payroll, and performance evaluations in one central hub.',
        icon: 'i-lucide-users',
        color: 'bg-zinc-900 dark:bg-indigo-500',
        benefits: ['Onboarding workflows', 'Payroll automation', 'Time tracking']
    },
    {
        title: 'Business Analytics',
        description: 'Deep business intelligence that turns enterprise data into actionable insights.',
        icon: 'i-lucide-bar-chart-3',
        color: 'bg-indigo-600',
        benefits: ['Custom dashboards', 'AI-powered forecasting', 'Exportable reports']
    }
]
</script>

<style>
@keyframes bounce-slow {

    0%,
    100% {
        transform: translateY(0);
    }

    50% {
        transform: translateY(-10px);
    }
}

.animate-bounce-slow {
    animation: bounce-slow 4s ease-in-out infinite;
}
</style>
