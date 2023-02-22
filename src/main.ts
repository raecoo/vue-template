import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import { createApp as createAppBridge } from "@shopify/app-bridge";
// import { getSessionToken } from "@shopify/app-bridge-utils";


import "./assets/main.css";

const app = createApp(App);

const shopifyHost = new URLSearchParams(location.search).get("host");

const appBridge = createAppBridge({
  apiKey: "e93b19f4d725828de9f94e7619d834e6",
  host: window.btoa(shopifyHost + "/admin") || "",
  forceRedirect: false,
});

app.config.globalProperties.$appBridge = appBridge;

app.use(createPinia());
app.use(router);

app.mount("#app");
