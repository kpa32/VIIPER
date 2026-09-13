package device

type CreateOptions struct {
	IDVendor       *uint16
	IDProduct      *uint16
	Manufacturer   *string
	ProductName    *string
	SerialNumber   *string
	DeviceSpecific string
}
