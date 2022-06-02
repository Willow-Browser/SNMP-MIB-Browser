from os import system
from PyQt5.QtGui import QStandardItemModel
from MIB_Browser.objects.oid_storage import OID_Object


class AddBaseOids(object):
    @classmethod
    def get_base_oids(self) -> QStandardItemModel:
        tree = QStandardItemModel()
        root_node = tree.invisibleRootItem()

        iso = OID_Object("iso", ".1", "SNMPv2-SMI")
        org = OID_Object("org", ".1.3", "SNMPv2-SMI")
        dod = OID_Object("dod", ".1.3.6", "SNMPv2-SMI")
        internet = OID_Object("internet", ".1.3.6.1", "SNMPv2-SMI")
        mgmt = OID_Object("mgmt", ".1.3.6.1.2", "SNMPv2-SMI")
        mib_2 = OID_Object("mib-2", ".1.3.6.1.2.1", "SNMPv2-SMI")
        system = OID_Object("system", ".1.3.6.1.2.1.1", "SNMPv2-SMI")
        sysDescr = OID_Object("sysDescr", ".1.3.6.1.2.1.1.1", "SNMPv2-SMI")
        sysObjectID = OID_Object("sysObjectID", ".1.3.6.1.2.1.1.2", "SNMPv2-SMI")

        system.appendRows([sysDescr, sysObjectID])
        mib_2.appendRow(system)
        mgmt.appendRow(mib_2)
        internet.appendRow(mgmt)
        dod.appendRow(internet)
        org.appendRow(dod)
        iso.appendRow(org)
        root_node.appendRow(iso)

        return tree
