package resource

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"github.com/bencoronard/demo-go-common-libs/orm"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type resourceService interface {
	listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error)
	retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (*resource, error)
	createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (uint, error)
	updateResource(ctx context.Context, id uint, dto *resource, claims jwt.MapClaims) error
	deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error
}

type resourceServiceImpl struct {
	t orm.TransactionManager
	r resourceRepo
}

func NewResourceService(t orm.TransactionManager, r resourceRepo) resourceService {
	return &resourceServiceImpl{
		t: t,
		r: r,
	}
}

func (s *resourceServiceImpl) listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error) {
	var slice dto.Slice[resource]

	sub, err := claims.GetSubject()
	if err != nil {
		return slice, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return slice, err
	}

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		slice, err = s.r.findAllByCreatedBy(ctx, tx, page, uint(createdBy))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return slice, err
	}

	return slice, nil
}

func (s *resourceServiceImpl) retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (*resource, error) {
	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return nil, err
	}

	var resource *resource

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		resource, err = s.r.findByIdAndCreatedBy(ctx, tx, id, uint(createdBy))
		if err != nil {
			return err
		}
		if resource == nil {
			return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resource, err
}

func (s *resourceServiceImpl) createResource(ctx context.Context, dto *resource, claims jwt.MapClaims) (uint, error) {
	sub, err := claims.GetSubject()
	if err != nil {
		return 0, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}

	var resource resource
	resource.CreatedBy = uint(createdBy)
	resource.TextField = dto.TextField
	resource.NumberField = dto.NumberField
	resource.BooleanField = dto.BooleanField

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		if err := s.r.save(ctx, tx, &resource); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return resource.ID, err
}

func (s *resourceServiceImpl) updateResource(ctx context.Context, id uint, dto *resource, claims jwt.MapClaims) error {
	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		resource, err := s.r.findByIdAndCreatedBy(ctx, tx, id, uint(createdBy))
		if err != nil {
			return err
		}
		if resource == nil {
			return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
		}

		resource.TextField = dto.TextField
		resource.NumberField = dto.NumberField
		resource.BooleanField = dto.BooleanField

		if err := s.r.save(ctx, tx, resource); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *resourceServiceImpl) deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error {
	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		resource, err := s.r.findByIdAndCreatedBy(ctx, tx, id, uint(createdBy))
		if err != nil {
			return err
		}
		if resource == nil {
			return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
		}

		if err := s.r.delete(ctx, tx, resource); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
