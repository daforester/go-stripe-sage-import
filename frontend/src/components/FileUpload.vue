<script lang="ts" setup>

import { OpenFile } from "../../wailsjs/go/main/App";

async function pickFile() {
  const path = await OpenFile();
  console.log(path);
}

</script>

<template>
  <button @click="pickFile">Load Stripe File</button>
</template>

<style scoped>
button {
  background: green;
  color: white;
  cursor: pointer;
  font-weight: bold;
  height: 100vh;
  width: 100vw;
}
</style>
