<script setup lang="ts">
import { ref } from "vue";
import AgentSelection from "@/components/partials/AgentSelection.vue";
import AgentModifyPanel from "./partials/AgentModifyPanel.vue";

const openModal = ref(false);
</script>

<template>
  <div class="bg-blue-500">
    <div
      class="flex h-16 w-full items-center justify-between border-b md:justify-start lg:border-none"
    >
      <AgentSelection />
      <button
        type="button"
        class="mx-3 inline-block rounded-md bg-slate-800 py-2 px-4 text-base font-medium text-gray-50 shadow-sm"
        @click="openModal = true"
      >
        Hello
      </button>
      <AgentModifyPanel :open-modal="openModal" />
    </div>
  </div>
</template>
