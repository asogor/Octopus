package octopus

import (
	"regexp"
	"log"
	"fmt"
)

type Tracker interface {
	GetCurrentState()(state FunnelState)
	WriteRune(r rune, size int, err error)
}

type TrackerGroup interface {
	WriteRune(r rune, size int, err error)
	addTracker(id FunnelId, exp *regexp.Regexp, length int, action Action)(err error)
}

type tracker struct {
	id FunnelId
	exp *regexp.Regexp
	input runeChan
	action Action
	context *Context
	state FunnelState
}

type trackerGroup struct {
	context *Context
	trackers map[FunnelId] Tracker
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

func NewFunnelGroup(context *Context)(fg TrackerGroup) {
	r := trackerGroup{context,make(map[FunnelId]Tracker)}
	return &r
}

func (fg *trackerGroup)WriteRune(r rune, size int, err error){
	for _ , tracker := range fg.trackers {
		if(tracker.GetCurrentState() == STATE_RUN) {
			tracker.WriteRune(r,size,err)
		}
	}
}

func (fg *trackerGroup)addTracker(id FunnelId, exp *regexp.Regexp, length int, action Action)(err error){
	if(exp == nil){
		panic("Regexp is nil")
	}
	if _, got := fg.trackers[id]; !got {
		fg.trackers[id] = NewFunnel(id,exp,length,action,fg.context)
	}else{
		return fmt.Errorf("Funnel already exist: %v",id)
	}
	return nil
}

func NewFunnel(id FunnelId, exp *regexp.Regexp, length int, action Action,context *Context)(r Tracker) {
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
	m := &tracker{id,exp,runeInput,action,context,STATE_RUN}
	go m.run()
	return m
}

func (m *tracker) run() {
	loc := m.exp.FindReaderIndex(m.input);
	if(loc != nil){
		m.action.execute(m.context,&Result{m.id,RC_GOT_MATCH,loc})
	}
	m.action.execute(m.context,&Result{m.id,RC_COMPLETE,nil})
}

func (m *tracker) GetCurrentState()(state FunnelState) {
	return m.state
}

func (m *tracker) WriteRune(r rune, size int, err error) {
	entry := &Entry{r,size,err}
	m.input.input <- entry
}
