export const ROUTE_PERMISSIONS = [
  { path: "/inventory/movement", groups: ["staff", "admin"], roles: [] },
  {
    path: "/inventory/master/**",
    groups: ["inventory", "admin"],
    roles: [],
  },
];
