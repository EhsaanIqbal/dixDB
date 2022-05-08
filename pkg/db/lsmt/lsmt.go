package lsmt

import "sync"

type LSMT struct {
	rwm            sync.RWMutex
	bt             *Node
	treeInFlush    *Node
	flushThreshold int
	drwm           sync.RWMutex
	files          []File
}
