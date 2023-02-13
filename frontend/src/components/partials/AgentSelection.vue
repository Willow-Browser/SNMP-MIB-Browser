<script setup lang="ts">
import {
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
} from "@headlessui/vue";
import Check from "~icons/mdi/check";
import MenuSwap from "~icons/mdi/menu-swap";
import { ref, watch } from "vue";
import { EventsEmit } from "../../../wailsjs/runtime/runtime";

const people = [
  { id: 1, name: "Agent 1", unavailable: false },
  { id: 2, name: "Agent 2", unavailable: false },
  { id: 3, name: "Agent 3", unavailable: false },
  { id: 4, name: "Agent 4", unavailable: true },
  { id: 5, name: "Agent 5", unavailable: false },
];
const selectedPerson = ref(people[0]);

watch(selectedPerson, () => {
  EventsEmit("selectedAgent", selectedPerson.value);
});
</script>

<template>
  <p class="px-2 font-semibold text-gray-800">Agent:</p>
  <div class="w-72">
    <Listbox v-model="selectedPerson">
      <div class="relative mt-1">
        <ListboxButton
          class="relative w-full cursor-default rounded-lg bg-white py-2 pl-3 pr-10 text-left shadow-md focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
        >
          <span class="block truncate">{{ selectedPerson.name }}</span>
          <span
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2"
          >
            <MenuSwap
              height="20"
              width="20"
              class="text-gray-400"
              aria-hidden="true"
            />
          </span>
        </ListboxButton>
        <Transition
          enter-active-class="transition duration-100 ease-out"
          enter-from-class="opacity-0"
          enter-to-class="opacity-100"
          leave-active-class="transition duration-100 ease-in"
          leave-from-class="opacity-100"
          leave-to-class="opacity-0"
        >
          <ListboxOptions
            class="absolute z-50 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
          >
            <ListboxOption
              v-for="person in people"
              :key="person.id"
              v-slot="{ active, selected }"
              :value="person"
              :disabled="person.unavailable"
            >
              <li
                :class="[
                  active ? 'bg-amber-100 text-amber-900' : 'text-gray-900',
                  'relative cursor-default select-none py-2 pl-10 pr-4 text-left',
                ]"
              >
                <span
                  :class="[
                    selected ? 'font-medium' : 'font-normal',
                    'block truncate',
                  ]"
                  >{{ person.name }}</span
                >
                <span
                  v-if="selected"
                  class="absolute inset-y-0 left-0 flex items-center pl-3 text-amber-600"
                >
                  <Check height="20" width="20" aria-hidden="true" />
                </span>
              </li>
            </ListboxOption>
          </ListboxOptions>
        </Transition>
      </div>
    </Listbox>
  </div>
</template>
