package main

import (
	"github.com/boltdb/bolt"
	"fmt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//type Blockchain struct {
//	blocks []*Block
//}

//func (bc *Blockchain) AddBlock(data string) {
//	prevBlock := bc.blocks[len(bc.blocks)-1]
//	newBlock := NewBlock(data, prevBlock.Hash)
//	bc.blocks = append(bc.blocks, newBlock)
//}

//func NewBlockchain() *Blockchain {
//	return &Blockchain{[]*Block{NewGenesisBlock()}}
//}

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				fmt.Errorf(err.Error())
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		fmt.Errorf(err.Error())
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			fmt.Errorf(err.Error())
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}
