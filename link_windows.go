package xlink

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"github.com/asdine/storm/v3"
	"github.com/valyala/fasthttp"
	"github.com/vela-ssoc/vela-kit/lua"
	"net"
	"reflect"
	"time"
)

var typeof = reflect.TypeOf((*Mgt)(nil)).String()

type Mgt struct {
	lua.SuperVelaData
	ln     net.Listener
	handle func(ctx *fasthttp.RequestCtx)
	db     storm.Node
	path   string
}

func (mgt *Mgt) Listen() (net.Listener, error) {
	ln, err := winio.ListenPipe(mgt.path, &winio.PipeConfig{
		SecurityDescriptor: "D:P(A;;GA;;;AU)", //allow all user
	})

	if err != nil {
		return nil, err
	}

	if ln == nil {
		return nil, fmt.Errorf("win pipe listen fail")
	}

	mgt.ln = ln

	xEnv.Spawn(200, func() {
		defer mgt.Close()
		err = fasthttp.Serve(mgt.ln, mgt.Handle)
		xEnv.Errorf("named pipe server stop %v", err)
	})

	return ln, nil
}

func (mgt *Mgt) Close() error {
	mgt.handle = Forbidden
	mgt.V(lua.VTClose, time.Now())
	return nil
}

func (mgt *Mgt) Type() string {
	return typeof
}

func (mgt *Mgt) Name() string {
	return "vela.link.server"
}

func (mgt *Mgt) Handle(ctx *fasthttp.RequestCtx) {
	mgt.handle(ctx)
}

func (mgt *Mgt) Start() error {
	mgt.handle = xEnv.R().Handler()
	mgt.V(lua.VTRun, time.Now())
	return nil
}
