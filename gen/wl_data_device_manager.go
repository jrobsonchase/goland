package gen

// The wl_data_device_manager is a singleton global object that
// provides access to inter-client data transfer mechanisms such as
// copy-and-paste and drag-and-drop.  These mechanisms are tied to
// a wl_seat and this interface lets a client get a wl_data_device
// corresponding to a wl_seat.
type WlDataDeviceManager interface {
	// Create a new data source.
	CreateDataSource(Id WlNewId)
	// Create a new data device for a given seat.
	GetDataDevice(Id WlNewId, Seat WlObject)
}
