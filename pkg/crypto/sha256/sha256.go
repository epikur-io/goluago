package sha256

import (
	"crypto/sha256"

	"github.com/epikur-io/go-lua"
)

func Open(l *lua.State) {
	open := func(l *lua.State) int {
		lua.NewLibrary(l, library)
		return 1
	}
	lua.Require(l, "goluago/crypto/sha256", open, false)
	l.Pop(1)
}

var library = []lua.RegistryFunction{
	{"digest", digest},
}

func digest(l *lua.State) int {
	message := lua.CheckString(l, 1)

	h := sha256.New()
	h.Write([]byte(message))

	l.PushString(string(h.Sum(nil)))
	return 1
}
