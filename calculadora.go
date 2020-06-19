package main

import "fmt"

func main() {
	enca := `
	
	*******************	
	    CALCULADORA
	*******************
	`

	fmt.Println(enca)
	fmt.Println(`
		1-> Suma		
		2-> Resta
		3-> Multiplicación
		4-> Divesion
		
		`)
	fmt.Println("Que Operación desea realizar:")
	var ope int
	fmt.Scanln(&ope)
	var num int
	var num1 int
	var extra int
	
	extra= 5

	//Procedimiento para la suma

	if ope == 1 {
		fmt.Println("Ingrese Su Primer Numero Para La Suma:")
		fmt.Scanln(&num)
		fmt.Println("Ingrese Su Segundo Numero Para La Suma:")
		fmt.Scanln(&num1)
		resul := `
		***********************
		*resultado de la Suma *
		***********************	
		`
		resultado := num + num1 + extra
		fmt.Println(resul)
		fmt.Println("El resultado de la Suma es: ", resultado)

	}
	//Procedimiento para la Resta
	if ope == 2 {
		fmt.Println("Ingrese Su Primer Numero Para La Resta:")
		fmt.Scanln(&num)
		fmt.Println("Ingrese Su Segundo Numero Para La Resta:")
		fmt.Scanln(&num1)
		resul := `
		************************
		*resultado de la Resta *
		************************
		`
		resultado := num - num1 + extra
		fmt.Println(resul)
		fmt.Println("El resultado de la resta es: ", resultado)

	}
	if ope == 3 {
		fmt.Println("Ingrese Su Primer Numero Para La Multiplicación:")
		fmt.Scanln(&num)
		fmt.Println("Ingrese Su Segundo Numero Para La Multiplicación:")
		fmt.Scanln(&num1)
		resul := `
		********************************
		*resultado de la Multiplicación*
		********************************
		`
		resultado := num * num1 + extra
		fmt.Println(resul)
		fmt.Println("El resultado de la Multiplicacióna es: ", resultado)

	}
	if ope == 4 {
		fmt.Println("Ingrese Su Primer Numero Para La División:")
		fmt.Scanln(&num)
		fmt.Println("Ingrese Su Segundo Numero Para La División:")
		fmt.Scanln(&num1)
		resul := `
		***************************
		*resultado de la División *
		***************************
		`
		resultado := num / num1 + extra
		fmt.Println(resul)
		fmt.Println("El resultado de la División es: ", resultado)

	}
	var espera string
	fmt.Scanln(&espera)
}
