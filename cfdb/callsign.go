

package cfdb

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

// Add indexes to callsign table
func DBIndexCallsign(){
	DBAddUniqueIndex("callsign", "upper(callsign)", "callsign_upper")
}

// Database record for a callsign
type Callsign struct {
	CallsignID int64  ` db:"callsign_id" `
	Callsign   string  ` db:"callsign" `

}


func GetCallsigns(){

}


