package xlink

import (
	"github.com/valyala/fasthttp"
	"github.com/vela-ssoc/vela-kit/vela"
)

func (mgt *Mgt) define(r vela.Router) {

	r.GET("/api/v1/arr/agent/notice", xEnv.Then(func(ctx *fasthttp.RequestCtx) error {
		return nil
	}))

}
