package main

import (
	"fmt"
	libvirt "github.com/libvirt/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Conn error.")
	}
	defer conn.Close()

	nets, err := conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE)
	if err != nil {
		fmt.Println("List network error.")
	}

	fmt.Printf("%d list network devices:\n", len(nets))
	for _, net := range nets {
		name, err := net.GetName()
		if err == nil {
			fmt.Printf(" %s\n", name)
		}
		net.Free()
	}

}
