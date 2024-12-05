package ex01

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	exerciseerrors "iaindavis.dev/learning/learning-go/ch9/exercises/ex01/errors"
	exercisesource "iaindavis.dev/learning/learning-go/ch9/exercises/exercisesource"
)

type Employee exercisesource.Employee
type ErrInvalid exerciseerrors.ErrInvalid

const data = exercisesource.Data

var validID = exercisesource.ValidID

func ValidateEmployee(e Employee) error {
	validationErrors := make([]error, 0, 5)
	if len(e.ID) == 0 {
		// return exerciseerrors.ErrMissingID
		validationErrors = append(validationErrors, exerciseerrors.ErrMissingID)
	}
	if !validID.MatchString(e.ID) {
		validationErrors = append(validationErrors, exerciseerrors.ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		validationErrors = append(validationErrors, exerciseerrors.ErrMissingFirstName)
	}
	if len(e.LastName) == 0 {
		validationErrors = append(validationErrors, exerciseerrors.ErrMissingLastName)
	}
	if len(e.Title) == 0 {
		validationErrors = append(validationErrors, exerciseerrors.ErrMissingTitle)
	}
	if len(validationErrors) > 0 {
		return errors.Join(validationErrors...)
	}
	return nil
}

/*
Create a sentinel error to represent an invalid ID. In main, use
`errors.Is` to check for the sentinel error, and print a message when
it's found
*/
func Exercise_9_1() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			var validationError exerciseerrors.ErrInvalid
			if errors.As(err, &validationError) {
				fmt.Println(validationError.Error())
			}
			fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}
