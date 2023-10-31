package xlink

import (
	"github.com/vela-ssoc/vela-kit/lua"
)

func (mgt *Mgt) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	}

	return lua.LNil
}

func (mgt *Mgt) call(L *lua.LState) int {
	v := L.NewVelaData(mgt.Name(), typeof)
	if v.IsNil() {
		v.Set(mgt)
	}

	if mgt.IsRun() {
		L.Push(v)
		return 1
	}

	mgt.Start()

	L.Push(v)
	return 1
}
