package steps

import (
	"fmt"
	"github.com/nelsongp/chainofresponsabiliti/patient"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("payment Done")
	}
	fmt.Println("Cashier getting money from patient")
	p.PaymentDone = true
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
