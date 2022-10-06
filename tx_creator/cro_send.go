// package token

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/rlp"
// )

// func cro_send() {
//     client, err := ethclient.Dial("http://localhost:8545")
//     if err != nil {
//         log.Fatal(err)
//     }

//     privateKey, err := crypto.HexToECDSA("xxxxx")
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

//     value := big.NewInt(100000000) // in wei (1 basecro)
//     gasLimit := uint64(21000)                // in units
//     // gasPrice, err := client.SuggestGasPrice(context.Background())
//     // if err != nil {
//     //     log.Fatal(err)
//     // }

//     toAddress := common.HexToAddress("0x378c50D9264C63F3F92B806d4ee56E9D86FfB3Ec")
//     var data []byte
//     tx := types.NewTransaction(nonce, toAddress, value, gasLimit, big.NewInt(7), data)
//     // tx := types.NewTx(&types.DynamicFeeTx{
//     //     ChainID: big.NewInt(777),
//     //     Nonce: nonce,
//     //     GasFeeCap: big.NewInt(91097072255),
//     //     GasTipCap: big.NewInt(1000000000),
//     //     Gas: 21000,
//     //     To: &toAddress,
//     //     Value: value,
//     //     Data: data,
//     // })

//     chainID, err := client.NetworkID(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }
//     signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//     if err != nil {
//         log.Fatal(err)
//     }

//     raw, err := rlp.EncodeToBytes(signedTx)

// 	fmt.Printf("err:%v hash:%x\n0x%x", err, signedTx.Hash().Hex(), raw)
// }