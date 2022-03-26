package steps

import "github.com/nelsongp/chainofresponsabiliti/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
