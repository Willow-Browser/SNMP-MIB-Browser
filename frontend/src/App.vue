<script lang="ts" setup>
import HelloWorld from "@/components/HelloWorld.vue";
import OidTree from "@/components/OidTree.vue";
</script>

<template>
  <!-- <DiscordSideBar /> -->
  <OidTree />
  <div class="overflow-hidden">
    <div class="logo-box">
      <img class="logo vite" src="./assets/vite.svg" />
      <img class="logo electron" src="./assets/electron.svg" />
      <img class="logo vue" src="./assets/vue.svg" />
    </div>
    <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" /> -->
    <HelloWorld msg="Hello Vue 3 + TypeScript + Vite" />
    <div class="static-public">
      Place static files in the <code>/public/</code> folder
      <img style="width: 77px" :src="'./node.png'" />
    </div>
  </div>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.logo-box {
  display: flex;
  width: 100%;
  justify-content: center;
}

.static-public {
  display: flex;
  align-items: center;
  justify-content: center;
}
.static-public code {
  background-color: #eee;
  padding: 2px 4px;
  margin: 0 4px;
  border-radius: 4px;
  color: #304455;
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: 0.75s;
}

.logo.vite:hover {
  filter: drop-shadow(0 0 2em #747bff);
}

.logo.electron:hover {
  filter: drop-shadow(0 0 2em #9feaf9);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #249b73);
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
