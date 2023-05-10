<template>
  <v-snackbar
    v-model="snackbar.isOpen"
    :timeout="2000"
    location="top"
  >
    {{ snackbar.message }}
    <template v-slot:actions>
      <v-btn
        color="blue"
        variant="text"
        @click="toggleSnackbar"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script lang="ts" setup>
import { useCoreStore } from "@/stores/core";
import { storeToRefs } from "pinia";

const coreStore = useCoreStore();
const snackbar = storeToRefs(coreStore).snackbar.value;
const { toggleSnackbar } = coreStore;

</script>
