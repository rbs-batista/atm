package routes

import (
	controllers "atm/Controllers"
	"net/http"
)

func RouteHandling() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/sacar", controllers.Withdraw)
}
