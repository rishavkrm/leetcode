	type ListNode2 struct {
		Val  int
		Next *ListNode2
		Prev *ListNode2
	}

	type LRUCache struct {
		Dict     map[int]int
		Address  map[int]*ListNode2
		Head     *ListNode2
		Tail     *ListNode2
		CurrSize int
		MaxSize  int
	}

	func Constructor(capacity int) LRUCache {
		return LRUCache{
			MaxSize: capacity,
			Dict:    make(map[int]int),
			Address: make(map[int]*ListNode2),
		}
	}

	func (this *LRUCache) Get(key int) int {
		val, exists := this.Dict[key]
		if exists {
			// find the address
			node := this.Address[key]
			if node.Prev == nil { // already head
				return val
			} else {
				if node.Next == nil {
					this.Tail = node.Prev
				}
				node.Prev.Next = node.Next
                if node.Next != nil{
                    node.Next.Prev = node.Prev
                }
				node.Next = this.Head
				node.Next.Prev = node.Prev
				node.Prev = nil
				this.Head.Prev = node
				this.Head = node
			}
			return val
		}
		return -1
	}

	func (this *LRUCache) Put(key int, value int) {
		_, exists := this.Dict[key]
		this.Dict[key] = value
		if exists {
			// find the address
			node := this.Address[key]
			if node.Prev == nil { // already head
				return
			} else {
				if node.Next == nil {
					this.Tail = node.Prev
				}
				node.Prev.Next = node.Next
                if node.Next != nil{
                    node.Next.Prev = node.Prev
                }
				node.Next = this.Head
				node.Prev = nil
				this.Head.Prev = node
				this.Head = node
			}
			return
		} else {
			if this.MaxSize == this.CurrSize {
				currTail := this.Tail
				if currTail != nil {
					this.Tail = this.Tail.Prev
					if this.Tail != nil {
						this.Tail.Next = nil
					}
					delete(this.Dict, currTail.Val)
				}

			} else {
				this.CurrSize += 1
			}
			curr := &ListNode2{
				Val:  key,
				Prev: nil,
				Next: this.Head,
			}
			if this.Head != nil {
				this.Head.Prev = curr
			}
			this.Address[key] = curr
			this.Head = curr
			if this.Tail == nil {
				this.Tail = curr
			}

		}
		return
	}

	/**
	 * Your LRUCache object will be instantiated and called as such:
	 * obj := Constructor(capacity);
	 * param_1 := obj.Get(key);
	 * obj.Put(key,value);
	 */