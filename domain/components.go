package domain

import (
	"encoding/json"
)

type ComponentType string

const (
	EndpointComponent ComponentType = "endpoint"
	ActionComponent   ComponentType = "action"
	ModelComponent    ComponentType = "model"
)

var StringToComponentTypeMap = map[string]ComponentType{
	"endpoint": EndpointComponent,
	"action":   ActionComponent,
	"model":    ModelComponent,
}

var ComponentTypeToStringMap = map[ComponentType]string{
	EndpointComponent: "endpoint",
	ActionComponent:   "action",
	ModelComponent:    "model",
}

type Component struct {
	ID   int             `gorm:"column:id;primaryKey"`
	Type ComponentType   `gorm:"column:type"`
	Name string          `gorm:"column:name"`
	Data json.RawMessage `gorm:"column:data;type:jsonb"`
}

type ComponentData interface {
	GetType() ComponentType
}

type Components []Component

// Endpoint represents an endpoint component.
type Endpoint struct {
	Name   string
	URL    string
	Method string
}

func (e Endpoint) GetType() ComponentType {
	return EndpointComponent
}

// Action represents an action component.
type Action struct {
	Name string
}

func (a Action) GetType() ComponentType {
	return ActionComponent
}

// Model represents a model component.
type Model struct {
	Name string
}

func (m Model) GetType() ComponentType {
	return ModelComponent
}
