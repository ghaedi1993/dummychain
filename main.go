package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting our blockchain")
	var blockchain = CreateBlockchain(5)
	fmt.Printf("the number of blocks are : %d \n", len(blockchain.chain))

	startTime := time.Now()
	blockchain.addBlock("Javad", "Ali", 20)
	blockchain.addBlock("Javad", "Ali", 20)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	lb := blockchain.fetchLastBlock()
	fmt.Printf("the hash after mine is %s \n", lb.hash)
	fmt.Printf("the number of block are : %d \n", len(blockchain.chain))
	fmt.Printf("Elapsed time: %v with mineCount of : %d \n", elapsedTime, lb.mineCount)
}
