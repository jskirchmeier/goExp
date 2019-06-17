package behaviors

// various behaviors the objects in our adventure can exhibit
// commands act on different objects

// Lock is something that can be unlocked, and never locked back up :)
// a key may or may not be needed
// TODO: does this need to be a descriptor, or does the object implementing it take care of it?
// for now the second
type Lockable interface {
	// the state of the lock
	Locked() bool

	// Unlock changes the state of the lock, if locked
	// key may or may not be required
	// returns nil if unlocked successfully, or was unlocked to start with, or an err string
	// describing the key needed
	Unlock(key interface{}) (err string)
}

// Openable is something that can be opened, and never closed again (for now)
// could be a door in a room, a cabinet, a drawbridge, ....
// can be used in combination with lock, the command that will open the thing will take care of it
type Openable interface {
	// the state of the item
	Opened() bool

	// Open changes the state of the item
	Open() (err string)
}

// Describer is the basic mechanism to describe an object to the user
// should include any information needed by the user to select it if needed (Red Door, North passageway, etc.)
type Describer interface {
	Description() string
}
