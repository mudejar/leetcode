package main

func main() {

}

type Trie struct {
	isWord   bool
	children []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isWord:   false,
		children: make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.children[c] == nil {
			newNode := Constructor()
			cur.children[c] = &newNode
		}
		cur = cur.children[c]
	}

	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}

	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		c := int(prefix[i] - 'a')
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
