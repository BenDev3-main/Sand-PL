package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sand-lang/src/environment"
	"sand-lang/src/evaluator"
	"sand-lang/src/lexer"
	"sand-lang/src/parser"
)

func main() {
	// 1. Verificamos si pasaste un archivo (ej: ./sand.exe hola.snd)
	if len(os.Args) < 2 {
		fmt.Println("Usage: sand <file>.sand")
		return
	}

	filename := os.Args[1]

	// 2. Leemos el contenido del archivo
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// 3. Inicializamos el Lexer y el Parser
	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()

	// 4. Chequeamos errores de sintaxis
	if len(p.Errors()) != 0 {
		fmt.Println("Syntax Errors:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		return
	}

	// 5. Configuramos el entorno y ejecutamos
	env := environment.New()
	// Registramos stdio para que funcione stdio.logln
	env.Set("stdio", evaluator.StdioPackage)

	evaluator.Eval(program, env)
}

