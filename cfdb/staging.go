
package cfdb





var SCHEMA_STAGING  string = `
CREATE TABLE staging (
	flightid int,
	callsign char(20),
	lat float,
	lon float,
	alt_ft int,
	aero char(20),
	spd_kt int,
	true_hdg int,
	dist_nm int,
	ts datetime,
	total_sec int,
	file_name varchar(255)
)
WITH (
	  OIDS=FALSE
);
`


func CreateStagingTable(file_name string) error {


	return nil
}


// cvs columns are
//  fid,callsign,lat,lon,alt_ft,model,spd_kts,hdg,dist_nm,update,tot_secs
type DBLogRow struct {
	FlightId string
	Aero     string
	CallSign string
	Lat      string
	Lon      string
	AltFt    string
	SpdKt    string
	TrueHdg  int
	DistNm   int64
	Update   string
	TotalSec int64
	LineNo	int64
}
