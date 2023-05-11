import { createRouter, createWebHistory } from "vue-router";

import FormReport from "@/views/Dashboard/Report/form.vue";
import ListReport from "@/views/Dashboard/Report/index.vue";
import DashboardView from "@/views/Dashboard/index.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: DashboardView,
      children: [
        {
          path: "",
          name: "create-report",
          component: ListReport
        }
      ]
    },
    {
      path: "/dashboard",
      name: "dashboard",
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
              component: FormReport
            },
            {
              path: "evaluate/:id",
              name: "report-evaluate",
              component: FormReport
            }
          ]
        }
      ]
    }
  ]
});

export default router;
