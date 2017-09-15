package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

const (
	VariantTypeNil = iota

	// atomic types
	VariantTypeBool
	VariantTypeInt
	VariantTypeReal
	VariantTypeString

	// math types
	VariantTypeVector2 // 5
	VariantTypeRect2
	VariantTypeVector3
	VariantTypeTransform2d
	VariantTypePlane
	VariantTypeQuat // 10
	VariantTypeRect3
	VariantTypeBasis
	VariantTypeTransform

	// misc types
	VariantTypeColor
	VariantTypenodePath // 15
	VariantTyperId
	VariantTypeObject
	VariantTypeDictionary
	VariantTypeArray // 20

	// ARRAYS
	VariantTypePoolByteArray
	VariantTypePoolIntArray
	VariantTypePoolRealArray
	VariantTypePoolStringArray
	VariantTypePoolVector2Array // 25
	VariantTypePoolVector3Array
	VariantTypePoolColorArray
)

type Variant struct {
	variant *C.godot_variant
}

//variant type conversions
func variantAsBool(variant *C.godot_variant) bool {
	godotBool := C.godot_variant_as_bool(variant)
	return bool(godotBool)
}

func boolAsVariant(value bool) *C.godot_variant {
	var variant C.godot_variant
	godotBool := boolAsGodotBool(value)
	C.godot_variant_new_bool(&variant, godotBool)

	return &variant
}

func variantAsUInt(variant *C.godot_variant) uint64 {
	godotUInt := C.godot_variant_as_uint(variant)
	return uint64(godotUInt)
}

func uIntAsVariant(value uint64) *C.godot_variant {
	var variant C.godot_variant
	cUInt64 := uIntAsCUInt64(value)
	C.godot_variant_new_uint(&variant, cUInt64)

	return &variant
}

func variantAsInt(variant *C.godot_variant) int64 {
	godotInt := C.godot_variant_as_int(variant)
	return int64(godotInt)
}

func intAsVariant(value int64) *C.godot_variant {
	var variant C.godot_variant
	cInt64 := intAsCInt64(value)
	C.godot_variant_new_int(&variant, cInt64)

	return &variant
}

func variantAsReal(variant *C.godot_variant) float64 {
	godotFloat := C.godot_variant_as_real(variant)
	return float64(godotFloat)
}

func realAsVariant(value float64) *C.godot_variant {
	var variant C.godot_variant
	cDouble := realAsCReal(value)
	C.godot_variant_new_real(&variant, cDouble)

	return &variant
}

func variantAsString(variant *C.godot_variant) string {
	godotString := C.godot_variant_as_string(variant)

	return godotStringAsString(&godotString)
}

func stringAsVariant(value string) *C.godot_variant {
	var variant C.godot_variant
	godotString := stringAsGodotString(value)
	C.godot_variant_new_string(&variant, godotString)

	return &variant
}

func variantAsNodePath(variant *C.godot_variant) *NodePath {
	nodePath := C.godot_variant_as_node_path(variant)

	return &NodePath{
		nodePath: &nodePath,
	}
}

func nodePathAsVariant(nodePath *NodePath) *C.godot_variant {
	var variant C.godot_variant
	C.godot_variant_new_node_path(&variant, nodePath.nodePath)

	return &variant
}

//TODO: Store these in another file? Possibly types.go or typeAsType.go?
//Non-variant type conversions
func basisAsGodotBasis(value Basis) *C.godot_basis {
	return value.basis
}

func godotBasisAsBasis(value C.godot_basis) *Basis {
	return &Basis{basis: &value}
}

func boolAsGodotBool(value bool) C.godot_bool {
	return C.godot_bool(value)
}

func godotBoolAsBool(value C.godot_bool) bool {
	return bool(value)
}

func uIntAsCUInt64(value uint64) C.uint64_t {
	return C.uint64_t(value)
}

func intAsCInt64(value int64) C.int64_t {
	return C.int64_t(value)
}

func intAsGodotInt(value int64) C.godot_int {
	return C.godot_int(value)
}

func godotIntAsInt(value C.godot_int) int {
	return int(value)
}

func realAsCReal(value float64) C.double {
	return C.double(value)
}

func realAsGodotReal(value float64) C.godot_real {
	return C.godot_real(value)
}

func godotRealAsReal(value C.godot_real) float64 {
	return float64(value)
}

func stringAsGodotString(value string) *C.godot_string {
	var godotString C.godot_string
	cString := C.CString(value)
	C.godot_string_new_data(&godotString, cString, C.int(len(value)))

	return &godotString
}

func godotStringAsString(value *C.godot_string) string {
	godotCString := C.godot_string_c_str(value)

	return C.GoString(godotCString)
}
func vec2AsGodotVec2(value Vector2) *C.godot_vector2 {
	return value.vector2
}

func godotVec2AsVec2(value C.godot_vector2) *Vector2 {
	return &Vector2{vector2: &value}
}

func vec3AsGodotVec3(value Vector3) *C.godot_vector3 {
	return value.vector3
}

func godotVec3AsVec3(value C.godot_vector3) *Vector3 {
	return &Vector3{vector3: &value}
}

func vec3AxisAsGodotVec3Axis(value Vec3Axis) C.godot_vector3_axis {
	return C.godot_vector3_axis(value)
}

func godotVec3AxisAsVec3Axis(value C.godot_vector3_axis) Vec3Axis {
	return Vec3Axis(value)
}

/*
TODO:
godot_array GDAPI godot_variant_as_array(const godot_variant *p_self);
godot_color GDAPI godot_variant_as_color(const godot_variant *p_self);
godot_dictionary GDAPI godot_variant_as_dictionary(const godot_variant *p_self);
godot_object GDAPI *godot_variant_as_object(const godot_variant *p_self);
godot_plane GDAPI godot_variant_as_plane(const godot_variant *p_self);
godot_pool_byte_array GDAPI godot_variant_as_pool_byte_array(const godot_variant *p_self);
godot_pool_color_array GDAPI godot_variant_as_pool_color_array(const godot_variant *p_self);
godot_pool_int_array GDAPI godot_variant_as_pool_int_array(const godot_variant *p_self);
godot_pool_real_array GDAPI godot_variant_as_pool_real_array(const godot_variant *p_self);
godot_pool_string_array GDAPI godot_variant_as_pool_string_array(const godot_variant *p_self);
godot_pool_vector2_array GDAPI godot_variant_as_pool_vector2_array(const godot_variant *p_self);
godot_pool_vector3_array GDAPI godot_variant_as_pool_vector3_array(const godot_variant *p_self);
godot_quat GDAPI godot_variant_as_quat(const godot_variant *p_self);
godot_rect2 GDAPI godot_variant_as_rect2(const godot_variant *p_self);
godot_rect3 GDAPI godot_variant_as_rect3(const godot_variant *p_self);
godot_rid GDAPI godot_variant_as_rid(const godot_variant *p_self);
godot_transform GDAPI godot_variant_as_transform(const godot_variant *p_self);
godot_transform2d GDAPI godot_variant_as_transform2d(const godot_variant *p_self);

PRIORITY:
godot_basis GDAPI godot_variant_as_basis(const godot_variant *p_self);
godot_vector2 GDAPI godot_variant_as_vector2(const godot_variant *p_self);
godot_vector3 GDAPI godot_variant_as_vector3(const godot_variant *p_self);
*/
