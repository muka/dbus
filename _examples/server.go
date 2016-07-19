package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"os"
)

const (
	intfName = "com.github.guelfey.Demo"
	name     = "com.github.guelfey.Demo"
	path     = "/com/github/guelfey/Demo"
)

type foo string

func (f foo) Foo() (string, *dbus.Error) {
	fmt.Println(f)
	return string(f), nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	reply, err := conn.RequestName("com.github.guelfey.Demo",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}

	f := foo("Bar!")

	conn.Export(f, path, intfName)
	conn.Export(introspect.NewIntrospectable(&introspect.Node{
		Name: path,
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			{
				Name:    intfName,
				Methods: introspect.Methods(f),
			},
		},
	}), path,
		"org.freedesktop.DBus.Introspectable")

	fmt.Printf("Listening on %s / %s ...", intfName, path)
	select {}
}
