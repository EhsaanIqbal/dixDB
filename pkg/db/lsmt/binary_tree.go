package lsmt

import (
	"fmt"
)

type KVPair struct {
	Key   string
	Value string
}

type Node struct {
	KV    KVPair
	Left  *Node
	Right *Node
	Size  int
}

func NewBinaryTree(kv []KVPair) *Node {
	size := len(kv)
	if size == 0 {
		return nil
	}

	root := &Node{
		KV:   kv[size/2],
		Left: NewBinaryTree(kv[0 : size/2]),
		Size: size,
	}

	if rightIndex := size/2 + 1; rightIndex < size {
		root.Right = NewBinaryTree(kv[rightIndex:size])
	}

	return root
}

func Insert(n **Node, kv KVPair) {
	if *n == nil {
		*n = &Node{KV: kv, Size: 1}
	} else if kv.Key < (*n).KV.Key {
		Insert(&(*n).Left, kv)
		(*n).Size++
	} else if kv.Key > (*n).KV.Key {
		Insert(&(*n).Right, kv)
		(*n).Size++
	} else {
		(*n).KV.Value = kv.Value
	}
}

func Search(n *Node, key string) (KVPair, error) {
	if n == nil {
		return KVPair{}, fmt.Errorf("key %s does not exist", key)
	} else if n.KV.Key == key {
		return n.KV, nil
	} else if key <= n.KV.Key {
		return Search(n.Left, key)
	} else {
		return Search(n.Right, key)
	}
}
