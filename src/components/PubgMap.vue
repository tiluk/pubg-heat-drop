<script setup lang="ts">
import * as L from 'leaflet'
import { nextTick, onMounted, ref, watch } from 'vue'
import 'leaflet.heat'
import 'leaflet/dist/leaflet.css'

const props = defineProps<{
  selectedMap: string
}>()
const leafletMap = ref<L.Map>()
const heat = ref<L.HeatLayer>()
const latlngs = ref<L.LatLng[]>([])
const playerCount = ref<number>(100)

const toMapPath = (map: string) => `src/assets/maps/${map.toLocaleLowerCase()}/Low_Res.png`

const initMap = () => {
  const bounds: L.LatLngBoundsLiteral = [
    [0, 0],
    [1000, 1000]
  ]
  leafletMap.value = L.map('pubgmap', {
    crs: L.CRS.Simple,
    minZoom: 0,
    maxZoom: 2,
    attributionControl: false,
    zoomControl: false,
    maxBoundsViscosity: 0.9,
    maxBounds: bounds,
    doubleClickZoom: false
  })
  L.imageOverlay(toMapPath(props.selectedMap), bounds).addTo(leafletMap.value)
  leafletMap.value.fitBounds(bounds)

  heat.value = L.heatLayer([], {
    radius: 25
  }).addTo(leafletMap.value)

  leafletMap.value.on('click', (e: L.LeafletMouseEvent) => {
    e.latlng.alt = 25 / playerCount.value
    latlngs.value.push(e.latlng)
    heat.value?.setLatLngs(latlngs.value)
  })
}

onMounted(() => {
  nextTick(() => {
    initMap()
  })
})

watch(
  () => props.selectedMap,
  () => {
    leafletMap.value?.remove()
    heat.value?.remove()
    initMap()
  }
)
</script>

<template>
  <div id="pubgmap" class="mapContainer"></div>
</template>

<style scoped lang="scss">
.mapContainer {
  height: 100vh;
  overflow: hidden;
  background-color: black;
}
</style>
