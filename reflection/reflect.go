package reflection

import (
	"errors"
	"reflect"
	"testing"
	"unsafe"
)

// Gets a pointer of a private or public field in a testing.M struct
func GetFieldPointerOfM(m *testing.M, fieldName string) (unsafe.Pointer, error) {
	val := reflect.Indirect(reflect.ValueOf(m))
	member := val.FieldByName(fieldName)
	if member.IsValid() {
		ptrToY := unsafe.Pointer(member.UnsafeAddr())
		return ptrToY, nil
	}
	return nil, errors.New("field can't be retrieved")
}

// Gets a pointer of a private or public field in a testing.T struct
func GetFieldPointerOfT(t *testing.T, fieldName string) (unsafe.Pointer, error) {
	val := reflect.Indirect(reflect.ValueOf(t))
	member := val.FieldByName(fieldName)
	if member.IsValid() {
		ptrToY := unsafe.Pointer(member.UnsafeAddr())
		return ptrToY, nil
	}
	return nil, errors.New("field can't be retrieved")
}

// Gets a pointer of a private or public field in a testing.B struct
func GetFieldPointerOfB(b *testing.B, fieldName string) (unsafe.Pointer, error) {
	val := reflect.Indirect(reflect.ValueOf(b))
	member := val.FieldByName(fieldName)
	if member.IsValid() {
		ptrToY := unsafe.Pointer(member.UnsafeAddr())
		return ptrToY, nil
	}
	return nil, errors.New("result can't be retrieved")
}