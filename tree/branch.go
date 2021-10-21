package tree

const (
	BranchTerminator string = "*"
	BranchRoot       string = "root"
)

type Branch struct {
	key   string
	nodes map[string]Node
}

func NewBranch(key string) *Branch {
	return &Branch{key, make(map[string]Node)}
}

func (b *Branch) Name() string {
	return b.key
}

func (b *Branch) IsLeaf() bool {
	return len(b.nodes) == 0
}

func (b *Branch) IsRoot() bool {
	return b.key == BranchRoot
}

func (b *Branch) HasChild(child string) bool {
	return !IsNilNode(b.nodes[child])
}

func (b *Branch) Next(child string) Node {
	return b.nodes[child]
}

func (b *Branch) Children() map[string]Node {
	return b.nodes
}

// Append adds a Node `n` at `child`.
// If a Node at `key` already exists and mode == ModePreserve,
// that Node is returned and the added Node is discarded.
func (b *Branch) Append(key string, n Node, mode PreserveMode) Node {
	if node, ok := b.nodes[key]; ok && mode.Preserve() {
		return node
	}

	b.nodes[key] = n
	return n
}
