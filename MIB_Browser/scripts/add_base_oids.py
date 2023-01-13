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
        directory = OID_Object("directory", ".1.3.6.1.1", "SNMPv2-SMI")
        mgmt = OID_Object("mgmt", ".1.3.6.1.2", "SNMPv2-SMI")
        mib_2 = OID_Object("mib-2", ".1.3.6.1.2.1", "SNMPv2-SMI")
        system = OID_Object("system", ".1.3.6.1.2.1.1", "SNMPv2-MIB")
        sysDescr = OID_Object("sysDescr", ".1.3.6.1.2.1.1.1", "SNMPv2-MIB")
        sysObjectID = OID_Object("sysObjectID", ".1.3.6.1.2.1.1.2", "SNMPv2-MIB")
        sysUpTime = OID_Object("sysUpTime", ".1.3.6.1.2.1.1.3", "SNMPv2-MIB")
        sysContact = OID_Object("sysContact", ".1.3.6.1.2.1.1.4", "SNMPv2-MIB")
        sysName = OID_Object("sysName", ".1.3.6.1.2.1.1.5", "SNMPv2-MIB")
        sysLocation = OID_Object("sysLocation", ".1.3.6.1.2.1.1.6", "SNMPv2-MIB")
        sysServices = OID_Object("sysServices", ".1.3.6.1.2.1.1.7", "SNMPv2-MIB")
        experimental = OID_Object("experimental", ".1.3.6.1.3", "SNMPv2-SMI")
        private = OID_Object("private", ".1.3.6.1.4", "SNMPv2-SMI")
        enterprises = OID_Object("enterprises", ".1.3.6.1.4.1", "SNMPv2-SMI")
        security = OID_Object("security", ".1.3.6.1.5", "SNMPv2-SMI")
        snmpV2 = OID_Object("snmpV2", ".1.3.6.1.6", "SNMPv2-SMI")
        snmpDomains = OID_Object("snmpDomains", ".1.3.6.1.6.1", "SNMPv2-SMI")
        snmpProxys = OID_Object("snmpProxys", ".1.3.6.1.6.2", "SNMPv2-SMI")
        snmpModules = OID_Object("snmpModules", ".1.3.6.1.6.3", "SNMPv2-SMI")

        system.appendRows([sysDescr, sysObjectID, sysUpTime, sysContact, sysName, sysLocation, sysServices])
        mib_2.appendRow(system)
        mgmt.appendRow(mib_2)
        internet.appendRows([directory, mgmt, experimental, private, security, snmpV2])
        private.appendRow(enterprises)
        snmpV2.appendRows([snmpDomains, snmpProxys, snmpModules])
        dod.appendRow(internet)
        org.appendRow(dod)
        iso.appendRow(org)
        root_node.appendRow(iso)

        return tree
