package machine

import (
	"encoding/json"
	"time"
)

var parseLayouts = []string{"2006-01-02T15:04:05.9999Z", "2006-01-02 15:04:05", "2006-01-02", time.RFC3339Nano, time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano}

var formatLayout = "2006-01-02 15:04:05"

func parseTime(layouts []string, value string) (time.Time, error) {
	var err error
	for _, v := range layouts {
		var t time.Time
		t, err = time.Parse(v, value)

		// continue to the next layout
		if err != nil {
			continue
		}

		// parsed successfully
		return t, nil
	}

	return time.Time{}, err
}

type jsonMachine struct {
	Fqdn                                  string      `json:"fqdn"`
	Os                                    string      `json:"os,omitempty"`
	Arch                                  string      `json:"arch,omitempty"`
	RAM                                   int         `json:"ram,omitempty"`
	Cores                                 int         `json:"cores,omitempty"`
	Diskspace                             int         `json:"diskspace,omitempty"`
	Vmhost                                string      `json:"vmhost,omitempty"`
	Description                           string      `json:"description,omitempty"`
	OsRelease                             string      `json:"os_release,omitempty"`
	Uptime                                int         `json:"uptime,omitempty"`
	DeviceTypeID                          DeviceType  `json:"device_type_id,omitempty"`
	Serialnumber                          string      `json:"serialnumber,omitempty"`
	OwnerID                               int         `json:"owner_id,omitempty"`
	AutoUpdate                            bool        `json:"auto_update,omitempty"`
	SwitchURL                             string      `json:"switch_url,omitempty"`
	MrtgURL                               string      `json:"mrtg_url,omitempty"`
	ConfigInstructions                    string      `json:"config_instructions,omitempty"`
	SwCharacteristics                     string      `json:"sw_characteristics,omitempty"`
	BusinessPurpose                       string      `json:"business_purpose,omitempty"`
	BusinessCriticality                   string      `json:"business_criticality,omitempty"`
	BusinessNotification                  string      `json:"business_notification,omitempty"`
	UnattendedUpgrades                    bool        `json:"unattended_upgrades,omitempty"`
	UnattendedUpgradesBlacklistedPackages string      `json:"unattended_upgrades_blacklisted_packages,omitempty"`
	UnattendedUpgradesReboot              bool        `json:"unattended_upgrades_reboot,omitempty"`
	UnattendedUpgradesTime                string      `json:"unattended_upgrades_time,omitempty"`
	UnattendedUpgradesRepos               string      `json:"unattended_upgrades_repos,omitempty"`
	PendingUpdates                        int         `json:"pending_updates,omitempty"`
	PendingSecurityUpdates                int         `json:"pending_security_updates,omitempty"`
	PendingUpdatesSum                     int         `json:"pending_updates_sum,omitempty"`
	PendingUpdatesPackageNames            string      `json:"pending_updates_package_names,omitempty"`
	SeverityClass                         string      `json:"severity_class,omitempty"`
	UcsRole                               string      `json:"ucs_role,omitempty"`
	BackupType                            BackupType  `json:"backup_type,omitempty"`
	BackupBrand                           BackupBrand `json:"backup_brand,omitempty"`
	BackupLastFullSize                    int64       `json:"backup_last_full_size,omitempty"`
	BackupLastIncSize                     int64       `json:"backup_last_inc_size,omitempty"`
	BackupLastDiffSize                    int64       `json:"backup_last_diff_size,omitempty"`
	Nics                                  []Nic       `json:"nics,omitempty"`

	CreateMachine bool `json:"create_machine,string"`

	// time.Time fields need special handling for omitempty / unmarshalling if multiple formats can occur
	ServicedAt string `json:"serviced_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`

	BackupLastFullRun string `json:"backup_last_full_run,omitempty"`
	BackupLastIncRun  string `json:"backup_last_inc_run,omitempty"`
	BackupLastDiffRun string `json:"backup_last_diff_run,omitempty"`
}

func (m *Machine) UnmarshalJSON(buf []byte) error {
	var jm jsonMachine

	err := json.Unmarshal(buf, &jm)
	if err != nil {
		return err
	}

	m.Fqdn = jm.Fqdn
	m.Os = jm.Os
	m.Arch = jm.Arch
	m.RAM = jm.RAM
	m.Cores = jm.Cores
	m.Diskspace = jm.Diskspace
	m.Vmhost = jm.Vmhost
	m.Description = jm.Description

	if jm.ServicedAt != "" {
		m.ServicedAt, err = parseTime(parseLayouts, jm.ServicedAt)
		if err != nil {
			return err
		}
	}

	if jm.DeletedAt != "" {
		m.DeletedAt, err = parseTime(parseLayouts, jm.DeletedAt)
		if err != nil {
			return err
		}
	}

	if jm.CreatedAt != "" {
		m.CreatedAt, err = parseTime(parseLayouts, jm.CreatedAt)
		if err != nil {
			return err
		}
	}

	if jm.UpdatedAt != "" {
		m.UpdatedAt, err = parseTime(parseLayouts, jm.UpdatedAt)
		if err != nil {
			return err
		}
	}

	m.OsRelease = jm.OsRelease
	m.Uptime = jm.Uptime
	m.DeviceTypeID = jm.DeviceTypeID
	m.Serialnumber = jm.Serialnumber
	m.OwnerID = jm.OwnerID
	m.AutoUpdate = jm.AutoUpdate
	m.SwitchURL = jm.SwitchURL
	m.MrtgURL = jm.MrtgURL
	m.ConfigInstructions = jm.ConfigInstructions
	m.SwCharacteristics = jm.SwCharacteristics
	m.BusinessPurpose = jm.BusinessPurpose
	m.BusinessCriticality = jm.BusinessCriticality
	m.BusinessNotification = jm.BusinessNotification
	m.UnattendedUpgrades = jm.UnattendedUpgrades
	m.UnattendedUpgradesBlacklistedPackages = jm.UnattendedUpgradesBlacklistedPackages
	m.UnattendedUpgradesReboot = jm.UnattendedUpgradesReboot
	m.UnattendedUpgradesTime = jm.UnattendedUpgradesTime
	m.UnattendedUpgradesRepos = jm.UnattendedUpgradesRepos
	m.PendingUpdates = jm.PendingUpdates
	m.PendingSecurityUpdates = jm.PendingSecurityUpdates
	m.PendingUpdatesSum = jm.PendingUpdatesSum
	m.PendingUpdatesPackageNames = jm.PendingUpdatesPackageNames
	m.SeverityClass = jm.SeverityClass
	m.UcsRole = jm.UcsRole
	m.BackupType = jm.BackupType
	m.BackupBrand = jm.BackupBrand

	if jm.BackupLastFullRun != "" {
		m.BackupLastFullRun, err = parseTime(parseLayouts, jm.BackupLastFullRun)
		if err != nil {
			return err
		}
	}

	if jm.BackupLastIncRun != "" {
		m.BackupLastIncRun, err = parseTime(parseLayouts, jm.BackupLastIncRun)
		if err != nil {
			return err
		}
	}

	if jm.BackupLastDiffRun != "" {
		m.BackupLastDiffRun, err = parseTime(parseLayouts, jm.BackupLastDiffRun)
		if err != nil {
			return err
		}
	}

	m.BackupLastFullSize = jm.BackupLastFullSize
	m.BackupLastIncSize = jm.BackupLastIncSize
	m.BackupLastDiffSize = jm.BackupLastDiffSize
	m.Nics = jm.Nics

	return nil
}

func (m Machine) MarshalJSON() ([]byte, error) {
	var jm jsonMachine

	jm.Fqdn = m.Fqdn
	jm.Os = m.Os
	jm.Arch = m.Arch
	jm.RAM = m.RAM
	jm.Cores = m.Cores
	jm.Diskspace = m.Diskspace
	jm.Vmhost = m.Vmhost
	jm.Description = m.Description

	if !m.ServicedAt.IsZero() {
		jm.ServicedAt = m.ServicedAt.Format(formatLayout)
	}

	if !m.DeletedAt.IsZero() {
		jm.DeletedAt = m.DeletedAt.Format(formatLayout)
	}

	if !m.CreatedAt.IsZero() {
		jm.CreatedAt = m.CreatedAt.Format(formatLayout)
	}

	if !m.UpdatedAt.IsZero() {
		jm.UpdatedAt = m.UpdatedAt.Format(formatLayout)
	}

	jm.OsRelease = m.OsRelease
	jm.Uptime = m.Uptime
	jm.DeviceTypeID = m.DeviceTypeID
	jm.Serialnumber = m.Serialnumber
	jm.OwnerID = m.OwnerID
	jm.AutoUpdate = m.AutoUpdate
	jm.SwitchURL = m.SwitchURL
	jm.MrtgURL = m.MrtgURL
	jm.ConfigInstructions = m.ConfigInstructions
	jm.SwCharacteristics = m.SwCharacteristics
	jm.BusinessPurpose = m.BusinessPurpose
	jm.BusinessCriticality = m.BusinessCriticality
	jm.BusinessNotification = m.BusinessNotification
	jm.UnattendedUpgrades = m.UnattendedUpgrades
	jm.UnattendedUpgradesBlacklistedPackages = m.UnattendedUpgradesBlacklistedPackages
	jm.UnattendedUpgradesReboot = m.UnattendedUpgradesReboot
	jm.UnattendedUpgradesTime = m.UnattendedUpgradesTime
	jm.UnattendedUpgradesRepos = m.UnattendedUpgradesRepos
	jm.PendingUpdates = m.PendingUpdates
	jm.PendingSecurityUpdates = m.PendingSecurityUpdates
	jm.PendingUpdatesSum = m.PendingUpdatesSum
	jm.PendingUpdatesPackageNames = m.PendingUpdatesPackageNames
	jm.SeverityClass = m.SeverityClass
	jm.UcsRole = m.UcsRole
	jm.BackupType = m.BackupType
	jm.BackupBrand = m.BackupBrand

	if !m.BackupLastFullRun.IsZero() {
		jm.BackupLastFullRun = m.BackupLastFullRun.Format(formatLayout)
	}

	if !m.BackupLastIncRun.IsZero() {
		jm.BackupLastIncRun = m.BackupLastIncRun.Format(formatLayout)
	}

	if !m.BackupLastDiffRun.IsZero() {
		jm.BackupLastDiffRun = m.BackupLastDiffRun.Format(formatLayout)
	}

	jm.BackupLastFullSize = m.BackupLastFullSize
	jm.BackupLastIncSize = m.BackupLastIncSize
	jm.BackupLastDiffSize = m.BackupLastDiffSize
	jm.Nics = m.Nics

	jm.CreateMachine = m.CreateMachine

	return json.Marshal(jm)
}

// Machine represents a IDB machine entry.
type Machine struct {
	// Fully qualified domain name of the machine.
	Fqdn string

	// Operating system.
	Os string

	// Machine architecture.
	Arch string

	// Size of the RAM in Mebibytes
	RAM int

	// Number of CPU cores.
	Cores int

	// Disk space in Mebibytes
	Diskspace int

	// Vmhost
	Vmhost string

	// Textual description of the machine.
	Description string

	// Date of last service.
	ServicedAt time.Time

	// Deletion date.
	DeletedAt time.Time

	// Creation date.
	CreatedAt time.Time

	// Update date.
	UpdatedAt time.Time

	// Operating system release information.
	OsRelease string

	// Machine uptime.
	Uptime int

	// Type of the machine (physical, virtual, switch, ...)
	DeviceTypeID DeviceType

	// Device serial number.
	Serialnumber string

	// ID of the owner in IDB.
	OwnerID int

	// ???
	AutoUpdate bool

	// ???
	SwitchURL string

	// ???
	MrtgURL string

	// ???
	ConfigInstructions string

	// ???
	SwCharacteristics string

	// ???
	BusinessPurpose string

	// ???
	BusinessCriticality string

	// ???
	BusinessNotification string

	// ???
	UnattendedUpgrades bool

	// ???
	UnattendedUpgradesBlacklistedPackages string

	// ???
	UnattendedUpgradesReboot bool

	// ???
	UnattendedUpgradesTime string

	// ???
	UnattendedUpgradesRepos string

	// ???
	PendingUpdates int

	// ???
	PendingSecurityUpdates int

	// ???
	PendingUpdatesSum int

	// ???
	PendingUpdatesPackageNames string

	// ???
	SeverityClass string

	// ???
	UcsRole string

	// Backup type (if this machine has backups, etc.. See BackupType documentation).
	BackupType BackupType

	// Which product is used for the backup.
	BackupBrand BackupBrand

	// Last complete full backup date.
	BackupLastFullRun time.Time

	// Last complete incremental backup date.
	BackupLastIncRun time.Time

	// Last complete differiential backup date.
	BackupLastDiffRun time.Time

	// Sum of the size of all completed full backups
	BackupLastFullSize int64

	// Sum of the size of all completed incremental backups
	BackupLastIncSize int64

	// Sum of the size of all completed differential backups
	BackupLastDiffSize int64

	// Network interfaces
	Nics []Nic

	// Create machine if not existing already. This will be set by the Update method of idbclient.Idb
	// and is only exported to be visible for marshalling.
	CreateMachine bool
}

// Backup is a convienience method to fill the fields needed for a update of the backup data.
// Times are omitted by marshalling when their IsZero() method returns true. Sizes can be 0 to omit them.
func (m *Machine) Backup(fqdn string, brand BackupBrand, lastFull, lastInc, lastDiff time.Time, sizeFull, sizeInc, sizeDiff int64) (err error) {
	m.Fqdn = fqdn
	m.BackupBrand = brand
	m.BackupLastFullRun = lastFull
	m.BackupLastIncRun = lastInc
	m.BackupLastDiffRun = lastDiff

	m.BackupLastFullSize = sizeFull
	m.BackupLastIncSize = sizeInc
	m.BackupLastDiffSize = sizeDiff

	return nil
}

// Equal tests for equality of machine objects
func Equal(m1, m2 *Machine) bool {
	e := m1.Fqdn == m2.Fqdn
	e = e && m1.Os == m2.Os
	e = e && m1.Arch == m2.Arch
	e = e && m1.RAM == m2.RAM
	e = e && m1.Cores == m2.Cores
	e = e && m1.Diskspace == m2.Diskspace
	e = e && m1.Vmhost == m2.Vmhost
	e = e && m1.Description == m2.Description
	e = e && m1.ServicedAt == m2.ServicedAt
	e = e && m1.DeletedAt == m2.DeletedAt
	e = e && m1.CreatedAt == m2.CreatedAt
	e = e && m1.UpdatedAt == m2.UpdatedAt
	e = e && m1.OsRelease == m2.OsRelease
	e = e && m1.Uptime == m2.Uptime
	e = e && m1.DeviceTypeID == m2.DeviceTypeID
	e = e && m1.Serialnumber == m2.Serialnumber
	e = e && m1.OwnerID == m2.OwnerID
	e = e && m1.AutoUpdate == m2.AutoUpdate
	e = e && m1.SwitchURL == m2.SwitchURL
	e = e && m1.MrtgURL == m2.MrtgURL
	e = e && m1.ConfigInstructions == m2.ConfigInstructions
	e = e && m1.SwCharacteristics == m2.SwCharacteristics
	e = e && m1.BusinessPurpose == m2.BusinessPurpose
	e = e && m1.BusinessCriticality == m2.BusinessCriticality
	e = e && m1.BusinessNotification == m2.BusinessNotification
	e = e && m1.UnattendedUpgrades == m2.UnattendedUpgrades
	e = e && m1.UnattendedUpgradesBlacklistedPackages == m2.UnattendedUpgradesBlacklistedPackages
	e = e && m1.UnattendedUpgradesReboot == m2.UnattendedUpgradesReboot
	e = e && m1.UnattendedUpgradesTime == m2.UnattendedUpgradesTime
	e = e && m1.UnattendedUpgradesRepos == m2.UnattendedUpgradesRepos
	e = e && m1.PendingUpdates == m2.PendingUpdates
	e = e && m1.PendingSecurityUpdates == m2.PendingSecurityUpdates
	e = e && m1.PendingUpdatesSum == m2.PendingUpdatesSum
	e = e && m1.PendingUpdatesPackageNames == m2.PendingUpdatesPackageNames
	e = e && m1.SeverityClass == m2.SeverityClass
	e = e && m1.UcsRole == m2.UcsRole
	e = e && m1.BackupType == m2.BackupType
	e = e && m1.BackupBrand == m2.BackupBrand
	e = e && m1.BackupLastFullRun == m2.BackupLastFullRun
	e = e && m1.BackupLastIncRun == m2.BackupLastIncRun
	e = e && m1.BackupLastDiffRun == m2.BackupLastDiffRun
	e = e && m1.BackupLastFullSize == m2.BackupLastFullSize
	e = e && m1.BackupLastIncSize == m2.BackupLastIncSize
	e = e && m1.BackupLastDiffSize == m2.BackupLastDiffSize

	if len(m1.Nics) == len(m2.Nics) {
		for i, v := range m1.Nics {
			e = e && v == m2.Nics[i]
		}
	} else {
		e = false
	}

	return e
}

// Nic represents a network interface
type Nic struct {
	// IPAddress of the interface
	IPAddress IPAddress `json:"ip_address"`

	// Interface name, eg. "eth0"
	Name string `json:"name"`
}

// IPAddress is a pair of address and netmask
type IPAddress struct {
	// IPv4 Address in dotted decimal form
	Addr string `json:"addr,omitempty"`

	// IPv4 Netmask in dotted decimal form
	Netmask string `json:"netmask,omitempty"`

	// IPv6 Address
	AddrV6 string `json:"addr_v6,omitempty"`
	
	// IPv6 Prefix
	NetmaskV6 string `json:"netmask_v6,omitempty"`
}
