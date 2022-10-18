package tree

// 前缀树节点
type Node struct {
	childer map[rune]*Node
	isEnd bool	
}


// 前缀树根节点
type PrefixTree struct {
	root *Node;
}


func (self *Node) IsEnd() bool {
	return self.isEnd
}

func (self *Node) SetEnd() {
	self.isEnd = true
}

func NewNode() *Node {
	return &Node{childer: make(map[rune]*Node)}
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		root: &Node{
			childer: make(map[rune]*Node),
		},
	}
}

// 插入到前缀树
func (self *PrefixTree) Insert(word string) {
	node := self.root
	
	for _, v := range word {
		if _, ok := node.childer[v]; !ok {
			node.childer[v] = NewNode()
		}

		node = node.childer[v]
	}

	node.SetEnd()
}

// 搜索前缀
func (self *PrefixTree) SearchPrefix(word string) *Node {
	node := self.root
	
	for _, v := range word {
		if _, ok := node.childer[v]; !ok {
			return nil
		}

		node = node.childer[v]
	}
	
	return node
}

// 查询字符串是否存在
func (self *PrefixTree) Search(word string) bool {
	node := self.SearchPrefix(word)
	return node != nil && node.IsEnd()
}




