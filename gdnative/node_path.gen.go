package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include <gdnative/node_path.h>
*/
import "C"

type NodePath struct {
	base *C.godot_node_path
}

func (t *NodePath) getBase() *C.godot_node_path {
	return t.base
}