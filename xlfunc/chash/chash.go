/**
一致性HASH算法，可以用于解决分库分表分配问题
*/

package chash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// HashValues hash values sorted
type HashValues []uint32

func (hashValues HashValues) Len() int { return len(hashValues) }
func (hashValues HashValues) Swap(i, j int) {
	hashValues[i], hashValues[j] = hashValues[j], hashValues[i]
}
func (hashValues HashValues) Less(i, j int) bool { return hashValues[i] < hashValues[j] }

// HashNode node information
type HashNode struct {
	// name of node
	Name string
	// other attributes of node
	Data map[string]interface{}
}

type CHash struct {
	// sorted hash value array
	HashValues HashValues
	// node array, storing node related information
	HashNodes map[uint32]*HashNode
	// number of virtual nodes
	VirtualNumber int
}

func NewCHash() *CHash {
	cHash := new(CHash)
	cHash.HashValues = make([]uint32, 0)
	cHash.HashNodes = make(map[uint32]*HashNode, 0)
	cHash.VirtualNumber = 100
	return cHash
}

// InitNodes initialize nodes
func (cHash *CHash) InitNodes(nodes []HashNode) {
	for _, node := range nodes {
		cHash.AddNode(node)
	}
}

// ClearNodes delete all nodes
func (cHash *CHash) ClearNodes(nodes []HashNode) {
	for _, node := range nodes {
		cHash.DelNode(node)
	}
}

// GetHashValue get hash value
func (cHash *CHash) GetHashValue(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// AddNode add node
func (cHash *CHash) AddNode(node HashNode) {
	// expand each node to the virtual number
	for i := 0; i < cHash.VirtualNumber; i++ {
		key := fmt.Sprintf("%s_%d", node.Name, i)
		hashCode := cHash.GetHashValue(key)
		cHash.HashNodes[hashCode] = &node
		cHash.HashValues = append(cHash.HashValues, hashCode)
		sort.Sort(cHash.HashValues)
	}
}

// DelNode delete node
func (cHash *CHash) DelNode(node HashNode) {
	// delete all the virtual number of node expansions
	for i := 0; i < cHash.VirtualNumber; i++ {
		name := fmt.Sprintf("%s_%d", node.Name, i)
		hashCode := cHash.GetHashValue(name)
		delete(cHash.HashNodes, hashCode)
		index := cHash.BSearchNodeByKey(hashCode, 0, len(cHash.HashValues)-1)
		if -1 != index {
			cHash.HashValues = append(cHash.HashValues[:index], cHash.HashValues[index+1:]...)
		}
	}
}

// BSearchNodeByKey search node by hash value
func (cHash *CHash) BSearchNodeByKey(key uint32, min int, max int) int {
	mid := (min + max) / 2
	if mid > min {
		if key == cHash.HashValues[mid] {
			return mid
		} else if key > cHash.HashValues[mid] {
			return cHash.BSearchNodeByKey(key, mid, max)
		} else {
			return cHash.BSearchNodeByKey(key, min, mid)
		}
	}
	return -1
}

// Match match the node to belong by name
func (cHash *CHash) Match(name string) *HashNode {
	hashCode := cHash.GetHashValue(name)
	for i := 0; i < len(cHash.HashValues); i++ {
		if hashCode < cHash.HashValues[i] {
			if i == 0 {
				matchIndex := len(cHash.HashValues) - 1
				return cHash.HashNodes[cHash.HashValues[matchIndex]]
			}
			return cHash.HashNodes[cHash.HashValues[i-1]]
		}
	}
	return nil
}

// MatchCode match the node to belong by hash value
func (cHash *CHash) MatchCode(value uint32) *HashNode {
	for i := 0; i < len(cHash.HashValues); i++ {
		if value < cHash.HashValues[i] {
			if i == 0 {
				matchIndex := len(cHash.HashValues) - 1
				return cHash.HashNodes[cHash.HashValues[matchIndex]]
			}
			return cHash.HashNodes[cHash.HashValues[i-1]]
		}
	}
	return nil
}
