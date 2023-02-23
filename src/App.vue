<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import Header from "./components/Header.vue";
import HelloWorld from "./components/HelloWorld.vue";

</script>

<template>
  <AppProvider>
    <Header msg="sss" ></Header>
    <RouterView />
    <FooterHelp>Footer help content</FooterHelp>
  </AppProvider>
</template>

<style scoped>

</style>
