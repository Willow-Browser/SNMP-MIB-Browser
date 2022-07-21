import imp
import os
import sys
from PyQt5.QtWidgets import QTreeView
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


def parse_mibs(mib: str, tree_view: QTreeView):
    head_tail = os.path.split(mib)
    code_gen = JsonCodeGen()
    # debug.setLogger(debug.Debug('compiler'))

    mib_stubs = code_gen.baseMibs

    # dstDirectory = os.path.join('.')

    searchers = [StubSearcher(*mib_stubs)]

    fileWriter = TestFileWrite(head_tail[0]).setOptions(suffix='.json')

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

        # TODO : actually write the mib parsing logic

        _ = mib_compiler.compile(
            mib_name,
            **dict(noDeps=False,
                   genTexts=True,
                   textFilter=True and (lambda symbol, text: text) or None))

        xyz = 1
    except error.PySmiError:
        sys.stderr.write('ERROR: %s\r\n' % sys.exc_info()[1])
        sys.exit(70)

    for some_class in classes:
        print(some_class)

    pass


class TestFileWrite(AbstractWriter):
    def __init__(self, path) -> None:
        self._path = decode(os.path.normpath(path))
        self.oids = []

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

        for key, y in newData.items():
            if key != 'imports':
                z = 1
                try:
                    if y['class'] not in classes:
                        classes.append(y['class'])
                    if y['class'] == 'moduleidentity':
                        oid = OID_Object(y['name'], y['oid'], mibname)
                        oid.description = y['description']
                        self.oids.append(oid)
                    elif y['class'] == 'objectidentity':
                        oid = OID_Object(y['name'], y['oid'], mibname)
                        if 'description' in y:
                            oid.description = y['description']
                        self.oids.append(oid)
                    elif y['class'] == 'objecttype':
                        oid = OID_Object(y['name'], y['oid'], mibname)
                        oid.description = y['description']
                        oid.access = y['maxaccess']
                        oid.status = y['status']
                        self.oids.append(oid)
                except KeyError:
                    xy = 5
        return
