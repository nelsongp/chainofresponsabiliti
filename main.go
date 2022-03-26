package main

import (
	"github.com/nelsongp/chainofresponsabiliti/patient"
	"github.com/nelsongp/chainofresponsabiliti/steps"
)

func main() {
	cashier := &steps.Cashier{}

	//set next for medican department
	medical := &steps.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &steps.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &steps.Reception{}
	reception.SetNext(doctor)

	patient := &patient.Patient{Name: "abc", MedicineDone: true}
	//Patient visiting
	reception.Execute(patient)
}
