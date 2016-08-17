package server

import (
	"errors"
	"net/http"
	"os"

	"encoding/json"
	"gopkg.in/yaml.v2"

	"github.com/gorilla/mux"

	"io/ioutil"
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

type FileInfo struct {
	FileName string ` json:"filename"  `
	Date     string ` json:"date"  `
	Size     int64  ` json:"size"  `
}

type FilesPayload struct {
	Success bool       ` json:"success" `
	Files   []FileInfo ` json:"files" `
}

// /ajax/csv-files - Lists abailable csv files
func AX_CSVFiles(resp http.ResponseWriter, req *http.Request) {

	// Check directory exists
	if _, err := os.Stat(Config.CSVDir); os.IsNotExist(err) {
		SendAjaxError(resp, req, errors.New("cvs dir `"+Config.CSVDir+"` does not exist"))
		return
	}

	payload := new(FilesPayload)
	payload.Success = true
	payload.Files = make([]FileInfo, 0, 0)

	// Read files list into payload
	files, err := ioutil.ReadDir(Config.CSVDir)
	if err != nil {
		SendAjaxError(resp, req, err)
		return
	}
	for _, f := range files {
		payload.Files = append(payload.Files, FileInfo{FileName: f.Name(),
			Size: f.Size(),
			Date: f.Name()[8 : len(f.Name())-4]})
	}

	SendAjaxPayload(resp, req, payload)
}
