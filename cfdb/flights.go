

package cfdb


var SCHEMA_FLIGHT  string = `
CREATE TABLE flights (
	flight_id int,
	flight_path geometry(Line,3857),
	CONSTRAINT flight_id PRIMARY KEY (flight_id)
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