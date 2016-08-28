package server

import (

	"net/http"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/gorilla/mux"

	"github.com/freeflightsim/crossfeed-logger/cfdb"
	"github.com/freeflightsim/crossfeed-logger/cvslog"

)

// SendAjaxPayload is the function that sends the http reply
// machine encoded `payload` formatted,  ie a serialiser
// html.pages should not hit here, but otherwise expected is
// a reply with the "bites" in the particular machine readable format
// and correct mime type etc eg json, yaml and xml.hell in m$.excel
func SendAjaxPayload(resp http.ResponseWriter, request *http.Request, payload interface{}) {

	// pretty returns indents data and readable (notably json) is ?pretty=1 in url
	pretty := true // ON for NOW debug time request.URL.Query().Get("pretty") == "0"

	// Determine which encoding from the mux/router
	vars := mux.Vars(request)
	enc := vars["ext"]
	if enc == "" {
		enc = "json"
	}
	// TODO validate encoding and serialiser
	// eg yaml, json/js, csv etc (eg error for html maybe)

	// TODO map[string] = encoding

	// Lets get ready to encode folks... by default text/plain
	var bites []byte
	var err error
	var content_type string = "text/plain"

	if enc == "yaml" {
		// encode  in yaml - pretty igonered as yaml is pretty ;-)
		bites, err = yaml.Marshal(payload)
		content_type = "text/yaml"

	} else if enc == "json" || enc == "js" {
		// encode in json
		if pretty {
			bites, err = json.MarshalIndent(payload, "", "  ")
		} else {
			bites, err = json.Marshal(payload)
		}
		content_type = "application/json"

	} else {
		// serialise an error
		bites = []byte("OOPs no `.ext` recognised ")
	}

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", content_type)
	resp.Write(bites)
}

type ErrorPayload struct {
	Success bool   ` json:"success" `
	Error   string ` json:"error" `
}

func SendAjaxError(resp http.ResponseWriter, request *http.Request, err error) {
	SendAjaxPayload(resp, request, ErrorPayload{Success: true, Error: err.Error()})
}

type InfoPayload struct {
	Success bool   ` json:"success" `
	Info    string ` json:"info" `
}

// /ajax/info - sys info TODO
func AX_Info(resp http.ResponseWriter, req *http.Request) {

	payload := new(InfoPayload)
	payload.Success = true

	SendAjaxPayload(resp, req, payload)
}



type LogFilesPayload struct {
	Success bool       ` json:"success" `
	Files   []cvslog.FileInfo ` json:"files" `
}

// /ajax/csvlogs - Lists available csv files
func AX_CSVListFiles(resp http.ResponseWriter, req *http.Request) {

	// Check directory exists
	var err error

	payload := new(LogFilesPayload)
	payload.Success = true

	payload.Files, err = cvslog.ListFiles(Config.CSVDir)
	if err != nil {
		SendAjaxError(resp, req, err)
		return
	}

	SendAjaxPayload(resp, req, payload)
}




// /ajax/csvlogs/stage/{file_name} - Lists available csv files
func AX_CSVStageFile(resp http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	file_name := vars["file_name"]

	fmt.Println("file_name=", file_name)

	var err error
	payload := new(LogFilesPayload)
	payload.Success = true

	_, err = cvslog.StageFile(Config.CSVDir + "/" + file_name)

	//payload.Files, err = cflog.CSVList()
	if err != nil {
		SendAjaxError(resp, req, err)
		return
	}

	SendAjaxPayload(resp, req, payload)
}

// /ajax/csvlogs/import - Imports cvs staged log
func AX_CSVImport(resp http.ResponseWriter, req *http.Request) {

	payload := map[string]interface{} {"success": true}

	payload["errors"] = cfdb.ImportStaging()

	SendAjaxPayload(resp, req, payload)
}


// /ajax/admin/db/create - Creates database tables
func AX_DBCreateAll(resp http.ResponseWriter, req *http.Request) {

	 payload := map[string]interface{} {"success": true}

	payload["errors"] = cfdb.DBCreateTables()

	SendAjaxPayload(resp, req, payload)
}

// /ajax/admin/db/info - Info re db
func AX_DBInfo(resp http.ResponseWriter, req *http.Request) {

	payload := map[string]interface{} {"success": true}

	payload["info"] = cfdb.DBInfo()

	SendAjaxPayload(resp, req, payload)
}


// /ajax/callsigns - All Callsigns
func AX_Callsigns(resp http.ResponseWriter, req *http.Request) {

	payload := map[string]interface{} {"success": true}

	payload["callsigns"] = cfdb.GetCallsigns()

	SendAjaxPayload(resp, req, payload)
}


// /ajax/callsign/{callsign} - Callsign specific
func AX_Callsign(resp http.ResponseWriter, req *http.Request) {

	payload := map[string]interface{} {"success": true}

	payload["info"] = cfdb.DBInfo()

	SendAjaxPayload(resp, req, payload)
}