package medium

// https://leetcode.cn/problems/lru-cache/
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

type LRUCache struct {
	cache map[int]*Node
	root  *Node
	tail  *Node
	cap   int
}

type Node struct {
	k   int
	v   int
	pre *Node
	pro *Node
}

func Constructor(capacity int) LRUCache {
	node := &Node{}
	return LRUCache{
		cache: make(map[int]*Node, capacity),
		root:  node,
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		pre := n.pre
		pro := n.pro
		if pre != this.root {
			if n == this.tail {
				this.tail = pre
			}

			pre.pro = pro
			if pro != nil {
				pro.pre = pre
			}

			rp := this.root.pro
			this.root.pro = n
			n.pre = this.root
			n.pro = rp
			rp.pre = n
		}
		return n.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		pre := n.pre
		pro := n.pro
		if n.pre != this.root {
			if n == this.tail {
				this.tail = pre
			}

			pre.pro = pro
			if pro != nil {
				pro.pre = pre
			}

			rp := this.root.pro
			this.root.pro = n
			n.pre = this.root
			n.pro = rp
			rp.pre = n
		}
		n.v = value
	} else {
		rp := this.root.pro
		node := Node{
			k:   key,
			v:   value,
			pre: this.root,
			pro: rp,
		}
		this.root.pro = &node
		if rp != nil {
			rp.pre = &node
		} else {
			this.tail = &node
		}
		if len(this.cache) == this.cap {
			tail := this.tail
			this.tail = tail.pre
			delete(this.cache, tail.k)
		}
		this.cache[key] = &node
	}
}
