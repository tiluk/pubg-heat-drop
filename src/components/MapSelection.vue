<script setup lang="ts">
const maps = [
  {
    name: 'Erangel',
    image: '/maps/erangel/Preview.png'
  },
  {
    name: 'Miramar',
    image: '/maps/miramar/Preview.png'
  },
  {
    name: 'Sanhok',
    image: '/maps/sanhok/Preview.png'
  },
  {
    name: 'Vikendi',
    image: '/maps/vikendi/Preview.png'
  }
]

const toPreviewPath = (map: string) => `src/assets/maps/${map.toLocaleLowerCase()}/Preview.png`
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>Map Selection:</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-card
        v-for="map in maps"
        :key="map.name"
        class="mx-2 my-2 map-card"
        :image="toPreviewPath(map.name)"
        max-width="340"
        :title="map.name"
      >
        <template v-slot:actions>
          <v-btn text="Switch To Map" variant="elevated" color="primary" block></v-btn>
        </template>
      </v-card>
    </v-row>
  </v-container>
</template>

<style scoped lang="scss">
.map-card {
  color: white;
  mix-blend-mode: difference;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.75);
}
</style>
