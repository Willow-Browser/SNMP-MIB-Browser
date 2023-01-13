<script lang="ts" setup>
import { ParseMib } from "../../wailsjs/go/main/App";
</script>

<template>
  <button
    class="relative flex items-center justify-center h-12 w-12 mt-2 mb-2 mx-auto bg-gray-400 hover:bg-green-600 dark:bg-gray-800 text-green-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg group"
    @click="ParseMib"
  >
    <slot></slot>
    <span
      class="absolute w-auto p-2 m-2 min-w-max left-14 rounded-md shadow-md text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100"
    >
      tooltip 💡
    </span>
  </button>
</template>
