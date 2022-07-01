import imp
import os
import sys
from PyQt5.QtWidgets import QTreeView
from pysmi.reader import getReadersFromUrls
from pysmi.reader.localfile import FileReader
from pysmi.codegen import JsonCodeGen
from pysmi.compiler import MibCompiler
from pysmi.parser import SmiV2Parser
from pysmi.writer import CallbackWriter, FileWriter
from pysmi import debug


def parse_mibs(mib: str, tree_view: QTreeView):
    head_tail = os.path.split(mib)
    code_gen = JsonCodeGen()
    debug.setLogger(debug.Debug('compiler'))

    fileWriter = CallbackWriter(lambda *x: None)

    fileWriter = FileWriter(head_tail[0]).setOptions(suffix='.json')

    parser = SmiV2Parser()

    mib_compiler = MibCompiler(
        parser,
        code_gen,
        fileWriter
    )

    x, _ = os.path.splitext(head_tail[1])

    mib_compiler.addSources(FileReader(head_tail[0]))

    processed = mib_compiler.compile(x, **dict(noDeps=True))

    mib_compiler.buildIndex(processed)

    pass
