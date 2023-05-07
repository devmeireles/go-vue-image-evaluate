import { createRouter, createWebHistory } from "vue-router";

import DashboardView from "@/views/Dashboard/index.vue";
import ListReport from "@/views/Dashboard/Report/index.vue";
import CreateReport from "@/views/Dashboard/Report/create.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/dashboard",
      name: "home",
      component: DashboardView,
      children: [
        {
          path: "",
          name: "dashboard",
          component: ListReport
        },
        {
          path: "report",
          name: "report",
          children: [
            {
              path: "",
              name: "report-index",
              component: ListReport
            },
            {
              path: "create",
              name: "report-create",
              component: CreateReport
            }
          ]
        }
      ]
    }
  ]
});

export default router;
