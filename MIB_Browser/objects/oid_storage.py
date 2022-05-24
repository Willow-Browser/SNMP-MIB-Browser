from PyQt5.QtGui import QStandardItem, QColor, QIcon


class OID_Object(QStandardItem):
    def __init__(self, text) -> None:
        # super important to remember, but you can set the icon
        super.__init__()
        self.setEditable(False)
        self.setForeground(QColor(0, 0, 0))
        self.setText(text)
        self.name = text
        self.oid = ""
        self.mib = ""
        self.syntax = ""
        self.access = ""
        self.status = ""
        self.def_val = ""
        self.indexes = ""
        self.description = ""
