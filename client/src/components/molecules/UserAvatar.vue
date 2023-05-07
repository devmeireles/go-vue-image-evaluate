<template>
  <div class="d-flex flex-row items-center">
    <v-avatar :color="color" :size="size">
      <v-img v-if="img" :src="img"></v-img>
      <span v-if="initials && !img" class="text-white text-body">{{
        initials
      }}</span>
    </v-avatar>

    <div class="d-flex flex-row align-self-center pl-2">
      <p :class="[captionSize, captionWeight]">{{ content }}</p>
    </div>
  </div>
</template>

<script lang="ts">
  export default {
    name: "UserAvatar",
    props: {
      img: String,
      initials: String,
      size: {
        type: Number,
        default: 40
      },
      color: {
        type: String,
        default: "info"
      },
      content: String,
      captionSize: String,
      captionWeight: String
    }
  };
</script>
