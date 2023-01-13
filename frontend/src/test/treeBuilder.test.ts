import { describe, expect, it } from "vitest";
import { oidstorage } from "../../wailsjs/go/models";
import { OidTree, TreeSorter } from "../utils/treeBuilder";

describe("class test suite", () => {
  it("sequential children", () => {
    const oids = Array<oidstorage.Oid>();

    const firstOid = new oidstorage.Oid('{"name": "iso","oid": ".1"}');
    const secondOid = new oidstorage.Oid('{"name": "org","oid": ".1.3"}');
    const thirdOid = new oidstorage.Oid('{"name": "dod","oid": ".1.3.6"}');

    oids.push(firstOid, thirdOid, secondOid);

    const obj = new TreeSorter(oids);

    const tree = obj.createOidTree();

    const expectedTree: OidTree = {
      name: "iso",
      oid: ".1",
      children: [
        {
          name: "org",
          oid: ".1.3",
          children: [
            {
              name: "dod",
              oid: ".1.3.6",
            },
          ],
        },
      ],
    };

    expect(tree).toEqual(expectedTree);
  });

  it("multiple children at single level", () => {
    const oids = Array<oidstorage.Oid>();

    const firstOid = new oidstorage.Oid('{"name": "iso","oid": ".1"}');
    const secondOid = new oidstorage.Oid('{"name": "org","oid": ".1.3"}');
    const thirdOid = new oidstorage.Oid('{"name": "dod","oid": ".1.3.6"}');
    const thirdOidChild1 = new oidstorage.Oid(
      '{"name": "thirdOidChild1","oid": ".1.3.6.1"}'
    );
    const thirdOidChild2 = new oidstorage.Oid(
      '{"name": "thirdOidChild2","oid": ".1.3.6.2"}'
    );
    const thirdOidChild3 = new oidstorage.Oid(
      '{"name": "thirdOidChild3","oid": ".1.3.6.3"}'
    );

    oids.push(
      firstOid,
      thirdOid,
      secondOid,
      thirdOidChild2,
      thirdOidChild1,
      thirdOidChild3
    );

    const obj = new TreeSorter(oids);

    const tree = obj.createOidTree();

    const expectedTree: OidTree = {
      name: "iso",
      oid: ".1",
      children: [
        {
          name: "org",
          oid: ".1.3",
          children: [
            {
              name: "dod",
              oid: ".1.3.6",
              children: [
                {
                  name: "thirdOidChild1",
                  oid: ".1.3.6.1",
                },
                {
                  name: "thirdOidChild2",
                  oid: ".1.3.6.2",
                },
                {
                  name: "thirdOidChild3",
                  oid: ".1.3.6.3",
                },
              ],
            },
          ],
        },
      ],
    };

    expect(tree).toEqual(expectedTree);
  });

  // it("realistic example", () => {
  //   let oids = Array<oidstorage.Oid>();

  //   let firstOid = new oidstorage.Oid('{"name": "iso","oid": ".1"}');
  //   let secondOid = new oidstorage.Oid('{"name": "org","oid": ".1.3"}');
  //   let thirdOid = new oidstorage.Oid('{"name": "dod","oid": ".1.3.6"}');

  //   oids.push(firstOid, thirdOid, secondOid);

  //   assert.equal(firstOid.name, "iso");
  //   assert.equal(firstOid.oid, ".1");

  //   let tree = sortTree(oids);

  //   const expectedTree: OidTree = {
  //     name: "iso",
  //     oid: ".1",
  //     children: [
  //       {
  //         name: "org",
  //         oid: ".1.3",
  //         children: [
  //           {
  //             name: "dod",
  //             oid: ".1.3.6",
  //           },
  //         ],
  //       },
  //     ],
  //   };

  //   expect(tree).toEqual(expectedTree);
  // });
});

describe("oid number counter test suite", () => {
  it("first test", () => {
    const oid = ".1.3.6.5";
    const expected = 4;

    expect(TreeSorter.getNumber(oid)).toEqual(expected);
  });

  it("no leading period", () => {
    const oid = "1.3.6.5";
    const expected = 4;

    expect(TreeSorter.getNumber(oid)).toEqual(expected);
  });

  it("multi character oid number with leading period", () => {
    const oid = ".1.3.6.5.12345";
    const expected = 5;

    expect(TreeSorter.getNumber(oid)).toEqual(expected);
  });

  it("multi character oid number", () => {
    const oid = "1.3.6.5.12345";
    const expected = 5;

    expect(TreeSorter.getNumber(oid)).toEqual(expected);
  });
});

describe("isDirectChild test suite", () => {
  it("simple direct child", () => {
    const expected = true;
    const parent = ".1.3";
    const child = ".1.3.6";

    expect(TreeSorter.isDirectChild(parent, child)).toEqual(expected);
  });

  it("multi character direct child", () => {
    const expected = true;
    const parent = ".1.3";
    const child = ".1.3.656";

    expect(TreeSorter.isDirectChild(parent, child)).toEqual(expected);
  });

  it("multi character direct child 2", () => {
    const expected = true;
    const parent = ".1.31";
    const child = ".1.31.656";

    expect(TreeSorter.isDirectChild(parent, child)).toEqual(expected);
  });

  it("not direct parent 1", () => {
    const expected = false;
    const parent = ".1.3.6.1";
    const child = ".1.3.5.1.5";

    expect(TreeSorter.isDirectChild(parent, child)).toEqual(expected);
  });

  it("not direct parent 2", () => {
    const expected = false;
    const parent = ".1.3.6.1";
    const child = ".1.3.5.1.5.5.1235";

    expect(TreeSorter.isDirectChild(parent, child)).toEqual(expected);
  });
});
