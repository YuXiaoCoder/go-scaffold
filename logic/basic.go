package logic

import (
	"fmt"
	"reflect"
)

// deepGetFields 递归获取结构体字段
func deepGetFields(ifaceType reflect.Type) (fields []reflect.StructField) {
	fields = make([]reflect.StructField, 0)

	for i := 0; i < ifaceType.NumField(); i++ {
		field := ifaceType.Field(i)
		// 判断是否为内嵌结构体
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			fields = append(fields, deepGetFields(field.Type)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

// copyStruct 拷贝结构体
func copyStruct(srcPtr, dstPtr interface{}) (err error) {
	srcT := reflect.TypeOf(srcPtr)  // 获取源对象的反射类型对象
	srcV := reflect.ValueOf(srcPtr) // 获取源对象的原始值对象
	dstT := reflect.TypeOf(dstPtr)  // 获取目标对象的反射类型对象
	dstV := reflect.ValueOf(dstPtr) // 获取目标对象的原始值对象

	// 参数必须是结构体的指针
	if srcT.Kind() != reflect.Ptr || dstT.Kind() != reflect.Ptr || srcT.Elem().Kind() != reflect.Struct || dstT.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("params must be a struct pointer")
	}

	// 参数不能为空指针
	if srcV.IsNil() || dstV.IsNil() {
		return fmt.Errorf("params must not be nil")
	}

	srcE := srcV.Elem()  // 获取指针对应的值
	destE := dstV.Elem() // 获取指针对应的值

	// 获取所有的字段
	srcFields := deepGetFields(srcE.Type())
	for _, v := range srcFields {
		// 若为匿名字段，则返回
		if v.Anonymous {
			continue
		}

		srcValue := srcE.FieldByName(v.Name)
		destValue := destE.FieldByName(v.Name)

		// 若不存在此字段，则返回
		if !destValue.IsValid() {
			continue
		}

		// 若类型一致且允许修改，则设置
		if srcValue.Type() == destValue.Type() && destValue.CanSet() {
			destValue.Set(srcValue)
			continue
		}

		// 若指针对应的值的类型与目标字段类型一致，则设置
		if srcValue.Kind() == reflect.Ptr && !srcValue.IsNil() && srcValue.Type().Elem() == destValue.Type() {
			destValue.Set(srcValue.Elem())
			continue
		}

		// 若指针对应的值的类型与源字段类型一致，则设置
		if destValue.Kind() == reflect.Ptr && srcValue.Type() == destValue.Type().Elem() {
			destValue.Set(reflect.New(srcValue.Type()))
			destValue.Elem().Set(srcValue)
			continue
		}
	}
	return nil
}
