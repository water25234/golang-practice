package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	type User struct {
		id   string
		name string
	}
	var executors = make([]*User, 2)

	executor := &User{}
	setUnExportedStrField(executor, "id", "123456")
	setUnExportedStrField(executor, "name", "IT")
	executors[0] = executor

	executor = &User{}
	setUnExportedStrField(executor, "id", "654321")
	setUnExportedStrField(executor, "name", "MK")
	executors[1] = executor

	fmt.Println(executors)

	// var executors = make([]*aEntity.Info, 2)

	// executor := &aEntity.Info{}
	// setUnExportedStrField(executor, "id", "123456")
	// setUnExportedStrField(executor, "name", "IT")
	// setUnExportedStrField(executor, "userID", "qwe431")
	// setUnExportedStrField(executor, "authorizeTimeMs", "1234567890")
	// executors[0] = executor

	// executor = &aEntity.Info{}
	// setUnExportedStrField(executor, "id", "654321")
	// setUnExportedStrField(executor, "name", "CP")
	// setUnExportedStrField(executor, "userID", "zxc345")
	// setUnExportedStrField(executor, "authorizeTimeMs", "0987654321")
	// executors[1] = executor

	// return executors
}

func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

func setUnExportedStrField(ptr interface{}, fieldName string, newFieldVal interface{}) (err error) {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)

	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()

	nv := reflect.ValueOf(newFieldVal)
	if v.Kind() != nv.Kind() {
		return fmt.Errorf("expected kind %v, got kind: %v", v.Kind(), nv.Kind())
	}
	v.Set(nv)
	return nil
}
