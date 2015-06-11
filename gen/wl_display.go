package gen

// The core global object.  This is a special singleton object.  It
// is used for internal Wayland protocol features.
type WlDisplay interface {
	// The error event is sent out when a fatal (non-recoverable)
	// error has occurred.  The object_id argument is the object
	// where the error occurred, most often in response to a request
	// to that object.  The code identifies the error and is defined
	// by the object interface.  As such, each interface defines its
	// own set of error codes.  The message is an brief description
	// of the error, for (debugging) convenience.
	Error(ObjectId WlObject, Code WlUint, Message WlString)
	// This event is used internally by the object ID management
	// logic.  When a client deletes an object, the server will send
	// this event to acknowledge that it has seen the delete request.
	// When the client receive this event, it will know that it can
	// safely reuse the object ID.
	DeleteId(Id WlUint)
	// The sync request asks the server to emit the 'done' event
	// on the returned wl_callback object.  Since requests are
	// handled in-order and events are delivered in-order, this can
	// be used as a barrier to ensure all previous requests and the
	// resulting events have been handled.
	// The object returned by this request will be destroyed by the
	// compositor after the callback is fired and as such the client must not
	// attempt to use it after that point.
	// The callback_data passed in the callback is the event serial.
	Sync(Callback WlNewId)
	// This request creates a registry object that allows the client
	// to list and bind the global objects available from the
	// compositor.
	GetRegistry(Registry WlNewId)
}

// These errors are global and can be emitted in response to any
// server request.
type WlDisplayError uint32

const (
	WlDisplayInvalidObject WlDisplayError = 0
	WlDisplayInvalidMethod WlDisplayError = 1
	WlDisplayNoMemory      WlDisplayError = 2
)
