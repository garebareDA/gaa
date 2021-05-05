package engine

type Engine struct {
	routes Route
}

type ContextFunc func(*Context)

func (e Engine) New() {

}

func (e Engine) Get(path string, header ContextFunc) {
	if header == nil {
		panic("http: nil header")
	}

}

func (e Engine) Post(path string, c ContextFunc) {

}
