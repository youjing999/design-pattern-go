package main

import "fmt"

//OutpatientSystem 门诊系统
type OutpatientSystem struct {
}

func (out *OutpatientSystem) RegistrationOperation() {
	fmt.Println("门诊系统...挂号")
}

// DoctorSystem 医生系统
type DoctorSystem struct {
}

func (doctor *DoctorSystem) treatmentOperation() {
	fmt.Println("医生系统...治疗")
}

// PharmacySystem 药房系统
type PharmacySystem struct {
}

func (p *PharmacySystem) dispensingOperation() {
	fmt.Println("药房系统...发药")
}

type Facade struct {
	outpatient     *OutpatientSystem
	doctorSystem   *DoctorSystem
	pharmacySystem *PharmacySystem
}

func NewFacade() *Facade {
	return &Facade{
		new(OutpatientSystem),
		new(DoctorSystem),
		new(PharmacySystem),
	}
}

func (facade *Facade) start() {
	facade.outpatient.RegistrationOperation()
	facade.doctorSystem.treatmentOperation()
	facade.pharmacySystem.dispensingOperation()
}
func main() {
	facade := NewFacade()
	facade.start()
}
