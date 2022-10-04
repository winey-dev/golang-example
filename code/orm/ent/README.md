# ENT
FaceBook에서 만든 Golang ORM 이다.

Code Generate 방식이다. 


## INSTALL 
먼저 ent 명령어를 설치 해야하 한다.
`go install entgo.io/ent/cmd/ent:latest`


## Schema 생성 
먼저 Schema 생성을 위한 ent module을 init 한다.
```
/* ${A} ${B} 는 그냥 본인이 원하는 DIRECTORY이다.*/
$ ent init User Pet
$ ls ent/schema
 user.go pet.go
```

위에서 생성된 user.go pet.go에 schema를 정의한다.

## Code Generate 
```
$ go generate ent/
```
위 명령어 수행 후 user/pet 에 대한 CRUD Go Code가 생성된것을 확인한다.


