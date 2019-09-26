package balance

import (

	"time"
	"errors"
	"math/rand"
)

func init() {
	//RegisterBalancer(, &RandomBalance{})
	mBalanceManger["random"]=&RandomBalance{}
}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts)==0{
		err=errors.New("no instance")
		return
	}
	rand.Seed(time.Now().UnixNano())
	return insts[rand.Intn(len(insts))],nil
}
