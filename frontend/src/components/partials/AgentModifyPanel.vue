<script setup lang="ts">
import Divider from "./Divider.vue";
import { ref } from "vue";
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  TransitionChild,
  TransitionRoot,
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
} from "@headlessui/vue";
import WindowClose from "~icons/mdi/window-close";
// import LinkVariant from "~icons/mdi/link-variant";
import MenuSwap from "~icons/mdi/menu-swap";
import Check from "~icons/mdi/check";
import { EventsEmit } from "../../../wailsjs/runtime/runtime";

// TODO : create a component to split all of the dropdowns up

const DEFAULT_AGENT_ADDRESS = "";
const DEFAULT_AGENT_PORT = "";

const agentAddress = ref(DEFAULT_AGENT_ADDRESS);
const agentPort = ref(DEFAULT_AGENT_PORT);
const readCommunity = ref("");
const writeCommunity = ref("");
const usmUserName = ref("");
const authKey = ref("");
const privKey = ref("");

const agentTypes = [
  { id: 1, name: "Version 1", unavailable: true },
  { id: 2, name: "Version 2", unavailable: true },
  { id: 3, name: "Version 3", unavailable: false },
];
const selectedAgentType = ref(agentTypes[2]);

const authTypes = [
  { id: 1, name: "noAuthNoPriv", unavailable: false },
  { id: 2, name: "authNoPriv", unavailable: true },
  { id: 3, name: "authPriv", unavailable: true },
];
const selectedAuthType = ref(authTypes[0]);

function submit(payload: MouseEvent) {
  console.log(payload);

  const obj = {
    agentAddress: agentAddress.value,
    agentPort: agentPort.value,
    agentType: selectedAgentType.value,
    readCommunity: readCommunity.value,
    writeCommunity: writeCommunity.value,
    authType: selectedAuthType.value,
    usmUserName: usmUserName.value,
    authKey: authKey.value,
    privKey: privKey.value,
  };

  EventsEmit("createAgent", obj);
}

function onClick() {
  open.value = !open.value;
  agentAddress.value = DEFAULT_AGENT_ADDRESS;
  agentPort.value = DEFAULT_AGENT_PORT;
}

function IsAuthKeyDisabled() {
  if (selectedAuthType.value.id === 1) {
    return true;
  }

  return false;
}

function IsPrivKeyDisabled() {
  if (selectedAuthType.value.id === 3) {
    return false;
  }

  return true;
}

const open = ref(true);
</script>

<template>
  <input
    id="agent-modify-modal"
    type="checkbox"
    class="fixed h-0 w-0 appearance-none opacity-0"
    @change="onClick()"
  />
  <!-- TODO : do not close the modal by clicking outside the message box -->
  <div class="text-left text-black">
    <TransitionRoot as="template" :show="open">
      <Dialog as="div" class="relative z-10" @close="open = false">
        <div class="fixed inset-0" />

        <div class="fixed inset-0 overflow-hidden">
          <div class="absolute inset-0 overflow-hidden">
            <div
              class="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10 sm:pl-16"
            >
              <TransitionChild
                as="template"
                enter="transform transition ease-in-out duration-300"
                enter-from="translate-x-full"
                enter-to="translate-x-0"
                leave="transform transition ease-in-out duration-300"
                leave-from="translate-x-0"
                leave-to="translate-x-full"
              >
                <DialogPanel class="pointer-events-auto w-screen max-w-md">
                  <form
                    class="flex h-full flex-col divide-y divide-gray-200 bg-white shadow-xl"
                  >
                    <div class="h-0 flex-1 overflow-y-auto">
                      <div class="bg-indigo-700 py-6 px-4 sm:px-6">
                        <div class="flex items-center justify-between">
                          <DialogTitle class="text-lg font-medium text-white">
                            Modify Manager
                          </DialogTitle>
                          <div class="ml-3 flex h-7 items-center">
                            <button
                              type="button"
                              class="rounded-md bg-indigo-700 text-indigo-200 hover:text-white focus:outline-none focus:ring-2 focus:ring-white"
                              @click="open = false"
                            >
                              <span class="sr-only">Close Panel</span>
                              <WindowClose aria-hidden="true" class="h-6 w-6" />
                            </button>
                          </div>
                        </div>
                        <!-- <div class="mt-1">
                          <p class="text-sm text-indigo-300">
                            Get started by filling in the information below to
                            create your new project.
                          </p>
                        </div> -->
                      </div>
                      <div class="flex flex-1 flex-col justify-between">
                        <div class="divide-y divide-gray-200 px-4 sm:px-6">
                          <div class="space-y-6 pt-6 pb-5">
                            <div>
                              <label
                                for="project-name"
                                class="block text-sm font-medium text-gray-900"
                              >
                                Agent Address
                              </label>
                              <div class="mt-1">
                                <input
                                  id="project-name"
                                  v-model="agentAddress"
                                  type="text"
                                  name="project-name"
                                  class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                />
                              </div>
                            </div>
                            <div>
                              <label
                                for="description"
                                class="block text-sm font-medium text-gray-900"
                              >
                                Agent Port
                              </label>
                              <div class="mt-1">
                                <input
                                  id="description"
                                  v-model="agentPort"
                                  name="description"
                                  type="text"
                                  class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                />
                              </div>
                            </div>
                            <div>
                              <label
                                for="selection"
                                class="block text-sm font-medium text-gray-900"
                              >
                                Agent Type
                              </label>
                              <Listbox
                                id="selection"
                                v-model="selectedAgentType"
                              >
                                <div class="relative mt-1">
                                  <ListboxButton
                                    class="relative block w-full cursor-default rounded-md border-[1px] border-gray-300 bg-white py-2 pl-3 pr-10 text-left shadow-sm focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
                                  >
                                    <span class="block truncate">
                                      {{ selectedAgentType.name }}</span
                                    >
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
                                  <transition
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
                                        v-for="agentType in agentTypes"
                                        :key="agentType.id"
                                        v-slot="{ active, selected }"
                                        :value="agentType"
                                        :disabled="agentType.unavailable"
                                      >
                                        <li
                                          :class="[
                                            active
                                              ? 'bg-amber-100 text-amber-900'
                                              : 'text-gray-900',
                                            'relative cursor-default select-none py-2 pl-10 pr-4 text-left',
                                          ]"
                                        >
                                          <span
                                            :class="[
                                              selected
                                                ? 'font-medium'
                                                : 'font-normal',
                                              'block truncate',
                                            ]"
                                            >{{ agentType.name }}</span
                                          >
                                          <span
                                            v-if="selected"
                                            class="absolute inset-y-0 left-0 flex items-center pl-3 text-amber-600"
                                          >
                                            <Check
                                              height="20"
                                              width="20"
                                              aria-hidden="true"
                                            />
                                          </span>
                                        </li>
                                      </ListboxOption>
                                    </ListboxOptions>
                                  </transition>
                                </div>
                              </Listbox>
                            </div>
                            <Divider />
                            <div
                              v-show="
                                selectedAgentType.id === 1 ||
                                selectedAgentType.id === 2
                              "
                            >
                              <div>
                                <label
                                  for="readCommunity"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  Read Community
                                </label>
                                <div class="mt-1">
                                  <input
                                    id="readCommunity"
                                    v-model="readCommunity"
                                    name="readCommunity"
                                    type="text"
                                    class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                  />
                                </div>
                              </div>
                              <div>
                                <label
                                  for="writeCommunity"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  Write Community
                                </label>
                                <div class="mt-1">
                                  <input
                                    id="writeCommunity"
                                    v-model="writeCommunity"
                                    name="writeCommunity"
                                    type="text"
                                    class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                  />
                                </div>
                              </div>
                            </div>
                            <div
                              v-show="selectedAgentType.id === 3"
                              class="space-y-6 pb-5"
                            >
                              <div>
                                <label
                                  for="usmUserName"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  USM User Name
                                </label>
                                <div class="mt-1">
                                  <input
                                    id="usmUserName"
                                    v-model="usmUserName"
                                    name="usmUserName"
                                    type="text"
                                    class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                  />
                                </div>
                              </div>
                              <div>
                                <label
                                  for="selection"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  Auth Type
                                </label>
                                <Listbox
                                  id="selection"
                                  v-model="selectedAuthType"
                                >
                                  <div class="relative mt-1">
                                    <ListboxButton
                                      class="relative block w-full cursor-default rounded-md border-[1px] border-gray-300 bg-white py-2 pl-3 pr-10 text-left shadow-sm focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
                                    >
                                      <span class="block truncate">
                                        {{ selectedAuthType.name }}</span
                                      >
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
                                    <transition
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
                                          v-for="authType in authTypes"
                                          :key="authType.id"
                                          v-slot="{ active, selected }"
                                          :value="authType"
                                          :disabled="authType.unavailable"
                                        >
                                          <li
                                            :class="[
                                              active
                                                ? 'bg-amber-100 text-amber-900'
                                                : 'text-gray-900',
                                              'relative cursor-default select-none py-2 pl-10 pr-4 text-left',
                                            ]"
                                          >
                                            <span
                                              :class="[
                                                selected
                                                  ? 'font-medium'
                                                  : 'font-normal',
                                                'block truncate',
                                              ]"
                                              >{{ authType.name }}</span
                                            >
                                            <span
                                              v-if="selected"
                                              class="absolute inset-y-0 left-0 flex items-center pl-3 text-amber-600"
                                            >
                                              <Check
                                                height="20"
                                                width="20"
                                                aria-hidden="true"
                                              />
                                            </span>
                                          </li>
                                        </ListboxOption>
                                      </ListboxOptions>
                                    </transition>
                                  </div>
                                </Listbox>
                              </div>
                              <div>
                                <label
                                  for="authKey"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  Auth Key
                                </label>
                                <div class="mt-1">
                                  <input
                                    id="authKey"
                                    v-model="authKey"
                                    name="authKey"
                                    type="text"
                                    :disabled="IsAuthKeyDisabled()"
                                    class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 disabled:border-gray-500 disabled:bg-gray-300 sm:text-sm"
                                  />
                                </div>
                              </div>
                              <div>
                                <label
                                  for="privKey"
                                  class="block text-sm font-medium text-gray-900"
                                >
                                  Priv Key
                                </label>
                                <div class="mt-1">
                                  <input
                                    id="privKey"
                                    v-model="privKey"
                                    name="privKey"
                                    type="text"
                                    :disabled="IsPrivKeyDisabled()"
                                    class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 disabled:border-gray-500 disabled:bg-gray-300 sm:text-sm"
                                  />
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="flex flex-shrink-0 justify-end px-4 py-4">
                      <button
                        type="button"
                        class="rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                        @click="open = false"
                      >
                        Cancel
                      </button>
                      <button
                        type="button"
                        class="ml-4 inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                        @click="submit"
                      >
                        Save
                      </button>
                    </div>
                  </form>
                </DialogPanel>
              </TransitionChild>
            </div>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>
