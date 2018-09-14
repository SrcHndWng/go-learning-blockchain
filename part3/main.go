package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

var bc *Blockchain

func main() {
	defer func() {
		if bc != nil {
			bc.db.Close()
		}
	}()

	app := cli.NewApp()
	app.Name = "go-learning-blockchain"
	app.Description = "implemention to learn blockchain"
	app.Version = "1.3.1"

	app.Commands = []cli.Command{
		{
			Name:      "addblock",
			ShortName: "a",
			Usage:     "addblock BLOCK_DATA : add a block to the blockchain",
			ArgsUsage: "data string to add to blockchain",
			Before:    blockchain,
			Action:    addBlock,
		},
		{
			Name:      "printchain",
			ShortName: "p",
			Usage:     "printchain : print all the blocks of the blockchain",
			Before:    blockchain,
			Action:    printChain,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func blockchain(c *cli.Context) error {
	bc = NewBlockchain()
	return nil
}

func addBlock(c *cli.Context) error {
	data := c.Args().First()
	if data == "" {
		fmt.Println("addblock must take a BLOCK_DATA arg.")
		return nil
	}

	bc.AddBlock(data)
	fmt.Println("Success!")
	return nil
}

func printChain(c *cli.Context) error {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return nil
}
