package app

import "github.com/google/uuid"

func GetLastSliceElement[T interface{}](slice []T) *T {
	return &slice[len(slice)-1]
}

func CreateAnonymousUsername() string {
	return "Anon-" + uuid.New().String()[0:5]
}
