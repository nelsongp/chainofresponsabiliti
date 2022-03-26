package steps

import (
	"fmt"
	"github.com/nelsongp/chainofresponsabiliti/patient"
)

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *patient.Patient) {
	if p.RegistrationDone {
		fmt.Println("pacient registration is already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registern patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}
func (r *Reception) SetNext(next Department) {
	r.next = next
}
