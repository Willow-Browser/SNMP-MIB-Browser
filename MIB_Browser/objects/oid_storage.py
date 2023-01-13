from PyQt5.QtGui import QStandardItem, QColor, QIcon
from enum import Enum
from typing import TypeVar


# OID_Object = TypeVar("OID_Object")


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
    def __init__(self, text="", oid="", mib="") -> None:
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
        self.oid_type = ""

    def set_name(self, text: str) -> "OID_Object":
        self.name = text
        self.setText(text)
        return self

    def set_oid(self, oid: str) -> "OID_Object":
        self.oid = oid
        return self

    def set_mib(self, mib: str) -> "OID_Object":
        self.mib = mib
        return self

    def set_description(self, description: str) -> "OID_Object":
        self.description = description
        return self

    def set_max_access(self, access: str) -> "OID_Object":
        self.access = access
        return self

    def set_status(self, status: str) -> "OID_Object":
        self.status = status
        return self

    def set_type(self, oid_type: str) -> "OID_Object":
        # TODO : set icon based on this
        self.oid_type = oid_type
        return self
