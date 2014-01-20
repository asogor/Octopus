package octopus

import (
	"regexp"
	"log"
	"fmt"
)

type Funnel interface {
	GetCurrentState()(state FunnelState)
	WriteRune(r rune, size int, err error)
}

type FunnelGroup interface {
	WriteRune(r rune, size int, err error)
	AddFunnel(id FunnelId, exp *regexp.Regexp, length int, action Action)(err error)
}

type funnel struct {
	id FunnelId
	exp *regexp.Regexp
	input runeChan
	action Action
	context *Context
	state FunnelState
}

type funnelGroup struct {
	context *Context
	funnels map[FunnelId] Funnel
}

type Entry struct {
	r rune
	size int
	err error
}

type runeChan struct {
	input chan *Entry
}

func (rc runeChan)ReadRune() (r rune, size int, err error) {
	entry := <- rc.input
	log.Printf("Read %#v",entry)
	return entry.r,entry.size,entry.err
}

func NewFunnelGroup(context *Context)(fg FunnelGroup) {
	r := funnelGroup{context,make(map[FunnelId]Funnel)}
	return &r
}

func (fg *funnelGroup)WriteRune(r rune, size int, err error){
	for _ , funnel := range fg.funnels {
		if(funnel.GetCurrentState() == STATE_RUN) {
			funnel.WriteRune(r,size,err)
		}
	}
}

func (fg *funnelGroup)AddFunnel(id FunnelId, exp *regexp.Regexp, length int, action Action)(err error){
	if(exp == nil){
		panic("Regexp is nil")
	}
	if(fg.funnels[id] == nil){
		fg.funnels[id] = NewFunnel(id,exp,length,action,fg.context)
	}else{
		return fmt.Errorf("Funnel already exist: %v",id)
	}
	return nil
}

func NewFunnel(id FunnelId, exp *regexp.Regexp, length int, action Action,context *Context)(r Funnel) {
	input := make (chan *Entry, length)
	runeInput := runeChan{input}
	if(exp == nil){
		panic("Regexp is nil")
	}
	if(input == nil){
		panic("Input is nil")
	}
	if(action == nil){
		panic("Action is nil")
	}
	if(context == nil){
		panic("Context is nil")
	}
	m := &funnel{id,exp,runeInput,action,context,STATE_RUN}
	go m.run()
	return m
}

func (m *funnel) run() {
	loc := m.exp.FindReaderIndex(m.input);
	if(loc != nil){
		m.action.execute(m.context,&Result{m.id,RC_GOT_MATCH,loc})
	}
	m.action.execute(m.context,&Result{m.id,RC_COMPLETE,nil})
}

func (m *funnel) GetCurrentState()(state FunnelState) {
	return m.state
}

func (m *funnel) WriteRune(r rune, size int, err error) {
	entry := &Entry{r,size,err}
	m.input.input <- entry
}
