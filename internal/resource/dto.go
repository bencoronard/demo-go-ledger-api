package resource

type resourceDTO struct {
	ID           uint    `json:"id"`
	Version      int     `json:"version"`
	TextField    *string `json:"textField"`
	NumberField  *int    `json:"numberField"`
	BooleanField *bool   `json:"booleanField"`
}
