package exerciseerrors

var ErrInvalidID ErrInvalid = newErrInvalid(InvalidField, "Employee.ID")
var ErrMissingID ErrInvalid = newErrInvalid(MissingField, "Employee.ID")
var ErrMissingFirstName ErrInvalid = newErrInvalid(MissingField, "Employee.FirstName")
var ErrMissingLastName ErrInvalid = newErrInvalid(MissingField, "Employee.LastName")
var ErrMissingTitle ErrInvalid = newErrInvalid(MissingField, "Employee.Title")
