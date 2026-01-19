package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"gorm.io/gorm"
)

type resourceRepo interface {
	findAllByCreatedBy(ctx context.Context, page dto.Pageable, createdBy uint) (dto.Slice[resource], error)
	findByIdAndCreatedBy(ctx context.Context, id uint, createdBy uint) (*resource, error)
	save(ctx context.Context, ent *resource) (*resource, error)
	delete(ctx context.Context, ent *resource) error
}

type resourceRepoImpl struct {
	db *gorm.DB
}

func NewResourceRepo(db *gorm.DB) resourceRepo {
	return &resourceRepoImpl{db: db}
}

func (r *resourceRepoImpl) findAllByCreatedBy(ctx context.Context, page dto.Pageable, createdBy uint) (dto.Slice[resource], error) {
	query := gorm.G[resource](r.db).Where("created_by = ?", createdBy).Limit(page.GetLimit() + 1).Offset(page.GetOffset())

	for _, sort := range page.Sort {
		orderClause := fmt.Sprintf("%s %s", sort.Property, sort.Direction)
		query = query.Order(orderClause)
	}

	ents, err := query.Find(ctx)
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	return *dto.NewSlice(ents, &page, len(ents)), nil
}

func (r *resourceRepoImpl) findByIdAndCreatedBy(ctx context.Context, id uint, createdBy uint) (*resource, error) {
	ent, err := gorm.G[resource](r.db).Where("id = ?", id).Where("created_by = ?", createdBy).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &ent, nil
}

func (r *resourceRepoImpl) save(ctx context.Context, ent *resource) (*resource, error) {
	if ent.ID == 0 {
		return ent, gorm.G[resource](r.db).Create(ctx, ent)
	}

	rowsAffected, err := gorm.G[resource](r.db).Where("id = ?", ent.ID).Select("*").Updates(ctx, *ent)
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrOptimisticLockFailure
	}

	return ent, nil
}

func (r *resourceRepoImpl) delete(ctx context.Context, ent *resource) error {
	_, err := gorm.G[resource](r.db).Where("id = ?", ent.ID).Delete(ctx)
	return err
}
