package libs

type lib struct {
	a string
}

type TLib = *lib

func NewLib(s string) TLib {
	return &lib{a: s}
}

func (l *lib) String() string {
	return l.a
}
