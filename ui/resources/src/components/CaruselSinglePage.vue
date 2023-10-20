<script setup>
import { ref } from "vue";
import { defineProps } from "vue";
const props = defineProps({ imgLink: String });
</script>

<style scoped>
.single-page-carusel {
  display: flex;
  width: 100%;
  height: 100%;
  background-color: rgb(243, 29, 207);
}

.carusel-page-title {
  position: absolute;
  background-color: blue;
}

img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>

<template>
  <div class="single-page-carusel">
    <div class="carusel-page-title">CARUSEL SINGLE PAGE</div>
    <img :src="imgLink" alt="caruselImage" />
  </div>
</template>
