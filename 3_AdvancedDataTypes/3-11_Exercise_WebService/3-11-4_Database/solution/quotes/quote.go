package quotes

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
)

// Quote represents a quote, inlcuding its author and an optional source. The ID is a unique key.
type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}

// Serialize returns a gob encoding of quote q.
func (q Quote) Serialize() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(q)
	if err != nil {
		return nil, errors.Wrapf(err, "Serialize: encoding failed for %v", q)
	}
	return b.Bytes(), nil
}

// Deserialize takes a byte slice that contains a gob-encoded quote
// and turns it back into a Quote.
func (q *Quote) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(q)
	if err != nil {
		return errors.Wrapf(err, "Deserialize: decoding failed for %s", b)
	}
	return nil
}

// String implements the stringer interface. A Quote can now be used
// everywhere a string is expected.
func (q *Quote) String() string {
	s := fmt.Sprintf("\"%s\"\n\n(%s", q.Text, q.Author)
	if q.Source != "" {
		s += fmt.Sprintf(", %s", q.Source)
	}
	s += fmt.Sprint(")\n")
	return s
}
