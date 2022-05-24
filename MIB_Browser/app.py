import imp
import os
import sys
import typing
from PyQt5.QtWidgets import QApplication, QMainWindow, QWidget, QDialog, QMessageBox, QTreeWidgetItem, QTreeWidget, QFileSystemModel
from PyQt5.QtGui import QStandardItemModel, QStandardItem, QColor
from PyQt5.QtCore import pyqtSlot, pyqtSignal

from PyQt5.uic import loadUi

from main_window import Ui_MainWindow


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
        tree_model = QStandardItemModel()
        root_node = tree_model.invisibleRootItem()

        prj_a = StandardItem('Project A')
        file_a_py = StandardItem('file_a.py')
        prj_a.appendRow(file_a_py)

        root_node.appendRow(prj_a)
        self.treeView.setModel(tree_model)
        self.treeView.doubleClicked.connect(self.get_value)

    def get_value(self, val):
        print(val.data())
