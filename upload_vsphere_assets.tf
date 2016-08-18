provider "vsphere" {
 user = "satiar"
 password = "@SujataPant!"
 vsphere_server = "10.247.81.78"
 allow_unverified_ssl = true
}
resource "vsphere_file" "arpita-ubuntu-vmdk" {
 datacenter = "FTC"
 datastore = "san1"
 source_file = "/home/aocole/ubuntu1404.vmdk"
 destination_file = "satiartest/ubuntu1404.vmdk"
}
resource "vsphere_file" "arpita-cidata" {
 datacenter = "FTC"
 datastore = "san1"
 source_file = "/home/satiar/cidata.iso"
 destination_file = "satiartest/cidata.iso"
}
