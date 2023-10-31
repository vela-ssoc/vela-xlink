package xlink

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

func Constructor(env vela.Environment, callback func(interface{})) {
	xEnv = env
}
