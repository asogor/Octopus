/**
 */
package octopus

import (
	//	"io"
	"fmt"
	"errors"
	"regexp"
)

type funnelTemplate struct {
	id   FunnelId
	expr regexp.Regexp
	aid  ActionId
}

type builder struct {
	funnelGroup map[FunnelGroupKey]FunnelGroupBuilder
	actions     map[ActionId]Action
}

type funnelGroupBuilder struct {
	builder *builder
	funnels map[FunnelId]*funnelDef
}

type funnelDef struct {
	id         FunnelId
	expression *regexp.Regexp
	actionId   ActionId
}

type Builder interface {
	AddAction(action Action) (err error)
	NewFunnelGroup(key FunnelGroupKey) (builder FunnelGroupBuilder, err error)
	NewTrackerGroup(context *Context, bufferSize int) (tracker TrackerGroup, err error)
}

type FunnelGroupBuilder interface {
	AddFunnel(id FunnelId, expression *regexp.Regexp, aid ActionId) (err error)
	newTrackerGroup(context *Context,bufferSize int) (tracker TrackerGroup)
}

type Id interface {
	toString() (id string)
}

func NewBuilder() (b Builder) {
	i := &builder{make(map[FunnelGroupKey]FunnelGroupBuilder), make(map[ActionId]Action)}
	return i
}

func (m *builder) AddAction(action Action) (err error) {
	if _, got := m.actions[action.getId()]; !got {
		m.actions[action.getId()] = action
	} else {
		return errors.New("Duplicate Action")
	}
	return nil
}

func (m *builder) NewFunnelGroup(key FunnelGroupKey) (b FunnelGroupBuilder, err error) {
	if _, got := m.funnelGroup[key]; !got {
		funnelGroup := funnelGroupBuilder{m, make(map[FunnelId]*funnelDef)}
		m.funnelGroup[key] = &funnelGroup
		return &funnelGroup, nil
	}

	return nil, errors.New("Duplicate Action")
}

func  (m *builder) NewTrackerGroup(context *Context,bufferSize int) (tracker TrackerGroup, err error) {
	if funnelGroup , got := m.funnelGroup[context.path]; got {
		return funnelGroup.newTrackerGroup(context,bufferSize), nil
	}

	return nil, errors.New("FunnelGroupKey Not Found: " + fmt.Sprintf("%b", context.path))
}

func (builder *funnelGroupBuilder) AddFunnel(id FunnelId, expression *regexp.Regexp, aid ActionId) (err error) {
	if _, got := builder.funnels[id]; !got {
		f := funnelDef{id, expression, aid}
		builder.funnels[id] = &f
	}

	return errors.New("Duplicate Funnel")
}

func (builder *funnelGroupBuilder) newTrackerGroup(context *Context, bufferSize int) (tracker TrackerGroup) {
	tg := NewTrackerGroup(context)
	for _ , value := range builder.funnels {
		tg.addTracker(value.id,value.expression,bufferSize,builder.builder.actions[value.actionId])
	}
	return tg
}
