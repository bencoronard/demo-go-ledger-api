package resource

type resourceDTO struct {
	ID           int64
	Version      int64
	TextField    *string
	NumberField  *int
	BooleanField *bool
}
