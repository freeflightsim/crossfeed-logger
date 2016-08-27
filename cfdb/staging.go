
package cfdb


import (
	"fmt"
)


var SCHEMA_STAGING  string = `
CREATE TABLE staging (
    s_id serial NOT NULL,
    flightid bigint,
	callsign varchar(20),
	callsign_id int,
	lat float,
	lon float,
	alt_ft int,
	aero varchar(20),
	aero_id int,
	spd_kt int,
	true_hdg int,
	dist_nm int,
	ts timestamp,
	total_sec int,
	source varchar(255),
	source_info varchar(255),
	row_error varchar(255),
	imported boolean,
	CONSTRAINT idx_s_id PRIMARY KEY (s_id)
)
WITH (
	  OIDS=FALSE
);
`
// Add indexes to staging
func DBIndexStaging(){
	DBAddIndex("staging", "callsign_id", "")
	DBAddIndex("staging", "aero_id", "")
}





func ImportStaging() []string {
	fmt.Println("================\nImporting callsigns\n")
	//var err error
	//var sql string
	ret := make([]string, 0)

	importFlights()
	return ret
	importAeros()
	importCallsigns()



	return ret
}

// cvs columns are
//  fid,callsign,lat,lon,alt_ft,model,spd_kts,hdg,dist_nm,update,tot_secs
type DBStagingRow struct {
	ID int64  ` db:"s_id" `
	FlightID int64  ` db:"flightid" `
	Aero     string  ` db:"aero" `
	AeroID     string  ` db:"aero_id" `
	Callsign string  ` db:"callsign" `
	CallsignID string  ` db:"callsign_id" `
	Lat      float32 ` db:"lat" `
	Lon      float32 ` db:"lon" `
	AltFt    int ` db:"alt_ft" `
	SpdKt    string ` db:"spd_kt" `
	TrueHdg  int32 ` db:"true_hdg" `
	DistNm   int32 ` db:"dist_nm" `
	Update   string ` db:"ts" `
	TotalSec int64 ` db:"total_sec" `
	Source	string ` db:"source" `
	SourceInfo	string ` db:"source_info" `
	RowError *string ` db:"row_error" `
	Imported *bool ` db:"imported" `
}


func importFlight(f_id int64) []string {

	fmt.Println("================\nFlight" + "ID" + "\n")
	var err error
	var sql string
	ret := make([]string, 0)

	rows := []DBStagingRow{}
	sql = "select * from staging "
	sql += " where flightid = $1"
	err = Dbx.Select(&rows, sql, f_id)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}
	fmt.Println("rows=", len(rows))

	// Check if flight already exists
	var flight *Flight
	flight, err = GetFlight(f_id)
	if flight == nil {
		flight = new(Flight)
		flight.FlightID = f_id
		sql := " insert into flight("
		sql += "flight_id"
		sql += ") values ("
		sql += "$1 "
		sql += ")"
		_, err = Dbx.Exec(sql, flight.FlightID)
		if err != nil {
			fmt.Println("Err=", err, sql)
			ret = append(ret, err.Error() )
		}
	}

	return ret
}

func importFlights() []string {

	fmt.Println("================\nImporting flights\n")
	var err error
	var sql string
	ret := make([]string, 0)

	// get unstaged Flights
	var row DBStagingRow
	sql = "select distinct(flightid) from staging "
	sql += " where imported is null "
	sql += " and aero_id is not null "
	sql += " and callsign_id is not null "
	sql += " limit 1 "
	err = Dbx.Get(&row, sql)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}
	fmt.Println("row=", row, row.Callsign)

	importFlight(row.FlightID)
	return ret
}




func importAeros() [] string {

	fmt.Println("================\nImporting aeros\n")
	var err error
	var sql string
	ret := make([]string, 0)

	/// get unidentifies aero from staging
	aeros := []Aero{}
	sql = "select distinct(aero) as aero from staging "
	sql += " where imported is null  and aero_id is null "
	err = Dbx.Select(&aeros, sql)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}
	fmt.Println("aeros", aeros)

	// insert unidentified aero's
	sqli := "insert into aero(aero)values($1);"
	for idx, a := range aeros {
		fmt.Println(a)
		_, err = Dbx.Exec(sqli, a.Aero)
		if err != nil {
			fmt.Println("Err=", err, sql)
			ret = append(ret, err.Error() )
		}
		if idx == 4 {
			break
		}
	}

	// Update existing aero_ids
	sql = "  update staging "
	sql += " set aero_id = aero.aero_id "
	sql += " from aero "
	sql += " where upper(aero.aero) = upper(staging.aero) "
	sql += " and staging.aero_id is null and staging.imported is null  "
	_, err = Dbx.Exec(sql)
	fmt.Println(sql)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}

	return ret

}


func importCallsigns() [] string {

	fmt.Println("================\nImporting callsigns\n")
	var err error
	var sql string
	ret := make([]string, 0)

	/// get unidentifies callsigns from staging
	callsigns := []Callsign{}
	sql = "select distinct(callsign) as callsign from staging "
	sql += " where imported is null  and callsign_id is null "
	err = Dbx.Select(&callsigns, sql)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}
	fmt.Println("callsigns", callsigns)

	// insert unidentified callsigns
	sqli := "insert into callsign(callsign)values($1);"
	for idx, cs := range callsigns {
		fmt.Println(cs)
		_, err = Dbx.Exec(sqli, cs.Callsign)
		if err != nil {
			fmt.Println("Err=", err, sql)
			ret = append(ret, err.Error() )
		}
		if idx == 4 {
			break
		}
	}

	// Update existing callsigns
	sql = "  update staging "
	sql += " set callsign_id = callsign.callsign_id "
	sql += " from callsign "
	sql += " where upper(callsign.callsign) = upper(staging.callsign) "
	sql += " and staging.callsign_id is null and staging.imported is null  "
	_, err = Dbx.Exec(sql)
	fmt.Println(sql)
	if err != nil {
		fmt.Println("Err=", err, sql)
		ret = append(ret, err.Error() )
	}

	return ret

}



