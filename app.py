import sys
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog, QMessageBox, QTreeWidgetItem, QTreeWidget, QFileSystemModel
from PyQt5.QtGui import QStandardItemModel, QStandardItem, QColor

from PyQt5.uic import loadUi

from main_window import Ui_MainWindow


class StandardItem(QStandardItem):
    def __init__(self, text):
        super().__init__()
        self.setEditable(False)
        self.setForeground(QColor(0, 0, 0))
        self.setText(text)


class Window(QMainWindow, Ui_MainWindow):
    def __init__(self, parent=None) -> None:
        super().__init__(parent)
        self.setupUi(self)
        self.connect_signal_slots()

    def connect_signal_slots(self):
        data = {"Project A": ["file_a.py", "file_a.txt", "something.xls"],
                "Project B": ["file_b.csv", "photo.jpg"],
                "Project C": []}

        items = []
        for key, values in data.items():
            item = QTreeWidgetItem([key])
            for value in values:
                ext = value.split(".")[-1].upper()
                child = QTreeWidgetItem([value, ext])
                item.addChild(child)
            items.append(item)

        self.treeWidget.setHeaderLabels(["Name", "Type"])
        self.treeWidget.insertTopLevelItems(0, items)

        treeModel = QStandardItemModel()
        root_node = treeModel.invisibleRootItem()

        prj_a = StandardItem('Project A')

        file_a_py = StandardItem('file_a.py')
        prj_a.appendRow(file_a_py)

        root_node.appendRow(prj_a)
        self.treeView.setModel(treeModel)
        self.treeView.doubleClicked.connect(self.get_value)

    def get_value(self, val):
        x = 1
        print(val.data())


if __name__ == '__main__':
    app = QApplication(sys.argv)
    win = Window()
    win.show()
    sys.exit(app.exec())
