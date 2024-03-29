<script lang="ts" setup>
import { computed, ref } from "vue";
// import Folder from "~icons/mdi/folder";
import Key from "~icons/mdi/key-variant";
import Leaf from "~icons/mdi/leaf";
import Pen from "~icons/mdi/fountain-pen-tip";
import Table from "~icons/mdi/table-large";
import TableRow from "~icons/mdi/table-row";
import LightningBolt from "~icons/mdi/lightning-bolt";
import PlusCircle from "~icons/mdi/plus-circle-outline";
import FolderOutline from "~icons/mdi/folder-outline";
import { Icon } from "@iconify/vue";
import { OidTree } from "../../utils/treeBuilder";
import { EventsEmit, EventsOn } from "../../../wailsjs/runtime/runtime";
import { SendGetRequestWithOid } from "../../../wailsjs/go/main/App";

const showChildren = ref(false);
const isSelected = ref(false);

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

function isModuleIdentifier(): boolean {
  switch (props.node.type) {
    case "ObjectIdentity":
    case "ModuleIdentity":
      return true;
    default:
      return false;
  }
}

// TODO : conformance OIDs

function isObjectType(): boolean {
  return props.node.type === "ObjectType";
}

function isIndex(): boolean {
  return props.node.is_index;
}

function isTable(): boolean {
  if (props.node.syntax) {
    return props.node.syntax.includes("SEQUENCE OF");
  }
  return false;
}

function isRow(): boolean {
  return props.node.is_row;
}

function isReadOnly(): boolean {
  return props.node.access === "read-only";
}

function isReadWrite(): boolean {
  return props.node.access === "read-write";
}

function isReadCreate(): boolean {
  return props.node.access === "read-create";
}

function isNotificationType(): boolean {
  return props.node.type === "NotificationType";
}

function printType() {
  console.log(props.node);

  toggleChildren();
}

EventsOn("deselectItems", () => {
  if (isSelected.value) {
    isSelected.value = false;
  }
});

EventsOn("sendSelectedOids", () => {
  if (isSelected.value) {
    const oidString = props.node.oid + ".0";
    SendGetRequestWithOid(oidString);
  }
});

function onClick(payload: MouseEvent) {
  if (!payload.ctrlKey) {
    EventsEmit("deselectItems");
  }

  isSelected.value = true;
}
</script>

<template>
  <div>
    <div class="pb-1">
      <div
        :style="indent"
        :class="cursorClass()"
        class="flex text-gray-900"
        @dblclick="printType()"
      >
        <Icon
          v-if="hasChildren()"
          :icon="
            showChildren ? 'mdi:minus-box-outline' : 'mdi:plus-box-outline'
          "
          height="20"
          width="20"
          @click="toggleChildren"
        />
        <div :class="calculatePadding()" class="flex">
          <FolderOutline
            v-if="isModuleIdentifier()"
            class="folder"
            height="20"
            width="20"
          />
          <Key v-else-if="isIndex()" class="key" height="20" width="20" />
          <Table v-else-if="isTable()" class="table" height="20" width="20" />
          <TableRow v-else-if="isRow()" class="row" height="20" width="20" />
          <Leaf
            v-else-if="isObjectType() && isReadOnly()"
            class="leaf"
            height="20"
            width="20"
          />
          <Pen
            v-else-if="isObjectType() && isReadWrite()"
            class="pen"
            height="20"
            width="20"
          />
          <PlusCircle
            v-else-if="isObjectType() && isReadCreate()"
            class="plus"
            height="20"
            width="20"
          />
          <LightningBolt
            v-else-if="isNotificationType()"
            class="lightning"
            height="20"
            width="20"
          />
          <p
            class="ml-1 select-none pr-1"
            :class="isSelected ? 'bg-blue-600 text-white' : ''"
            @click="onClick"
          >
            {{ node.name }}
          </p>
        </div>
      </div>
    </div>
    <div v-show="showChildren">
      <TreeMenu
        v-for="oid in node.children"
        :key="oid.oid"
        :node="oid"
        :depth="depth + 1"
      ></TreeMenu>
    </div>
  </div>
</template>
