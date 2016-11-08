//go:generate stringer -type=BackupBrand

package machine

// BackupBrand selects which backup software is used.
type BackupBrand int

const (
	BackupBrandNone BackupBrand = iota
	BackupBrandBacula
	BackupBrandSEP
	BackupBrandBackupPC
	BackupBrandFile
	BackupBrandEnd
)
