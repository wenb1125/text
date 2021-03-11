package client

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"xianfengChain/chain"
)

/**
 * 客户端（命令行窗口工具），主要用户实现与用户进行动态交互
 * 		①将帮助信息等输出到控制台
 * 		②读取参数并解析，根据解析结果调用对应的项目功能
 */
type Client struct {
	Chain chain.BlockChain
}

func (client *Client) Run() {
	if len(os.Args) == 1 {//用户没有输入任何指令
		client.Help()
		return
	}
	//1.解析命令行参数
	command := os.Args[1]
	//2.确定用户输入的命令
	switch command {
	case CREATECHAIN:
		flag.NewFlagSet(CREATECHAIN,flag.ExitOnError)

	case GENERATEGENSIS:
		generateGensis := flag.NewFlagSet(GENERATEGENSIS,flag.ExitOnError)
		gensis := generateGensis.String("gensis","","创世区块中的自定义数据")
		generateGensis.Parse(os.Args[2:])
		//1.先判断是否已存在创世区块
		hashBig := new(big.Int)
		hashBig.SetBytes(client.Chain.LastBlock.Hash[:])
		if hashBig.Cmp(big.NewInt(0)) == 1 {//创世区块hash值不为0，即有值
			fmt.Println("抱歉，创世区块已存在，无法覆盖写入")
			return
		}
		//2.如果创世区块不存在，才去调用creategenesis
		client.Chain.CreateGenesis([]byte(*gensis))
		fmt.Println("恭喜，创世区块创建成功并写入数据")
	case ADDNEWBLOCK:
		fmt.Println("调用生成新区块的功能")
	case GETLASTBLOCK:
		fmt.Println("获取最新区块功能")
	case GETALLBLOCK:
		fmt.Println("获取所有区块功能")
	case GETBLOCKCOUNT:
		fmt.Println("获取区块的数量")
	case HELP:
		fmt.Println("获取使用说明")
	default:
		fmt.Println("go run main.go : Unknown subcommand.")
		fmt.Println("Use go run main.go help for more information.")
	}
	//3.根据不同的命令，调用blockChain的对应功能
	//4.根据调用的结果，将功能调用的结果信息输出到控制台，提供给用户


}

/**
 * 该方法用于向控制台输出项目的使用说明
 */
func (client *Client) Help() {
	fmt.Println("----------Welcome to XianfengChain Project----------")
	fmt.Println()
	fmt.Println("使用说明: ")
	fmt.Println("\tgo run main.go command [arguments]")
	fmt.Println()
	fmt.Println("当前可以使用的功能: ")
	fmt.Println()
	fmt.Println("\tcreatechain   \t创建一条区块链")
	fmt.Println("\tgenerategensis\t生成创世区块,接收一个参数gensis表示创世区块的数据")
	fmt.Println()
	fmt.Println("更多帮助请使用: go run main.go help")

}
