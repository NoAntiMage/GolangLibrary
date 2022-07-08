package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type BaseReponse struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Flag   bool   `json:"flag"`
}

func main() {
	m := map[string]interface{}{"msg": "helloWujimaster", "status": 1, "flag": true}
	fmt.Println(m)

	s, _ := Map2Json(m)
	fmt.Println(s)

	var br BaseReponse
	Json2Struct(&br, s)
	fmt.Println(br)

	s, _ = Struct2Json(br)
	fmt.Println(s)

	var br2 BaseReponse
	Map2Struct(&br2, m)
	fmt.Println(br2)

	var br3 BaseReponse = BaseReponse{Msg: "helloworld", Status: 2, Flag: false}
	m3 := Struct2Map(br3)
	fmt.Println(m3)

	var br4 BaseReponse = BaseReponse{Msg: "greeting", Status: 3, Flag: true}
	m4 := Struct2MapV2(br4)
	fmt.Println(m4)
}

func Json2Map(jsonStr string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return m, nil
}

func Map2Json(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonByte), nil
}

// C style: return Pstructer
func Json2Struct(Pstructer interface{}, jsonStr string) {
	err := json.Unmarshal([]byte(jsonStr), Pstructer)
	if err != nil {
		fmt.Println(err)
	}
}

func Struct2Json(structer interface{}) (string, error) {
	jsonBytes, err := json.Marshal(structer)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonBytes), err
}

func Map2Struct(Pstructer interface{}, m map[string]interface{}) error {
	s, err := Map2Json(m)
	if err != nil {
		return err
	}

	Json2Struct(Pstructer, s)
	return nil
}

func Struct2Map(structer interface{}) map[string]interface{} {
	s, err := Struct2Json(structer)
	if err != nil {
		fmt.Println(err)
	}

	m, err := Json2Map(s)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

// the json comment in struct is ignored
func Struct2MapV2(structer interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	t := reflect.TypeOf(structer)
	v := reflect.ValueOf(structer)

	fieldNum := t.NumField()

	for i := 0; i < fieldNum; i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}
