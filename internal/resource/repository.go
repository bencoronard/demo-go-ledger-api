package resource

import (
	"fmt"

	"github.com/bencoronard/demo-go-common-libs/dto"
	"gorm.io/gorm"
)

type resourceRepo interface {
	findByIdAndCreatedBy(tx *gorm.DB, id uint, createdBy uint) (resource, error)
	findAllByCreatedBy(tx *gorm.DB, page dto.Pageable, createdBy uint) (dto.Slice[resource], error)
	save(tx *gorm.DB, ent *resource) error
	delete(tx *gorm.DB, ent resource) error
}

type resourceRepoImpl struct{}

func NewResourceRepo() resourceRepo {
	return &resourceRepoImpl{}
}

func (r *resourceRepoImpl) findByIdAndCreatedBy(tx *gorm.DB, id uint, createdBy uint) (resource, error) {
	ent := resource{ID: id}

	if err := tx.Where(&resource{CreatedBy: createdBy}).First(&ent).Error; err != nil {
		return resource{}, err
	}
	return ent, nil
}

func (r *resourceRepoImpl) findAllByCreatedBy(tx *gorm.DB, page dto.Pageable, createdBy uint) (dto.Slice[resource], error) {
	var ents []resource

	query := tx.Where(&resource{CreatedBy: createdBy}).Limit(page.Limit() + 1).Offset(page.Offset())
	for _, sort := range page.Sort {
		query = query.Order(fmt.Sprintf("%s %s", sort.Property, sort.Direction))
	}

	if err := query.Find(&ents).Error; err != nil {
		return dto.Slice[resource]{}, err
	}

	return dto.NewSlice(ents, page, len(ents)), nil
}

func (r *resourceRepoImpl) save(tx *gorm.DB, ent *resource) error {
	if ent.ID == 0 {
		return tx.Create(ent).Error
	}

	result := tx.Model(ent).Updates(ent)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrOptimisticLockFailure
	}

	return nil
}

func (r *resourceRepoImpl) delete(tx *gorm.DB, ent resource) error {
	result := tx.Model(&ent).Delete(&ent)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrOptimisticLockFailure
	}

	return nil
}
