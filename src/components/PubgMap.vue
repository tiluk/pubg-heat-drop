<script setup lang="ts">
import { onMounted } from 'vue'

const props = defineProps<{
  map: string
}>()

const toMapPath = (map: string) => `src/assets/maps/${map.toLocaleLowerCase()}/Low_Res.png`

onMounted(() => {})
</script>

<template>
  <v-img :src="toMapPath(props.map)"></v-img>
</template>

<style scoped lang="scss">
.v-img {
  height: 100vh;
}
</style>
