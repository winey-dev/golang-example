package data

import (
	"fmt"
	"testing"
)

func TestNewLocation(t *testing.T) {
	locs := NewLocation()
	for _, loc := range locs {
		fmt.Printf("%+v\n", loc)
	}

	fmt.Println("make location data: ", len(locs))

}

func TestNewResourceValue(t *testing.T) {
	fvs := NewResourceFieldValue()
	for _, fv := range fvs {
		fv.UpdateValue()
		fmt.Printf("%-32s:%0.3f\n", fv.ItemName, fv.ItemValue)
	}

	fmt.Println("make field data: ", len(fvs))
}
func TestNewResourceTagValue(t *testing.T) {
	fvs := NewResourceTagValue()
	for _, fv := range fvs {
		fv.UpdateValue()
		fmt.Printf("%-6s[%-12s]:%0.3f\n", fv.ItemName, fv.FieldName, fv.ItemValue)
	}

	fmt.Println("make field data: ", len(fvs))
}
