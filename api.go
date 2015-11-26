package gluajson

import (
	"encoding/json"

	"github.com/yuin/gopher-lua"
)

var api = map[string]lua.LGFunction{
	"decode": apiDecode,
	"encode": apiEncode,
}

func apiDecode(L *lua.LState) int {
	str := L.CheckString(1)

	var value interface{}
	err := json.Unmarshal([]byte(str), &value)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(fromJSON(L, value))
	return 1
}

func apiEncode(L *lua.LState) int {
	value := L.CheckAny(1)

	indent := ""
	if L.GetTop() == 2 {
		indent = L.CheckString(2)
	}

	var data []byte
	var err error
	visited := make(map[*lua.LTable]bool)

	if indent != "" {
		data, err = toIndentedJSON(value, visited, indent)
	} else {
		data, err = toJSON(value, visited)
	}

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(string(data)))
	return 1
}


