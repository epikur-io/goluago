package hex

import (
	"encoding/hex"

	"github.com/epikur-io/go-lua"
)

func Open(l *lua.State) {
	open := func(l *lua.State) int {
		lua.NewLibrary(l, library)
		return 1
	}
	lua.Require(l, "goluago/encoding/hex", open, false)
	l.Pop(1)
}

var library = []lua.RegistryFunction{
	{"encode", encode},
	{"decode", decode},
}

func encode(l *lua.State) int {
	value := lua.CheckString(l, 1)
	encoded := hex.EncodeToString([]byte(value))

	l.PushString(encoded)
	return 1
}

func decode(l *lua.State) int {
	value := lua.CheckString(l, 1)
	decoded, err := hex.DecodeString(value)

	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	l.PushString(string(decoded))
	return 1
}
