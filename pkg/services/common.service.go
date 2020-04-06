package services

import (
	"math/rand"
	"time"
)

type CommonService struct {
}

func (commonService *CommonService) Init() {
}

func (c CommonService) GetRandom() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}

func (c CommonService) GetRandomFromList(list []string) string {
	rand.Seed(time.Now().Unix())
	return list[rand.Intn(len(list))]
}
