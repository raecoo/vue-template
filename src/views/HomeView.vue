<script setup lang="ts">
import TheWelcome from "../components/TheWelcome.vue";

import {
  Layout,
  LayoutSection,
  AnnotatedSection,
  TextField,
} from "@ownego/polaris-vue";
</script>

<template>
  <Page fullWidth="true" subtitle="Perfect for any pet">
    <!-- <template #pageTitle>
    <Badge status="success">Paid</Badge>
  </template>
  <template #primaryAction>
    <Button primary :connectedDisclosure="{ accessibilityLabel: 'Other save actions', actions: [{content: 'Save as new'}] }">
      Save
    </Button>
  </template> -->

    <Card sectioned>
      <template #title>Credit card</template>
      <p>Credit card information</p>
    </Card>

  </Page>
</template>
