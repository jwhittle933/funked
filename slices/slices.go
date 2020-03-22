package slices

func FilterIntSlice(src []int, cb func(int) bool) []int {
	ch := make(chan int, 1)
	done := make(chan struct{}, 1)
	dist := make([]int, 0)

	go func() {
		for i := 0; i < len(src); i++ {
			go func(i int) {
				if cb(src[i]) {
					ch <- src[i]
				}
			}(i)
		}
		done <- struct{}{}
	}()

	for {
		select {
		case val := <-ch:
			//
			break
		case <-done:
			//
		}
	}

	return dist
}

func FilterInt32Slice(src []int32, cb func(int32) bool) {
	for i := 0; i < len(src); i++ {
		//
	}
}

func FilterInt64Slice(src []int64, cb func(int32) bool) {
	for i := 0; i < len(src); i++ {
		//
	}
}

func FilterStringSlice(src []string, cb func(string) bool) {
	for i := 0; i < len(src); i++ {
		//
	}
}
