import os
import sys
from PyQt5.QtWidgets import QTreeView
from PyQt5.QtCore import QAbstractItemModel
from PyQt5.QtGui import QStandardItemModel
from pysmi.reader import getReadersFromUrls
from pysmi.searcher import AnyFileSearcher, PyFileSearcher, PyPackageSearcher, StubSearcher
from pysmi.borrower import AnyFileBorrower, PyFileBorrower
from pysmi.reader.localfile import FileReader
from pysmi.codegen import JsonCodeGen
from pysmi.compiler import MibCompiler
from pysmi.parser import SmiV2Parser
from pysmi.writer import CallbackWriter, FileWriter
from pysmi.writer.base import AbstractWriter
from pysmi.compat import encode, decode
from pysmi import debug
from pysmi import error
import json

from MIB_Browser.objects.oid_storage import OID_Object

data = {}
classes = []


def _iter_items(tree_view: QTreeView):
    def recurse(parent: QAbstractItemModel):
        for row in range(parent.rowCount()):
            for column in range(parent.columnCount()):
                child = parent.child(row, column)
                yield child
                if child.hasChildren():
                    yield from recurse(child)
    root = tree_view.model().invisibleRootItem()
    if root is not None:
        yield from recurse(root)


def iterItems2(tree_view: QTreeView):
    def recurse(parent: QAbstractItemModel):
        if root is not None:
            for row in range(parent.rowCount()):
                for column in range(parent.columnCount()):
                    child = parent.child(row, column)
                    yield child
                    if child.hasChildren():
                        for item in recurse(child):
                            yield item

    root = tree_view.model().invisibleRootItem()
    return recurse(root)


def iterItems3(tree_view: QTreeView):
    def recurse(root: QStandardItemModel):
        if root is not None:
            stack = [root]
            while stack:
                parent = stack.pop(0)
                for row in range(parent.rowCount()):
                    for column in range(parent.columnCount()):
                        child = parent.child(row, column)
                        yield child
                        if child.hasChildren():
                            stack.append(child)

    root = tree_view.model().invisibleRootItem()
    return recurse(root)


def parse_mibs(mib: str, tree_view: QTreeView):
    head_tail = os.path.split(mib)
    code_gen = JsonCodeGen()
    # debug.setLogger(debug.Debug('compiler'))

    mib_stubs = code_gen.baseMibs

    # dstDirectory = os.path.join('.')

    searchers = [StubSearcher(*mib_stubs)]

    fileWriter = FakeFileWrite(head_tail[0]).setOptions(suffix='.json')

    parser = SmiV2Parser()

    mib_compiler = MibCompiler(
        parser,
        code_gen,
        fileWriter
    )

    mib_name, _ = os.path.splitext(head_tail[1])

    try:
        mib_compiler.addSources(
            FileReader(head_tail[0]).setOptions(**dict(fuzzyMatching=True)))

        mib_compiler.addSources(FileReader(
            os.path.join(os.path.expanduser("~"), ".pysnmp", "mibs")))

        mib_compiler.addSearchers(*searchers)

        _ = mib_compiler.compile(
            mib_name,
            **dict(noDeps=False,
                   genTexts=True,
                   textFilter=True and (lambda _, text: text) or None))
    except error.PySmiError:
        sys.stderr.write('ERROR: %s\r\n' % sys.exc_info()[1])
        sys.exit(70)

    items = _iter_items(tree_view)
    add_oids = []  # try this again with private added
    new_items = []

    for item in items:
        new_items.append(item)

    def is_direct_child(parent_oid: str, child_oid: str) -> bool:
        parent = parent_oid.split(".")
        child = child_oid.split(".")

        if len(parent) < len(child) and (len(parent) + 1) == len(child):
            for i in range(len(parent)):
                if parent[i] != child[i]:
                    return False

            return True
        return False

    def is_duplicate(oid: str) -> bool:
        for item in new_items:
            if item.oid == oid:
                return True
        return False

    for oid in fileWriter.oids:
        if not is_duplicate(oid.oid):
            for item in new_items:
                if is_direct_child(item.oid, oid.oid):
                    # surprised this works
                    item.appendRow(oid)
                    # add_oids.append(oid)

    for some_class in classes:
        print(some_class)

    for new_oid in add_oids:
        root = tree_view.model().invisibleRootItem()
        if root is not None:
            # for (row)
            pass
        pass

    # tree_view.setModel(items)

    pass


class FakeFileWrite(AbstractWriter):
    def __init__(self, path) -> None:
        self._path = decode(os.path.normpath(path))
        self.oids = []
        self.mibname = ""

    def __str__(self) -> str:
        return '%s{"%s"}' % (self.__class__.__name__, self._path)

    def getData(self, mibname):
        filename = os.path.join(self._path, decode(mibname))

        f = None

        try:
            f = open(filename)
            data = f.read()
            f.close()
            return data
        except (OSError, IOError, UnicodeEncodeError):
            if f:
                f.close()
            return ''

    def putData(self, mibname, data, comments=(), dryRun=False):
        newData = json.loads(data)
        self.mibname = mibname

        for key, y in newData.items():
            if key != 'imports':
                try:
                    if y['class'] not in classes:
                        classes.append(y['class'])
                    if y['class'] == 'moduleidentity':
                        self._add_module_identity(y)
                    elif y['class'] == 'objectidentity':
                        self._add_object_identity(y)
                    elif y['class'] == 'objecttype':
                        self._add_oid_object(y)
                except KeyError:
                    xy = 5
        return

    def _add_module_identity(self, oid: dict) -> None:
        new_oid = (OID_Object(oid['name'], oid['oid'], self.mibname)
                   .set_description(oid['description']))
        self.oids.append(new_oid)

    def _add_object_identity(self, oid: dict) -> None:
        new_oid = OID_Object(oid['name'], oid['oid'], self.mibname)
        if 'description' in oid:
            new_oid.description = oid['description']
        self.oids.append(new_oid)

    def _add_oid_object(self, oid: dict) -> None:
        new_oid = (OID_Object(oid['name'], oid['oid'], self.mibname)
                   .set_description(oid['description'])
                   .set_max_access(oid['maxaccess'])
                   .set_status(oid['status'])
                   .set_type(oid['nodetype']))
        # TODO : add syntax as well
        self.oids.append(new_oid)
