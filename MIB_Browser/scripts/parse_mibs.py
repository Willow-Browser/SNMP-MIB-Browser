import imp
import os
import sys
from this import d
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

data = {}


def parse_mibs(mib: str, tree_view: QTreeView):
    head_tail = os.path.split(mib)
    code_gen = JsonCodeGen()
    debug.setLogger(debug.Debug('compiler'))

    mib_stubs = JsonCodeGen.baseMibs

    dstDirectory = os.path.join('.')

    mibBorrowers = [('http://mibs.snmplabs.com/json/notexts/@mib@', False),
                    ('http://mibs.snmplabs.com/json/fulltexts/@mib@', True)]

    borrowers = [AnyFileBorrower(x[1], genTexts=mibBorrowers[x[0]][1]).setOptions(exts=['.json'])
                 for x in enumerate(getReadersFromUrls(*[m[0] for m in mibBorrowers], **dict(lowcaseMatching=False)))]

    searchers = [AnyFileSearcher(dstDirectory).setOptions(exts=['.json']), StubSearcher(*mib_stubs)]

    fileWriter = CallbackWriter(func)

    fileWriter = TestFileWrite(head_tail[0]).setOptions(suffix='.json')

    parser = SmiV2Parser()

    mib_compiler = MibCompiler(
        parser,
        code_gen,
        fileWriter
    )

    mib_name, _ = os.path.splitext(head_tail[1])

    try:
        mib_compiler.addSources(FileReader(head_tail[0]))

        mib_compiler.addSearchers(*searchers)

        mib_compiler.addBorrowers(*borrowers)

        _ = mib_compiler.compile(
            mib_name,
            **dict(noDeps=True,
                   genTexts=True,
                   textFilter=True and (lambda symbol, text: text) or None))
    except error.PySmiError:
        sys.stderr.write('ERROR: %s\r\n' % sys.exc_info()[1])
        sys.exit(70)

    pass


def func(mibname, contents, cbCtx):
    x = 1
    pass


class TestFileWrite(AbstractWriter):
    def __init__(self, path) -> None:
        self._path = decode(os.path.normpath(path))

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

        # newData is a dictionary that contains what we should need

        x = 1
        return
