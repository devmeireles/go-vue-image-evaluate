import type { App } from "vue";
import { createI18n, type LocaleMessages, type VueMessageType } from "vue-i18n";

import { useCoreStore } from "@/stores/core";
import portuguese from "@/locales/pt.json";
import english from "@/locales/en.json";

const messages: LocaleMessages<VueMessageType> = {
  en: { ...english },
  pt: { ...portuguese }
};

export function installI18n(app: App<Element>) {
  const i18n = createI18n({
    locale: useCoreStore().language,
    fallbackLocale: "en",
    messages
  });
  app.use(i18n);
  return i18n;
}
