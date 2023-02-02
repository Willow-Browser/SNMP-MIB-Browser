<script lang="ts" setup>
import TreeMenu from "@/components/partials/TreeMenu.vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { GetCurrentOids } from "../../wailsjs/go/main/App";
import { OidTree, TreeSorter } from "../utils/treeBuilder";
import { ref, reactive, onBeforeMount } from "vue";

const updateCounter = ref(0);

onBeforeMount(ReloadMibTree);

async function ReloadMibTree() {
  otherOidTree.oidTree = new TreeSorter(await GetCurrentOids()).createOidTree();
  updateCounter.value++;
}

const otherOidTree = reactive({
  oidTree: { name: "oids loading...", oid: "place2" } as OidTree,
});

EventsOn("mibsLoaded", ReloadMibTree);
</script>

<template>
  <div
    class="fixed top-16 left-0 flex h-screen w-1/3 flex-col bg-gray-100 shadow-lg"
  >
    <div
      id="title"
      class="flex h-14 bg-gradient-to-r from-indigo-400 to-white text-left"
    >
      <h2 class="flex items-center text-2xl text-black">SNMP Oids</h2>
    </div>
    <div class="scrollbar flex justify-start overflow-auto text-left">
      <TreeMenu class="" :node="otherOidTree.oidTree" :depth="0"></TreeMenu>
    </div>
  </div>
</template>
