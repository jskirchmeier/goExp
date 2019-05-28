package simple

// can be used in all other things
// implements the Thing interface
type Identity struct {
	ID     int    `yaml:"id"`
	Nombre string `yaml:"name"`
}

func (t *Identity) Identify() int {
	return t.ID
}

func (t *Identity) Name() string {
	return t.Nombre
}
