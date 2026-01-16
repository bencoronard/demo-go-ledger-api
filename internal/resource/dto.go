package resource

type resourceDTO struct {
	ID           uint
	Version      int
	TextField    *string
	NumberField  *int
	BooleanField *bool
}
