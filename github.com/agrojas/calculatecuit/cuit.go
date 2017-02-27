package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Cuit struct {
	genderLabel    string    `json:"genderLabel"`
	dni            string    `json:"dni"`
	genderCode     string    `json:"genderCode"`
	validationCode int       `json:"validationCode"`
	cuit           string    `json:"cuit"`
	time           time.Time `json:"datetime"`
}

func (c *Cuit) SetGenderCode(genderCode string) {
	c.genderCode = genderCode
}

func CompleteCuit(cuit *Cuit) {
	genderCode, ok := getGenderCode(cuit.genderLabel)
	if !ok {
		panic(ok)
	}
	cuit.SetGenderCode(genderCode)
	log.Println("CompleteCuit:", cuit)
}

func getGenderCode(gender string) (string, bool) {
	var GENDER = map[string]string{
		"M": "20",
		"F": "27",
		"C": "30",
	}
	var genderCode, ok = GENDER[gender]
	if !ok {
		log.Fatal("Invalid gender")
	}
	log.Println("Gender Code:", genderCode)
	return genderCode, ok
}

/*
CUIL / CUIT
CUIL/T: Son 11 números en total: :geek:
XY - 12345678 - Z
XY: Indican el tipo (Masculino, Femenino o una empresa)
12345678: Número de DNI
Z: Código Verificador
Algoritmo:
Se determina XY con las siguientes reglas
Masculino:20
Femenio:27
Empresa:30
Se múltiplica XY 12345678 por un número de forma separada:

X * 5
Y * 4
1 * 3
2 * 2
3 * 7
4 * 6
5 * 5
6 * 4
7 * 3
8 * 2
Se suman dichos resultados. El resultado obtenido se divide por 11. De esa división se obtiene un Resto que determina Z
Si el resto es 0= Entoces Z=0	Si el resto es 1= Entonces se aplica la siguiente regla:
*Si es hombre: Z=9 y XY pasa a ser 23 *Si es mujer: Z=4 y XY pasa a ser 23
Caso contrario XY pasa a ser (11- Resto).
*/
func GetValidationCode(genderCode, dni string) (int, bool) {
	args := []string{genderCode, dni}
	var cuit_concat = strings.Join(args, "")
	cuit_numbers_for_position := [10]int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
	fmt.Printf("cuit_concat", cuit_concat)
	fmt.Printf("cuit_numbers_for_position", cuit_numbers_for_position)
	for pos, char := range cuit_concat {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos, cuit_numbers_for_position[pos])
	}
	return 0, true
}
