package design

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Consistent struct {
	hashSortedNodes  []uint32          // 排序的hash虚拟节点
	circle           map[uint32]string // hash对应的节点
	nodes            map[string]bool   // 已绑定的节点
	sync.RWMutex                       // 读写锁
	virtualNodeCount int               // 虚拟节点数量
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Add(node string, virtualNodeCount int) error {
	if node == "" {
		return nil
	}
	c.RLock()
	defer c.RUnlock()

	if c.circle == nil {
		c.circle = make(map[uint32]string)
	}

	if c.nodes == nil {
		c.nodes = make(map[string]bool)
	}

	if _, ok := c.nodes[node]; ok {
		return nil
	}

	// 设置真实节点
	c.nodes[node] = true

	// 增加虚拟节点
	for i := 0; i < virtualNodeCount; i++ {
		hash := c.hashKey(node + strconv.Itoa(i))
		c.circle[hash] = node
		c.hashSortedNodes = append(c.hashSortedNodes, hash)
	}

	// 虚拟节点排序
	sort.Slice(c.hashSortedNodes, func(i, j int) bool {
		return c.hashSortedNodes[i] < c.hashSortedNodes[j]
	})

	return nil
}

func (c *Consistent) GetNode(key string) string {
	c.RLock()
	defer c.RUnlock()
	hash := c.hashKey(key)
	i := c.getPosition(hash)

	return c.circle[c.hashSortedNodes[i]]
}

func (c *Consistent) getPosition(hash uint32) int {
	// 找到第一个大于等于hash的虚拟节点
	i := sort.Search(len(c.hashSortedNodes), func(i int) bool {
		return c.hashSortedNodes[i] >= hash
	})

	if i < len(c.hashSortedNodes) {
		// 如果找到，直接返回
		if i == len(c.hashSortedNodes)-1 {
			// 如果是最后一个节点，返回第一个节点
			return 0
		} else {
			return i
		}
	} else {
		// 如果没有找到，说明hash值比所有节点的hash值都大，返回第一个节点
		return len(c.hashSortedNodes) - 1
	}
}
