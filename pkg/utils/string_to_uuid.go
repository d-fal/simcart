package utils

import (
	"unsafe"

	"github.com/google/uuid"
)

func StringToUUID(target string) (*uuid.UUID, error) {
	to, err := uuid.Parse(target)
	if err != nil {
		return nil, err
	}
	return &to, nil
}

func StringToByte(str string) []byte {

	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

func StringToPointer(str string) *string {
	return (*string)(unsafe.Pointer(&str))
}
