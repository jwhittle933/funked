package genserver

type Tuple []byte
type List []byte

var (
	Empty []byte = nil
)

type Server interface {
	Init(state []byte) ([]byte, error)
	HandleCall(key string, data []byte) []byte
	HandleCast(key string, data []byte) []byte
}

type Link struct {
	s    Server
	call chan bool
	cast chan bool
}

// Start starts a GenServer process without links
func Start(s Server, state []byte) (*Link, error) {
	return nil, nil
}

// StartLink starts a GenServer on the current process
func StartLink(s Server, state []byte) (*Link, error) {
	_, err := s.Init(state)
	if err != nil {
		return nil, err
	}

	call := make(chan bool)
	cast := make(chan bool)

	// this goroutine needs to be
	// referenced by a process
	go func() {
		for {
			select {
			case <-call:
				break
			case <-cast:
				break
			}
		}
	}()

	return &Link{s, call, cast}, nil
}

// Call invokes the call handler synchronously
func (l *Link) Call(name string, data []byte) {
	l.s.HandleCall(name, data)
}

// Cast invokes the cast handler asynchronously
func (l *Link) Cast(name string, data []byte) {
	go func() {
		l.s.HandleCast(name, data)
	}()
}
