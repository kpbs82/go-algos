package longest_word

import "log"


type Node struct {
	value rune
	link *Trie
}


type Trie struct {
	childNodes []Node
	endOfString bool
}

var Dict *Trie

func findRuneLink(links []Node, value rune) (*Trie, bool) {
	for _, link := range links {
		if link.value == value {
			return link.link, true
		}
	}
	return nil, false
}

func NewTrie() *Trie {
	return &Trie{endOfString: true, childNodes: make([]Node, 0)}
}


func (r *Trie) PrefixExists(s string) bool {
	//check := true
	var check bool
	i := r
	for _, runeValue := range s {
		_, ok := findRuneLink(i.childNodes, runeValue)
		if ok {
			check = true
			log.Println("Found - ", string(runeValue))
			for j:=0; j<len(i.childNodes); j++ {
				log.Println("child node - ", string(i.childNodes[j].value))
			}
			return  check
		}
	}


	return check
}


func (r *Trie) traverse(s string) bool {
	i := r

	for _, runeValue := range s {
		ti, ok := findRuneLink(i.childNodes, runeValue)
		if !ok {
			return false
		}
		i = ti
		log.Println("Found prefix - ", string(runeValue))
		i.traverse(string(i.childNodes[0].value))
	}
	return true
}

func (r *Trie) Exists(s string) bool {
	log.Println("check if words exists for prefix - ", s)
	check := true
	i := r
	for _, runeValue := range s {
		ti, ok := findRuneLink(i.childNodes, runeValue)
		if !ok {
			return false
		}
		i = ti
		log.Println("Found prefix - ", string(runeValue))
		i.traverse(string(i.childNodes[0].value))

	}
	if !i.endOfString {
		i.endOfString = true
		check = false
	}
	return check
}


func (r *Trie) Add(s string) bool {
	check := true
	i := r
	for _, runeValue := range s {
		ti, ok := findRuneLink(i.childNodes, runeValue)
		if !ok {
			ti = new(Trie)
			ti.childNodes = make([]Node, 0)
			i.childNodes = append(i.childNodes,Node{value: runeValue, link: ti})
		}
		i = ti
	}
	if !i.endOfString {
		i.endOfString = true
		check = false
	}
	return check
}

// Count returns the number of words in a trie.
func (r *Trie) Count() int {
	count := 0
	for _, child := range r.childNodes {
		if child.link.endOfString == true {
			count++
		}
		count += child.link.Count()
	}

	return count
}

func BuildTrie() {
	Dict = NewTrie()

	Dict.Add("algol")
	Dict.Add("fortran")
	Dict.Add("fox")
	Dict.Add("flop")
	Dict.Add("fun")
	Dict.Add("simula")

	log.Println("Dict count - ", Dict.Count())
}