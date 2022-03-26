package steps

import (
	"fmt"
	"github.com/nelsongp/chainofresponsabiliti/patient"
)

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *patient.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to the patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medicinal giviing to the pacient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
