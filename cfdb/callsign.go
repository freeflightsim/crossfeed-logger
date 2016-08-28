

package cfdb

import (
	"fmt"
)

var SCHEMA_CALLSIGN  string = `
CREATE TABLE callsign (
	callsign_id serial NOT NULL,
	callsign varchar(20),
	CONSTRAINT idx_callsign_id PRIMARY KEY (callsign_id)
)
WITH (
	  OIDS=FALSE
);

`
var VIEW_CALLSIGN string = `
CREATE OR REPLACE VIEW v_callsign as
	select callsign.callsign_id, callsign.callsign,
	count(flight.callsign_id) as flights_count
	from callsign
	inner join flight on flight.callsign_id = callsign.callsign_id
	group by callsign.callsign_id
`

// Add indexes to callsign table
func DBIndexCallsign(){
	DBAddUniqueIndex("callsign", "upper(callsign)", "callsign_upper")
}

// Database record for a callsign
type CallsignRow struct {
	CallsignID int64  ` db:"callsign_id" `
	Callsign   string  ` db:"callsign" json:"callsign" `
}

// View for a callsign
type CallsignView struct {
	//CallsignID int64  ` db:"callsign_id" `
	Callsign   string  ` db:"callsign" json:"callsign" `
	FlightsCount    int  ` db:"flights_count" json:"flights_count" `
}

func GetCallsigns() []CallsignView{

	var recs =  []CallsignView{}

	sql := "select callsign, flights_count "
	sql += " from v_callsign "
	sql += " order by callsign asc "

	err := Dbx.Select(&recs, sql)
	if err != nil {
		fmt.Println("getcall", err)
	}
	return recs

}




