
package cvslog

import (
	"os"
	"errors"
	"io/ioutil"

)

// Path to crossfeed-dailies/csv
var CSVDir string


type FileInfo struct {
	FileName string ` json:"filename"  `
	Date     string ` json:"date"  `
	Size     int64  ` json:"size"  `
}


// Return error if csv path missing
func DirExists() error {
	if _, err := os.Stat(CSVDir); os.IsNotExist(err) {
		return errors.New("cvs dir `"+ CSVDir + "` does not exist")
	}
	return nil
}

// Return error if csv file_name missing
func FileExists(file_name string) error {

	if errd := DirExists(); errd != nil {
		return errd
	}
	if _, err := os.Stat(CSVDir + "/" + file_name); os.IsNotExist(err) {
		return errors.New("cvs file `"+ file_name + "` does not exist")
	}
	return nil
}

// Lists All available lof giles
func CSVList() ([]FileInfo, error) {

	// list of file to return
	fileinfos := make([]FileInfo, 0)

	if errd := DirExists(); errd != nil {
		return fileinfos, errd
	}

	// Read files from os
	files, err := ioutil.ReadDir(CSVDir)
	if err != nil {
		return fileinfos, err
	}

	// copy into our struct
	for _, f := range files {
		fileinfos = append(fileinfos, FileInfo{FileName: f.Name(),
			Size: f.Size(),
			Date: f.Name()[8 : len(f.Name())-4]})
	}
	return fileinfos, nil
}
