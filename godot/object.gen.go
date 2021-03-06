package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

// ObjectConnectFlags is an enum for ConnectFlags values.
type ObjectConnectFlags int

const (
	ObjectConnectDeferred ObjectConnectFlags = 1
	ObjectConnectOneshot  ObjectConnectFlags = 4
	ObjectConnectPersist  ObjectConnectFlags = 2
)

//func NewObjectFromPointer(ptr gdnative.Pointer) Object {
func newObjectFromPointer(ptr gdnative.Pointer) Object {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Object{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class for all non built-in types. Everything not a built-in type starts the inheritance chain from this class. Objects do not manage memory, if inheriting from one the object will most likely have to be deleted manually (call the [method free] function from the script or delete from C++). Some derivatives add memory management, such as [Reference] (which keeps a reference count and deletes itself automatically when no longer referenced) and [Node], which deletes the children tree when deleted. Objects export properties, which are mainly useful for storage and editing, but not really so much in programming. Properties are exported in [method _get_property_list] and handled in [method _get] and [method _set]. However, scripting languages and C++ have simpler means to export them. Objects also receive notifications ([method _notification]). Notifications are a simple way to notify the object about simple events, so they can all be handled together.
*/
type Object struct {
	owner gdnative.Object
}

func (o *Object) BaseClass() string {
	return "Object"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Object) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *Object) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Returns the given property. Returns [code]null[/code] if the [code]property[/code] does not exist.
	Args: [{ false property String}], Returns: void
*/
func (o *Object) X_Get(property gdnative.String) {
	//log.Println("Calling Object.X_Get()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "_get")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the object's property list as an [Array] of dictionaries. Dictionaries must contain: name:String, type:int (see TYPE_* enum in [@GlobalScope]) and optionally: hint:int (see PROPERTY_HINT_* in [@GlobalScope]), hint_string:String, usage:int (see PROPERTY_USAGE_* in [@GlobalScope]).
	Args: [], Returns: Array
*/
func (o *Object) X_GetPropertyList() gdnative.Array {
	//log.Println("Calling Object.X_GetPropertyList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "_get_property_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        The virtual method called upon initialization.
	Args: [], Returns: void
*/
func (o *Object) X_Init() {
	//log.Println("Calling Object.X_Init()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "_init")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Notify the object internally using an ID.
	Args: [{ false what int}], Returns: void
*/
func (o *Object) X_Notification(what gdnative.Int) {
	//log.Println("Calling Object.X_Notification()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(what)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "_notification")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets a property. Returns [code]true[/code] if the [code]property[/code] exists.
	Args: [{ false property String} { false value Variant}], Returns: bool
*/
func (o *Object) X_Set(property gdnative.String, value gdnative.Variant) gdnative.Bool {
	//log.Println("Calling Object.X_Set()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(property)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "_set")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Adds a user-defined [code]signal[/code]. Arguments are optional, but can be added as an [Array] of dictionaries, each containing "name" and "type" (from [@GlobalScope] TYPE_*).
	Args: [{ false signal String} {[] true arguments Array}], Returns: void
*/
func (o *Object) AddUserSignal(signal gdnative.String, arguments gdnative.Array) {
	//log.Println("Calling Object.AddUserSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)
	ptrArguments[1] = gdnative.NewPointerFromArray(arguments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "add_user_signal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Calls the [code]method[/code] on the object and returns a result. Pass parameters as a comma separated list.
	Args: [{ false method String}], Returns: Variant
*/
func (o *Object) Call(method gdnative.String) gdnative.Variant {
	//log.Println("Calling Object.Call()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "call")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Calls the [code]method[/code] on the object during idle time and returns a result. Pass parameters as a comma separated list.
	Args: [{ false method String}], Returns: Variant
*/
func (o *Object) CallDeferred(method gdnative.String) gdnative.Variant {
	//log.Println("Calling Object.CallDeferred()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "call_deferred")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Calls the [code]method[/code] on the object and returns a result. Pass parameters as an [Array].
	Args: [{ false method String} { false arg_array Array}], Returns: Variant
*/
func (o *Object) Callv(method gdnative.String, argArray gdnative.Array) gdnative.Variant {
	//log.Println("Calling Object.Callv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(method)
	ptrArguments[1] = gdnative.NewPointerFromArray(argArray)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "callv")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the object can translate strings.
	Args: [], Returns: bool
*/
func (o *Object) CanTranslateMessages() gdnative.Bool {
	//log.Println("Calling Object.CanTranslateMessages()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "can_translate_messages")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Connects a [code]signal[/code] to a [code]method[/code] on a [code]target[/code] object. Pass optional [code]binds[/code] to the call. Use [code]flags[/code] to set deferred or one shot connections. See [code]CONNECT_*[/code] constants. A [code]signal[/code] can only be connected once to a [code]method[/code]. It will throw an error if already connected. To avoid this, first use [method is_connected] to check for existing connections.
	Args: [{ false signal String} { false target Object} { false method String} {[] true binds Array} {0 true flags int}], Returns: enum.Error
*/
func (o *Object) Connect(signal gdnative.String, target ObjectImplementer, method gdnative.String, binds gdnative.Array, flags gdnative.Int) gdnative.Error {
	//log.Println("Calling Object.Connect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)
	ptrArguments[1] = gdnative.NewPointerFromObject(target.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromString(method)
	ptrArguments[3] = gdnative.NewPointerFromArray(binds)
	ptrArguments[4] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "connect")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Disconnects a [code]signal[/code] from a [code]method[/code] on the given [code]target[/code].
	Args: [{ false signal String} { false target Object} { false method String}], Returns: void
*/
func (o *Object) Disconnect(signal gdnative.String, target ObjectImplementer, method gdnative.String) {
	//log.Println("Calling Object.Disconnect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)
	ptrArguments[1] = gdnative.NewPointerFromObject(target.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "disconnect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Emits the given [code]signal[/code].
	Args: [{ false signal String}], Returns: Variant
*/
func (o *Object) EmitSignal(signal gdnative.String) gdnative.Variant {
	//log.Println("Calling Object.EmitSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "emit_signal")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Deletes the object from memory.
	Args: [], Returns: void
*/
func (o *Object) Free() {
	//log.Println("Calling Object.Free()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "free")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns a [Variant] for a [code]property[/code].
	Args: [{ false property String}], Returns: Variant
*/
func (o *Object) Get(property gdnative.String) gdnative.Variant {
	//log.Println("Calling Object.Get()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the object's class as a [String].
	Args: [], Returns: String
*/
func (o *Object) GetClass() gdnative.String {
	//log.Println("Calling Object.GetClass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_class")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns an [Array] of dictionaries with information about signals that are connected to the object. Inside each [Dictionary] there are 3 fields: - "source" is a reference to signal emitter. - "signal_name" is name of connected signal. - "method_name" is a name of method to which signal is connected.
	Args: [], Returns: Array
*/
func (o *Object) GetIncomingConnections() gdnative.Array {
	//log.Println("Calling Object.GetIncomingConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_incoming_connections")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false property NodePath}], Returns: Variant
*/
func (o *Object) GetIndexed(property gdnative.NodePath) gdnative.Variant {
	//log.Println("Calling Object.GetIndexed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_indexed")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the object's unique instance ID.
	Args: [], Returns: int
*/
func (o *Object) GetInstanceId() gdnative.Int {
	//log.Println("Calling Object.GetInstanceId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_instance_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the object's metadata for the given [code]name[/code].
	Args: [{ false name String}], Returns: Variant
*/
func (o *Object) GetMeta(name gdnative.String) gdnative.Variant {
	//log.Println("Calling Object.GetMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_meta")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the object's metadata as a [PoolStringArray].
	Args: [], Returns: PoolStringArray
*/
func (o *Object) GetMetaList() gdnative.PoolStringArray {
	//log.Println("Calling Object.GetMetaList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_meta_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the object's methods and their signatures as an [Array].
	Args: [], Returns: Array
*/
func (o *Object) GetMethodList() gdnative.Array {
	//log.Println("Calling Object.GetMethodList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_method_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the list of properties as an [Array] of dictionaries. Dictionaries contain: name:String, type:int (see TYPE_* enum in [@GlobalScope]) and optionally: hint:int (see PROPERTY_HINT_* in [@GlobalScope]), hint_string:String, usage:int (see PROPERTY_USAGE_* in [@GlobalScope]).
	Args: [], Returns: Array
*/
func (o *Object) GetPropertyList() gdnative.Array {
	//log.Println("Calling Object.GetPropertyList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_property_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the object's [Script] or [code]null[/code] if one doesn't exist.
	Args: [], Returns: Reference
*/
func (o *Object) GetScript() ReferenceImplementer {
	//log.Println("Calling Object.GetScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_script")

	// Call the parent method.
	// Reference
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newReferenceFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ReferenceImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Reference" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ReferenceImplementer)
	}

	return &ret
}

/*
        Returns an [Array] of connections for the given [code]signal[/code].
	Args: [{ false signal String}], Returns: Array
*/
func (o *Object) GetSignalConnectionList(signal gdnative.String) gdnative.Array {
	//log.Println("Calling Object.GetSignalConnectionList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_signal_connection_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the list of signals as an [Array] of dictionaries.
	Args: [], Returns: Array
*/
func (o *Object) GetSignalList() gdnative.Array {
	//log.Println("Calling Object.GetSignalList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "get_signal_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if a metadata is found with the given [code]name[/code].
	Args: [{ false name String}], Returns: bool
*/
func (o *Object) HasMeta(name gdnative.String) gdnative.Bool {
	//log.Println("Calling Object.HasMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "has_meta")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the object contains the given [code]method[/code].
	Args: [{ false method String}], Returns: bool
*/
func (o *Object) HasMethod(method gdnative.String) gdnative.Bool {
	//log.Println("Calling Object.HasMethod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "has_method")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the given user-defined [code]signal[/code] exists.
	Args: [{ false signal String}], Returns: bool
*/
func (o *Object) HasUserSignal(signal gdnative.String) gdnative.Bool {
	//log.Println("Calling Object.HasUserSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "has_user_signal")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if signal emission blocking is enabled.
	Args: [], Returns: bool
*/
func (o *Object) IsBlockingSignals() gdnative.Bool {
	//log.Println("Calling Object.IsBlockingSignals()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "is_blocking_signals")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the object inherits from the given [code]type[/code].
	Args: [{ false type String}], Returns: bool
*/
func (o *Object) IsClass(aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Object.IsClass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "is_class")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if a connection exists for a given [code]signal[/code], [code]target[/code], and [code]method[/code].
	Args: [{ false signal String} { false target Object} { false method String}], Returns: bool
*/
func (o *Object) IsConnected(signal gdnative.String, target ObjectImplementer, method gdnative.String) gdnative.Bool {
	//log.Println("Calling Object.IsConnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)
	ptrArguments[1] = gdnative.NewPointerFromObject(target.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "is_connected")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the [code]queue_free[/code] method was called for the object.
	Args: [], Returns: bool
*/
func (o *Object) IsQueuedForDeletion() gdnative.Bool {
	//log.Println("Calling Object.IsQueuedForDeletion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "is_queued_for_deletion")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Notify the object of something.
	Args: [{ false what int} {False true reversed bool}], Returns: void
*/
func (o *Object) Notification(what gdnative.Int, reversed gdnative.Bool) {
	//log.Println("Calling Object.Notification()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(what)
	ptrArguments[1] = gdnative.NewPointerFromBool(reversed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "notification")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *Object) PropertyListChangedNotify() {
	//log.Println("Calling Object.PropertyListChangedNotify()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "property_list_changed_notify")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set property into the object.
	Args: [{ false property String} { false value Variant}], Returns: void
*/
func (o *Object) Set(property gdnative.String, value gdnative.Variant) {
	//log.Println("Calling Object.Set()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(property)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If set to true, signal emission is blocked.
	Args: [{ false enable bool}], Returns: void
*/
func (o *Object) SetBlockSignals(enable gdnative.Bool) {
	//log.Println("Calling Object.SetBlockSignals()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set_block_signals")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false property NodePath} { false value Variant}], Returns: void
*/
func (o *Object) SetIndexed(property gdnative.NodePath, value gdnative.Variant) {
	//log.Println("Calling Object.SetIndexed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(property)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set_indexed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Define whether the object can translate strings (with calls to [method tr]). Default is true.
	Args: [{ false enable bool}], Returns: void
*/
func (o *Object) SetMessageTranslation(enable gdnative.Bool) {
	//log.Println("Calling Object.SetMessageTranslation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set_message_translation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set a metadata into the object. Metadata is serialized. Metadata can be [i]anything[/i].
	Args: [{ false name String} { false value Variant}], Returns: void
*/
func (o *Object) SetMeta(name gdnative.String, value gdnative.Variant) {
	//log.Println("Calling Object.SetMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set_meta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set a script into the object, scripts extend the object functionality.
	Args: [{ false script Reference}], Returns: void
*/
func (o *Object) SetScript(script ReferenceImplementer) {
	//log.Println("Calling Object.SetScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(script.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "set_script")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Translate a message. Only works if message translation is enabled (which it is by default). See [method set_message_translation].
	Args: [{ false message String}], Returns: String
*/
func (o *Object) Tr(message gdnative.String) gdnative.String {
	//log.Println("Calling Object.Tr()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(message)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Object", "tr")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

// ObjectImplementer is an interface that implements the methods
// of the Object class.
type ObjectImplementer interface {
	Class
	X_Get(property gdnative.String)
	X_GetPropertyList() gdnative.Array
	X_Init()
	X_Notification(what gdnative.Int)
	X_Set(property gdnative.String, value gdnative.Variant) gdnative.Bool
	AddUserSignal(signal gdnative.String, arguments gdnative.Array)
	Call(method gdnative.String) gdnative.Variant
	CallDeferred(method gdnative.String) gdnative.Variant
	Callv(method gdnative.String, argArray gdnative.Array) gdnative.Variant
	CanTranslateMessages() gdnative.Bool
	Disconnect(signal gdnative.String, target ObjectImplementer, method gdnative.String)
	EmitSignal(signal gdnative.String) gdnative.Variant
	Free()
	Get(property gdnative.String) gdnative.Variant
	GetClass() gdnative.String
	GetIncomingConnections() gdnative.Array
	GetIndexed(property gdnative.NodePath) gdnative.Variant
	GetInstanceId() gdnative.Int
	GetMeta(name gdnative.String) gdnative.Variant
	GetMetaList() gdnative.PoolStringArray
	GetMethodList() gdnative.Array
	GetPropertyList() gdnative.Array
	GetScript() ReferenceImplementer
	GetSignalConnectionList(signal gdnative.String) gdnative.Array
	GetSignalList() gdnative.Array
	HasMeta(name gdnative.String) gdnative.Bool
	HasMethod(method gdnative.String) gdnative.Bool
	HasUserSignal(signal gdnative.String) gdnative.Bool
	IsBlockingSignals() gdnative.Bool
	IsClass(aType gdnative.String) gdnative.Bool
	IsConnected(signal gdnative.String, target ObjectImplementer, method gdnative.String) gdnative.Bool
	IsQueuedForDeletion() gdnative.Bool
	Notification(what gdnative.Int, reversed gdnative.Bool)
	PropertyListChangedNotify()
	Set(property gdnative.String, value gdnative.Variant)
	SetBlockSignals(enable gdnative.Bool)
	SetIndexed(property gdnative.NodePath, value gdnative.Variant)
	SetMessageTranslation(enable gdnative.Bool)
	SetMeta(name gdnative.String, value gdnative.Variant)
	SetScript(script ReferenceImplementer)
	Tr(message gdnative.String) gdnative.String
}
