

package cvslog

import(
	"bufio"
	"os"
	"encoding/csv"
	"fmt"
	"io"
	"time"

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
	Lines int64
	Started string
	Ended string
}

var sql_insert string = `
insert into staging(
	flightid, callsign, aero,
	lat, lon, alt_ft, spd_kt, true_hdg,
	dist_nm, ts, total_sec,
	file_name, line_no
) values (
	$1, $2, $3,
	$4, $5, $6, $7, $8,
	$9, $10, $11,
	$12, $13
)
`


func ImportFile(file_name string)(*ImportInfo, error){

	if _, err := os.Stat(file_name); os.IsNotExist(err) {
		return nil, err
	}

	f, erro := os.Open(file_name)
	if erro != nil {
		return nil, erro
	}

	info := &ImportInfo{Lines: -1, Started: time.Now().UTC().String()}
	fmt.Println("Start=", info)

	//db_table, errt := cfdb.CreateStagingTable(file_name)
	cfdb.Dbx.Ping()

	stmt, err := cfdb.Dbx.Prepare(sql_insert)
	if err != nil {
		return info, err
	}

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {

		info.Lines++
		if info.Lines == 10 {
			//return info, nil
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


		//r := DBLogRow{}
		//r.FlightId = record[0]
		//r.CallSign = record[1]
		//r.Lat = record[2]
		//r.Lon = record[3]
		//r.AltFt = record[4]
		//r.Aero = record[5]
		//r.SpdKt = record[6]
		//r.TrueHdg = record[7]

		res := stmt.QueryRow(rec[0], rec[1], rec[5],
							rec[2], rec[3], rec[4],rec[6],
							rec[7], rec[8], rec[9],
							info.Lines, file_name)
		fmt.Println("res=", res)
		//for value := range record {
		//	fmt.Printf("  %v\n", record[value])
		//}
	}
	info.Ended = time.Now().UTC().String()
	fmt.Println("End=",info)


	return info, nil
}