import { defineStore } from "pinia";

export const useCoreStore = defineStore({
  id: "core",
  state: () => ({
    isOpenSidebar: true,
    snackbar: {
      isOpen: false,
      message: ""
    },
    theme: "light",
    language: "en"
  }),
  getters: {
    sidebarState: (state) => state.isOpenSidebar,
    snackbarState: (state) => state.snackbar,
    currentTheme: (state) => state.theme,
    currentLanguage: (state) => state.language
  },
  actions: {
    toggleSidebar() {
      this.isOpenSidebar = !this.isOpenSidebar;
    },
    toggleSnackbar() {
      this.snackbar.isOpen = !this.snackbar.isOpen;
      this.snackbar.message = "";
    },
    openSnackbar(message: string) {
      this.snackbar.isOpen = true;
      this.snackbar.message = message;
    },
    toggleTheme() {
      this.theme = this.theme === "light" ? "dark" : "light";
    },
    toggleLanguage() {
      this.language = this.language === "en" ? "pt" : "en";
      console.log(this.language);
    }
  }
  // persist: true
});
