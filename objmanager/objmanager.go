package objmanager

import (
	"sync"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

// IntrospectData The introspection data for the org.freedesktop.DBus.ObjectManager interface.
var IntrospectData = introspect.Interface{
	Name: "org.freedesktop.DBus.ObjectManager",
	Methods: []introspect.Method{
		{
			Name: "GetManagedObjects",
			Args: []introspect.Arg{
				{"objpath_interfaces_and_properties", "e{oe{se{sv}}}", "in"},
			},
		},
	},
	Signals: []introspect.Signal{
		{
			Name: "InterfacesAdded",
			Args: []introspect.Arg{
				{"object_path", "o", "out"},
				{"interfaces_and_properties", "e{se{sv}}", "out"},
			},
		},
		{
			Name: "InterfacesRemoved",
			Args: []introspect.Arg{
				{"object_path", "o", "out"},
				{"interfaces", "as", "out"},
			},
		},
	},
}

// IntrospectDataString The introspection data for the org.freedesktop.DBus.ObjectManager
// interface, as  a string.
const IntrospectDataString = `
	<interface name="org.freedesktop.DBus.ObjectManager">
		<method name="GetManagedObjects">
			<arg name="objpath_interfaces_and_properties" direction="in" type="e{oe{se{sv}}}"/>
		</method>
		<signal name="InterfacesAdded">
			<arg name="object_path" type="s"/>
			<arg name="interfaces_and_properties" type="e{se{sv}}"/>
		</signal>
		<signal name="InterfacesRemoved">
			<arg name="object_path" type="o"/>
			<arg name="interfaces" type="as"/>
		</signal>
	</interface>
`

// Object an object retrieved
type Object struct {
}

// ObjectManager allows call to retrieve list of objects
type ObjectManager struct {
	m    map[string]map[string]*Object
	mut  sync.RWMutex
	conn *dbus.Conn
	path dbus.ObjectPath
}

// New returns a new ObjectManager structure that intercat with the D-Bus interface.
func New(conn *dbus.Conn, path dbus.ObjectPath) *ObjectManager {
	objman := &ObjectManager{conn: conn, path: path}
	conn.Export(objman, path, "org.freedesktop.DBus.ObjectManager")
	return objman
}

// GetManagedObjects implements org.freedesktop.DBus.ObjectManager.GetManagedObjects.
// func (objman *ObjectManager) GetManagedObjects() (dbus.Variant, *dbus.Error) {}
