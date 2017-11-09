package quotes

import (
	"github.com/coreos/bbolt"
)

// DB is a quote database.
type DB struct {
	db *bolt.DB
}

const (
	quoteBucket = "standard"
)

// Open opens the database file at path and returns a DB or an error.
func Open(path string) (*DB, error) {

	// TODO:
	// Open the DB file at path using bolt.Open().
	// Pass 0600 as file mode, and nil as options.
	// Return a pointer to the open DB, or an error.

}

func (d *DB) Close() error {

	// TODO:
	// Close the database d.db and return any
	// error or nil.

}

// Create takes a quote and saves it to the database, using the author name
// as the key. If the author already exists, Create returns an error.
func (d *DB) Create(q *Quote) error {
	err := d.db.Update(func(tx *bolt.Tx) error {

		// TODO: Create a bucket if it does not exist already.
		// Use the constant quoteBucket as the bucket name.
		//
		// Remember to use []byte(...) to convert a string into a byte
		// slice if required.
		//
		// Ensure that the quote we want to save does not already exist.
		// Hint: Call bucket.Get and verify if the result has zero length.
		//
		// Serialize the quote, using the Serialize method from quote.go.
		//
		// Put the serialized quote into the bucket.
		//
		// Remember to check all errors.

	})

	// TODO: Check the error returned by d.db.Update. Return an error or nil.
}

// Get takes an author name and retrieves the corresponding quote from the DB.
func (d *DB) Get(author string) (*Quote, error) {
	q := &Quote{}
	err := d.db.View(func(tx *bolt.Tx) error {

		// TODO:
		// Get the bucket with the name as specified by the constant quoteBucket.
		// The bucket is available from the transaction object tx.
		//
		// Get the desired quote by the author's name.
		//
		// Again, remember to use []byte(...) to convert a string into a byte
		// slice if required.
		//
		// Deserialize the quote into q - use the Deserialize method from
		// quote.go for this.
		// Remember that we are within a closure that has access to q.
		//
		// Check and return any error that occurs.
	})
	// Check the error returned by d.db.View.
	// Return (nil, err) or (&q, nil), respectively.
}

// List lists all records in the DB.
func (d *DB) List() ([]*Quote, error) {
	// The database returns byte slices that we need to de-serialize
	// into Quote structures.
	structList := []*Quote{}

	// We use a View as we don't update anything.
	err := d.db.View(func(tx *bolt.Tx) error {

		// TODO:
		// Get the bucket from the transaction tx.
		//
		// Iterate over all elements of the bucket.
		// Hint: BoltDB has a ForEach method for this.
		//   * For each element, create a new *Quote and deserialize
		//     the element value into the *Quote.
		//   * Then append the *Quote to structList.
		//
		// Check and return any errors.
	})

	// TODO: Check the error returned by d.db.View().
	// Return (structList, nil) or (nil, err), respectively.
}
