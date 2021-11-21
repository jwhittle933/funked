package tree

type PreserveMode int

const (
	ModePreserve PreserveMode = 0
	ModeReplace  PreserveMode = 1
)

const (
	Root string = "root"
)

func (p PreserveMode) Preserve() bool {
	return p == ModePreserve
}

func (p PreserveMode) Replace() bool {
	return p == ModeReplace
}

type Trie interface {
	Find(path ...string) (depth int, n Node)
	Add(n Node, mode PreserveMode, path ...string)
	Depth() int
	Node
}

type Next interface {
	Next(child string) Node
}

type Childish interface {
	HasChild(name string) bool
	Children() map[string]Node
}

type Node interface {
	Name() string
	IsLeaf() bool
	IsRoot() bool
	Append(key string, n Node, mode PreserveMode) Node
	Next
	Childish
}

type Tree struct {
	root  Node
	depth int
}

func New(root Node) *Tree {
	return &Tree{root, 0}
}

func (t *Tree) Name() string {
	return t.root.Name()
}

func (t *Tree) IsLeaf() bool {
	return t.root.IsLeaf()
}

func (t *Tree) IsRoot() bool {
	return t.root.IsRoot()
}

func (t *Tree) Append(key string, n Node, mode PreserveMode) Node {
	return t.root.Append(key, n, mode)
}

func (t *Tree) Next(child string) Node {
	return t.root.Next(child)
}

func (t *Tree) Children() map[string]Node {
	return t.root.Children()
}

func (t *Tree) HasChild(child string) bool {
	return t.root.HasChild(child)
}

func (t *Tree) Find(path ...string) (int, Node) {
	depth := 0
	n := t.root

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

	if !t.isPathMatch(n, depth, path) {
		return 0, nil
	}

	return depth, n
}

func (t *Tree) Add(node Node, mode PreserveMode, path ...string) {
	if len(path) == 0 {
		return
	}

	n := t.root
	pathLen := len(path)
	for i, p := range path {
		if i == pathLen-1 {
			n.Append(p, node, mode)
			break
		}

		if n.HasChild(p) {
			n = n.Next(p)
			continue
		}

		n.Append(p, node, mode)
	}

	t.setDepth(pathLen)
}

func (t *Tree) Depth() int {
	return t.depth
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
