package evaluator

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/object"
	"fmt"
)

func Eval(node ast.Node,env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node,env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression,env)
	case *ast.BlockStatement:
		return evalBlockStatement(node,env)
	case *ast.IfExpression:
		return evalIfExpression(node,env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right,env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left,env)
		right := Eval(node.Right,env)
		if isError(left) {
			return left
		}
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.ReturnStatement:
		val := Eval(node.Value,env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		val := Eval(node.Value,env)
		if isError(val){
			return val
		}
		env.Set(node.Name.Value,val)
	case *ast.Identifier:
		return evalIdentifier(node,env)
	case *ast.FunctionLiteral:
		params := node.Arguments
		body := node.Body
		return &object.Function{Arguments: params,Env: env,Body: body}
	case *ast.CallExpression:
		fn := Eval(node.Function,env)
		if isError(fn){
			return fn
		}
		args := evalExpressions(node.Arguments,env)
		if len(args) == 1 && isError(args[0]){
			return args[0]
		}
		return applyFunction(fn,args)
	case *ast.ArrayLiteral:
		elems := evalExpressions(node.Elements,env)
		if len(elems) == 1 && isError(elems[0]){
			return elems[0]
		}
		return &object.Array{Elements: elems}
	case *ast.IndexExpression:
		left := Eval(node.Left,env)
		if isError(left){
			return left
		}
		index := Eval(node.Index,env)
		if isError(index){
			return index
		}
		return evalIndexExpression(left,index)
	}
	return nil
}

func evalIndexExpression(left object.Object, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY && index.Type() == object.INTEGER:
		return evalArrayIndexExpression(left,index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(left object.Object, index object.Object) object.Object {
	arrayReturn := left.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int16(len(arrayReturn.Elements)-1)
	if idx<0 || idx>max {
		return NULL
	}
	return arrayReturn.Elements[idx]
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn,args)
		eval := Eval(fn.Body,extendedEnv)
		return unwrapReturnValue(eval)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func unwrapReturnValue(eval object.Object) object.Object {
	if returnVal ,ok := eval.(*object.ReturnValue);ok{
		return returnVal.Value
	}
	return eval
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for i,arg := range fn.Arguments{
		env.Set(arg.Value,args[i])
	}
	return env
}

func evalExpressions(arguments []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _,e := range arguments{
		eval := Eval(e,env)
		if isError(eval){
			return []object.Object{eval}
		}
		result = append(result,eval)
	}
	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {

	if val,ok := env.Get(node.Value);ok {
		return val
	}

	if builtin,ok := builtins[node.Value];ok{
		return builtin
	}

	return newError("identifier not found: " + node.Value)

}

func evalBlockStatement(node *ast.BlockStatement,env *object.Environment) object.Object {
	var result object.Object

	for _,statement := range node.Statements{
		result = Eval(statement,env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.ERROR {
				return result
			}
		}
	}
	return result
}

func evalIfExpression(node *ast.IfExpression,env *object.Environment) object.Object {
	condition := Eval(node.Condition,env)
	if isError(condition) {
		return condition
	}
	if ifTruthy(condition) {
		return Eval(node.True,env)
	} else if node.False != nil {
		return Eval(node.False,env)
	} else {
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
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.BOOLEAN && right.Type() == object.BOOLEAN:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch operator {
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
		left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return newError("unknown operator: -%s", right.Type())
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
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func evalProgram(program *ast.Program,env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement,env)
		switch result := result.(type){
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func newError(format string,a ...interface{}) *object.Error{
	return &object.Error{Message: fmt.Sprintf(format,a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}
