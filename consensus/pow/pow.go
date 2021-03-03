package pow

import "fmt"

type PoW struct {

}

func (pow PoW) Run() interface{} {
	fmt.Println("这是PoW方式的共识机制算法")
	return nil
}
