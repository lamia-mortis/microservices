<script setup>
import { ref } from "vue";
import IconItem from "./IconItem.vue";
</script>

<script>
export default {
  name: "MenuSideBar",
};
</script>

<style scoped>
.menu-side-bar-container {
  width: 250px;
  height: 100%;
  background-color: #0c2a32;
}

nav {
  display: flex;
  flex-direction: column;
  height: 100px;
}

nav:hover {
  background-color: #3c6c7f;
}

a {
  text-decoration: none;
  color: white;
  margin-left: auto;
  margin-right: auto;
  flex-grow: 2;
}

svg {
  margin: auto;
  padding-top: 10px;
  flex-grow: 1;
}
</style>

<template>
  <div class="menu-side-bar-container">
    <nav>
      <IconItem name="home" />
      <router-link to="/">Home</router-link>
    </nav>
    <nav>
      <IconItem name="transactions" />
      <router-link to="/transactions">Transactions</router-link>
    </nav>
    <nav>
      <IconItem name="blog" />
      <router-link to="/blog">Blog</router-link>
    </nav>
    <nav>
      <IconItem name="aboutPage" />
      <router-link to="/aboutpage">About us</router-link>
    </nav>
    <nav>
      <IconItem name="location" />
      <router-link to="/location">Location</router-link>
    </nav>
  </div>
</template>
