package dummy

// Dummy it's a dummy object
type Dummy struct {
	Name string
}

func NewDummy(name string) *Dummy {
	return &Dummy{Name: name}
}

func (d Dummy) String() string {
	return d.Name
}
