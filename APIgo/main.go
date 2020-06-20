package main
 
import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)
func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "form.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        num1 := r.FormValue("num1")
        num2 := r.FormValue("num2")
        accion := r.FormValue("accion")
        fmt.Fprintf(w, "Numero 1 = %s\n", num1)
        fmt.Fprintf(w, "Numero 2 = %s\n", num2)
        fmt.Fprintf(w, "Accion = %s\n", accion)
        
        n1, err := strconv.Atoi(num1)
        if err == nil {
            fmt.Println(num1)
        }

        n2,  err := strconv.Atoi(num2)
        if err == nil {
            fmt.Println(num2)
        }

        if accion == "1" {
            fmt.Fprintf(w, "Suma = %d\n", n1 + n2) 
        }

        if accion == "2" {
            fmt.Fprintf(w, "Resta = %d\n", n1 - n2) 
        }

        if accion == "3" {
            fmt.Fprintf(w, "Division = %d\n", n1 / n2) 
        }

        if accion == "4" {
            fmt.Fprintf(w, "Multiplicacion = %d\n", n1 * n2) 
        }
        
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
 
func main() {
    http.HandleFunc("/", hello)
 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
