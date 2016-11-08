//go:generate stringer -type=BackupType

package machine

// BackupType is the main backup setting (if backup is enabled).
type BackupType int

const (
	BackupTypeNo BackupType = iota
	BackupTypeYes
	BackupTypeNotNeeded
	BackupTypeNotResponsible
	BackupTypeEnd
)
