package usecase

import (
	"fmt"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
	Component "github.com/Behzad-Khokher/Go-Challenge/domain/components"
)

type CreateEndpointReq struct {
	Name   string
	URL    string
	Method string
}

type EditEndpointReq struct {
	id     int
	Name   *string
	URL    *string
	Method *string
}

type Endpoint interface {
	Create(req CreateEndpointReq) error
	Edit(req EditEndpointReq) error
	// Get(id int) (*domain.Endpoint, error)
}

type endpointUsecase struct {
	store Component.Store
}

func NewEndpointUsecase(cs Component.Store) Endpoint {

	return &endpointUsecase{store: cs}
}

func (u *endpointUsecase) Create(req CreateEndpointReq) error {

	endpoint := &domain.Endpoint{
		Name:   req.Name,
		URL:    req.URL,
		Method: req.Method,
	}

	componentEndpoint, err := EndpointToComponent(endpoint)
	if err != nil {
		return err
	}

	return u.store.Create(componentEndpoint)
}

func (u *endpointUsecase) Edit(req EditEndpointReq) error {
	// Fetch the existing component
	component, err := u.store.GetByID(req.id)
	if err != nil {
		return err
	}

	if component == nil {
		return fmt.Errorf("component with id %d not found", req.id)
	}

	// Convert component's Data to Endpoint
	endpoint, err := ComponentToEndpoint(component)
	if err != nil {
		return err
	}

	// Update endpoint fields based on EditEndpointReq
	if req.Name != nil {
		endpoint.Name = *req.Name
	}
	if req.URL != nil {
		endpoint.URL = *req.URL
	}
	if req.Method != nil {
		endpoint.Method = *req.Method
	}

	// Convert updated endpoint back to component
	updatedComponent, err := EndpointToComponent(endpoint)
	if err != nil {
		return err
	}

	// Update the component in the store
	return u.store.Update(updatedComponent)
}

// func (u *endpointUsecase) Get(id int) (*domain.Endpoint, error) {
// 	return nil, nil
// }
