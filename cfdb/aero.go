


package cfdb

var SCHEMA_AERO  string = `
CREATE TABLE aero (
	aero_id serial NOT NULL,
	aero varchar(20),
	CONSTRAINT idx_aero_id PRIMARY KEY (aero_id)
)
WITH (
	  OIDS=FALSE
);
`

// Add indexes to callsign table
func DBIndexAero(){
	DBAddIndex("aero", "upper(aero)", "aero_upper")
}

// Database record for a callsign
type Aero struct {
	AeroID int64  ``
	Aero   string

}

