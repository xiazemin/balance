package balance

import "github.com/golang/go/src/fmt"

type Instance struct {
	Ip   string
	Port int
}

func NewInstance(ip string, port int) *Instance {
	return &Instance{
		Ip:   ip,
		Port: port,
	}
}

func (this *Instance) String() string {
	return this.Ip + ":" + fmt.Sprintf("%d", this.Port)
}

type IBalance interface {
	DoBalance(insts []*Instance, key ...string) (inst *Instance, err error)
}

var mBalanceManger map[string]IBalance

func init()  {
	mBalanceManger=make(map[string]IBalance)
}

func Get(name string)IBalance  {
	return mBalanceManger[name]
}
