package main

import (
	"fmt"

	driver "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Id         string `gorm:"primaryKeyl;column:id"`
	FirstName  string `gorm:"column:first_name"`
	SecondName string `gorm:"column:second_name"`
}

func (p Person) TableName() string {
	return "person"
}

func main() {
	fmt.Println("vim-go")
	dbcfg := driver.NewConfig()
	dbcfg.User = "root"
	dbcfg.Passwd = "admin1234"
	dbcfg.Net = "tcp"
	dbcfg.DBName = "semina"
	dbcfg.Addr = "127.0.0.1:3306"

	conn, err := gorm.Open(mysql.Open(dbcfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		fmt.Println("db connect failed. ", err)
		return
	}

	p := []Person{
		Person{
			Id:         "lsmin0703",
			FirstName:  "Lee",
			SecondName: "Seungmin",
		},
		Person{
			Id:         "smlee",
			FirstName:  "Lee",
			SecondName: "Seungmin",
		},
	}

	err = conn.AutoMigrate(&Person{})
	if err != nil {
		fmt.Println("migrate failed. ", err)
		return
	}

	db := conn.Create(&p)
	if db.Error != nil {
		fmt.Println("insert person failed. ", db.Error)
	}

	var ret []Person

	db = conn.Find(&ret)
	fmt.Println(db)
	if db.Error != nil {
		fmt.Println("find person failed. ", db.Error)
	}
	fmt.Println(ret)
}
