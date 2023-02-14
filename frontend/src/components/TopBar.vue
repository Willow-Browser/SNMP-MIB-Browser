<script setup lang="ts">
import AgentSelection from "@/components/partials/AgentSelection.vue";
import AgentModifyPanel from "./partials/AgentModifyPanel.vue";
</script>

<template>
  <div class="bg-blue-500">
    <div
      class="flex h-16 w-full items-center justify-between border-b md:justify-start lg:border-none"
    >
      <AgentSelection />
      <label
        type="button"
        class="mx-3 inline-block cursor-pointer rounded-md bg-slate-800 py-2 px-4 text-base font-medium text-gray-50 shadow-sm"
        for="agent-modify-modal"
      >
        Hello
      </label>
      <AgentModifyPanel />
    </div>
  </div>
</template>
