package evaluator

import (
	"fmt"
	"os"
)

var StdioPackage = &PackageObject{
	Name: "stdio",
	Methods: map[string]interface{}{
		"log": func(args ...interface{}) interface{} {
			for _, arg := range args {
				fmt.Print(arg)
			}
			return nil
		},
		"logln": func(args ...interface{}) interface{} {
			for _, arg := range args {
				fmt.Println(arg)
			}
			return nil
		},
		"openfile": func(args ...interface{}) interface{} {
			if len(args) < 1 {
				return "ERROR: path argument required"
			}
			path, ok := args[0].(string)
			if !ok {
				return "ERROR: path must be a string"
			}

			f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				return "ERROR: " + err.Error()
			}
			return &FileObject{File: f}
		},
	},
}
