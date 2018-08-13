package main

import (
	"github.com/heshoots/fgcmvp/pkg"
	"net/http"
)

func main() {
	pkg.InitDB()
	router := pkg.Router()
	http.ListenAndServe(":3000", router)
}
