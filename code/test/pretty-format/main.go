package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	user := NewTable("user")
	user.SetColumns([]string{"uid", "name", "email"})
	for i := 0; i < 10; i++ {
		user.SetRecord([]interface{}{
			fmt.Sprintf("id_%d", i),
			fmt.Sprintf("name_%d", i),
			fmt.Sprintf("email_%d", i),
		})
	}

	user.Print()
}

/*
Table
*/
const (
	tap = 2
)

type Table struct {
	name            string
	columns         []string
	values          [][]interface{}
	recordMaxLength []int
}

type TableI interface {
	SetColumns(cols []string)
	SetRecord(values []interface{}) error
	Print()
}

func NewTable(name string) TableI {
	return &Table{
		name: name,
	}
}
func (t *Table) SetColumns(cols []string) {
	for _, col := range cols {
		t.columns = append(t.columns, col)
		t.recordMaxLength = append(t.recordMaxLength, len(col)+tap)
	}
}

func (t *Table) SetRecord(values []interface{}) error {
	if len(values) == 0 {
		return errors.New("empty value")
	}
	if len(t.columns) != len(values) {
		return errors.New("invalued value length")
	}

	record := make([]interface{}, len(t.columns))

	for i, value := range values {
		length, data := getValue(value)

		if t.recordMaxLength[i] < length+tap {
			t.recordMaxLength[i] = length + tap
		}
		record[i] = data
	}
	t.values = append(t.values, record)
	return nil
}

func (t *Table) Print() {
	fmt.Printf("table : %s\n", t.name)
	var div int
	for i := 0; i < len(t.columns); i++ {
		fmt.Printf("%-*s ", t.recordMaxLength[i], t.columns[i])
		div += t.recordMaxLength[i]
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", div))

	for _, valueLine := range t.values {
		for i, v := range valueLine {
			fmt.Printf("%-*v ", t.recordMaxLength[i], v)
		}
		fmt.Println()
	}
	fmt.Println()

}

func getValue(value interface{}) (int, string) {
	if value == nil {
		return 0, ""
	}

	reflect.ValueOf(value)
	to := reflect.TypeOf(value)
	switch to.Kind() {
	case reflect.String:
		return len(value.(string)), value.(string)
	case reflect.Struct:
		vo := reflect.ValueOf(value)
		_, ok := vo.Interface().(time.Time)
		if ok {
			tFormat := value.(time.Time).Format(time.RFC3339)
			return len(tFormat), tFormat
		}
	case reflect.Float64, reflect.Float32:
		f := fmt.Sprintf("%0.3f", value)
		return len(f), f
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f := fmt.Sprintf("%d", value)
		return len(f), f
	}

	f := fmt.Sprintf("%+v", value)
	return len(f), f

}
