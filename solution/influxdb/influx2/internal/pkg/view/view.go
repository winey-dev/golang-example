package view

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Column struct {
	ColumnName     string
	MaxValueLength int
	DataType       reflect.Kind
}

func FindColum(cols []*Column, name string) (*Column, bool) {
	for _, col := range cols {
		if col.ColumnName == name {
			return col, true
		}
	}
	return nil, false

}
func View(result *api.QueryTableResult) {
	var cols []*Column

	for result.Next() {
		if result.TableChanged() {
			cols = make([]*Column, 0)
			fmt.Printf("table changed position : %d\n", result.TablePosition())
			for _, col := range result.TableMetadata().Columns() {
				if ok := IsIgnoreField(col.Name()); !ok {
					cols = append(cols, &Column{ColumnName: col.Name()})
				}
			}
		}
		values := result.Record().Values()
		SetValueLen(cols, values)
		viewTable(cols, values)
	}
}

func viewTable(cols []*Column, values map[string]interface{}) {
	// header view
	var repeat int
	for _, col := range cols {
		fmt.Printf("%*s ", col.MaxValueLength, col.ColumnName)
		repeat += col.MaxValueLength + 1 + len(col.ColumnName)
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", repeat))

	// table view
}

func SetValueLen(cols []*Column, values map[string]interface{}) {
	for k, v := range values {
		if IsIgnoreField(k) {
			continue
		}

		if col, ok := FindColumn(cols, k); !ok {
			continue
		}
		maxLength, dataType := GetValueLen(v)

		col.MaxValueLength, col.DataType = GetValueLen(v)
	}
}

func GetValueLen(arg interface{}) (int, reflect.Kind) {

}

func IsIgnoreField(name string) bool {
	if name == "table" || name == "_start" || name == "_stop" || name == "result" {
		return true
	}
	return false
}
