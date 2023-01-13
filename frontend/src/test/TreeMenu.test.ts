import { describe, expect, it, test } from "vitest";
import { mount } from "@vue/test-utils";
import TreeMenu from "../components/partials/TreeMenu.vue";
import { OidTree } from "../utils/treeBuilder";

test("mount component", async () => {
  expect(TreeMenu).toBeTruthy();

  const oidTree: OidTree = {
    name: "iso",
    oid: ".1",
    type: "ObjectIdentity",
    children: [
      {
        name: "org",
        oid: ".1.3",
        type: "ObjectIdentity",
        children: [
          {
            name: "dod",
            oid: ".1.3.6",
            type: "ModuleIdentity",
          },
        ],
      },
    ],
  };

  const wrapper = mount(TreeMenu, {
    props: {
      node: oidTree,
      depth: 1,
    },
  });

  // TODO : test for specific icon in the component

  expect(wrapper.text()).toContain("iso");
});
