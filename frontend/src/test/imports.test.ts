import { describe, expect, test } from "vitest";

describe("import vue components", async () => {
  test("normal imports as expected", async () => {
    const cmp = await import("../components/partials/TreeMenu.vue");
    expect(cmp).toBeDefined();
  });

  test("template string imports as expected", async () => {
    // eslint-disable-next-line @typescript-eslint/quotes
    const cmp = await import(`../components/partials/TreeMenu.vue`);
    expect(cmp).toBeDefined();
  });

  test("dynamic imports as expected", async () => {
    const name = "TreeMenu";
    const cmp = await import(`../components/partials/${name}.vue`);
    expect(cmp).toBeDefined();
  });
});
