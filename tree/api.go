package tree

type PreserveMode int

const (
	ModePreserve PreserveMode = 0
	ModeReplace  PreserveMode = 1
)

const (
	RootNode string = "root"
)

func (p PreserveMode) Preserve() bool {
	return p == ModePreserve
}

func (p PreserveMode) Replace() bool {
	return p == ModeReplace
}

func (p PreserveMode) Uint8() uint8 {
	return uint8(p)
}

type Node interface {
	Name() string
	IsLeaf() bool
	IsRoot() bool
}

type Next interface {
	Next(child string) Trie
}

type Childish interface {
	HasChild(name string) bool
	Children() map[string]Trie
}

type Trie interface {
	Find(path ...string) (depth int, n Node)
	Add(tree Trie, mode uint8, path ...string)
	Append(key string, tree Trie, mode uint8) Node
	Depth() int
	Node
	Next
	Childish
}

type Tree struct {
	key   string
	depth int
	nodes map[string]Trie
}

func New(key string) *Tree {
	return &Tree{key, 0, make(map[string]Trie)}
}

func NewRoot() *Tree {
	return New(RootNode)
}

func (t *Tree) Name() string {
	return t.key
}

func (t *Tree) IsLeaf() bool {
	return len(t.nodes) == 0
}

func (t *Tree) IsRoot() bool {
	return t.key == RootNode
}

func (t *Tree) HasChild(child string) bool {
	return !IsNilNode(t.nodes[child])
}

func (t *Tree) Next(child string) Trie {
	return t.nodes[child]
}

func (t *Tree) Children() map[string]Trie {
	return t.nodes
}

func (t *Tree) Append(key string, tree Trie, mode uint8) Node {
	if found, ok := t.nodes[key]; ok && PreserveMode(mode).Preserve() {
		return found
	}

	t.nodes[key] = tree
	if t.depth == 0 {
		t.depth = 1
	}

	return tree
}

func (t *Tree) Find(path ...string) (int, Node) {
	var (
		depth int  = 0
		n     Trie = t
	)

	if len(path) == 0 {
		return depth, nil
	}

	for _, p := range path {
		if n.HasChild(p) {
			n = n.Next(p)
			depth++
			continue
		}

		break
	}

	return depth, n
}

func (t *Tree) FindExact(path ...string) (int, Node) {
	depth, n := t.Find(path...)
	if t.isPathMatch(n, depth, path) {
		return depth, n
	}

	return 0, nil
}

func (t *Tree) Add(tree Trie, mode uint8, path ...string) {
	var (
		n Trie = t
	)

	if len(path) == 0 {
		return
	}

	pathLen := len(path)
	for i, p := range path {
		if i == pathLen-1 {
			n.Append(p, tree, mode)
			break
		}

		if n.HasChild(p) {
			n = n.Next(p)
			continue
		}

		n.Append(p, tree, mode)
	}

	t.setDepth(pathLen)
}

func (t *Tree) Depth() int {
	return t.depth
}

func (t *Tree) Bredth() int {
	return len(t.Children())
}

func (t *Tree) setDepth(delta int) {
	if delta > t.depth {
		t.depth = delta
	}
}

func (t *Tree) isPathMatch(n Node, depth int, path []string) bool {
	pathLen := len(path) - 1
	lastPath := path[pathLen]

	if pathLen+1 != depth || n.Name() != lastPath {
		return false
	}

	return true
}

func IsNilNode(n Node) bool {
	return n == nil
}

func IsNilTrie(t Trie) bool {
	return t == nil
}
