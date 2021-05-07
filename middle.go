package gaa

type Constructor func(Handle) Handle

type Middle struct {
	constructors []Constructor
}

func MiddleNew(constructors ...Constructor) Middle {
	return Middle{append(([]Constructor(nil)), constructors...)}
}

func (m Middle) Append(constructors ...Constructor) Middle {
	new := make([]Constructor, 0, len(m.constructors)+len(constructors))
	new = append(new, m.constructors...)
	new = append(new, constructors...)

	return Middle{constructors: new}
}

func (m Middle) Then(h Handle) Handle {
	for i := range m.constructors {
		h = m.constructors[len(m.constructors)-1-i](h)
	}
	return h
}
