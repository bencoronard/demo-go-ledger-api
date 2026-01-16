package resource

import (
	"time"

	"gorm.io/plugin/optimisticlock"
)

func toDTO(r resource) resourceDTO {
	return resourceDTO{
		ID:           r.ID,
		Version:      int(r.Version.Int64),
		TextField:    r.TextField,
		NumberField:  r.NumberField,
		BooleanField: r.BooleanField,
	}
}

func toEntity(dto resourceDTO, createdBy uint, createdAt, lastUpdated time.Time) resource {
	return resource{
		ID:           dto.ID,
		Version:      optimisticlock.Version{Int64: int64(dto.Version), Valid: true},
		CreatedBy:    createdBy,
		CreatedAt:    createdAt,
		LastUpdated:  lastUpdated,
		TextField:    dto.TextField,
		NumberField:  dto.NumberField,
		BooleanField: dto.BooleanField,
	}
}
