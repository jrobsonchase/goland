package gen

// The wl_data_source object is the source side of a wl_data_offer.
// It is created by the source client in a data transfer and
// provides a way to describe the offered data and a way to respond
// to requests to transfer the data.
type WlDataSource interface {
	// Sent when a target accepts pointer_focus or motion events.  If
	// a target does not accept any of the offered types, type is NULL.
	// Used for feedback during drag-and-drop.
	Target(MimeType WlString)
	// Request for data from the client.  Send the data as the
	// specified mime type over the passed file descriptor, then
	// close it.
	Send(MimeType WlString, Fd WlFd)
	// This data source has been replaced by another data source.
	// The client should clean up and destroy this data source.
	Cancelled()
	// This request adds a mime type to the set of mime types
	// advertised to targets.  Can be called several times to offer
	// multiple types.
	Offer(MimeType WlString)
	// Destroy the data source.
	Destroy()
}
