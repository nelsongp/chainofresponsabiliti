package steps

import "fmt"

// Este sera otro de los eslabones de la cadena
type TeamLeader struct {
	//practicamente es quien es el siguiente en la cadena a realizar
	supervisor Responsable
}

// lo creamos para cuando creemos un colega nuevo le inyectamos quien el siguiente en la cadena
func NewTeamLeader(s Responsable) TeamLeader {
	return TeamLeader{supervisor: s}
}

// recibimos en el Handle si es un Merge request
func (cl TeamLeader) Handle(process string) {
	//evaluamos si el process es un merge request si es asi nos qeudamos aca
	if process == "MR" {
		fmt.Println("Procesando Merge.... OK!!")
		return
	}
	//si no es merge request vamos al siguiente manejador de la cadena
	cl.supervisor.Handle(process)
}
