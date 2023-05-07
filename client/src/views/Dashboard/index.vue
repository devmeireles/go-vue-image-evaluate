<!-- eslint-disable vue/multi-word-component-names -->

<template>
  <Sidebar :isOpenSidebar="store.sidebarState" />
  <HeaderBar
    :toggleSidebar="store.toggleSidebar"
    :toggleTheme="store.toggleTheme"
    :toggleLanguage="store.toggleLanguage"
    :user-data="user.currentUser"
  />
  <v-main class="main">
    <DashboardPage>
      <router-view />
    </DashboardPage>
  </v-main>
</template>

<script lang="ts" setup>
  import Sidebar from "@/components/organisms/Sidebar.vue";
  import HeaderBar from "@/components/organisms/HeaderBar.vue";
  import DashboardPage from "@/components/organisms/DashboardPage.vue";
  import { useUserStore } from "@/stores/user";
  import { useCoreStore } from "@/stores/core";

  const store = useCoreStore();
  const user = useUserStore();
</script>

<style lang="scss" scoped>
  .main {
    // background-color: #f4f4f4;
  }
</style>
