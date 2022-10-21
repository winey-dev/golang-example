package data

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

var ReqResp = []string{
	"request_count",
	"response_succ",
	"response_fail",
}

var Code = []string{
	"400",
	"403",
	"404",
	"500",
	"501",
}

type ReqRespValue struct {
	ItemName  string
	Code      string
	ItemValue int
}

type ReqRespFaker struct {
	SuccAvg float64 `faker:"boundary_start=0.7, boundary_end=1"`
	Count   int     `faker:"boundary_start=0, boundary_end=30"`
}

func NewReqRespValue() []ReqRespValue {
	var items []ReqRespValue
	rrf := ReqRespFaker{}
	faker.FakeData(&rrf)

	//Count Make
	count := ReqRespValue{
		ItemName:  ReqResp[0],
		ItemValue: rrf.Count,
	}
	items = append(items, count)

	if count.ItemValue == 0 {
		return items
	}

	succCount := float64(count.ItemValue) * rrf.SuccAvg

	succ := ReqRespValue{
		ItemName:  ReqResp[1],
		ItemValue: int(math.Floor(succCount)),
	}

	items = append(items, succ)
	remainCount := count.ItemValue - succ.ItemValue

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Code), func(i, j int) {
		Code[i], Code[j] = Code[j], Code[i]
	})

	codeIndex := 0
	for remainCount > 0 {
		var fail ReqRespValue
		failCount := rand.Intn(remainCount)
		if failCount == 0 {
			failCount = 1
		}
		if codeIndex == len(Code)-1 {
			fail = ReqRespValue{
				ItemName:  ReqResp[2],
				Code:      Code[codeIndex],
				ItemValue: remainCount,
			}
			break
		} else {
			fail = ReqRespValue{
				ItemName:  ReqResp[2],
				Code:      Code[codeIndex],
				ItemValue: failCount,
			}
		}
		remainCount = remainCount - failCount
		codeIndex++
		items = append(items, fail)
	}
	return items
}
