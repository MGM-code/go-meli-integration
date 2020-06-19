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
	var bucle int
	bucle = 1

	for bucle < 2 {

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
			4-> Division
			5-> Salir
			
			`)
		fmt.Println("Que Operación desea realizar:")
		var ope int
		fmt.Scanln(&ope)
		var num int
		var num1 int
		var num2 int
		var num3 int
		//var extra int
		
		//extra= 5

		//Procedimiento para la suma

		if ope == 1 {
			fmt.Println("Ingrese Su Primer Numero Para La Suma:")
			fmt.Scanln(&num)
			fmt.Println("Ingrese Su Segundo Numero Para La Suma:")
			fmt.Scanln(&num1)
			fmt.Println("Ingrese Su Primer Numero Para La Segunda Suma:")
			fmt.Scanln(&num2)
			fmt.Println("Ingrese Su Segundo Numero Para La Segunda Suma:")
			fmt.Scanln(&num3)

			resul := `
			***********************
			*resultado de la Suma *
			***********************	
			`
			resultado := Add(num, num1)
			go Add(num2, num3)
			fmt.Println(resul)
			fmt.Println("El resultado de la Suma es: ", resultado)
		}
		//Procedimiento para la Resta
		if ope == 2 {
			fmt.Println("Ingrese Su Primer Numero Para La Resta:")
			fmt.Scanln(&num)
			fmt.Println("Ingrese Su Segundo Numero Para La Resta:")
			fmt.Scanln(&num1)
			fmt.Println("Ingrese Su Primer Numero Para La Segunda Resta:")
			fmt.Scanln(&num2)
			fmt.Println("Ingrese Su Segundo Numero Para La Segunda Resta:")
			fmt.Scanln(&num3)
			resul := `
			************************
			*resultado de la Resta *
			************************
			`
			resultado := Res(num, num1)
			go Res(num2, num3)
			fmt.Println(resul)
			fmt.Println("El resultado de la resta es: ", resultado)

		}
		if ope == 3 {
			fmt.Println("Ingrese Su Primer Numero Para La Multiplicación:")
			fmt.Scanln(&num)
			fmt.Println("Ingrese Su Segundo Numero Para La Multiplicación:")
			fmt.Scanln(&num1)
			fmt.Println("Ingrese Su Primer Numero Para La Segunda Multiplicación:")
			fmt.Scanln(&num2)
			fmt.Println("Ingrese Su Segundo Numero Para La Segunda Multiplicación:")
			fmt.Scanln(&num3)
			resul := `
			********************************
			*resultado de la Multiplicación*
			********************************
			`
			resultado := Mul(num, num1)
			go Mul (num2, num3)
			fmt.Println(resul)
			fmt.Println("El resultado de la Multiplicacióna es: ", resultado)

		}
		if ope == 4 {
			fmt.Println("Ingrese Su Primer Numero Para La División:")
			fmt.Scanln(&num)
			fmt.Println("Ingrese Su Segundo Numero Para La División:")
			fmt.Scanln(&num1)
			fmt.Println("Ingrese Su Primer Numero Para La Segunda División:")
			fmt.Scanln(&num2)
			fmt.Println("Ingrese Su Segundo Numero Para La Segunda División:")
			fmt.Scanln(&num3)
			resul := `
			***************************
			*resultado de la División *
			***************************
			`
			resultado := Div(num, num1)
			go Div(num2, num3)
			fmt.Println(resul)
			fmt.Println("El resultado de la División es: ", resultado)

		}
		if ope == 5 {
			fmt.Println("Fin. Presione para salir")
			bucle = 2
		}
	}
	var espera string
	fmt.Scanln(&espera)
}
   
