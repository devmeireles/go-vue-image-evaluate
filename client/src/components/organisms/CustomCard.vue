<template>
  <div>
    <v-card elevation="0" class="custom-card">
      <v-card-title class="title pt-7 mx-5 d-flex justify-space-between">
        <div>{{ title }}</div>
        <div v-if="createRoute">
          <v-btn
            color="success"
            variant="outlined"
            :onclick="() => $router.push(createRoute)"
          >
            {{ $t('actions.create') }}
            <v-icon
              end
              icon="mdi-plus"
            ></v-icon>
          </v-btn>
        </div>
      </v-card-title>
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
    },
    createRoute: {
      type: String,
      default: null,
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
