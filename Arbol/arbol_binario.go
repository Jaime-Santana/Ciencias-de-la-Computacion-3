package main

import (
	"fmt"
	"strconv"
	"strings"
	
)

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

var cola[] string 
var colaA[] *Arbol

var arbol *Arbol

func NewArbol( v string ) *Arbol {
	t := new( Arbol)
	t.Derecha=nil
	t.Izquierda=nil
	t.Valor=v
	return t
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInorden(t.Izquierda)

	fmt.Print(t.Valor)
    fmt.Print(" ")
	RecorrerInorden(t.Derecha)
}

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}
  fmt.Print(t.Valor)
  fmt.Print(" - ")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
}

func RecorrerPosorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerPosorden(t.Izquierda)
	RecorrerPosorden(t.Derecha)
  fmt.Print(t.Valor)
  fmt.Print(" - ")
}

func ordenarEntrada(cadena1 string){
	var char string
	cadena:=strings.Fields(cadena1)
	for x :=0; x<len(cadena); x++{
		char=cadena[x]
		cola=append(cola, char)
	}

	fmt.Println(cola)
}

func obtenerRespuesta(t *Arbol) int64{
	if(t.Valor=="+" ){
		return obtenerRespuesta(t.Izquierda) + obtenerRespuesta(t.Derecha)
	}else if(t.Valor=="-" ){
		return obtenerRespuesta(t.Izquierda) - obtenerRespuesta(t.Derecha)
	}else if(t.Valor=="*" ){
		return obtenerRespuesta(t.Izquierda) * obtenerRespuesta(t.Derecha)
	}else if(t.Valor=="/"){
		return obtenerRespuesta(t.Izquierda) / obtenerRespuesta(t.Derecha)
	}else{
		num, _:=strconv.ParseInt(t.Valor, 10, 0)
		return num
	}
}

func llenarArbol() *Arbol{
	for len(cola)>0 {
		elem :=cola[0]
		cola = append(cola[:0], cola[1:]...)
		if(elem=="+" || elem=="-" || elem=="*" || elem=="/"){
			arbol = NewArbol(elem)
			leng :=len(colaA)
			arbol.Izquierda=(colaA[leng-2])
			arbol.Derecha=(colaA[leng-1])
			colaA = colaA[:len(colaA)-2]
			colaA=append(colaA, arbol)
		}else{
			arbol := NewArbol(elem)
			colaA=append(colaA, arbol)
		}
	}
	return arbol
}



func main() {
	
  fmt.Println("Ingrese la cadena en postfijo, los opradores y operandos debe ir separados por un espacio")

  ordenarEntrada("1 2 + 3 4 - *")
  arbol :=llenarArbol()
  RecorrerInorden(arbol)
  fmt.Println("")
  fmt.Println(obtenerRespuesta(arbol))

}
