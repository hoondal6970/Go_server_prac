package main

import (
	"net/http"

	"github.com/hoondal6970/Go_server_prac/web3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
