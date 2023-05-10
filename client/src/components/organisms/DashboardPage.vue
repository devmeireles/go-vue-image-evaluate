<template>
  <v-container>
    <slot />
    <Snackbar />
  </v-container>
</template>

<script lang="ts">
import Snackbar from "@/components/molecules/Snackbar.vue";
import { defineComponent } from "vue";

export default defineComponent({
  name: "DashboardPage",
  components: { Snackbar },
});
</script>
