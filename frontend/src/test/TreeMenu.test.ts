import { describe, expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import TreeMenu from "../components/partials/TreeMenu.vue";
import { OidTree } from "../utils/treeBuilder";

describe("Test TreeMenu", () => {
  test("mount component", async () => {
    expect(TreeMenu).toBeTruthy();

    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "ObjectIdentity",
      access: "",
      syntax: "",
      is_index: false,
      is_row: false,
      children: [
        {
          name: "org",
          oid: ".1.3",
          type: "ObjectIdentity",
          access: "",
          syntax: "",
          is_index: false,
          is_row: false,
          children: [
            {
              name: "dod",
              oid: ".1.3.6",
              type: "ModuleIdentity",
              access: "",
              syntax: "",
              is_index: false,
              is_row: false,
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

  test("Leaf icon used for ReadOnly", () => {
    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "ObjectType",
      access: "read-only",
      syntax: "",
      is_index: false,
      is_row: false,
    };

    const wrapper = mount(TreeMenu, {
      props: {
        node: oidTree,
        depth: 1,
      },
    });

    expect(wrapper.get("svg").classes()).toStrictEqual(["leaf"]);
  });

  test("Folder icon used for ObjectIdentity", () => {
    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "ObjectIdentity",
      access: "",
      syntax: "",
      is_index: false,
      is_row: false,
    };

    const wrapper = mount(TreeMenu, {
      props: {
        node: oidTree,
        depth: 1,
      },
    });

    expect(wrapper.get("svg").classes()).toStrictEqual(["folder"]);
  });

  test("Folder icon used for ModuleIdentity", () => {
    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "ModuleIdentity",
      access: "",
      syntax: "",
      is_index: false,
      is_row: false,
    };

    const wrapper = mount(TreeMenu, {
      props: {
        node: oidTree,
        depth: 1,
      },
    });

    expect(wrapper.get("svg").classes()).toStrictEqual(["folder"]);
  });

  test("Pen icon used for ReadWrite OID", () => {
    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "ObjectType",
      access: "read-write",
      syntax: "",
      is_index: false,
      is_row: false,
    };

    const wrapper = mount(TreeMenu, {
      props: {
        node: oidTree,
        depth: 0,
      },
    });

    expect(wrapper.get("svg").classes()).toStrictEqual(["pen"]);
  });

  test("LightningBolt icon used for NotificationType", () => {
    const oidTree: OidTree = {
      name: "iso",
      oid: ".1",
      type: "NotificationType",
      access: "",
      syntax: "",
      is_index: false,
      is_row: false,
    };

    const wrapper = mount(TreeMenu, {
      props: {
        node: oidTree,
        depth: 0,
      },
    });

    expect(wrapper.get("svg").classes()).toStrictEqual(["lightning"]);
  });
});
