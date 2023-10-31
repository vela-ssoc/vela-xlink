package xlink

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

func Constructor(env vela.Environment, callback func(interface{})) {
	xEnv = env

	mgt := &Mgt{
		db:     xEnv.Storm("VELA_LINK_MGT_DB"),
		path:   "\\\\.\\pipe\\ssc",
		handle: Forbidden,
	}

	callback(mgt)
	mgt.define(xEnv.R())

	if _, err := mgt.Listen(); err != nil {
		xEnv.Error("windows manager pipe listen fail %v", err)
	}
	xEnv.Set("xlink", lua.NewExport("vela.xlink.export", lua.WithFunc(mgt.call)))
}
