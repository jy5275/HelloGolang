package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// MarshalTest 就是把 struct 转为 json 格式的字符串 (好吧其
// 实是[]byte), struct 中 Header 字段特殊处理了一下也不知道为啥
func MarshalTest() {
	str := `{"precomputed": true}`
	by := []byte(str)
	fmt.Printf("%s\n", by)
	h := json.RawMessage(str)
	c := struct {
		Header *json.RawMessage //`json:"header"`
		Body   string           //`json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	// 此步 c 作为 struct 被转为 json 格式字符串
	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
}

func UnmarshalTest() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}
	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)

	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		// 此步 c.Point 中的 RawMessage 被转为 struct YCbCr dst
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
