<script setup lang="ts">
import { RouterView } from "vue-router";
import Header from "./components/Header.vue";
</script>

<template>
  <AppProvider>
    <Header />
    <RouterView />
    <FooterHelp>Copy Right is here</FooterHelp>
  </AppProvider>
</template>

<style scoped></style>
