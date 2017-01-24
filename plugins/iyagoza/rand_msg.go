package iyagoza

import (
	"math/rand"
	"time"
)

type randMsg struct {
	r    *rand.Rand
	msgs []string
}

func newRandMsg(msgs []string) *randMsg {
	return &randMsg{
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
		msgs: msgs,
	}
}

func (r *randMsg) GetMessage() string {
	return r.msgs[r.r.Intn(len(r.msgs))]
}
