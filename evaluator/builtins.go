package evaluator

import (
	"github.com/hamzamogni/math_parser/object"
	"math"
)

var builtins = map[string]*object.Builtin{
	"cos":  mathBuiltinSingleArg("cos", math.Cos),
	"acos": mathBuiltinSingleArg("acos", math.Acos),
	"sin":  mathBuiltinSingleArg("sin", math.Sin),
	"asin": mathBuiltinSingleArg("asin", math.Asin),
	"tan":  mathBuiltinSingleArg("tan", math.Tan),
	"atan": mathBuiltinSingleArg("atan", math.Atan),
	"exp":  mathBuiltinSingleArg("exp", math.Exp),
	"log":  mathBuiltinSingleArg("log", math.Log),
	"abs":  mathBuiltinSingleArg("abs", math.Abs),

	"sqrt": mathBuiltinSingleArg("sqrt", math.Sqrt),
	"pow":  mathBuiltinTwoArgs("pow", math.Pow),
}

func mathBuiltinSingleArg(name string, function func(arg float64) float64) *object.Builtin {
	return &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("%s accepts 1 argument. got=%d", name, len(args))
			}

			if args[0].Type() != object.NUMBER_OBJ {
				return newError("%s only accepts numbers. got=%s", name, args[0].Type())
			}
			return &object.Number{Value: function(args[0].(*object.Number).Value)}
		},
	}
}

func mathBuiltinTwoArgs(name string, function func(arg1, arg2 float64) float64) *object.Builtin {
	return &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("%s accepts 2 arguments. got=%d", name, len(args))
			}

			if args[0].Type() != object.NUMBER_OBJ || args[1].Type() != object.NUMBER_OBJ {
				return newError("%s only accepts numbers. got=%s and %s", name, args[0].Type(), args[1].Type())
			}
			return &object.Number{Value: function(args[0].(*object.Number).Value, args[1].(*object.Number).Value)}
		},
	}
}
