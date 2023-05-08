<template>
  <div>
    <v-card elevation="0" class="custom-card">
      <v-card-title class="title mx-5">{{ title }}</v-card-title>
      <div class="body pt-5 px-5">
        <slot />
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "CustomCard",
  props: {
    title: {
      type: String,
      default: null
    }
  }
});
</script>

<style lang="scss" scoped>
.custom-card {
  border-radius: 20px;
  box-shadow: 1px 0 20px #00000014 !important;

  .title {
    padding: 15px;
  }
}
</style>
