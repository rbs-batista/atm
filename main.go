package main

import (
	routes "atm/Routes"
	"fmt"
	"net/http"
	"strconv"
)

var port int = 8000

func main() {

	routes.RouteHandling()
	fmt.Printf("Servidor iniciado na porta :%d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
