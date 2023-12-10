package steps

import "fmt"

// Este sera el primer eslabon de la cadena el cual el primero en entrar en interaccion
type College struct {
	//practicamente es quien es el siguiente en la cadena a realizar
	supervisor Responsable
}

// lo creamos para cuando creemos un colega nuevo le inyectamos quien el siguiente en la cadena
func NewCollege(s Responsable) College {
	return College{supervisor: s}
}

// ek metodo Handle recibe el proceso a revisar en este caso si es un code review
func (c College) Handle(process string) {
	//evaluamos si el process es un code review si es asi nos qeudamos aca
	if process == "CR" {
		fmt.Println("Revisando.... que bien!!")
		return
	}
	//si no es code review vamos al siguiente manejador de la cadena
	c.supervisor.Handle(process)
}
