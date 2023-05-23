package unsafe

import (
	"reflect"
	"unsafe"
)

type Accessor struct {
	startAddr unsafe.Pointer
	Field     map[string]FieldMeta
}

type FieldMeta struct {
	offset uintptr
	typ    reflect.Type
}

func NewUnsafeAccessor(entity any) *Accessor {
	typ := reflect.TypeOf(entity).Elem()
	numField := typ.NumField()
	fields := make(map[string]FieldMeta, numField)
	for i := 0; i < numField; i++ {
		field := typ.Field(i)
		fields[field.Name] = FieldMeta{
			offset: field.Offset,
			typ:    field.Type,
		}
	}

	val := reflect.ValueOf(entity)

	return &Accessor{
		startAddr: val.UnsafePointer(),
		Field:     fields,
	}
}

func (a *Accessor) GetField(name string) any {

	addr := unsafe.Pointer(uintptr(a.startAddr) + a.Field[name].offset)
	return reflect.NewAt(a.Field[name].typ, addr).Elem().Interface()

}
