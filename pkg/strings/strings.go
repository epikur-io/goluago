package strings

import (
	"strings"

	"github.com/epikur-io/go-lua"
	"github.com/epikur-io/goluago/util"
)

func Open(l *lua.State) {
	strOpen := func(l *lua.State) int {
		lua.NewLibrary(l, stringLibrary)
		return 1
	}
	lua.Require(l, "goluago/strings", strOpen, false)
	l.Pop(1)
}

var stringLibrary = []lua.RegistryFunction{
	{"split", split},
	{"trim", trim},
	{"replace", replace},
}

func split(l *lua.State) int {
	str := lua.CheckString(l, 1)
	sep := lua.CheckString(l, 2)

	strArr := strings.Split(str, sep)

	return util.DeepPush(l, strArr)
}

func trim(l *lua.State) int {
	str := lua.CheckString(l, 1)
	l.PushString(strings.TrimSpace(str))
	return 1
}

func replace(l *lua.State) int {
	s := lua.CheckString(l, 1)
	old := lua.CheckString(l, 2)
	new := lua.CheckString(l, 3)
	n := lua.CheckInteger(l, 4)

	l.PushString(strings.Replace(s, old, new, n))
	return 1
}
