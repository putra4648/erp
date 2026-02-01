// constants/permissions.ts
export const ROUTE_PERMISSIONS = [
  { path: "/inventory/adjustment/**", groups: ["IT"], roles: ["ADMIN"] },
  {
    path: "/inventory/master/**",
    groups: ["inventory", "admin"],
    roles: [],
  },
];
