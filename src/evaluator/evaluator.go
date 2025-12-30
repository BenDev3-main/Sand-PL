package evaluator

import (
	"fmt"
	"os"
	"sand-lang/src/ast"
	"sand-lang/src/environment"
)

// Definimos constantes para tipos internos
const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	ERROR_OBJ   = "ERROR"
)

func Eval(node ast.Node, env *environment.Environment) interface{} {
	switch n := node.(type) {

	// 1. Punto de entrada: El programa completo
	case *ast.Program:
		return evalProgram(n, env)

	// 2. Declaración de variables: var x = ...
	case *ast.VarStatement:
		val := Eval(n.Value, env)
		env.Set(n.Name.Value, val)
		return val

	// 3. Expresiones sueltas
	case *ast.ExpressionStatement:
		return Eval(n.Expression, env)

	// 4. Literales (Valores puros)
	case *ast.IntegerLiteral:
		return n.Value

	case *ast.FloatLiteral:
		return n.Value

	case *ast.StringLiteral:
		return n.Value

	// 5. Identificadores: cuando usás el nombre de una variable
	case *ast.Identifier:
		val, ok := env.Get(n.Value)
		if !ok {
			return fmt.Errorf("identifier not found: %s", n.Value)
		}
		return val

	// 6. LLAMADAS CON PUNTO: stdio.logln(...)
	case *ast.MethodCallExpression:
		obj := Eval(n.Object, env)
		return evalMethodCall(obj, n.Method.Value, n.Arguments, env)

	// 7. LLAMADAS NORMALES: print(...)
	case *ast.CallExpression:
		function := Eval(n.Function, env)
		args := evalExpressions(n.Arguments, env)
		return applyFunction(function, args)

	}
	return nil
}

// --- Funciones de Soporte (Las que faltaban) ---

func evalProgram(program *ast.Program, env *environment.Environment) interface{} {
	var result interface{}
	for _, statement := range program.Statements {
		result = Eval(statement, env)
	}
	return result
}

func evalIdentifier(node *ast.Identifier, env *environment.Environment) interface{} {
	val, ok := env.Get(node.Value)
	if !ok {
		return "ERROR: variable no encontrada: " + node.Value
	}
	return val
}

func evalPrefixExpression(operator string, right interface{}) interface{} {
	switch operator {
	case "-":
		if val, ok := right.(int64); ok {
			return -val
		}
		return "ERROR: tipo desconocido para prefijo -"
	default:
		return nil
	}
}

func evalInfixExpression(operator string, left, right interface{}) interface{} {
	// Solo si ambos son números
	lVal, okL := left.(int64)
	rVal, okR := right.(int64)

	if okL && okR {
		switch operator {
		case "+":
			return lVal + rVal
		case "-":
			return lVal - rVal
		case "*":
			return lVal * rVal
		case "/":
			return lVal / rVal
		}
	}
	return "ERROR: operación no soportada"
}

func isError(obj interface{}) bool {
	if obj != nil {
		if s, ok := obj.(string); ok {
			// Un chequeo simple de error para este nivel
			return len(s) > 6 && s[:6] == "ERROR:"
		}
	}
	return false
}

// Dentro de evalFileMethods en evaluator.go
func evalFileMethods(fileObj *FileObject, method string, args []interface{}) interface{} {
	switch method {
	case "read":
		// Usamos el archivo interno de Go
		content, err := os.ReadFile(fileObj.File.Name())
		if err != nil {
			return "ERROR: " + err.Error()
		}
		return string(content)

	case "write":
		if len(args) < 1 {
			return "ERROR: falta contenido"
		}
		fileObj.File.WriteString(args[0].(string))
		return nil

	case "close":
		fileObj.File.Close()
		return "archivo cerrado"
	}
	return "ERROR: método no soportado"
}

// evalExpressions convierte la lista de AST Expressions en valores reales (interface{})
func evalExpressions(exps []ast.Expression, env *environment.Environment) []interface{} {
	var result []interface{}
	for _, e := range exps {
		evaluated := Eval(e, env)
		result = append(result, evaluated)
	}
	return result
}

// applyFunction ejecuta la función (ya sea nativa de Go o definida en Sand)
func applyFunction(fn interface{}, args []interface{}) interface{} {
	switch f := fn.(type) {
	case func(...interface{}) interface{}:
		return f(args...)
	default:
		return fmt.Errorf("not a function: %T", fn)
	}
}

// evalMethodCall busca y ejecuta un método dentro de un PackageObject (como stdio)
func evalMethodCall(obj interface{}, method string, args []ast.Expression, env *environment.Environment) interface{} {
	packageObj, ok := obj.(*PackageObject)
	if !ok {
		return fmt.Errorf("object is not a package")
	}

	fn, ok := packageObj.Methods[method]
	if !ok {
		return fmt.Errorf("method %s not found in package %s", method, packageObj.Name)
	}

	// Evaluamos los argumentos antes de pasarlos a la función nativa
	evaluatedArgs := evalExpressions(args, env)

	// Si es una función nativa de nuestro mapa de métodos
	if nativeFn, ok := fn.(func(...interface{}) interface{}); ok {
		return nativeFn(evaluatedArgs...)
	}

	return fmt.Errorf("invalid method type")
}
