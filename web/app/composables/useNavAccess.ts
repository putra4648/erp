import { computed } from "vue"


export const useNavAccess = () => {
    const { data } = useAuth()

    const userGroups = computed(() => data.value?.groups || [])
    const userRoles = computed(() => data.value?.roles || [])

    const canAccess = (config: { groups?: string[], roles?: string[] }) => {
        // Jika tidak ada batasan, izinkan
        if (!config.groups && !config.roles) return true

        // Cek apakah user memiliki salah satu Group yang dibutuhkan
        const hasGroup = config.groups
            ? config.groups.some(g => userGroups.value.includes(g))
            : true

        // Cek apakah user memiliki salah satu Role yang dibutuhkan
        const hasRole = config.roles
            ? config.roles.some(r => userRoles.value.includes(r))
            : true

        // User harus memenuhi kriteria Group DAN Role (Logic AND)
        // Atau bisa disesuaikan menjadi OR tergantung kebutuhan bisnis Anda
        return hasGroup && hasRole
    }

    return { canAccess }
}