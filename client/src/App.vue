<template>
  <v-app :theme="coreStore.currentTheme">
    <router-view />
  </v-app>
</template>

<script lang="ts">
  import { defineComponent } from "vue";
  import { useCoreStore } from "@/stores/core";

  export default defineComponent({
    name: "App",
    setup() {
      const coreStore = useCoreStore();

      return {
        coreStore
      };
    }
  });
</script>
