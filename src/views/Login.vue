<template>
  <div>
    <h1>Login</h1>
    <button @click="login">Log in with Shopify</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  methods: {
    login() {
      const redirectUrl = window.location.origin + '/auth/callback'
      const returnUrl = window.location.origin + '/dashboard'

      this.$appBridge.dispatch('Redirect', {
        path: '/auth',
        query: {
          redirectUrl,
          returnUrl,
        },
      })
    },
  },
})
</script>
