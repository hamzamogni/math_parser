package evaluator

import (
	"fmt"
	"parser_lexer/ast"
	"parser_lexer/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}

		right := Eval(node.Right)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *ast.Identifier:
		return evalIdentifier(node)
	case *ast.CallExpression:
		function := Eval(node.Function)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	}

	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return nil
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func evalExpressions(exps []ast.Expression) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e)
		result = append(result, evaluated)
	}

	return result
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "+":
		return evalPrefixPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalPrefixPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: +%s", right.Type())
	}

	value := right.(*object.Number).Value
	return &object.Number{Value: value}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	if left.Type() != right.Type() {
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	if left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ {
		return evalNumberInfixExpression(operator, left, right)
	}

	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Number).Value
	rightVal := right.(*object.Number).Value

	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "**":
		return &object.Number{Value: builtins["pow"].Fn(left, right).(*object.Number).Value}
	case "/":
		if right.Inspect() == "0.00" {
			return newError("division by zero")
		}
		return &object.Number{Value: leftVal / rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIdentifier(node *ast.Identifier) object.Object {
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("Unexpected identifier: " + node.Value)
}
