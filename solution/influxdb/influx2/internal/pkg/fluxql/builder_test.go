package fluxql

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	query := NewBuilder().
		From("bucket").
		RangeString("-5h").
		Measurement("smaple").
		Field("CPU_USEAGE").
		Filter("node_name", "worker-1").
		Filter("pod_name", "pod-1").
		Filters([]FilterKeyValue{
			FilterKeyValue{Key: "a", Value: "b", Op: OR},
			FilterKeyValue{Key: "c", Value: "d", Op: END},
		}).
		Filters([]FilterKeyValue{
			FilterKeyValue{Key: "d", Value: "e", Op: AND},
			FilterKeyValue{Key: "f", Value: "g", Op: END},
		}).
		KeepColumns([]string{"a", "b", "c"}).
		Builder().
		String()
	t.Logf("make query:\n%s", query)
}
