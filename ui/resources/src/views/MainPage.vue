<script setup>
import { ref } from "vue";
import MenuSideBar from "../components/MenuSideBar";
import UpperMenuBar from "../components/UpperMenuBar";
import CaruselComponent from "../components/CaruselComponent";
</script>

<style scoped>
.main-page-container {
  display: flex;
  height: 100vh;
}
</style>

<template>
  <div class="main-page-container">
    <MenuSideBar />
    <UpperMenuBar />
    <router-view />
  </div>
</template>
