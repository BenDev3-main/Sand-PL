package environment

type Environment struct {
	store map[string]interface{}
}

func New() *Environment {
	s := make(map[string]interface{})
	return &Environment{store: s}
}

func (e *Environment) Set(name string, val interface{}) interface{} {
	e.store[name] = val
	return val
}

func (e *Environment) Get(name string) (interface{}, bool) {
	val, ok := e.store[name]
	return val, ok
}
