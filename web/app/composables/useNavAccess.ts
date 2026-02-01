export const useNavAccess = () => {
  const { data } = useAuth();

  const canAccess = (config: { groups?: string[]; roles?: string[] }) => {
    const userGroups = data.value?.user?.groups || [];
    const userRoles = data.value?.user?.roles || [];

    // 1. Jika rute tidak butuh proteksi apa pun, izinkan (true)
    if (!config.groups && !config.roles) return true;

    // 2. Cek Groups (jika diminta)
    const hasGroup =
      config.groups && config.groups.length > 0
        ? config.groups.some((g) => userGroups.includes(g))
        : true; // Jika config tidak minta group, anggap lulus pengecekan group

    // 3. Cek Roles (jika diminta)
    const hasRole =
      config.roles && config.roles.length > 0
        ? config.roles.some((r) => userRoles.includes(r))
        : true; // Jika config tidak minta role, anggap lulus pengecekan role

    // 4. Return: Harus punya group yang sesuai DAN role yang sesuai
    return hasGroup && hasRole;
  };

  return { canAccess };
};
