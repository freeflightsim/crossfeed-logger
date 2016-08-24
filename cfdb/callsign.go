

package cfdb

var SCHEMA_CALLSIGN  string = `
CREATE TABLE flights (
	callsign_id int64.
	callsign varchar(20),
	CONSTRAINT flight_id PRIMARY KEY (callsign_id)
)
WITH (
	  OIDS=FALSE
);
`

// Database record for a callsign
type CallSign struct {
	ID int64  ``
	CallSign   string

}


func GetCallSigns(){

}


