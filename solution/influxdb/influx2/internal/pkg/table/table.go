package table

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	tap = 2
)

type Table struct {
	Position    int
	StartTime   string
	EndTime     string
	ColumnsName []string
	MaxValueLen []int
	Values      [][]interface{}
}

func NewTable(pos int) *Table {
	return &Table{
		Position: pos,
	}
}

func (t *Table) Show() {
	fmt.Printf("Table Position : %d\n", t.Position)
	var repeated int

	for i := 0; i < len(t.ColumnsName); i++ {
		fmt.Printf("%-*s ", t.MaxValueLen[i], t.ColumnsName[i])

		repeated += t.MaxValueLen[i]
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", repeated+10))

	for _, valueLine := range t.Values {
		for i, v := range valueLine {
			fmt.Printf("%-*v ", t.MaxValueLen[i], v)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Store(result *api.QueryTableResult) []*Table {
	var tables []*Table
	var table *Table
	for result.Next() {
		if result.TableChanged() {
			if table != nil {
				tables = append(tables, table)
			}
			table = NewTable(result.TablePosition())
			for _, col := range result.TableMetadata().Columns() {
				if ok := IsIgnoreField(col.Name()); !ok {
					table.ColumnsName = append(table.ColumnsName, col.Name())
					table.MaxValueLen = append(table.MaxValueLen, len(col.Name())+tap)
				}
			}
		}

		valueLine := make([]interface{}, len(table.ColumnsName))
		for i, k := range table.ColumnsName {
			value := result.Record().ValueByKey(k)
			valueLen, valueStr := GetValueLen(value)
			if table.MaxValueLen[i] < valueLen+tap {
				table.MaxValueLen[i] = valueLen + tap
			}
			valueLine[i] = valueStr
		}

		table.Values = append(table.Values, valueLine)
	}
	if table != nil {
		tables = append(tables, table)
	}
	return tables
}

func GetValueLen(arg interface{}) (int, string) {
	if arg == nil {
		return 0, ""
	}
	reflect.ValueOf(arg)
	to := reflect.TypeOf(arg)
	switch to.Kind() {
	case reflect.String:
		return len(arg.(string)), arg.(string)
	case reflect.Struct:
		vo := reflect.ValueOf(arg)
		_, ok := vo.Interface().(time.Time)
		if ok {
			tFormat := arg.(time.Time).Format(time.RFC3339)
			return len(tFormat), tFormat
		}
	case reflect.Float64, reflect.Float32:
		f := fmt.Sprintf("%0.3f", arg)
		return len(f), f
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f := fmt.Sprintf("%d", arg)
		return len(f), f
	}

	f := fmt.Sprintf("%+v", arg)
	return len(f), f
}

func IsIgnoreField(name string) bool {
	/*
		if name == "table" || name == "_start" || name == "_stop" || name == "result" {
			return true
		}
	*/
	return false
}
