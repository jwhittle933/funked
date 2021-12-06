package control

import "github.com/jwhittle933/funked/boolean"

type ElseIfer interface {
	ElseIf(cond bool, do func()) ElseIfer
	Else(do func())
}

type IfFlow bool

func If(cond bool, do func()) ElseIfer {
	if cond {
		do()
		return IfFlow(true)
	}

	return IfFlow(false)
}

func Unless(cond bool, do func()) ElseIfer {
	if !cond {
		do()
		return IfFlow(true)
	}

	return IfFlow(false)
}

func (i IfFlow) ElseIf(cond bool, do func()) ElseIfer {
	if boolean.And(!bool(i), cond) {
		do()
		return IfFlow(true)
	}

	return IfFlow(false)
}

func (i IfFlow) Else(do func()) {
	if !i {
		do()
	}
}
