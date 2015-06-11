package gen

// The wl_keyboard interface represents one or more keyboards
// associated with a seat.
type WlKeyboard interface {
	// This event provides a file descriptor to the client which can be
	// memory-mapped to provide a keyboard mapping description.
	Keymap(Format WlUint, Fd WlFd, Size WlUint)
	// Notification that this seat's keyboard focus is on a certain
	// surface.
	Enter(Serial WlUint, Surface WlObject, Keys WlArray)
	// Notification that this seat's keyboard focus is no longer on
	// a certain surface.
	// The leave notification is sent before the enter notification
	// for the new focus.
	Leave(Serial WlUint, Surface WlObject)
	// A key was pressed or released.
	// The time argument is a timestamp with millisecond
	// granularity, with an undefined base.
	Key(Serial WlUint, Time WlUint, Key WlUint, State WlUint)
	// Notifies clients that the modifier and/or group state has
	// changed, and it should update its local state.
	Modifiers(Serial WlUint, ModsDepressed WlUint, ModsLatched WlUint, ModsLocked WlUint, Group WlUint)
	// Informs the client about the keyboard's repeat rate and delay.
	// This event is sent as soon as the wl_keyboard object has been created,
	// and is guaranteed to be received by the client before any key press
	// event.
	// Negative values for either rate or delay are illegal. A rate of zero
	// will disable any repeating (regardless of the value of delay).
	// This event can be sent later on as well with a new value if necessary,
	// so clients should continue listening for the event past the creation
	// of wl_keyboard.
	RepeatInfo(Rate WlInt, Delay WlInt)
	Release()
}

// This specifies the format of the keymap provided to the
// client with the wl_keyboard.keymap event.
type WlKeyboardKeymapFormat uint32

const (
	WlKeyboardNoKeymap WlKeyboardKeymapFormat = 0
	WlKeyboardXkbV1    WlKeyboardKeymapFormat = 1
)

// Describes the physical state of a key which provoked the key event.
type WlKeyboardKeyState uint32

const (
	WlKeyboardReleased WlKeyboardKeyState = 0
	WlKeyboardPressed  WlKeyboardKeyState = 1
)
