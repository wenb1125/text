package pos

import "fmt"

type PoS struct {

}

func (pos PoS) Run() interface{} {
	fmt.Println("这是PoS方式的共识机制算法")
	return nil
}