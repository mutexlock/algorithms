package _146_lru

//type ListNode struct {
//	Key  int
//	Val  int
//	Next *ListNode
//}
//type LRUCache struct {
//	Capacity int
//	Len      int
//	Head     *ListNode
//	Tail     *ListNode
//	ValueMap map[int]*ListNode
//}
//
//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{}
//	lru.Capacity = capacity
//	lru.ValueMap = make(map[int]*ListNode, capacity)
//	lru.Head = nil
//	return lru
//}
//
//func (this *LRUCache) Get(key int) int {
//	//放到队尾
//    node := this.ValueMap[key]
//
//    node.
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	//放到队尾
//	node := new(ListNode)
//	node.Val = value
//	node.Key = key
//	if this.Head == nil {
//		this.Head = node
//	}
//
//	this.Tail = node
//	if this.Len == this.Capacity {
//		delete(this.ValueMap, this.Head.Key)
//		this.Head = this.Head.Next
//	} else {
//		this.Len++
//	}
//	this.ValueMap[key] = node
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity)}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		// 存在的化需要把这个node移动到header.不需要关心是否是第一个。有哨兵机制
		this.remove(node)
		this.putHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		//　存在节点.更新数据。移动到head
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		// 不存在
		// 如果容器已经满了，需要删除链表尾部,从map中删除
		if len(this.m) >= this.cap {
			deletekey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, deletekey)
		}

		//创建新的节点,并放在head,添加到map中
		newNode := Node{key: key, value: value}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) putHead(node *Node) {
	originNext := this.header.next
	this.header.next = node
	node.next = originNext

	originNext.pre = node
	node.pre = this.header
}
