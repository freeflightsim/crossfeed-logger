
package cfdb





var SCHEMA_STAGING  string = `
CREATE TABLE staging (
    s_id serial NOT NULL,
    flightid bigint,
	callsign varchar(20),
	lat float,
	lon float,
	alt_ft int,
	aero varchar(20),
	spd_kt int,
	true_hdg int,
	dist_nm int,
	ts timestamp,
	total_sec int,
	source varchar(255),
	source_info varchar(255),
	row_error varchar(255),
	CONSTRAINT s_id PRIMARY KEY (s_id)
)
WITH (
	  OIDS=FALSE
);
`




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
	Timestamp   string
	TotalSec int64

}
