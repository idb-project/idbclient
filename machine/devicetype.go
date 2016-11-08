package machine

// DeviceType mapping from idb/config/application.yml.
// Let's hope that the mapping doesn't get changed.
type DeviceType int

const (
	DeviceTypeNone		DeviceType	= iota
	DeviceTypePhyiscal
	DeviceTypeVirtual
	DeviceTypeSwitch
)
