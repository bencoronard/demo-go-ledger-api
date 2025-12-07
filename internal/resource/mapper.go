package resource

import "time"

func toDTO(r resource) resourceDTO {
	return resourceDTO{
		ID:           r.ID,
		Version:      r.Version,
		TextField:    r.TextField,
		NumberField:  r.NumberField,
		BooleanField: r.BooleanField,
	}
}

func toEntity(dto resourceDTO, createdBy int64, createdAt, lastUpdated time.Time) resource {
	return resource{
		ID:           dto.ID,
		Version:      dto.Version,
		CreatedBy:    createdBy,
		CreatedAt:    createdAt,
		LastUpdated:  lastUpdated,
		TextField:    dto.TextField,
		NumberField:  dto.NumberField,
		BooleanField: dto.BooleanField,
	}
}
