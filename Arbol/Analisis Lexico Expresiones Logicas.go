package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	"regexp"
)

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

type Variables struct{
	nombre string
	valor int64
}


type variablesb struct{
	nombre string
	valor string
}

var cola[] string 
var colaA[] *Arbol
var Lista_var[] Variables 
var tabl[] variablesb 

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

func ordenarEntrada(cadena1 string) bool{
	//var char string
	numeros, _ := regexp.Compile("^[0-9]+$")
    //letras, _ := regexp.Compile("^[a-zA-Z]+$")
    simbolos, _ := regexp.Compile("^[+-/*:]+$")
	puntos, _ := regexp.Compile("^:=$")
	comparacion, _:=regexp.Compile("^[=<>]+$")
	logicos, _:=regexp.Compile("^[&!|]+$")
	boolF, _:=regexp.Compile("^[false]|[true]$")
	cadena:=strings.Fields(cadena1)
	for x :=0; x<len(cadena); x++{
		if(numeros.MatchString(cadena[x])){
			no := "numero"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
		if(simbolos.MatchString(cadena[x])){
			no := "simbolo"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
		if(puntos.MatchString(cadena[x])){
			no := "punto"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
		if(comparacion.MatchString(cadena[x])){
			no := "comparacion"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
		if(logicos.MatchString(cadena[x])){
			no := "operador logico"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
		if(boolF.MatchString(cadena[x])){
			no := "valor logico"
			va := string(cadena[x])
			s := variablesb{nombre:no,valor:va}
			tabl = append(tabl,s)
		}
	}
	for j := 0 ;j<(len(tabl)) ; j++ {
		fmt.Println(tabl[j])
	}
	
	
	return true;
}

func esLetra(carc string) bool{
	return carc[0]>64 && carc[0]<91
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

func asignar_variable(t *Arbol){
	val := obtenerRespuesta(t.Izquierda)
	nom := t.Derecha.Valor
	s := Variables{nombre:nom,valor:val}
	Lista_var = append(Lista_var,s)
} 

func llenarArbol() *Arbol{
	for len(cola)>0 {
		elem :=cola[0]
		cola = append(cola[:0], cola[1:]...)
		if(elem=="+" || elem=="-" || elem=="*" || elem=="/" || elem == ":"){
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

/*func main() {
  
  fmt.Println("Ingrese la cadena en preorden, los opradores y operandos debe ir separados por un espacio")
  ordenarEntrada("1 2 + 3 4 - * X :")
  arbol :=llenarArbol()
  RecorrerInorden(arbol)
  fmt.Println("")
  asignar_variable(arbol)
  fmt.Println(Lista_var)
}*/

func main() {
  fmt.Println("Ingrese la cadena en postfijo, los opradores y operandos debe ir separados por un espacio")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  ordenarEntrada(scanner.Text())
}