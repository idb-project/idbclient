package machine

import (
	"encoding/json"
	"testing"
	"time"
)

var testTime time.Time
var zeroTime time.Time
var marshalTests []struct {
	m Machine
	j string
}

var equalTests []struct {
	m1 Machine
	m2 Machine
}

var unequalTests []struct {
	m1 Machine
	m2 Machine
}

func init() {
	testTime, _ = time.Parse(time.RFC3339Nano, "2006-01-02T15:04:05.000000000Z")

	marshalTests = []struct {
		m Machine
		j string
	}{
		{Machine{Fqdn: "test0", ServicedAt: testTime, CreatedAt: testTime, DeletedAt: testTime, UpdatedAt: testTime, BackupLastFullRun: testTime, BackupLastIncRun: testTime, BackupLastDiffRun: testTime}, `{"fqdn":"test0","create_machine":"false","serviced_at":"2006-01-02 15:04:05","deleted_at":"2006-01-02 15:04:05","created_at":"2006-01-02 15:04:05","updated_at":"2006-01-02 15:04:05","backup_last_full_run":"2006-01-02 15:04:05","backup_last_inc_run":"2006-01-02 15:04:05","backup_last_diff_run":"2006-01-02 15:04:05"}`},

		{Machine{Fqdn: "test1"}, `{"fqdn":"test1","create_machine":"false"}`},

		{Machine{Fqdn: "test2", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, `{"fqdn":"test2","nics":[{"ip_address":{"addr":"127.0.0.1","netmask":"255.255.255.255"},"name":"lo"}],"create_machine":"false"}`},
	}

	equalTests = []struct {
		m1 Machine
		m2 Machine
	}{
		{Machine{Fqdn: "test3", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, Machine{Fqdn: "test3", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}},
	}

	unequalTests = []struct {
		m1 Machine
		m2 Machine
	}{
		{Machine{Fqdn: "test4", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, Machine{Fqdn: "test4", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.0", Netmask: "255.255.255.255"}, Name: "lo"}}}},

		{Machine{Fqdn: "test5", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, Machine{Fqdn: "test5_unqeual", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}},

		{Machine{Fqdn: "test6", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, Machine{Fqdn: "test6", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "0.0.0.0"}, Name: "lo"}}}},

		{Machine{Fqdn: "test8", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "lo"}}}, Machine{Fqdn: "test8", Nics: []Nic{Nic{IPAddress: IPAddress{Addr: "127.0.0.1", Netmask: "255.255.255.255"}, Name: "eth0"}}}},
	}

}

func TestMarshal(t *testing.T) {
	for _, v := range marshalTests {
		j, err := json.Marshal(v.m)
		if err != nil {
			t.Error(err)
		}

		if string(j) != v.j {
			t.Log("Got     :", string(j))
			t.Log("Expected:", v.j)
			t.Fail()
		}
	}
}

func TestUnmarshal(t *testing.T) {
	for _, v := range marshalTests {
		var m Machine
		err := json.Unmarshal([]byte(v.j), &m)
		if err != nil {
			t.Error(err)
		}

		if !Equal(&m, &v.m) {
			t.Fail()
			t.Log("Got     :", m)
			t.Log("Expected:", v.m)
		}
	}
}

func TestEqual(t *testing.T) {
	for _, v := range equalTests {
		if !Equal(&v.m1, &v.m2) {
			t.Fail()
		}
	}
}

func TestUnequal(t *testing.T) {
	for _, v := range unequalTests {
		if Equal(&v.m1, &v.m2) {
			t.Fail()
		}
	}
}
