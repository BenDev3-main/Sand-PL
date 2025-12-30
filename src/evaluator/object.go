package evaluator

import (
	"os"
)

// Object is the base interface for all Sand data types
type Object interface {
	Type() string
	Inspect() string
}

// FileObject handles the os.File pointer for Sand
type FileObject struct {
	File *os.File
}

func (f *FileObject) Type() string { return "FILE" }
func (f *FileObject) Inspect() string {
	if f.File != nil {
		return "<file:" + f.File.Name() + ">"
	}
	return "<file:nil>"
}

// PackageObject allows grouping functions under a namespace like 'stdio'
type PackageObject struct {
	Name    string
	Methods map[string]interface{}
}

func (p *PackageObject) Type() string    { return "PACKAGE" }
func (p *PackageObject) Inspect() string { return "<package:" + p.Name + ">" }
