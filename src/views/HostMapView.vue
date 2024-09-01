<script setup lang="ts">
import PubgMap from '../components/PubgMap.vue'
import MapSelection from '../components/MapSelection.vue'
import LobbySettings from '../components/LobbySettings.vue'
</script>

<template>
  <main>
    <div class="main">
      <LobbySettings id="lobby-settings" />
      <PubgMap id="pubg-map" />
      <MapSelection id="map-selection" />
    </div>
  </main>
</template>

<style scoped>
.main {
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  gap: 0;
}

#lobby-settings {
  grid-column: span 3;
}

#pubg-map {
  grid-column: span 6;
}

#map-selection {
  grid-column: span 3;
}
</style>
