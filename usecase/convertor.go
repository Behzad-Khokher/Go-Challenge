package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
)

func EndpointToComponent(endpoint *domain.Endpoint) (*domain.Component, error) {
	data, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	endpointType, ok := domain.StringToComponentTypeMap["endpoint"]
	if !ok {
		return nil, fmt.Errorf("invalid component type: %s", "endpoint")
	}

	return &domain.Component{
		Type: endpointType,
		Name: endpoint.Name,
		Data: data,
	}, nil
}

func ComponentToEndpoint(component *domain.Component) (*domain.Endpoint, error) {
	if component.Type != domain.EndpointComponent {
		return nil, fmt.Errorf("invalid component type: %s", component.Type)
	}

	endpoint := &domain.Endpoint{}
	err := json.Unmarshal(component.Data, endpoint)
	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
