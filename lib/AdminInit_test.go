package lib

import (
	"testing"
)

func Test_tsToShortdate(t *testing.T) {

	shortDate := TsToShortdate("2016-01-02 18:18:18 000")

	if shortDate == "2016-01-02" {
		t.Error("TsToShortdate error")
	} //end of if

}
