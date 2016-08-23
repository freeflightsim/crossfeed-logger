package cvslog

import (
	"github.com/jmoiron/sqlx"
)

var DbLog *sqlx.DB



// cvs columns are
//  fid,callsign,lat,lon,alt_ft,model,spd_kts,hdg,dist_nm,update,tot_secs
type LogRow struct {
	FlightId string
	Aero     string
	CallSign string
	Lat      float32
	Lon      float32
	AltFt    int64
	SpdKt    int
	HdgTrue  int
	DistNm   int64
	Update   string
	TotalSec int64
	LineNo	int64
}




