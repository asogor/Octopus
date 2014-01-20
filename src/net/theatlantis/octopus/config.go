/**
 */
package octopus

import (
//	"io"
//	"fmt"
	"regexp"
	"errors"
)

type funnelTemplate struct {
	id FunnelId
	expr regexp.Regexp
	aid ActionId
}

type builder struct {
	path map[PathKey] PathBuilder
	actions map[ActionId] Action
}

type pathBuilder struct {
	builder *builder
}

type Builder interface {
	AddAction(action Action) (err error)
    NewPath(key PathKey) (builder PathBuilder,err error)
}

type PathBuilder interface {
	CreateFunnel(id FunnelId,expression *regexp.Regexp, aid ActionId)
}

type Id interface {
	toString()(id string)
}

func NewBuilder() (b Builder) {
	i := new(builder)
	i.path = make(map[PathKey]PathBuilder)
	i.actions = make(map[ActionId]Action)
	return i
}

func (m *builder) AddAction(action Action) (err error) {
	if(m.actions[action.getId()] == nil){
		m.actions[action.getId()] = action
	}else{
		return errors.New("Duplicate Action")
	}
	return nil
}

func (m *builder) NewPath(key PathKey) (b PathBuilder,err error) {
	if(m.path[key] == nil){
		path := pathBuilder{m}
		m.path[key] = &path
		return &path, nil
	}

	return nil, errors.New("Duplicate Action")
}

func (builder *pathBuilder) CreateFunnel(id FunnelId,expression *regexp.Regexp, aid ActionId) {

}
