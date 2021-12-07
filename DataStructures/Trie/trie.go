package trie

//						      root
//					a	 		e    		o	isEnd = false
//					r	 		r	  		r	isEnd = true	   The word is or
//				a	   g	 	a			e	isEnd = false
//				g	   o 		g		g		o	isEnd = true  The word is oreo
//				o	   n		o	 o		a
//									 n		n

//If we want to add new word ORC
//Find pointer to the first O in the ORC
//Find pointer to the first C in the ORC
//Find pointer to the C - it doesnt exists, so we should add new nood and set isEnd = true

//NOTE each node holds array with pointers for every letter that exists 	a 		b 		c 		d 		e 		f
//																			nil		nil		nil	   pointer nil		nil

//Time complexity o(m) m - lenght of the word

const AlphabetSize = 26

//Structure for the node
type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool
}

//Structure for the trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert method
func (t *Trie) Insert(w string) {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			//Create a node
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.IsEnd = true
}

//Search method
func (t *Trie) Search(w string) bool {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[charIndex]
	}
	return currentNode.IsEnd
}
