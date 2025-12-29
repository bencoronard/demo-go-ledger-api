package resource

import (
	"context"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"gorm.io/gorm"
)

type resourceRepo interface {
	findAll(ctx context.Context, page dto.Pageable, createdBy int64) (dto.Slice[resource], error)
	findById(ctx context.Context, id int64, createdBy int64) (*resource, error)
	save(ctx context.Context, ent *resource) (*resource, error)
	delete(ctx context.Context, ent *resource) error
}

type resourceRepoImpl struct {
	db *gorm.DB
}

func NewResourceRepoImpl(db *gorm.DB) resourceRepo {
	return &resourceRepoImpl{db: db}
}

// delete implements [resourceRepo].
func (r *resourceRepoImpl) delete(ctx context.Context, ent *resource) error {
	panic("unimplemented")
}

// findAll implements [resourceRepo].
func (r *resourceRepoImpl) findAll(ctx context.Context, page dto.Pageable, createdBy int64) (dto.Slice[resource], error) {
	panic("unimplemented")
}

// findById implements [resourceRepo].
func (r *resourceRepoImpl) findById(ctx context.Context, id int64, createdBy int64) (*resource, error) {
	panic("unimplemented")
}

// save implements [resourceRepo].
func (r *resourceRepoImpl) save(ctx context.Context, ent *resource) (*resource, error) {
	panic("unimplemented")
}
