package types

import (
	"fmt"
	"io"
)

type VarExpr struct {
	Position

	Name string
}

func (obj VarExpr) Resolve(scope map[string]Value, ev *Evaluator) (Value, error) {
	val, ok := scope[obj.Name]
	if !ok {
		return nil, fmt.Errorf("%s: not in scope: %s", obj.Pos(), obj.Name)
	}
	return val, nil
}

func (obj VarExpr) hashValue(w io.Writer) {
	fmt.Fprintf(w, "%T", obj)
	fmt.Fprint(w, obj.Name)
}

type AttributeExpr struct {
	Position

	Base Expression
	Name string
}

func (obj AttributeExpr) Resolve(scope map[string]Value, ev *Evaluator) (Value, error) {
	val, err := obj.Base.Resolve(scope, ev)
	if err != nil {
		return nil, err
	}
	switch mapval := val.(type) {
	case MapValue:
		val, ok := mapval.values[obj.Name]
		if !ok {
			return nil, fmt.Errorf("%s: map has no attribute %s", mapval.Pos(), obj.Name)
		}
		return val, nil
	default:
		return nil, fmt.Errorf("%s: %T has no attributes", mapval.Pos(), mapval)
	}
}

func (obj AttributeExpr) hashValue(w io.Writer) {
	fmt.Fprintf(w, "%T", obj)
	fmt.Fprint(w, obj.Name)
}
