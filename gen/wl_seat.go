package gen

// A seat is a group of keyboards, pointer and touch devices. This
// object is published as a global during start up, or when such a
// device is hot plugged.  A seat typically has a pointer and
// maintains a keyboard focus and a pointer focus.
type WlSeat interface {
	// This is emitted whenever a seat gains or loses the pointer,
	// keyboard or touch capabilities.  The argument is a capability
	// enum containing the complete set of capabilities this seat has.
	Capabilities(Capabilities WlUint)
	// In a multiseat configuration this can be used by the client to help
	// identify which physical devices the seat represents. Based on
	// the seat configuration used by the compositor.
	Name(Name WlString)
	// The ID provided will be initialized to the wl_pointer interface
	// for this seat.
	// This request only takes effect if the seat has the pointer
	// capability.
	GetPointer(Id WlNewId)
	// The ID provided will be initialized to the wl_keyboard interface
	// for this seat.
	// This request only takes effect if the seat has the keyboard
	// capability.
	GetKeyboard(Id WlNewId)
	// The ID provided will be initialized to the wl_touch interface
	// for this seat.
	// This request only takes effect if the seat has the touch
	// capability.
	GetTouch(Id WlNewId)
}

// This is a bitmask of capabilities this seat has; if a member is
// set, then it is present on the seat.
type WlSeatCapability uint32

const (
	WlSeatPointer  WlSeatCapability = 1
	WlSeatKeyboard WlSeatCapability = 2
	WlSeatTouch    WlSeatCapability = 4
)
