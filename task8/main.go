package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"example.com/Goeth/count"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。
发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。
*/
//查询区块：实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
func queryBlock(client *ethclient.Client, blockNumber *big.Int) {
	block, _ := client.BlockByNumber(context.Background(), blockNumber) // 获取指定区块号的区块信息
	fmt.Println("Block Hash:", block.Hash().Hex())                      // 输出区块的哈希值
	fmt.Println("Block Timestamp:", block.Time())                       // 输出区块的时间戳
	fmt.Println("Block Transaction Count:", len(block.Transactions()))  // 输出区块的交易数量
}

// ETH转账
func transferEth(client *ethclient.Client, to string, amount *big.Int) {
	privateKey, err := crypto.HexToECDSA("308671caf1b5fea4749c0535bb35dec8726ec742625c8c7a4f3ae6d7e759f5b7") //加载私钥
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()                   //获取公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //断言公钥类型
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //获取账户地址
	fmt.Println("账户地址", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) //获取账户的未使用交易数量
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce:", nonce)

	value := big.NewInt(200000000000000000) // in wei (1 eth)
	//获取gas限制
	gasLimit := uint64(21000)                                     // in units
	gasPrice, err := client.SuggestGasPrice(context.Background()) //获取gas价格
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice:", gasPrice)

	//获取交易接收者地址
	toAddress := common.HexToAddress(to) //交易接收者地址
	var data []byte
	//创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data) //创建交易

	//获取网络ID
	chainID, err := client.NetworkID(context.Background()) //获取网络ID
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID:", chainID)
	//对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey) //对交易进行签名
	if err != nil {
		log.Fatal(err)
	}

	//发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//打印交易哈希
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

// 调用合约Count
func callContract(client *ethclient.Client) {
	countContract, err := count.NewCount(common.HexToAddress("0xa9abf621798bd6d9f9a9d0ec62cd5df7aeb6b106"), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := countContract.CountNmb(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

}

// 监听CountNumber事件来获取最新的id值
func listenContractEvents(client *ethclient.Client, contractAbi abi.ABI) {
	fmt.Println("\n=== 监听CountNumber事件获取ID ===")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xa9abf621798bd6d9f9a9d0ec62cd5df7aeb6b106")},
		FromBlock: big.NewInt(0), // 从最新区块开始
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("过滤日志失败:", err)
	}

	for _, vLog := range logs {
		// 解析CountNumber事件
		event := struct {
			Id *big.Int
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "CountNumber", vLog.Data)
		if err != nil {
			log.Printf("解析事件失败: %v", err)
			continue
		}
		fmt.Printf("事件中的ID: %s, 区块: %d, 交易: %s\n",
			event.Id.String(), vLog.BlockNumber, vLog.TxHash.Hex())
	}
}

// 实时监听CountNumber事件
func realTimeListenCountNumber(client *ethclient.Client) {
	fmt.Println("\n=== 开始实时监听CountNumber事件 ===")

	// 合约地址
	contractAddress := common.HexToAddress("0xa9abf621798bd6d9f9a9d0ec62cd5df7aeb6b106")

	// 解析合约ABI
	contractAbi, err := abi.JSON(strings.NewReader(count.CountMetaData.ABI))
	if err != nil {
		log.Fatal("解析ABI失败:", err)
	}

	// 创建过滤查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 创建日志通道
	logs := make(chan types.Log)

	// 订阅日志
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("订阅日志失败:", err)
	}

	fmt.Println("监听已启动，等待CountNumber事件...")

	// 无限循环监听事件
	for {
		select {
		case err := <-sub.Err():
			log.Fatal("订阅出错:", err)
		case vLog := <-logs:
			// 检查是否是CountNumber事件
			if len(vLog.Topics) > 0 {
				fmt.Printf("收到事件 - 区块: %d, 交易: %s\n", vLog.BlockNumber, vLog.TxHash.Hex())

				// 解析CountNumber事件
				event := struct {
					Id *big.Int
				}{}

				err := contractAbi.UnpackIntoInterface(&event, "CountNumber", vLog.Data)
				if err != nil {
					log.Printf("解析事件失败: %v", err)
					continue
				}

				fmt.Printf("🎉 CountNumber事件触发! ID值: %s\n", event.Id.String())
			}
		}
	}
}

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b2b6067c7ee146698bcc5861b07825a6") // 连接到 Sepolia 测试网络
	if err != nil {
		log.Fatal(err)
	}

	// blockNumber := big.NewInt(5671744) // 指定要查询的区块号
	// queryBlock(client, blockNumber)    // 调用查询区块函数

	//transferEth(client, "0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455", big.NewInt(200000000000000000))//转账0.2ETH
	callContract(client) //调用合约Count

	// 使用正确的Count合约ABI（包含CountNumber事件）
	// contractAbi, err := abi.JSON(strings.NewReader(count.CountMetaData.ABI))
	// if err != nil {
	// 	log.Fatal("解析ABI失败:", err)
	// }

	// 调用已存在的函数
	// listenContractEvents(client, contractAbi)
	realTimeListenCountNumber(client) //实时监听CountNumber事件
}
