

package cflog


// Database record for callsign
type Callsign struct {

	CallsignId int64  ` gorm:"primary_key" `
	Callsign   string

}

func(cs Callsign) TableName() string {
	return "callsign"
}


