package main

// interface
type DataStore interface {
	UserNameForID(userId string) (string, bool)
}

// struct
type SimpleDataStore struct {
	userData map[string]string
}

// methods
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// factory
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}
