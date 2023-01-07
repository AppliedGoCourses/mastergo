package quotes

import (
	"fmt"
)

// Quote represents a quote, inlcuding its author and an optional source. The ID is a unique key.
type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}

// Serialize returns a gob encoding of quote q.
func (q Quote) Serialize() ([]byte, error) {

	// TODO:
	// Create a bytes.Buffer
	// Create a new gob.Encoder based on this buffer
	// Encode q into the buffer
	// If any error occurred, return nil and the error
	// else return the buffer and no error

}

// Deserialize takes a byte slice that contains a gob-encoded quote
// and turns it back into a Quote.
func (q *Quote) Deserialize(b []byte) error {

	// TODO:
	// Create a new bytes.Buffer from b
	// Create a new gob.Decoder
	// Decode the buffer into q
	// Return any error or nil

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
