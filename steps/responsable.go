package steps

// Definimos la interfaz que van a ocupar todos los involucrados en la cadena
type Responsable interface {
	//en este caso todas van a regresar un string pero podemos hacer que regresen lo que necesitemos
	Handle(process string)
}
