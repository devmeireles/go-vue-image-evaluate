<template>
  <v-menu bottom min-width="300px" rounded offset-y>
    <template v-slot:activator="{ props }">
      <v-btn icon v-bind="props">
        <UserAvatar :img="image" :initials="initials" />
      </v-btn>
    </template>
    <v-card width="300px">
      <v-list>
        <v-list-item>
          <UserAvatar
            :img="image"
            :content="name"
            :size="60"
            caption-size="text-body-1"
            caption-weight="font-weight-medium"
          />
        </v-list-item>
      </v-list>

      <v-divider></v-divider>

      <v-list density="compact">
        <v-list-item
          v-for="item in menu"
          :key="item.name"
          :value="item"
          active-color="primary"
        >
          <RouterLink :to="`${item.route}`" class="menu-link">
            <v-list-item start>
              <v-icon :icon="item.icon"></v-icon>
            </v-list-item>
            <v-list-item-title>{{ $t(item.name) }}</v-list-item-title>
          </RouterLink>
        </v-list-item>
      </v-list>
    </v-card>
  </v-menu>
</template>

<script lang="ts">
import { routes } from "@/consts/routes";
import UserAvatar from "../molecules/UserAvatar.vue";

export default {
  name: "AvatarMenu",
  props: {
    name: String,
    image: String,
    initials: String
  },
  data: () => ({
    menu: [
      {
        name: "core.account",
        route: routes.dashboard.main,
        icon: "mdi-account-outline"
      },
      {
        name: "core.settings",
        route: routes.dashboard.main,
        icon: "mdi-cog-outline"
      },
      {
        name: "core.logout",
        route: routes.dashboard.main,
        icon: "mdi-logout"
      }
    ]
  }),
  components: { UserAvatar }
};
</script>
