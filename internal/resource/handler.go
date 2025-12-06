package resource

import (
	"net/http"
)

type ResourceHandler struct {
	s resourceService
}

func NewResourceHandler(s resourceService) *ResourceHandler {
	return &ResourceHandler{s: s}
}

func (h *ResourceHandler) ListResources(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func (h *ResourceHandler) CreateResource(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (h *ResourceHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (h *ResourceHandler) RetrieveResource(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (h *ResourceHandler) UpdateResource(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// dtos := make([]*resourceDTO, len(ents.Content))
// 	for i, r := range ents.Content {
// 		dtos[i] = &resourceDTO{
// 			ID:           r.ID,
// 			Version:      r.Version,
// 			TextField:    r.TextField,
// 			NumberField:  r.NumberField,
// 			BooleanField: r.BooleanField,
// 		}
// 	}

// 	return &dto.Slice[*resourceDTO]{
// 		Content:    dtos,
// 		TotalCount: ents.TotalCount,
// 		HasNext:    ents.HasNext,
// 	}, nil
