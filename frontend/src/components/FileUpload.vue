<script lang="ts" setup>

import { OpenFile } from "../../wailsjs/go/main/App";

async function pickFile() {
  const path = await OpenFile();
  console.log(path);
}

</script>

<template>
  <main>
    <button @click="pickFile">Pick File</button>
  </main>
</template>

<style scoped>

</style>
