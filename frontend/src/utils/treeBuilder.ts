/* eslint-disable security/detect-object-injection */
import { oidstorage } from "../../wailsjs/go/models";

export interface OidTree {
  name: string;
  oid: string;
  children?: Array<OidTree>;
}

export class TreeSorter {
  #newOids: Array<oidstorage.Oid>;

  constructor(newOids: Array<oidstorage.Oid>) {
    this.#newOids = newOids;
  }

  createOidTree(): OidTree {
    const tree: OidTree = { name: "", oid: "" };
    let i: number;

    for (i = 0; i < this.#newOids.length; i++) {
      const oid = this.#newOids[i];
      if (oid.name === "iso") {
        tree.name = oid.name;
        tree.oid = oid.oid;
        break;
      }
    }

    this.#newOids.splice(i, 1);

    this.recursiveTreeBuild(tree);

    return tree;
  }

  recursiveTreeBuild(parent: OidTree) {
    const childOidArray = Array<oidstorage.Oid>();

    this.#newOids.forEach((oid) => {
      if (TreeSorter.isDirectChild(parent.oid, oid.oid)) {
        childOidArray.push(oid);
      }
    });

    childOidArray.sort((a, b) => a.oid.localeCompare(b.oid));

    childOidArray.forEach((oid) => {
      const newChildNode: OidTree = {
        name: oid.name,
        oid: oid.oid,
      };

      if (parent.children === undefined) {
        parent.children = Array<OidTree>();
      }

      parent.children?.push(newChildNode);

      this.recursiveTreeBuild(newChildNode);
    });
  }

  static isDirectChild(parentOid: string, childOid: string): boolean {
    let isChild = false;

    if (
      TreeSorter.getNumber(parentOid) + 1 ===
      TreeSorter.getNumber(childOid)
    ) {
      const lastPeriodIndex = childOid.lastIndexOf(".");
      const childComparer = childOid.substring(0, lastPeriodIndex);

      if (childComparer === parentOid) {
        isChild = true;
      }
    }

    return isChild;
  }

  static getNumber(oid: string): number {
    type Subtraction = 0 | 1;

    let leadingSubtraction: Subtraction = 1;
    const leadingPeriod = oid.at(0) === ".";
    if (!leadingPeriod) {
      leadingSubtraction = 0;
    }

    const splitOids = oid.split(".");
    return splitOids.length - leadingSubtraction;
  }
}
