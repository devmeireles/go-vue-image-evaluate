<template>
  <v-container>
    <slot />
  </v-container>
</template>

<script lang="ts">
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "DashboardPage"
  });
</script>
