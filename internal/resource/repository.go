package resource

import (
	"context"
	"database/sql"
)

type resourceRepo interface {
	findAll(ctx context.Context, p any, createdBy int64) (any, error)
	findById(ctx context.Context, id int64, createdBy int64) (*resource, error)
	save(ctx context.Context, ent *resource) (*resource, error)
	delete(ctx context.Context, ent *resource) error
}

type resourceRepoImpl struct {
	db *sql.DB
}

func NewResourceRepoImpl(db *sql.DB) resourceRepo {
	return &resourceRepoImpl{db: db}
}

// delete implements resourceRepo.
func (r *resourceRepoImpl) delete(ctx context.Context, ent *resource) error {
	panic("unimplemented")
}

// findAll implements resourceRepo.
func (r *resourceRepoImpl) findAll(ctx context.Context, p any, createdBy int64) (any, error) {
	panic("unimplemented")
}

// findById implements resourceRepo.
func (r *resourceRepoImpl) findById(ctx context.Context, id int64, createdBy int64) (*resource, error) {
	panic("unimplemented")
}

// save implements resourceRepo.
func (r *resourceRepoImpl) save(ctx context.Context, ent *resource) (*resource, error) {
	panic("unimplemented")
}
