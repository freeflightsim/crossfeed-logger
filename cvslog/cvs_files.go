
package cvslog

import (

	"io/ioutil"

)


// Basic CSV file Info
type FileInfo struct {
	FileName string ` json:"filename"  `
	Date     string ` json:"date"  `
	Size     int64  ` json:"size"  `
}

// Lists log files available (later we need to recognise flag of imported)
func ListFiles(csv_dir string) ([]FileInfo, error) {

	// list of file to return
	fileinfos := make([]FileInfo, 0)

	// Read files from os
	files, err := ioutil.ReadDir(csv_dir)
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
