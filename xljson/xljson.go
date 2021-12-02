package xljson

import (
	"github.com/buger/jsonparser"
)

type JsonObject struct {
	Buff []byte
}

func (j JsonObject) GetJsonObject(keys ...string) (JsonObject, error) {
	buff, _, _, err := jsonparser.Get(j.Buff, keys...)
	return JsonObject{buff}, err
}

// GetJsonObjectArray 获取JsonObject对象
func (j JsonObject) GetJsonObjectArray(keys ...string) ([]JsonObject, error) {
	var err error
	var joArray []JsonObject
	_, err = jsonparser.ArrayEach(
		j.Buff,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			joArray = append(joArray, JsonObject{value})
		}, keys...)
	return joArray, err
}

// Get 获取对象
func (j JsonObject) Get(keys ...string) ([]byte, jsonparser.ValueType, int, error) {
	return jsonparser.Get(j.Buff, keys...)
}

// GetString 获取字符串
func (j JsonObject) GetString(keys ...string) (string, error) {
	return jsonparser.GetString(j.Buff, keys...)
}

// GetBoolean 获取bool值
func (j JsonObject) GetBoolean(keys ...string) (bool, error) {
	return jsonparser.GetBoolean(j.Buff, keys...)
}

// GetInt 获取int64值
func (j JsonObject) GetInt(keys ...string) (int64, error) {
	return jsonparser.GetInt(j.Buff, keys...)
}

// GetFloat 获取float值
func (j JsonObject) GetFloat(keys ...string) (float64, error) {
	return jsonparser.GetFloat(j.Buff, keys...)
}

// Each 遍历json中的array
func (j JsonObject) Each(cb func(jo JsonObject), keys ...string) (int, error) {
	return jsonparser.ArrayEach(
		j.Buff,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			cb(JsonObject{value})
		}, keys...)
}
