package resource

import (
	"context"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type resourceService interface {
	listResources(ctx context.Context, p any, claims jwt.MapClaims) (any, error)
	retrieveResource(ctx context.Context, id int64, claims jwt.MapClaims) (*resource, error)
	createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (int64, error)
	updateResource(ctx context.Context, dto *resource, claims jwt.MapClaims) error
	deleteResource(ctx context.Context, id int64, claims jwt.MapClaims) error
}

type resourceServiceImpl struct {
	r resourceRepo
}

func NewResourceServiceImpl(r resourceRepo) resourceService {
	return &resourceServiceImpl{r: r}
}

// listResources implements resourceService.
func (s *resourceServiceImpl) listResources(ctx context.Context, p any, claims jwt.MapClaims) (any, error) {

	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	createdBy, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return nil, err
	}

	ents, err := s.r.findAll(ctx, p, createdBy)
	if err != nil {
		return nil, err
	}

	return ents, nil
}

// createResource implements resourceService.
func (s *resourceServiceImpl) createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (int64, error) {
	panic("unimplemented")
}

// deleteResource implements resourceService.
func (s *resourceServiceImpl) deleteResource(ctx context.Context, id int64, claims jwt.MapClaims) error {
	panic("unimplemented")
}

// retrieveResource implements resourceService.
func (s *resourceServiceImpl) retrieveResource(ctx context.Context, id int64, claims jwt.MapClaims) (*resource, error) {
	panic("unimplemented")
}

// updateResource implements resourceService.
func (s *resourceServiceImpl) updateResource(ctx context.Context, dto *resource, claims jwt.MapClaims) error {
	panic("unimplemented")
}
