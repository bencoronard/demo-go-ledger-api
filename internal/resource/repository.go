package resource

import (
	"context"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"gorm.io/gorm"
)

type resourceRepo interface {
	findAll(ctx context.Context, page dto.Pageable, createdBy uint) (dto.Slice[resource], error)
	findById(ctx context.Context, id uint, createdBy uint) (*resource, error)
	save(ctx context.Context, ent *resource) (*resource, error)
	delete(ctx context.Context, ent *resource) error
}

type resourceRepoImpl struct {
	db *gorm.DB
}

func NewResourceRepo(db *gorm.DB) resourceRepo {
	return &resourceRepoImpl{db: db}
}

func (r *resourceRepoImpl) findAll(ctx context.Context, page dto.Pageable, createdBy uint) (dto.Slice[resource], error) {
	return dto.Slice[resource]{}, nil
}

func (r *resourceRepoImpl) findById(ctx context.Context, id uint, createdBy uint) (*resource, error) {
	return &resource{}, nil
}

func (r *resourceRepoImpl) save(ctx context.Context, ent *resource) (*resource, error) {
	return &resource{}, nil
}

func (r *resourceRepoImpl) delete(ctx context.Context, ent *resource) error {
	return nil
}
