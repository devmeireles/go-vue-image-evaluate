import { defineStore } from "pinia";

export const useUserStore = defineStore({
  id: "user",
  state: () => ({
    user: {
      initials: "GM",
      fullName: "Gabriel Meireles",
      email: "gabriel@temp.com",
      image: "https://cdn.vuetifyjs.com/images/john.jpg"
    }
  }),
  getters: {
    currentUser: (state) => state.user
  },
  actions: {}
});
