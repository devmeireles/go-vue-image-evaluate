<template>
  <v-app :theme="coreStore.currentTheme">
    <router-view />
  </v-app>
</template>

<script lang="ts">
import { useCoreStore } from "@/stores/core";
import { defineComponent } from "vue";

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
