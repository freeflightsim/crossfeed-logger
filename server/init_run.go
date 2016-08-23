

package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

)

// Lets go ;-))
func Run(conf ConfigOpts) {

	Config = conf

	// Setup www router
	router := mux.NewRouter()

	router.HandleFunc("/", H_Home)

	// TODO make this /ajax/* into a subrouter is gut feel
	router.HandleFunc("/ajax/info", AX_Info)

	router.HandleFunc("/ajax/csvlogs", AX_CSVListFiles)
	router.HandleFunc("/ajax/csvlogs/import/{file_name}", AX_CSVImportFile)

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving on " + conf.HTTPAddress)
	http.Handle("/", router)
	err := http.ListenAndServe(conf.HTTPAddress , nil)
	if err != nil {
		fmt.Println("HTTP Error: ", err)
	}
}