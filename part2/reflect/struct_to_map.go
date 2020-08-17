/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/8/5 6:12 下午
*/

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
	Age int64 `json:"age"`
}

func StructToMapTest() {

	c := new(Cat)
	c.Name = "leo"
	c.Age = 21

	result := StructToMap(*c)

	fmt.Println(result)
	jjj,_ := json.Marshal(result)
	fmt.Println(string(jjj))
}

// c ：如果一个 struct 指针，必须通过 * 取改地址内存区的实际保存值
func StructToMap(c interface{}) map[string]interface{} {
	ref := reflect.ValueOf(c)

	fmt.Println(ref)

	t := ref.Type()

	result := make(map[string]interface{})

	for i:=0; i<t.NumField(); i++ {
		n := t.Field(i).Name
		fmt.Println(n)

		v := ref.FieldByName(n)
		fmt.Println(v.Type().Name(), v)

		result[strings.ToLower(n)] = v.Interface()
	}
	return result
}