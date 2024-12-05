package exerciseerrors

import "fmt"

// enum
type InvalidityType string

const MissingField InvalidityType = "missing required field"
const InvalidField InvalidityType = "invalid value for field"

// struct
type ErrInvalid struct {
	field       string
	invalidType InvalidityType
}

// method
func (ei ErrInvalid) Error() string {
	return fmt.Sprintf("Validation error -- %s %s", ei.invalidType, ei.field)
}

// factory
func newErrInvalid(it InvalidityType, field string) ErrInvalid {
	return ErrInvalid{
		invalidType: it,
		field:       field,
	}
}
