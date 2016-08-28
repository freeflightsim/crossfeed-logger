

package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

)

// Lets go ;-))
func Run() {


	// Setup www router
	router := mux.NewRouter()

	router.HandleFunc("/", H_Home)

	// TODO make this /ajax/* into a subrouter is gut feel
	router.HandleFunc("/ajax/info", AX_Info)

	router.HandleFunc("/ajax/csvlogs", AX_CSVListFiles)
	router.HandleFunc("/ajax/csvlogs/stage/{file_name}", AX_CSVStageFile)
	router.HandleFunc("/ajax/csvlogs/import", AX_CSVImport)

	router.HandleFunc("/ajax/admin/db/info", AX_DBInfo)
	router.HandleFunc("/ajax/admin/db/create", AX_DBCreateAll)

	router.HandleFunc("/ajax/callsigns", AX_Callsigns)
	router.HandleFunc("/ajax/callsign/{callsign}", AX_Callsign)





	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving on " + Config.HTTPAddress)
	http.Handle("/", router)
	err := http.ListenAndServe(Config.HTTPAddress , nil)
	if err != nil {
		fmt.Println("HTTP Error: ", err)
	}
}