package octopus

import (
	"regexp"
	"testing"
)

func TestCreateSimple(t *testing.T) {
	re1, _ := regexp.Compile(`[Le]{3}`)
	builder := NewBuilder()
	ta := &CounterAction{42, 0, 0}
	builder.AddAction(ta)
	fg, _ := builder.NewFunnelGroup("test")
	fg.AddFunnel(1, re1, 42)
}
