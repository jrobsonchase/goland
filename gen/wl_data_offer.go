package gen

// A wl_data_offer represents a piece of data offered for transfer
// by another client (the source client).  It is used by the
// copy-and-paste and drag-and-drop mechanisms.  The offer
// describes the different mime types that the data can be
// converted to and provides the mechanism for transferring the
// data directly from the source client.
type WlDataOffer interface {
	// Sent immediately after creating the wl_data_offer object.  One
	// event per offered mime type.
	Offer(MimeType WlString)
	// Indicate that the client can accept the given mime type, or
	// NULL for not accepted.
	// Used for feedback during drag-and-drop.
	Accept(Serial WlUint, MimeType WlString)
	// To transfer the offered data, the client issues this request
	// and indicates the mime type it wants to receive.  The transfer
	// happens through the passed file descriptor (typically created
	// with the pipe system call).  The source client writes the data
	// in the mime type representation requested and then closes the
	// file descriptor.
	// The receiving client reads from the read end of the pipe until
	// EOF and then closes its end, at which point the transfer is
	// complete.
	Receive(MimeType WlString, Fd WlFd)
	// Destroy the data offer.
	Destroy()
}
