<script setup>
import { reactive } from "vue";
import CaruselSinglePage from "./CaruselSinglePage.vue";
import caruselPhoto1 from "../../assets/caruselPhoot1.png";
import caruselPhoto2 from "../../assets/caruselPhoto2.png";
import caruselPhoto3 from "../../assets/caruselPhoto3.png";

const images = [caruselPhoto1, caruselPhoto2, caruselPhoto3];

const state = reactive({
  currentCaruselImg: images[0],
  toggledBtnIndex: 0,
  unToggledBtnClass: "carusel-scroll-button",
});

function changeImg(index) {
  state.currentCaruselImg = images[index];
}

function changeButtonColor(event) {
  state.toggledBtnIndex = +event.currentTarget.id;
}
</script>

<style scoped>
.main-page-carusel {
  width: 100%;
  background-color: rgb(73, 32, 66);
}

.carusel-scroll-buttons-container {
  position: absolute;
  bottom: 50px;
  margin-left: 40%;
}

.carusel-scroll-button {
  border: 2px solid black;
  border-radius: 50%;
  height: 20px;
  width: 20px;
  cursor: pointer;
}

.carusel-button-clicked {
  border: 2px solid black;
  border-radius: 50%;
  height: 20px;
  width: 20px;
  cursor: pointer;
  background-color: rgb(201, 130, 0);
}
</style>

<template>
  <div class="main-page-carusel">
    <CaruselSinglePage :imgLink="state.currentCaruselImg" />
    <div class="carusel-scroll-buttons-container">
      <button
        v-for="(image, index) in images"
        :class="[
          state.toggledBtnIndex === index
            ? 'carusel-button-clicked'
            : 'carusel-scroll-button',
        ]"
        @click="
          (event) => {
            changeImg(index);
            changeButtonColor(event);
          }
        "
        :key="index"
        :id="index"
      ></button>
    </div>
  </div>
</template>
