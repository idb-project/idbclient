package idbclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/idb-project/idbclient/machine"
	"log"
	"net/http"
	"net/url"
	"path"
)

const idbVersion = 2

// ErrStatus is returned if a unexpected HTTP status was returned by the IDB.
type ErrStatus struct {
	status   int
	expected int
	machine  *machine.Machine
}

func newErrStatus(status, expected int, m *machine.Machine) *ErrStatus {
	return &ErrStatus{status, expected, m}
}

func (s *ErrStatus) Error() string {
	return fmt.Sprintf("IDB returned status %v, expected %v. Machine: %+v", s.status, s.expected, s.machine)
}

// Idb contains IDB client functionality.
type Idb struct {
	url   	*url.URL
	transport *http.Transport
	apiToken	string
	Debug     bool
}

// NewIdb creates a new Idb which uses the IDB found at url.
// If the IDB instance does not have a valid SSL certificate, insecureSkipVerify can be used to skip SSL verification.
func NewIdb(apiURL, apiToken string, insecureSkipVerify bool) (*Idb, error) {
	i := new(Idb)

	var err error
	i.url, err = url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	i.url.Path = fmt.Sprintf("/api/v%v", idbVersion)

	i.transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify}}
	i.apiToken = apiToken

	return i, nil
}

// joinBaseURL appends path components to Idb.url
func (i *Idb) joinBaseURL(p ...string) *url.URL {
	u := new(url.URL)
	*u = *i.url	
	u.Path = path.Join(append([]string{u.Path}, p...)...)
	return u
}

// request sends a request to the IDB and returns the response and possible errors.
func (i *Idb) request(r *http.Request) (*http.Response, error) {
	query := r.URL.Query()
	query.Add("idb_api_token", i.apiToken)
	r.URL.RawQuery = query.Encode()

	if i.Debug {
		log.Printf("Request: %+v\n", r)
	}
	client := &http.Client{Transport: i.transport}

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	if i.Debug {
		log.Printf("Response: %+v\n", response)
	}

	return response, err
}

func (i *Idb) decodeResponse(v interface{}, r *http.Response) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMachine submits new values for a machine in IDB.
func (i *Idb) UpdateMachine(m *machine.Machine, create bool) (*machine.Machine, error) {
	var body bytes.Buffer
	enc := json.NewEncoder(&body)

	m.CreateMachine = create

	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PUT", i.joinBaseURL("machines").String(), &body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := i.request(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, newErrStatus(response.StatusCode, http.StatusOK, m)
	}

	var newMachine machine.Machine

	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&newMachine)
	if err != nil {
		return nil, err
	}

	return &newMachine, err
}

// GetMachine retrieves a single machine identified by fqdn.
func (i *Idb) GetMachine(fqdn string) (*machine.Machine, error) {
	fqdn = url.QueryEscape(fqdn)
	u := i.joinBaseURL("machines")

	query := url.Values{}
	query.Add("fqdn",fqdn)
	u.RawQuery = query.Encode()

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := i.request(request)
	if err != nil {
		return nil, err
	}

	var newMachine machine.Machine
	newMachine.Fqdn = fqdn

	if response.StatusCode != http.StatusOK {
		return nil, newErrStatus(response.StatusCode, http.StatusOK, &newMachine)
	}

	err = i.decodeResponse(&newMachine, response)

	return &newMachine, err
}
