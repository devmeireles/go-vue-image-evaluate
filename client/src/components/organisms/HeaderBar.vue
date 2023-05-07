<template>
  <v-app-bar app elevation="1">
    <v-app-bar-nav-icon @click="toggleSidebar()"></v-app-bar-nav-icon>

    <v-spacer></v-spacer>

    <AvatarMenu
      :name="userData.fullName"
      :image="userData.image"
      :initials="userData.initials"
    />

    <v-btn icon @click="toggleTheme()">
      <v-icon>mdi-brightness-4</v-icon>
    </v-btn>

    <v-btn icon @click="toggleLanguage()">
      <v-icon>mdi-translate</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts">
import AvatarMenu from "@/components/organisms/AvatarMenu.vue";
export default {
  name: "HeaderBar",
  props: {
    toggleSidebar: {
      type: Function,
      default: () => { },
    },
    toggleTheme: {
      type: Function,
      default: () => { }
    },
    toggleLanguage: {
      type: Function,
      default: () => { },
    },
    userData: {
      type: Object,
      default: {}
    }
  },
  components: { AvatarMenu }
};
</script>
