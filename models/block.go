package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int    `json:"index"`     // 这块链在整个链中的位置
	Timestamp string `json:"timestamp"` // 块生成的时间
	BPM       int    `json:"bpm"`       // 每分钟的心跳数
	Hash      string `json:"hash"`      // 块通过哈希算法生成的散列值
	PrevHash  string `json:"prevHash"`  // 表示前一个块的散列值
}

var BlockChain []Block

// 计算给定数据的SHA256散列值:
func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// 生成块函数
func GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

// 校验块，判断一个块是否被篡改
func IsBlockValid(newBlock, oldBlock Block) bool {
	if (oldBlock.Index + 1) != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// 一般来说更长的链标书的数据状态是更新的，我们需要将本地过期的链切换成新的链
func ReplaceBlock(newBlocks []Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}
