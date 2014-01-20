package octopus

import (
	"testing"
	"regexp"
)



func TestCreateSimple(t *testing.T) {
	re1, _ := regexp.Compile(`[Le]{3}`)
	builder := NewBuilder()
	ta := &CounterAction{42,0,0}
	builder.AddAction(ta)
	path, _ := builder.NewPath("test")
	path.CreateFunnel(1,re1,42)
}
