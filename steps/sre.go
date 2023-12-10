package steps

import "fmt"

// Como es el final de la cadena practicamente ya no le inyectaremos a nadie mas
type Sre struct{}

// como es final de cadena tambien su inicializador es vacio
func NewSre() Sre {
	return Sre{}
}

// recibimos en el Handle si es un Deploy y aca terminamos igual
func (sre Sre) Handle(process string) {
	//evaluamos si el process es un merge request si es asi nos qeudamos aca
	if process == "DP" {
		fmt.Println("Deployando todo.... OK!!")
		return
	}
	fmt.Println("No existe esa opcion")
}
