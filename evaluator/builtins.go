package evaluator

import "MonkeyInterpreter/object"

var builtins = map[string]*object.Builtin{
	"len":&object.Builtin{
		Fn:func (args ...object.Object) object.Object{
			if len(args)!=1{
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg:=args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int16(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY{
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
}
