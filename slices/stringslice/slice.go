package stringslice

type Slice []string

func From(strs []string) Slice {
	return strs
}

func (s Slice) Filter(sfn BoolFn) Slice {
	return sfn.Filter(s)
}

func (s Slice) Some(sfn BoolFn) bool {
	return sfn.Some(s)
}

func (s Slice) Every(sfn BoolFn) bool {
	return sfn.Every(s)
}

func (s Slice) Map(sfn StringFn) Slice {
	return sfn.Map(s)
}

