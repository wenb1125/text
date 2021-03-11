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
		addBlock := flag.NewFlagSet(ADDNEWBLOCK,flag.ExitOnError)
		data := addBlock.String("data","","区块存储的自定义内容")
		addBlock.Parse(os.Args[2:])

		//args := os.Args[2:]
		////1.从参数中取出所有以-开头的参数选项
		////2.准备一个当前命令支持的所有参数的切片

		err := client.Chain.AddNewBlock([]byte(*data))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("恭喜，已成功创建新区块，并存储到文件中")
	case GETLASTBLOCK:
		set := os.Args[2:]
		if len(set) > 0 {
			fmt.Println("命令错误...")
			return
		}
		last := client.Chain.GetLastBlock()
		hashBig := new(big.Int)
		hashBig.SetBytes(last.Hash[:])
		if hashBig.Cmp(big.NewInt(0)) > 0 {
			fmt.Println("查询到最新区块")
			fmt.Println("最新区块高度：",last.Height)
			fmt.Println("最新区块内容：",string(last.Data))
			fmt.Printf("最新区块hash：%x\n",last.Hash)
			fmt.Printf("前一个区块hash：%x\n",last.PreHash)
			return
		}
		fmt.Println("抱歉，当前暂无最新区块")
		fmt.Println("请使用go run main.go generategensis生成创世区块")
	case GETALLBLOCK:
		if len(os.Args[2:]) > 0 {
			fmt.Println("抱歉，getallblock不接收参数...")
			return
		}

		allBlock, err := client.Chain.GetAllBlocks()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("成功获取到所有区块...")
		for _, block := range allBlock {
			fmt.Printf("区块高度: %d\nHash: %x\n数据: %s\n-------------------------\n",block.Height,block.Hash,block.Data)
		}
	case GETBLOCKCOUNT:
		blocks, err := client.Chain.GetAllBlocks()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("---查询成功,当前共有%d个区块---\n",len(blocks))
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
