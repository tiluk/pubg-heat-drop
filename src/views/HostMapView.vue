<script setup lang="ts">
import PubgMap from '../components/PubgMap.vue'
import MapSelection from '../components/MapSelection.vue'
import LobbySettings from '../components/LobbySettings.vue'
import { ref } from 'vue'

const currentMap = ref<string>('erangel')

const selectedMap = (map: string) => {
  currentMap.value = map
}
</script>

<template>
  <main>
    <v-container id="main">
      <v-row>
        <v-col cols="12" lg="3">
          <LobbySettings />
        </v-col>
        <v-col cols="12" lg="6">
          <PubgMap :selectedMap="currentMap" class="map" />
        </v-col>
        <v-col cols="12" lg="3">
          <MapSelection @map-selected="selectedMap" />
        </v-col>
      </v-row>
    </v-container>
  </main>
</template>

<style scoped lang="scss">
.main {
  height: 100vh;
  width: 100vw;
}

.map {
  border: 5px solid black;
}
</style>
