package api

type MongodbError struct {
	message string
}

func (m MongodbError) Error() string {
	return m.message
}

var (
	ErrUserExists = &MongodbError{"require:Create account user is already exists"}
)
