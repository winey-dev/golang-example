# GORM
현재 Golang ORM 진영 중 제일 많은 STAR를 기록 중임

## Schema 생성
GORM에서는  Embedding or Tag 조합으로 Scheme를 구성 할 수 있다.
Embedding을 할 경우 기본적으로 id/createdAt/deletedAt/updateAt fields가 생성 된다.

``` golang
type Person struct {
        gorm.Model
	Id         string `gorm:"primaryKeyl;column:id"`
	FirstName  string `gorm:"column:first_name"`
	SecondName string `gorm:"column:second_name"`
}

// ~/go/pkg/mod/gorm.io/mod/model.go 
package gorm

import "time"

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model or you may build your own model without it
//    type User struct {
//      gorm.Model
//    }
type Model struct {
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt DeletedAt `gorm:"index"`
}
```


``` golang
type Person struct {
	Id         string `gorm:"primaryKeyl;column:id"`
	FirstName  string `gorm:"column:first_name"`
	SecondName string `gorm:"column:second_name"`
}
```
