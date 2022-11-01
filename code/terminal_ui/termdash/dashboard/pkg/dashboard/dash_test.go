package dashboard

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	chart := Data.Charts[0]
	chart.InitValue(Data.Interval, Data.Range)
	chart.TableView()

	fmt.Println()
	fmt.Println()

	chart.RotateFloats()
	chart.TableView()
}
