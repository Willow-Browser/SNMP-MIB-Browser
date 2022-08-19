from fileinput import filename
import imp
import os
import sys
import typing
from pathlib import Path
from PyQt5.QtWidgets import QApplication, QMainWindow, QWidget, QDialog, QMessageBox, QTreeWidgetItem, QTreeWidget, QFileSystemModel, QFileDialog
from PyQt5.QtGui import QStandardItemModel, QStandardItem, QColor
from PyQt5.QtCore import pyqtSlot, pyqtSignal

from PyQt5.uic import loadUi
from MIB_Browser.scripts.add_base_oids import AddBaseOids
from MIB_Browser.scripts.parse_mibs import parse_mibs

from MIB_Browser.main_window import Ui_MainWindow


def run():
    app = Application(sys.argv)
    win = Window()
    win.show()
    sys.exit(app.exec())


class StandardItem(QStandardItem):
    def __init__(self, text):
        super().__init__()
        self.setEditable(False)
        self.setForeground(QColor(0, 0, 0))
        self.setText(text)


class Application(QApplication):
    def __init__(self, argv: typing.List[str]) -> None:
        super().__init__(argv)


class Window(QMainWindow, Ui_MainWindow):
    def __init__(self, parent=None) -> None:
        super().__init__(parent)
        self.setupUi(self)
        self.connect_signal_slots()

    def connect_signal_slots(self):
        self.treeView.setModel(AddBaseOids.get_base_oids())
        self.treeView.doubleClicked.connect(self.get_value)
        self.actionLoadMibs.triggered.connect(self.was_clicked)

    def get_value(self, val):
        print(val.data())

    def was_clicked(self):
        dlg = QFileDialog(self)
        dlg.setFileMode(QFileDialog.FileMode.ExistingFile)
        dlg.setNameFilter("All Files (*.*)")
        dlg.setViewMode(QFileDialog.ViewMode.Detail)
        # dlg.setDirectory(str(Path.home()))
        if dlg.exec_():
            for file in dlg.selectedFiles():
                parse_mibs(file, self.treeView)
        print("Hey it worked!")

    def transverse_tree(self):
        tree_list = []
