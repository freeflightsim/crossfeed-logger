

package cfdb

import (
	"fmt"
)


func Make4dPoint(lat float32, lon float32, alt int, spd string) (string) {

	return  fmt.Sprintf("ST_GeomFromText('POINT(%f %f %d %s)')", lat, lon, alt, spd)
}

func MakeLine(points []string) string {

	return ""
}