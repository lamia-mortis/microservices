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
  position: fixed;
  justify-content: space-between;
  width: 80%;
  height: 30px;
  margin-left: 220px;
}

nav {
  display: flex;
  margin-left: 10%;
  height: fit-content;
  margin-top: 4px;
}

span {
  display: inline-block;
  margin-top: 10px;
  margin-left: 5px;
}
</style>

<template>
  <div class="upper-menu-bar-container">
    <nav class="login-nav">
      <IconItem name="login" />
      <span class="login">LOGIN</span>
    </nav>
    <nav class="signup-nav">
      <IconItem name="signup" />
      <span>SIGN UP</span>
    </nav>
    <nav class="mailing-list-nav">
      <IconItem name="mail" />
      <span>MAILING LIST</span>
    </nav>
    <nav class="shar-nav">
      <IconItem name="share" />
      <span>SHARE</span>
    </nav>
  </div>
</template>
