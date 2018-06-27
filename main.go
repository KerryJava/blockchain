package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	//for _, block := range bc.blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	fmt.Println()
	//}
	iterator := bc.Iterator()

	count := 0

	for {

		block := iterator.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
			return
		}

		count++

		if count >= 10 {
			fmt.Println("count exceed")
			break
		}
	}
}
