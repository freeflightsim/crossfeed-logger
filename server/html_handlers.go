
package server

import (

	"fmt"
	"net/http"
)


// Home page - TODO
func H_Home(resp http.ResponseWriter, request *http.Request){

	fmt.Fprintf(resp, "Welcome to <b>cf-logger</b> ;-)")
}