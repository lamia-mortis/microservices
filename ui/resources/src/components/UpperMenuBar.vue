<script setup>
import { ref } from "vue";
import IconItem from "./IconItem.vue";
</script>

<script>
export default {
  name: "UpperMenuBar",
};
</script>

<style scoped>
.upper-menu-bar-container {
  display: flex;
  position: absolute;
  background-color: red;
  margin-left: 220px;
}
</style>

<template>
  <div class="upper-menu-bar-container">
    <nav>LOGIN</nav>
    <nav>SIGN UP</nav>
    <nav>MAILING LIST</nav>
    <nav>SHARE</nav>
  </div>
</template>
