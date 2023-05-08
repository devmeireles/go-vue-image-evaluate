import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { VDataTableServer } from "vuetify/labs/components";
import defaults from "./default";
import theme from "./theme";

// Styles
import "@/assets/styles/main.scss";
import "@mdi/font/css/materialdesignicons.css";
import "vuetify/styles";

export default createVuetify({
  VDataTableServer,
  components,
  directives,
  defaults,
  theme
});
