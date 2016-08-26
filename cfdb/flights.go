

package cfdb


var SCHEMA_FLIGHT  string = `
CREATE TABLE flight (
	flight_id int,
	flight_path geometry(LINESTRING,3857),
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
func GetFlightInfo(flight_id int64)(Flight, error){

	var fl Flight
	//var err error
	//err = db.Get(&fl, "select from flights where flght_id = ?", flight_id)

	return fl, nil

}