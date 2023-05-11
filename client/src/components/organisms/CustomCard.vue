<template>
  <v-card elevation="0" class="custom-card">
    <v-container class="loading d-flex align-center justify-center" v-if="isLoading">
      <v-progress-circular indeterminate></v-progress-circular>
    </v-container>
    <div v-else>
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
    </div>
  </v-card>
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
    },
    isLoading: {
      type: Boolean,
      default: false,
    }
  }
});
</script>

<style lang="scss" scoped>
.loading {
  height: 50vh;
}

.custom-card {
  border-radius: 20px;
  box-shadow: 1px 0 20px #00000014 !important;

  .title {
    padding: 15px;
  }
}
</style>
