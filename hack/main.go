package main

import (
	"fmt"
	"log"
	"os"

	"github.com/daveoy/gopci/pkg/pci"
)

func main() {
	if path := os.Getenv("PCI_IDS_FILE"); path == "" {
		fmt.Println("PCI_IDS_FILE is NOT set, Label fields will be empty")
	} else {
		fmt.Println("PCI_IDS_FILE is set to:", path)
	}

	devices, err := pci.Scan()
	if err != nil {
		log.Fatalf("pci.Scan() error: %v", err)
	}

	for _, d := range devices {
		fmt.Printf("Address: %+b\n", d.Address)

		fmt.Printf("Vendor: %+v %+v\n", d.Vendor.ID, d.Vendor.Label)
		fmt.Printf("Product: %+v %+v\n", d.Product.ID, d.Product.Label)

		if d.Subvendor != nil {
			fmt.Printf("Subvendor: %+v %+v\n", d.Subvendor.ID, d.Subvendor.Label)
		} else {
			fmt.Printf("Subvendor: <nil>\n")
		}
		if d.Subdevice != nil {
			fmt.Printf("Subdevice: %+v %+v\n", d.Subdevice.ID, d.Subdevice.Label)
		} else {
			fmt.Printf("Subdevice: <nil>\n")
		}

		fmt.Println()
	}
}
