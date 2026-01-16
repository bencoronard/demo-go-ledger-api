package resource

import (
	"context"
	"strconv"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"github.com/golang-jwt/jwt/v5"
)

type resourceService interface {
	listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error)
	retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (*resource, error)
	createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (uint, error)
	updateResource(ctx context.Context, dto *resource, claims jwt.MapClaims) error
	deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error
}

type resourceServiceImpl struct {
	r resourceRepo
}

func NewResourceService(r resourceRepo) resourceService {
	return &resourceServiceImpl{r: r}
}

func (s *resourceServiceImpl) listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error) {

	sub, err := claims.GetSubject()
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	ents, err := s.r.findAll(ctx, page, uint(createdBy))
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	return ents, nil
}

// createResource implements [resourceService].
func (s *resourceServiceImpl) createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (uint, error) {
	panic("unimplemented")
}

// deleteResource implements [resourceService].
func (s *resourceServiceImpl) deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error {
	panic("unimplemented")
}

// retrieveResource implements [resourceService].
func (s *resourceServiceImpl) retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (*resource, error) {
	panic("unimplemented")
}

// updateResource implements [resourceService].
func (s *resourceServiceImpl) updateResource(ctx context.Context, dto *resource, claims jwt.MapClaims) error {
	panic("unimplemented")
}
