import { defineStore } from "pinia";

export const useCoreStore = defineStore({
  id: "core",
  state: () => ({
    isOpenSidebar: true,
    theme: "light",
    language: "en"
  }),
  getters: {
    sidebarState: (state) => state.isOpenSidebar,
    currentTheme: (state) => state.theme,
    currentLanguage: (state) => state.language
  },
  actions: {
    toggleSidebar() {
      this.isOpenSidebar = !this.isOpenSidebar;
    },
    toggleTheme() {
      this.theme = this.theme === "light" ? "dark" : "light";
    },
    toggleLanguage() {
      this.language = this.language === "en" ? "pt" : "en";
      console.log(this.language);
    }
  },
  persist: true
});
