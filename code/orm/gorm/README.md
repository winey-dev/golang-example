# GORM
현재 Golang ORM 진영 중 제일 많은 STAR를 기록 중임

## Schema 생성
GORM에서는  Embedding or Tag 조합으로 Scheme를 구성 할 수 있다.
Embedding을 할 경우 기본적으로 id/create_at/delete_at fields가 생성 된다.


``` golang
type Person struct {
	Id         string `gorm:"primaryKeyl;column:id"`
	FirstName  string `gorm:"column:first_name"`
	SecondName string `gorm:"column:second_name"`
}
```
