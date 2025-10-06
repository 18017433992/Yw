package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	blockNumber := big.NewInt(5671744)//区块数

// 	header, err := client.HeaderByNumber(context.Background(), blockNumber)
// 	fmt.Println(header.Number.Uint64())     // 区块数 5671744
// 	fmt.Println(header.Time)                // 区块时间戳 1712798400
// 	fmt.Println(header.Difficulty.Uint64()) // 0
// 	fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	block, err := client.BlockByNumber(context.Background(), blockNumber)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(block.Number().Uint64())     // 5671744
// 	fmt.Println(block.Time())                // 1712798400
// 	fmt.Println(block.Difficulty().Uint64()) // 0
// 	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
// 	fmt.Println(len(block.Transactions()))   // 70
// 	count, err := client.TransactionCount(context.Background(), block.Hash())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		fmt.Println(count) // 70
//	}
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	chainID, err := client.ChainID(context.Background()) //获取chainID
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	blockNumber := big.NewInt(5671744)                                    //指定区块数
// 	block, err := client.BlockByNumber(context.Background(), blockNumber) //获取指定区块
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, tx := range block.Transactions() { //遍历区块中的交易
// 		fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
// 		fmt.Println(tx.Value().String())    // 100000000000000000
// 		fmt.Println(tx.Gas())               // 21000
// 		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
// 		fmt.Println(tx.Nonce())             // 245132
// 		fmt.Println(tx.Data())              // []
// 		fmt.Println(tx.To().Hex())          // 返回交易接收者地址

// 		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
// 			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
// 		} else {
// 			log.Fatal(err)
// 		}

// 		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash()) //获取交易回执
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(receipt.Status) // 1
// 		fmt.Println(receipt.Logs)   // []
// 		break
// 	}

// 	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5") //获取指定区块的哈希
// 	count, err := client.TransactionCount(context.Background(), blockHash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for idx := uint(0); idx < count; idx++ {
// 		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx) //获取指定区块的指定交易
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
// 		break
// 	}

//		txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5") //获取指定交易的哈希
//		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)                     //获取指定交易
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(isPending)
//		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
//	}
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	blockNumber := big.NewInt(5671744)
// 	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5") //指定区块的哈希

// 	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false)) //获取指定区块的回执
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))) //获取指定区块的回执
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true

// 	for _, receipt := range receiptByHash {
// 		fmt.Println(receipt.Status)                // 1
// 		fmt.Println(receipt.Logs)                  // []
// 		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
// 		fmt.Println(receipt.TransactionIndex)      // 0
// 		fmt.Println(receipt.ContractAddress.Hex()) // 获取合约地址
// 		break
// 	}

// 	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5") //获取指定交易的回执
// 	receipt, err := client.TransactionReceipt(context.Background(), txHash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(receipt.Status)                // 1
// 	fmt.Println(receipt.Logs)                  // []
// 	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
// 	fmt.Println(receipt.TransactionIndex)      // 0
// 	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
// }

// func main() {
// 	privateKey, err := crypto.GenerateKey() // 生成私钥  0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKeyBytes := crypto.FromECDSA(privateKey)    // 私钥转字节
// 	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])   // 去掉'0x',私钥字节
// 	publicKey := privateKey.Public()                   // 获取公钥
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //断言公钥类型
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

//		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
//		fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
//		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()        // 获取地址
//		fmt.Println(address)
//		hash := sha3.NewLegacyKeccak256()                      // 创建Keccak256哈希对象
//		hash.Write(publicKeyBytes[1:])                         // 写入公钥字节
//		fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:])) // 输出完整32位哈希值
//		fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))        // 原长32位，截去12位，保留后20位 // 地址
//	}
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de") //加载私钥
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()                   //获取公钥
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //断言公钥类型
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //获取账户地址
// 	fmt.Println("账户地址", fromAddress)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) //获取账户的未使用交易数量
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("nonce:", nonce)

// 	value := big.NewInt(200000000000000000) // in wei (1 eth)
// 	//获取gas限制
// 	gasLimit := uint64(21000)                                     // in units
// 	gasPrice, err := client.SuggestGasPrice(context.Background()) //获取gas价格
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("gasPrice:", gasPrice)

// 	//获取交易接收者地址
// 	toAddress := common.HexToAddress("0xe7c3c44e07bdD94ab87dfA8ca3456671C3D528A4") //交易接收者地址
// 	var data []byte
// 	//创建交易
// 	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data) //创建交易

// 	//获取网络ID
// 	chainID, err := client.NetworkID(context.Background()) //获取网络ID
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("chainID:", chainID)
// 	//对交易进行签名
// 	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey) //对交易进行签名
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//发送交易
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//打印交易哈希
// 	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
// }
//代币转账
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	value := big.NewInt(0) // in wei (0 eth)
// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")    // 接收者地址
// 	tokenAddress := common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238") //合约地址

// 	transferFnSignature := []byte("transfer(address,uint256)") //函数名

// 	hash := sha3.NewLegacyKeccak256() //创建Keccak256哈希对象

// 	hash.Write(transferFnSignature) //写入函数名

// 	methodID := hash.Sum(nil)[:4] //获取函数ID

// 	fmt.Println(hexutil.Encode(methodID))                       // 0xa9059cbb
// 	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32) //获取接收者地址
// 	fmt.Println(hexutil.Encode(paddedAddress))                  // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
// 	amount := new(big.Int)
// 	amount.SetString("11000000", 10)                        // 11 tokens
// 	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32) //获取金额
// 	fmt.Println(hexutil.Encode(paddedAmount))               // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
// 	var data []byte
// 	data = append(data, methodID...)
// 	data = append(data, paddedAddress...)
// 	data = append(data, paddedAmount...)

// 	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		To:   &toAddress,
// 		Data: data,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(gasLimit) // 23256
// 	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
//	}
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	account := common.HexToAddress("0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455")
// 	balance, err := client.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//将WEI转化为eth打印
// 	ETHBalance := new(big.Float).SetInt(balance)
// 	ETHBalance = ETHBalance.Quo(ETHBalance, big.NewFloat(math.Pow10(18)))
// 	fmt.Println("balance", balance)
// 	fmt.Println("ETHBalance", ETHBalance)

//		blockNumber := big.NewInt(5532993)
//		balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("balanceAt", balanceAt) // 25729324269165216042\n
//		fbalance := new(big.Float)
//		fbalance.SetString(balanceAt.String())
//		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
//		fmt.Println("ethValue", ethValue) //0
//		pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("pendingBalance", pendingBalance)
//	}
// func main() { //查询代币余额
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// 使用包含完整元数据的代币合约 (USDC on Sepolia)
// 	tokenAddress := common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238")
// 	instance, err := token.NewErc20(tokenAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	address := common.HexToAddress("0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455")
// 	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 现在可以直接调用 metadata 方法
// 	name, err := instance.Name(&bind.CallOpts{})
// 	if err != nil {
// 		log.Printf("Failed to get name: %v", err)
// 		name = "Unknown"
// 	}

// 	symbol, err := instance.Symbol(&bind.CallOpts{})
// 	if err != nil {
// 		log.Printf("Failed to get symbol: %v", err)
// 		symbol = "Unknown"
// 	}

//		decimals, err := instance.Decimals(&bind.CallOpts{})
//		if err != nil {
//			log.Printf("Failed to get decimals: %v", err)
//			decimals = 18 // Default to 18 if not available
//		}
//		fmt.Printf("name: %s\n", name)         // "name: Golem Network"
//		fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
//		fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
//		fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
//		fbal := new(big.Float)
//		fbal.SetString(bal.String())
//		value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
//		fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
//	}

// getBlockWithRetry 尝试获取区块数据，带重试机制和详细日志
// func getBlockWithRetry(client *ethclient.Client, hash common.Hash, maxRetries int) (*types.Block, error) {
// 	hashStr := hash.Hex()
// 	log.Printf("Attempting to fetch block: %s", hashStr)

// 	for i := 0; i < maxRetries; i++ {
// 		block, err := client.BlockByHash(context.Background(), hash)
// 		if err == nil && block != nil {
// 			log.Printf("Successfully fetched block %s on attempt %d", hashStr, i+1)
// 			return block, nil
// 		}

// 		if err != nil {
// 			log.Printf("Attempt %d/%d failed with error: %v", i+1, maxRetries, err)
// 		} else {
// 			log.Printf("Attempt %d/%d failed: block returned nil", i+1, maxRetries)
// 		}

// 		if i < maxRetries-1 {
// 			waitTime := time.Millisecond * time.Duration(500*(i+1))
// 			log.Printf("Waiting %v before retry...", waitTime)
// 			time.Sleep(waitTime)
// 		}
// 	}

// 	log.Printf("Failed to get block %s after %d attempts", hashStr, maxRetries)
// 	return nil, fmt.Errorf("failed to get block %s after %d attempts", hashStr, maxRetries)
// }

// func main() {
// 	// 使用 WebSocket 端点以支持订阅功能
// 	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	headers := make(chan *types.Header)
// 	sub, err := client.SubscribeNewHead(context.Background(), headers)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case header := <-headers:
// 			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
// 			time.Sleep(time.Second * 300)
// 			block, err := client.BlockByHash(context.Background(), header.Hash())
// 			if err != nil {
// 				log.Fatal(err)
// 			}

//				fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
//				fmt.Println(block.Number().Uint64())   // 3477413
//				fmt.Println(block.Time())              // 1529525947
//				fmt.Println(block.Nonce())             // 130524141876765836
//				fmt.Println(len(block.Transactions())) // 7
//			}
//		}
//	}
//通过abigent部署合约
// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// privateKey, err := crypto.GenerateKey()
// 	// privateKeyBytes := crypto.FromECDSA(privateKey)
// 	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
// 	// fmt.Println("Private Key:", privateKeyHex)
// 	privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	chainId, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // in wei
// 	auth.GasLimit = uint64(300000) // in units
// 	auth.GasPrice = gasPrice

// 	input := "1.0"
// 	address, tx, instance, err := store.DeployStore(auth, client, input)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(address.Hex())
// 	fmt.Println(tx.Hash().Hex())

//		_ = instance
//	}
//
// 通过ethclient部署
// const (
// 	// store合约的字节码
// 	contractBytecode = "608060405234801561000f575f80fd5b5060405161087538038061087583398181016040528101906100319190610193565b805f908161003f91906103e7565b50506104b6565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f90565b61032c610320565b6103378184846102fb565b505050565b5b8181101561035a5761034f5f82610324565b60018101905061033d565b5050565b601f82111561039f5761037081610241565b61037984610253565b81016020851015610388578190505b61039c61039485610253565b83018261033c565b50505b505050565b5f82821c905092915050565b5f6103bf5f19846008026103a4565b1980831691505092915050565b5f6103d783836103b0565b9150826002028217905092915050565b6103f0826101da565b67ffffffffffffffff8111156104095761040861006f565b5b6104138254610211565b61041e82828561035e565b5f60209050601f83116001811461044f575f841561043d578287015190505b61044785826103cc565b8655506104ae565b601f19841661045d86610241565b5f5b828110156104845784890151825560018201915060208501945060208101905061045f565b868310156104a1578489015161049d601f8916826103b0565b8355505b6001600288020188555050505b505050505050565b6103b2806104c35f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f80fd5b61005d600480360381019061005891906101d7565b6100ad565b60405161006a9190610211565b60405180910390f35b61007b6100c2565b604051610088919061029a565b60405180910390f35b6100ab60048036038101906100a691906102ba565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610325565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051610194929190610355565b60405180910390a15050565b5f80fd5b5f819050919050565b6101b6816101a4565b81146101c0575f80fd5b50565b5f813590506101d1816101ad565b92915050565b5f602082840312156101ec576101eb6101a0565b5b5f6101f9848285016101c3565b91505092915050565b61020b816101a4565b82525050565b5f6020820190506102245f830184610202565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026c8261022a565b6102768185610234565b9350610286818560208601610244565b61028f81610252565b840191505092915050565b5f6020820190508181035f8301526102b28184610262565b905092915050565b5f80604083850312156102d0576102cf6101a0565b5b5f6102dd858286016101c3565b92505060206102ee858286016101c3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033c57607f821691505b60208210810361034f5761034e6102f8565b5b50919050565b5f6040820190506103685f830185610202565b6103756020830184610202565b939250505056fea26469706673582212205aae308f77654b000c9d222eff2d9f2bd2ac18d990b10774842e4309d4e3e15664736f6c634300081a0033"
// )

// func main() {
// 	// 连接到以太坊网络（这里使用 Goerli 测试网络作为示例）
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 创建私钥（在实际应用中，您应该使用更安全的方式来管理私钥）
// 	privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("error casting public key to ECDSA")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	// 获取nonce
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 获取建议的gas价格
// 	gasPrice, err := client.SuggestGasPrice(context.Background()) //1000028
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 解码合约字节码
// 	data, err := hex.DecodeString(contractBytecode)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 创建交易
// 	tx := types.NewContractCreation(nonce, big.NewInt(0), 10000000, gasPrice, data)

// 	// 签名交易
// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 发送交易
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

// 	// 等待交易被挖矿
// 	receipt, err := waitForReceipt(client, signedTx.Hash())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex())
// }

// func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
// 	for {
// 		receipt, err := client.TransactionReceipt(context.Background(), txHash)
// 		if err == nil {
// 			return receipt, nil
// 		}
// 		if err != ethereum.NotFound {
// 			return nil, err
// 		}
// 		// 等待一段时间后再次查询
// 		time.Sleep(1 * time.Second)
// 	}
// }

// const (
// 	contractAddr = "0x534862Ff0152D7A3e86b9a9c41a7A5558723e064"
// )

// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_ = storeContract
// }

// const (
// 	contractAddr = "0x534862Ff0152D7A3e86b9a9c41a7A5558723e064"
// )

// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	storeContract, err := store.NewStore(common.HexToAddress("0x534862Ff0152D7A3e86b9a9c41a7A5558723e064"), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// privateKey, err := crypto.HexToECDSA("317c48b81e16f3a4f8d9f6382d497dcab35858b8751796b6f4a1bda2135945de")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	var key [32]byte
// 	// var value [32]byte

// 	copy(key[:], []byte("demo_save_key"))
// 	// copy(value[:], []byte("demo_save_value11111"))

// 	// opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// tx, err := storeContract.SetItem(opt, key, value)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println("tx hash:", tx.Hash().Hex())

//		callOpt := &bind.CallOpts{Context: context.Background()}
//		valueInContract, err := storeContract.Items(callOpt, key)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//将valueInContract转化成字符串
//		valueInContractStr := string(valueInContract[:])
//		fmt.Println("is value saving in contract equals to origin value:", valueInContractStr)
//	}

// var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b2b6067c7ee146698bcc5861b07825a6")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	contractAddress := common.HexToAddress("0x534862Ff0152D7A3e86b9a9c41a7A5558723e064")
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(6920583),
// 		// ToBlock:   big.NewInt(2394201),
// 		Addresses: []common.Address{
// 			contractAddress,
// 		},
// 		// Topics: [][]common.Hash{
// 		//  {},
// 		//  {},
// 		// },
// 	}

// 	logs, err := client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, vLog := range logs {
// 		fmt.Println(vLog.BlockHash.Hex())
// 		fmt.Println(vLog.BlockNumber)
// 		fmt.Println(vLog.TxHash.Hex())
// 		event := struct {
// 			Key   [32]byte
// 			Value [32]byte
// 		}{}
// 		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(common.Bytes2Hex(event.Key[:]))
// 		fmt.Println(common.Bytes2Hex(event.Value[:]))
// 		var topics []string
// 		for i := range vLog.Topics {
// 			topics = append(topics, vLog.Topics[i].Hex())
// 		}

// 		fmt.Println("topics[0]=", topics[0])
// 		if len(topics) > 1 {
// 			fmt.Println("indexed topics:", topics[1:])
// 		}
// 	}

// 	eventSignature := []byte("ItemSet(bytes32,bytes32)")
// 	hash := crypto.Keccak256Hash(eventSignature)
// 	fmt.Println("signature topics=", hash.Hex())
// }

var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func maibase() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b2b6067c7ee146698bcc5861b07825a6")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x534862Ff0152D7A3e86b9a9c41a7A5558723e064")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
