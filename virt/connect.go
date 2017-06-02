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

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Println("List domain error.")
	}

	fmt.Printf("%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf(" %s\n", name)
		}
		dom.Free()
	}

}
