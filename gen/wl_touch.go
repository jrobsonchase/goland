package gen

// The wl_touch interface represents a touchscreen
// associated with a seat.
// Touch interactions can consist of one or more contacts.
// For each contact, a series of events is generated, starting
// with a down event, followed by zero or more motion events,
// and ending with an up event. Events relating to the same
// contact point can be identified by the ID of the sequence.
type WlTouch interface {
	// A new touch point has appeared on the surface. This touch point is
	// assigned a unique @id. Future events from this touchpoint reference
	// this ID. The ID ceases to be valid after a touch up event and may be
	// re-used in the future.
	Down(Serial WlUint, Time WlUint, Surface WlObject, Id WlInt, X WlFixed, Y WlFixed)
	// The touch point has disappeared. No further events will be sent for
	// this touchpoint and the touch point's ID is released and may be
	// re-used in a future touch down event.
	Up(Serial WlUint, Time WlUint, Id WlInt)
	// A touchpoint has changed coordinates.
	Motion(Time WlUint, Id WlInt, X WlFixed, Y WlFixed)
	// Indicates the end of a contact point list.
	Frame()
	// Sent if the compositor decides the touch stream is a global
	// gesture. No further events are sent to the clients from that
	// particular gesture. Touch cancellation applies to all touch points
	// currently active on this client's surface. The client is
	// responsible for finalizing the touch points, future touch points on
	// this surface may re-use the touch point ID.
	Cancel()
	Release()
}
