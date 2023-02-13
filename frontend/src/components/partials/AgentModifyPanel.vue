<script setup lang="ts">
import { ref, watch } from "vue";
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import WindowClose from "~icons/mdi/window-close";
import LinkVariant from "~icons/mdi/link-variant";

const DEFAULT_AGENT_ADDRESS = "";
const DEFAULT_AGENT_PORT = "";

const agentAddress = ref(DEFAULT_AGENT_ADDRESS);
const agentPort = ref(DEFAULT_AGENT_PORT);

function submit(payload: MouseEvent) {
  console.log(payload);
}

function onClick(_e: Event) {
  open.value = !open.value;
  agentAddress.value = DEFAULT_AGENT_ADDRESS;
  agentPort.value = DEFAULT_AGENT_PORT;
}

const open = ref(true);
</script>

<template>
  <input
    id="agent-modify-modal"
    type="checkbox"
    class="fixed h-0 w-0 appearance-none opacity-0"
    @change="onClick"
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
