package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/xiazemin/balance/balance"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for i:=0;i<2;i++{
		balanceName="iphash"
		fmt.Println(balance.Get(balanceName).DoBalance(insts,"192.168.1.2"))
		fmt.Println(balance.Get(balanceName).DoBalance(insts,"192.168.2.2"))

	}
	for i:=0;i<20;i++{
		balanceName="roundrobin"
		fmt.Println(balance.Get(balanceName).DoBalance(insts,balanceName))

	}
	balanceName = "random"
	for {
		inst, err := balance.Get(balanceName).DoBalance(insts,balanceName)
		if err != nil {
			fmt.Println("do balance err:", err)
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}