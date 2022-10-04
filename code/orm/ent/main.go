package main

import (
	"context"
	"ent/generate/ent"
	"fmt"
	"os"

	driver "github.com/go-sql-driver/mysql"
)

func main() {
	dbcfg := driver.NewConfig()
	dbcfg.User = "root"
	dbcfg.Passwd = "admin1234"
	dbcfg.Net = "tcp"
	dbcfg.DBName = "semina"
	dbcfg.Addr = "127.0.0.1:3306"

	client, err := ent.Open("mysql", dbcfg.FormatDSN())
	if err != nil {
		fmt.Printf("db open failed. err=%v\n", err)
		os.Exit(1)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		fmt.Println("schema create failed.. ", err)
		os.Exit(1)
	}

	u, err := client.Person.Get(context.TODO(), "smlee")
	if err != nil {
		fmt.Println("get failed.. ", err)
		os.Exit(1)
	}

	fmt.Println(u.String())

}
