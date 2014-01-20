/**
 *
 */
package octopus

import (
//	"io"
//	"fmt"
)

/*
Represents a unique action instance
 */
type ActionId int
/*
Represents a unique funnel. A single tracking path can have multiple funnels.
 */
type FunnelId int
/*
PathKey differentiates different funnel groups.
 */
type PathKey string
/*
Represents a unique user session
 */
type SessionId string

type ResultCode int

type FunnelState int

const (
	RC_GOT_MATCH ResultCode = 1
	RC_COMPLETE ResultCode = 2
	STATE_RUN FunnelState = 1
	STATE_COMPLETE FunnelState = 2
)

type Context struct {
	session SessionId
	path PathKey
}

type Result struct {
    id FunnelId
	code ResultCode
	pos []int
}

type Action interface {
	getId()(oid ActionId)
	execute(c *Context,r *Result)
}

type CounterAction struct {
	id ActionId
	CompleteCounter int
	MatchCounter int
}

func (t *CounterAction) getId()(oid ActionId) {
	return t.id
}

func (t *CounterAction) execute(c *Context,r *Result) {
	if(r.code == RC_GOT_MATCH){
		t.MatchCounter++
	}
	if(r.code == RC_COMPLETE){
		t.CompleteCounter++
	}
}
