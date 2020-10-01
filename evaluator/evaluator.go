package evaluator

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/object"
)

func Eval(node ast.Node) object.Object{
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator,right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator,left,right)
	}
	return nil
}

func evalIfExpression(node *ast.IfExpression) object.Object {
	condition := Eval(node.Condition)
	if ifTruthy(condition){
		return Eval(node.True)
	}else if node.False != nil{
		return Eval(node.False)
	}else{
		return NULL
	}
}

func ifTruthy(condition object.Object) bool {
	switch condition {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER && right.Type() == object.INTEGER:
		return evalIntegerInfixExpression(operator,left,right)
	case left.Type() == object.BOOLEAN && right.Type() == object.BOOLEAN:
		return evalBooleanInfixExpression(operator,left,right)
	default:
		return NULL
	}
}

func evalBooleanInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch operator{
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator{
	case "+":
		return &object.Integer{Value: leftVal+rightVal}
	case "-":
		return &object.Integer{Value: leftVal-rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER{
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _,statement := range statements{
		result = Eval(statement)
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input{
		return TRUE
	}
	return FALSE
}
