package main
//package calculator

	
import "fmt"

func Add(x int, y int) int {
    return x + y
}

func Res(x int, y int) int {
    return x - y
}


func Div(x int, y int) int {
    return x / y
}

func Mul(x int, y int) int {
    return x * y
}

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
	//var extra int
	
	//extra= 5

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
		resultado := Add(num, num1)
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
		resultado := Res(num, num1)
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
		resultado := Mul(num, num1)
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
		resultado := Div(num, num1)
		fmt.Println(resul)
		fmt.Println("El resultado de la División es: ", resultado)

	}
	var espera string
	fmt.Scanln(&espera)
}
   
