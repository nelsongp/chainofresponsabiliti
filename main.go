package main

import (
	"fmt"
	"github.com/nelsongp/chainofresponsabiliti/steps"
	"log"
)

func main() {
	//creamos las implementaciones que vamos a ocupar osea nuestros eslabones
	//tambien esta capa incializadora la podemos manejar desde esl bootsstrap como en la arqui hexagonal
	//definimos el final que seria el sre
	sre := steps.NewSre()
	//defiminimos el teamleader el cual como siguiente paso es el sre
	teamLeader := steps.NewTeamLeader(sre)
	//definimos el collega que es el primer paso y como siguiente implementa al teamleader
	college := steps.NewCollege(teamLeader)

	//creamos nuestra implementacion de la cadena
	fmt.Println("Que quieres solicitar? CR(Code Review), MR (Merge Request), DP(Deploy)")
	option := ""
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatalf("cannot read your option")
	}

	//el colega como punto de partida maneja las opciones
	college.Handle(option)
}
