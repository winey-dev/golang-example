package main

import "fmt"

func main() {
	m1 := map[string]string{
		"error_code":  "500",
		"error_point": "internal_server",
		"app_name":    "em-ui",
		"pod_name":    "em-ui-0",
	}

	m2 := map[string]string{
		"app_name":    "em-ui",
		"pod_name":    "em-ui-0",
		"error_point": "internal_server",
		"error_code":  "500",
	}

	fmt.Println(m1, m2)

	for i := 0; i < 1000000; i++ {
		if fmt.Sprintf("%v", m1) != fmt.Sprintf("%v", m2) {
			fmt.Printf("not equals %v:%v\n", m1, m2)
		}
	}
}
