package main

import (
	"log"

	"github.com/honne23/open-pay/pkg/directory"
)

var (
	cfo directory.Manager
	cto directory.Manager
	ceo directory.Manager
)

func init() {
	cfo = directory.NewManager(nil, nil)
	cto = directory.NewManager(nil, nil)
	ceoEmployees := []directory.Member{&cfo, &cto}
	ceo = directory.NewManager(nil, &ceoEmployees)
}

func main() {
	productLead1 := directory.NewManager(nil, nil)
	productLead2 := directory.NewManager(nil, nil)

	juniorDev1 := directory.NewEmployee()
	juniorDev2 := directory.NewEmployee()
	cto.AddEmployees([]directory.Member{&productLead1, &productLead2})

	productLead2.AddEmployees([]directory.Member{&juniorDev1, &juniorDev2})

	log.Printf("CEO: %s\n\n", ceo.GetID())
	log.Printf("CFO: %s\n\n", cfo.GetID())
	log.Printf("CTO: %s\n\n", cto.GetID())
	log.Printf("productLead1: %s\n\n", productLead1.GetID())
	log.Printf("productLead2: %s\n\n", productLead2.GetID())
	log.Printf("juniorDev1: %s\n\n\n\n", juniorDev1.GetID())
	log.Printf("juniorDe2: %s\n\n\n\n", juniorDev2.GetID())

	common := directory.FindCommonManager(ceo, &juniorDev1, &juniorDev2)
	log.Printf("Common manager: %s", common.GetID())

}

/**
* Adding a child to a node should turn them into a manager
 */
