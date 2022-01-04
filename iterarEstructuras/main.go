package main

import (
	"fmt"
	"reflect"
)

type Fecha struct {
	dia, mes, anio int
}

type Persona struct {
	Nombre          string
	Apellido        string
	Altura          int
	FechaNacimiento Fecha
}

func main() {
	p := Persona{
		Nombre:   "Juan",
		Apellido: "Perez",
		Altura:   175,
	}

	v := reflect.ValueOf(p)
	tipoObtenidoDeReflection := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("El campo %s tiene como valor: %v\n", tipoObtenidoDeReflection.Field(i).Name, v.Field(i).Interface())
	}
}
