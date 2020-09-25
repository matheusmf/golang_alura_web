package main

import (
	"fmt"
	"loja/routes"
	"net/http"
	"time"
)

func main() {
	routes.CarregaRotas()
	fmt.Println(time.Now().Format("02/01/2006 15:04:05") + " - " + "Iniciando loja...")
	http.ListenAndServe(":8000", nil)
}
