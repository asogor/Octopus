package octopus

import (
	"regexp"
	"testing"
	"strings"
	"log"
)

func TestCreateSimple(t *testing.T) {
	re1, _ := regexp.Compile(`[Le]{3}`)
	re2, _ := regexp.Compile(`[Te]{3}`)
	builder := NewBuilder()
	ta := &CounterAction{42, 0, 0}
	tb := &CounterAction{43, 0, 0}
	builder.AddAction(ta)
	builder.AddAction(tb)
	fg, _ := builder.NewFunnelGroup("test")
	fg.AddFunnel(1, re1, 42)
	fg.AddFunnel(1, re2, 43)
	tg, tgerr := builder.NewTrackerGroup(&Context{"userId99","test"},10)

	if(tgerr != nil) {
		t.Errorf("Failed  %#v", tgerr)
	}

	reader := strings.NewReader(`Mrs. Leonora Teonora`)
	var r rune = 0
	var size int = 0
	var err error = nil
	for err == nil {
		r, size, err = reader.ReadRune()
		log.Printf("Write %#v %#v %#v", r, size, err)
		tg.WriteRune(r, size, err)
	}
}
