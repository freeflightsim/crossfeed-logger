

package cvslog

import(
	"bufio"
	"os"
	"encoding/csv"
	"fmt"
	"io"
	"time"
	"path/filepath"

	"github.com/freeflightsim/crossfeed-logger/cfdb"
)

// fid,callsign,lat,lon,alt_ft,model,spd_kts,hdg,dist_nm,update,tot_secs
const(
	C_FID = iota
	C_CALLSIGN
	C_LAT
	C_LON
	C_ALT_FT
	C_MODEL
	C_SPD_KT
	C_HDG_TRUE
	C_DIST_NM
	C_AIRBONE_SECS
)

type ImportInfo struct {
	Source string ` json:"source" `
	Lines int64 ` json:"lines" `
	Started string
	Ended string
}




func ImportFile(path_to_file string)(*ImportInfo, error){

	// check file exists, readable etc
	if _, err := os.Stat(path_to_file); os.IsNotExist(err) {
		return nil, err
	}
	f, erro := os.Open(path_to_file)
	if erro != nil {
		return nil, erro
	}

	// info on this import
	info := &ImportInfo{Lines: -1, Started: time.Now().UTC().String()}
	fmt.Println("Start=", info)
	// TODO add source record

	// check db is alive
	cfdb.Dbx.Ping()

	// Nuke existing sources for now
	source_name := filepath.Base(path_to_file)
	// Nuke existing records for this file
	_, errn := cfdb.Dbx.Exec("delete from staging where file_name = ' " +  source_name + " ' ")
	if errn != nil {
		return info, errn
	}


	// fid,callsign,lat,lon,alt_ft, model,spd_kts,hdg,dist_nm,update,tot_secs
	// 0   1        2   3   4      5      6       7   8       9      10
	var sql_insert string = `
		insert into staging(
			flightid, callsign,
			lat, lon, alt_ft,
			aero,
			spd_kt, true_hdg,
			dist_nm, ts, total_sec,
			file_name, line_no
		) values (
			$1, $2,
			$3,	$4, $5,
			$6,
			$7, $8,
			$9, $10, $11,
			$12, $13
		)
		`

	tx, errt := cfdb.Dbx.Begin()
	if errt != nil {
		return info, errt
	}
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {

		info.Lines++
		if info.Lines == 1 {
			continue // bye bye first line with headers
		}
		rec, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		//fmt.Println(len(record))
		if len(rec) == 0 {
			fmt.Println("WTF=", rec)
			continue
		}

		/*
		_, erri := cfdb.Dbx.Exec(sql_insert,
							rec[0], rec[1],
							rec[2],	rec[3], rec[4],
							rec[5],
							rec[6], rec[7],
							rec[8], rec[9], rec[10],
							source_name, info.Lines)
		*/
		_, erri := tx.Exec(sql_insert,
			rec[0], rec[1],
			rec[2],	rec[3], rec[4],
			rec[5],
			rec[6], rec[7],
			rec[8], rec[9], rec[10],
			source_name, info.Lines)

		if erri != nil {
			fmt.Println("erri=", erri)
		}
		if info.Lines % 1000 == 0 {
			fmt.Println("line=", info.Lines)
		}
		//for value := range record {
		//	fmt.Printf("  %v\n", record[value])
		//}
	}
	errf := tx.Commit()
	if errf != nil {
		fmt.Println("errf=", errf)
	}
	info.Ended = time.Now().UTC().String()
	fmt.Println("End=",info)


	return info, nil
}