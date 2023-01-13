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
    class="fixed top-0 left-0 w-1/3 h-screen flex flex-col bg-gray-100 shadow-lg"
  >
    <div
      id="title"
      class="bg-gradient-to-r from-indigo-400 to-white h-14 text-left flex"
    >
      <h2 class="text-2xl text-black flex items-center">SNMP Oids</h2>
    </div>
    <div class="flex justify-start text-left overflow-auto scrollbar">
      <TreeMenu class="" :node="otherOidTree.oidTree" :depth="0"></TreeMenu>
    </div>
  </div>
</template>
