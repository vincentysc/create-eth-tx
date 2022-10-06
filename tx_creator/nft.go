// package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/hex"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/rlp"
// 	"golang.org/x/crypto/sha3"
// )

// func main() {
//     client, err := ethclient.Dial("http://localhost:8545")
//     if err != nil {
//         log.Fatal(err)
//     }

//     privateKey, err := crypto.HexToECDSA("e6108e7918bef948e816f3a73fbeffe61d66d21a1dddf40687e5ddd416cb76e5")
//     if err != nil {
//         log.Fatal(err)
//     }

//     publicKey := privateKey.Public()
//     publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//     if !ok {
//         log.Fatal("error casting public key to ECDSA")
//     }

//     fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//     nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//     if err != nil {
//         log.Fatal(err)
//     }

//     value := big.NewInt(0) // in wei (0 eth)
//     gasPrice, err := client.SuggestGasPrice(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }

//     toAddress := common.HexToAddress("0x378c50D9264C63F3F92B806d4ee56E9D86FfB3Ec")
//     nftAddress := common.HexToAddress("0x7860a7112DFCE092E7a23A1d4ECe44b5eCDBad84")

//     transferFnSignature := []byte("mintToken(address,string,uint256,uint256)")
//     hash := sha3.NewLegacyKeccak256()
//     hash.Write(transferFnSignature)
//     methodID := hash.Sum(nil)[:4]
//     fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

//     paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
//     fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d


//     // tokenUri := []byte("https://cronoscan.com/images/logo.svg?v=22.9.2.0")
//     // paddedTokenUri := common.LeftPadBytes(tokenUri, 32)
    

//     tokenId := new(big.Int)
//     tokenId.SetString("2000000000000000000", 10) // 1000 tokens
//     paddedTokenId := common.LeftPadBytes(tokenId.Bytes(), 32)
//     fmt.Println(hexutil.Encode(paddedTokenId)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

//     amount := new(big.Int)
//     amount.SetString("1000000000000000000", 10) // 1000 tokens
//     paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
//     fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000


//     // "0xbe7da533000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001000000000000000000000000378c50d9264c63f3f92b806d4ee56e9d86ffb3ec000000000000000000000000000000000000000000000000000000000000003068747470733a2f2f63726f6e6f7363616e2e636f6d2f696d616765732f6c6f676f2e7376673f763d32322e392e322e3000000000000000000000000000000000"
//     // var data []byte
//     // data = append(data, methodID...)
//     // data = append(data, paddedAddress...)
//     // data = append(data, paddedTokenUri...)
//     // data = append(data, paddedTokenId...)
//     // data = append(data, paddedAmount...)

//     data, err := hex.DecodeString("e376562000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016100000000000000000000000000000000000000000000000000000000000000")
//     if err != nil {
//         panic(err)
//     }
//     fmt.Printf("% x", data)

//     gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
//         To:   &toAddress,
//         Data: data,
//     })
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(gasLimit) // 23256

//     tx := types.NewTransaction(nonce, nftAddress, value, gasLimit, gasPrice, data)

//     chainID, err := client.NetworkID(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }

//     signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//     if err != nil {
//         log.Fatal(err)
//     }

//     raw, err := rlp.EncodeToBytes(signedTx)

//     fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
//     fmt.Printf("\n0x%x", raw)
// }