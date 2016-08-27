

package cfdb

import (
	"fmt"
	"database/sql"

)

var SCHEMA_FLIGHT  string = `
CREATE TABLE flight (
	flight_id bigint,
	flight_path geometry(MultiPointZM, 4326),
	CONSTRAINT idx_flight_id PRIMARY KEY (flight_id)
)
WITH (
	  OIDS=FALSE
);
`


type Flight struct {
	FlightID int64  `db:"flight_id"  json:"flight_id" `

}


// Return info about a flight
func GetFlight(flight_id int64)(*Flight, error){

	fl := new(Flight)
	var err error
	q := " select flight_id "
	q += " from flight where flight_id = $1"
	err = Dbx.Get(fl, q, flight_id)
	//err = row.Scan(fl)
	if err == sql.ErrNoRows {
		fmt.Println("NOT FOUNC", err)
		return nil, nil
	}
	return fl, nil

}