package view

import (
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

func ViewDetail(result *api.QueryTableResult) {
	// key : value len
	var storeKey map[string]int
	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("change table")
			fmt.Println()
			storeKey = make([]string, 0, 0)
			for _, col := range result.TableMetadata().Columns() {
				if ok := viewColumn(col.Name()); ok {
					storeKey = append(storeKey, col.Name())
				}
			}
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------")
		}
		values := result.Record().Values()

		for _, key := range storeKey {
			viewValue(key, values[key])
		}
		fmt.Println()
	}
}

func viewValue(key string, value interface{}) {
	if key == "" {
		return
	}
	if key == "_time" {
		fmt.Printf("%-38v ", value)
	} else if key == "_measurement" || key == "_field" {
		fmt.Printf("%-32v ", value)
	} else {
		fmt.Printf("%-16v ", value)
	}

}

func viewColumn(name string) bool {
	if IsIgnoreField(name) {
		return false
	}

	if name == "_time" {
		fmt.Printf("%-38s ", name)
	} else if name == "_measurement" || name == "_field" {
		fmt.Printf("%-32s ", name)
	} else {
		fmt.Printf("%-16s ", name)
	}
	return true
}

func Vview(result *api.QueryTableResult) {
	var header bool
	var storeKey map[string]int
	for result.Next() {
		if result.TableChanged() {
			Clear(storeKey)
			fmt.Printf("table changed position : %d\n", result.TablePosition())
			for _, col := range result.TableMetadata().Columns() {
				if ok := viewColumn(col.Name()); ok {
					storeKey[col.Name()] = 0
				}
			}
		}
		values := result.Record().Values()

	}
}

func SetValueLen(storeKey map[string]int, values map[string]interface{}) {
	for k, v := range values {
		if IsIgnoreField(k) {
			continue
		}
		storeKey[k] = GetValueLen(v)
	}
}

func GetValueLen(arg interface{}) int {

}

func Clear(storeKey map[string]int) {
	for k, _ := range storeKey {
		delete(storeKey, k)
	}
}

func IsIgnoreField(name string) bool {
	if name == "table" || name == "_start" || name == "_stop" || name == "result" {
		return true
	}
	return false
}
