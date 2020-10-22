package time

import (
	"github.com/Shopify/go-lua"
	"time"
)

func Open(l *lua.State) {
	timeOpen := func(l *lua.State) int {
		lua.NewLibrary(l, timeLibrary)
		return 1
	}
	lua.Require(l, "goluago/time", timeOpen, false)
	l.Pop(1)
}

var timeLibrary = []lua.RegistryFunction{
	{"now", now},
	{"format", format},
	{"sleep", sleep},
	{"since", since},
	{"add", add},
}

func format(l *lua.State) int {
	epochNanoToFormat := int64(lua.CheckNumber(l, 1))
	layout := lua.CheckString(l, 2)
	zone := lua.CheckString(l, 3)

	loc, err := time.LoadLocation(zone)
	if err != nil {
		lua.Errorf(l, err.Error())
	}

	unixTime := time.Unix(0, epochNanoToFormat)
	timeInTimeZone := unixTime.In(loc)
	l.PushString(timeInTimeZone.Format(layout))

	return 1
}

func add(l *lua.State) int {
	start := int64(lua.CheckNumber(l, 1))
	startUnix := time.Unix(0, start)
	valueToAdd := lua.CheckMap(l, 2)

	inc := startUnix.Add((time.Hour * time.Duration(valueToAdd["hour"])) + (time.Minute * time.Duration(valueToAdd["minute"])) + (time.Second * time.Duration(valueToAdd["second"])))
	l.PushNumber(float64(inc.UnixNano()))
	return 1
}

func sleep(l *lua.State) int {
	ns := lua.CheckInteger(l, 1)
	time.Sleep(time.Nanosecond * time.Duration(ns))
	return 1
}

func now(l *lua.State) int {
	l.PushNumber(float64(time.Now().UnixNano()))
	return 1
}

func since(l *lua.State) int {
	start := lua.CheckNumber(l, 1)
	diff := float64(time.Now().UnixNano()) - start
	l.PushNumber(diff)
	return 1
}
