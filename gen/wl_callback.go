package gen

// Clients can handle the 'done' event to get notified when
// the related request is done.
type WlCallback interface {
	// Notify the client when the related request is done.
	Done(CallbackData WlUint)
}
