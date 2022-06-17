package cdd

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	unixTimeUTC:=time.UnixMilli(1654704000000) //gives unix time stamp in utc

	unitTimeInRFC3339 :=unixTimeUTC.Format("2006-01-02") // converts utc time to RFC3339 format

	fmt.Println("unix time stamp in UTC :--->",unixTimeUTC)
	fmt.Println("unix time stamp in unitTimeInRFC3339 format :->",unitTimeInRFC3339)
}
