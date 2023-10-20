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
  flex-grow: 1;
}

svg {
  margin: auto;
  flex-grow: 1;
}

span {
  margin-left: auto;
  margin-right: auto;
}
</style>

<template>
  <div class="menu-side-bar-container">
    <router-link to="/">
      <nav>
        <IconItem name="home" />
        <span>Home</span>
      </nav>
    </router-link>
    <router-link to="/transactions">
      <nav>
        <IconItem name="transactions" />
        <span>Transactions</span>
      </nav>
    </router-link>
    <router-link to="/blog">
      <nav>
        <IconItem name="blog" />
        <span>Blog</span>
      </nav>
    </router-link>
    <router-link to="/aboutpage">
      <nav>
        <IconItem name="aboutPage" />
        <span>About us</span>
      </nav>
    </router-link>
    <router-link to="/location">
      <nav>
        <IconItem name="location" />
        <span>Location</span>
      </nav>
    </router-link>
  </div>
</template>
