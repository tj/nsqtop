package nsqd

import "encoding/json"
import "net/http"
import "fmt"

// NSQD client.
type NSQD struct {
	Address string
}

// New client with the http address.
func New(addr string) *NSQD {
	return &NSQD{
		Address: addr,
	}
}

// Stats for topics and channels.
func (n *NSQD) Stats() (*Stats, error) {
	url := fmt.Sprintf("http://%s/stats?format=json", n.Address)

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var s *stats
	err = json.NewDecoder(req.Body).Decode(&s)
	return s.Data, err
}
