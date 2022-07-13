from PyQt5.QtGui import QStandardItem, QColor, QIcon
from enum import Enum


class OID_Type(Enum):
    MODULE_IDENTITY = 1
    TEXTUAL_CONVENTION = 2
    OBJECT_IDENTITY = 3
    OBJECT_TYPE = 4
    NOTIFICATION_TYPE = 5
    MODULE_COMPLIANCE = 6
    OBJECT_GROUP = 7
    NOTIFICATION_GROUP = 8


class OID_Object(QStandardItem):
    def __init__(self, text, oid, mib) -> None:
        # super important to remember, but you can set the icon
        super().__init__()
        self.setEditable(False)
        self.setForeground(QColor(0, 0, 0))
        self.setText(text)
        self.name = text
        self.oid = oid
        self.mib = mib
        self.syntax = ""
        self.access = ""
        self.status = ""
        self.def_val = ""
        self.indexes = ""
        self.description = ""
