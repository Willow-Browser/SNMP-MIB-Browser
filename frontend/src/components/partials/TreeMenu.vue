<script lang="ts" setup>
import { computed, ref } from "vue";
import PlusBoxOutline from "~icons/mdi/plus-box-outline";
import MinusBoxOutline from "~icons/mdi/minus-box-outline";
import { OidTree } from "../../utils/treeBuilder";

const showChildren = ref(false);

const props = defineProps<{
  node: OidTree;
  depth: number;
}>();

const indent = computed(() => {
  return { transform: `translate(${props.depth * 20}px)` };
});

function toggleChildren() {
  showChildren.value = !showChildren.value;
}

function cursorClass(): string {
  if (props.node.children !== undefined) {
    return "cursor-pointer";
  } else {
    return "cursor-default";
  }
}

function hasChildren(): boolean {
  return props.node.children !== undefined;
}

function calculatePadding(): string {
  let padding = "pl-1";
  if (!hasChildren()) {
    padding = "pl-6";
  }

  return padding;
}
</script>

<template>
  <div>
    <div class="pb-1 mb-1" @click="toggleChildren">
      <div :style="indent" :class="cursorClass()" class="flex text-gray-900">
        <PlusBoxOutline
          v-if="hasChildren() && !showChildren"
          height="20"
          width="20"
        />
        <MinusBoxOutline
          v-else-if="hasChildren() && showChildren"
          height="20"
          width="20"
        />
        <p :class="calculatePadding()">
          {{ node.name }}
        </p>
      </div>
    </div>
    <div v-if="showChildren">
      <TreeMenu
        v-for="oid in node.children"
        :key="oid.oid"
        :node="oid"
        :depth="depth + 1"
      ></TreeMenu>
    </div>
  </div>
</template>
