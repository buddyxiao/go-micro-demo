package service

import (
	"math/rand"
	"time"
)

var randS randService

type randService struct {
}

func Rand() randService {
	return randS
}

func (r randService) GetRand(int642 int64) int64 {
	rand.Seed(time.Now().UnixMilli())
	return rand.Int63n(int642)
}
