package simple

// can be used in all other things
// implements the Thing interface
type Id struct {
	IdValue int
}

func (t *Id) Id() int {
	return t.IdValue
}

type Name struct {
	NameValue string
}

func (t *Name) Name() string {
	return t.NameValue
}

// many items will have compound descriptions and will be built
// description may also change as the story evolves
// so this implementation is for simple cases
type Description struct {
	DescriptionValue string
}

func (t *Description) Description() string {
	return t.DescriptionValue
}
