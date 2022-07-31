package presence

import (
	"reflect"

	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

func Route(root ast.AstNode, path ...interface{}) (interface{}, error) {
	curr := root
	for _, ind := range path {
		if ii, ok := ind.(int); ok {
			for val, ok := curr.(*ast.ValueNode); ok; val, ok = curr.(*ast.ValueNode) {
				curr = val.Child
			}
			if arr, ok := curr.(*ast.ArrayNode); ok && ii < len(arr.ValueArray) {
				curr = &arr.ValueArray[ii]
			} else if ok && ii >= len(arr.ValueArray) {
				return curr.Evaluate(), &OutOfRangeError{ii, len(arr.ValueArray)}
			} else {
				return curr.Evaluate(), &InvalidRouteError{reflect.TypeOf(curr.Evaluate()).String(), "integer"}
			}
		} else if si, ok := ind.(string); ok {
			for val, ok := curr.(*ast.ValueNode); ok; val, ok = curr.(*ast.ValueNode) {
				curr = val.Child
			}
			if arr, ok := curr.(*ast.ObjectNode); ok {
				tmp, exist := arr.ValueObject[si]
				if !exist {
					keyList := make([]string, 0)
					for k := range arr.ValueObject {
						keyList = append(keyList, k)
					}
					return curr.Evaluate(), &NotFoundKeyError{si, keyList}
				}
				curr = &tmp
			} else {
				return curr.Evaluate(), &InvalidRouteError{reflect.TypeOf(curr.Evaluate()).String(), "string"}
			}
		}
	}
	return curr.Evaluate(), nil
}

type path interface {
	match(cases)
}

type ArrInd int

func (i ArrInd) match(c cases) { c.ArrInd(i) }

type ObjInd string

func (i ObjInd) match(c cases) { c.ObjInd(i) }

type cases struct {
	ArrInd func(ArrInd)
	ObjInd func(ObjInd)
}

func Access(an ast.AstNode, path ...path) interface{} {
	curr := an
	for _, ind := range path {
		ind.match(cases{
			ArrInd: func(i ArrInd) {
				for val, ok := curr.(*ast.ValueNode); ok; val, ok = curr.(*ast.ValueNode) {
					curr = val.Child
				}
				if arr, ok := curr.(*ast.ArrayNode); ok {
					curr = &arr.ValueArray[int(i)]
				}
			},
			ObjInd: func(s ObjInd) {
				for val, ok := curr.(*ast.ValueNode); ok; val, ok = curr.(*ast.ValueNode) {
					curr = val.Child
				}
				if obj, ok := curr.(*ast.ObjectNode); ok {
					tmp := obj.ValueObject[string(s)]
					curr = &tmp
				}
			},
		})
	}
	return curr.Evaluate()
}
