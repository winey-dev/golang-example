package main

import (
	"fmt"
	"go/token"
	"go/types"
	"strings"
)

func main() {
	fs := token.NewFileSet()
	m := map[string]string{
		"ITEM_VALUE_1": "10",
		"ITEM_VALUE_2": "5",
		"ITEM_VALUE_3": "3",
	}

	formula := "ITEM_VALUE_1 + ITEM_VALUE_2 * ITEM_VALUE_3"

	str := ExpressionToString(formula, m)
	fmt.Println(str)

	tv, err := types.Eval(fs, nil, token.NoPos, str)
	if err != nil {
		fmt.Printf("eval failed. err=%v\n", err)
		return
	}
	value := tv.Value.String()
	fmt.Printf("value:%s\n", value)
}

func ExpressionToString(format string, m map[string]string) string {
	for k, v := range m {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}
