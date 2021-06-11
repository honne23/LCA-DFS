package directory

import (
	"log"
	"testing"
)

/**
*


                               ┌──────────────┐
                               │              │
                               │              │
                               │     CEO      │
                               │              │
                               │              │
                               └──┬────────┬──┘
                                  │        │
                                  │        │
                                  │        │
          ┌────────────┐   ◄──────┘        └──────►  ┌──────────────┐
          │            │                             │              │
          │            │                             │              │
          │    CFO     │                             │     CTO      │
          │            │                             │              │
          │            │                             │              │
          └────────────┘                             └───┬─────┬────┘
                                                         │     │
                                                         │     │
                                                         │     │
                                                         │     │
                                  ┌──────────┐◄──────────┘     └───────────►  ┌───────────┐
                                  │          │                                │           │
                                  │          │                                │           │
                                  │   PM3    │                                │    PM2    │
                                  │          │                                │           │
                                  │          │                                │           │
                                  └──────────┘                                └──┬─────┬──┘
                                                                                 │     │
                                                                                 │     │
                                                                                 │     │
                                                            ┌──────────┐         │     │           ┌─────────┐
                                                            │          │         │     │           │         │
                                                            │          │ ◄───────┘     └────────►  │         │
                                                            │ Junior1  │                           │ Junior2 │
                                                            │          │                           │         │
                                                            └──────────┘                           └─────────┘
*/

func TestCommonManager(t *testing.T) {

	cfo := NewManager(nil, nil)
	cto := NewManager(nil, nil)
	ceoEmployees := []Member{&cfo, &cto}
	ceo := NewManager(nil, &ceoEmployees)
	productLead1 := NewManager(nil, nil)
	productLead2 := NewManager(nil, nil)

	juniorDev1 := NewEmployee()
	juniorDev2 := NewEmployee()
	cto.AddEmployees([]Member{&productLead1, &productLead2})

	productLead2.AddEmployees([]Member{&juniorDev1, &juniorDev2})
	log.Printf("CEO: %s\n\n", ceo.GetID())
	log.Printf("CFO: %s\n\n", cfo.GetID())
	log.Printf("CTO: %s\n\n", cto.GetID())
	log.Printf("productLead1: %s\n\n", productLead1.GetID())
	log.Printf("productLead2: %s\n\n", productLead2.GetID())
	log.Printf("juniorDev1: %s\n\n", juniorDev1.GetID())
	log.Printf("juniorDev2: %s\n\n", juniorDev2.GetID())
	lcm := FindCommonManager(ceo, &juniorDev1, &juniorDev2)
	if lcm.GetID() != productLead2.GetID() {
		t.Errorf("Failed Common Manager: Found %s, expected %s", lcm.GetID(), productLead2.GetID())
	}

}

/**
*
                                                           ┌──────────────┐
                                                           │              │
                                                           │              │
                                                           │     CEO      │
                                                           │              │
                                                           │              │
                                                           └──┬────────┬──┘
                                                              │        │
                                                              │        │
                                                              │        │
                                      ┌────────────┐   ◄──────┘        └──────►  ┌───────────┐
                                      │            │                             │           │
                                      │            │                             │           │
                                      │    CFO     │                             │   CTO     │
                                      │            │                             │           │
                                      │            │                             │           │
                                      └──┬────┬────┘                             └───┬────┬──┘
                                         │    │                                      │    │
                                         │    │                                      │    │
                                         │    │                                      │    │
                 ┌────────────┐          │    │     ┌──────────┐                     │    │
                 │            │ ◄────────┘    └────►│          │                     │    │
                 │            │                     │          │     ┌─────────┐◄────┘    └────► ┌────────────┐
                 │   AC1      │                     │   AC2    │     │         │                 │            │
                 │            │                     │          │     │         │                 │            │
                 │            │                     │          │     │  PM1    │                 │    PM2     │
                 └─────┬──────┘                     └──────────┘     │         │                 │            │
                       │                                             │         │                 │            │
                       │                                             └─────────┘                 └──────┬─────┘
                       │                                                                                │
                       │                                                                                │
                       │                                                                                │
┌─────────┐   ◄────────┘                                                                                │
│         │                                                                                             │
│         │                                                                                             │               ┌──────────────┐
│   JAC1  │                                                                                             └────────────►  │              │
│         │                                                                                                             │              │
│         │                                                                                                             │      JD2     │
└─────────┘                                                                                                             │              │
                                                                                                                        │              │
                                                                                                                        └──────────────┘
*/

func TestUnbalancedTree(t *testing.T) {

	cfo := NewManager(nil, nil)
	cto := NewManager(nil, nil)
	ceoEmployees := []Member{&cfo, &cto}
	ceo := NewManager(nil, &ceoEmployees)

	productLead1 := NewManager(nil, nil)
	productLead2 := NewManager(nil, nil)

	juniorDev1 := NewEmployee()
	juniorDev2 := NewEmployee()
	juniorDev3 := NewEmployee()

	productLead2.AddEmployees([]Member{&juniorDev1, &juniorDev2})

	productLead1.AddEmployees([]Member{&juniorDev3})

	cto.AddEmployees([]Member{&productLead1, &productLead2})

	ac1 := NewManager(nil, nil)
	ac2 := NewManager(nil, nil)

	cfo.AddEmployees([]Member{&ac1, &ac2})

	log.Printf("CEO: %s\n\n", ceo.GetID())
	log.Printf("CFO: %s\n\n", cfo.GetID())
	log.Printf("CTO: %s\n\n", cto.GetID())
	log.Printf("ac1: %s\n\n", ac1.GetID())
	log.Printf("ac2: %s\n\n", ac2.GetID())
	log.Printf("productLead1: %s\n\n", productLead1.GetID())
	log.Printf("productLead2: %s\n\n", productLead2.GetID())
	log.Printf("juniorDev1: %s\n\n", juniorDev1.GetID())
	log.Printf("juniorDev2: %s\n\n", juniorDev2.GetID())

	lcm := FindCommonManager(ceo, &ac2, &juniorDev2)
	if lcm.GetID() != ceo.GetID() {
		t.Errorf("Failed Common Manager: Found %s, expected %s", lcm.GetID(), ceo.GetID())
	}

}

/**
*
		┌──────────────┐
		│              │
		│              │
		│     CEO      │
		│              │
		│              │
		└──┬────────┬──┘
*/

func TestRootOnly(t *testing.T) {
	cfo := NewManager(nil, nil)
	cto := NewManager(nil, nil)
	ceoEmployees := []Member{&cfo, &cto}
	ceo := NewManager(nil, &ceoEmployees)
	productLead1 := NewManager(nil, nil)
	productLead2 := NewManager(nil, nil)
	cto.AddEmployees([]Member{&productLead1, &productLead2})
	lcm := FindCommonManager(ceo, &ceo, &productLead1)
	if lcm.GetID() != ceo.GetID() {
		t.Errorf("Failed Common Manager: Found %s, expected %s", lcm.GetID(), ceo.GetID())
	}
}