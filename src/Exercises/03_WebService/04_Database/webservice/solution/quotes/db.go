package quotes

import (
	"github.com/coreos/bbolt"
	"github.com/pkg/errors"
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
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Open: cannot open DB file "+path)
	}
	return &DB{
		db: db,
	}, nil
}

func (d *DB) Close() error {
	err := d.db.Close()
	if err != nil {
		return errors.Wrap(err, "Close: cannot close database")
	}
	return nil
}

// Create takes a quote and saves it to the database, using the author name
// as the key. If the author already exists, Create returns an error.
func (d *DB) Create(q *Quote) error {
	err := d.db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(quoteBucket))
		if err != nil {
			return errors.Wrapf(err, "Create: cannot open or create bucket %s", []byte(quoteBucket))
		}

		// Ensure we do not accidentally update a record
		v := bucket.Get([]byte(q.Author))
		if len(v) > 0 {
			return errors.Errorf("Author %s already exists", q.Author)
		}

		// Turn the Quote into []byte and save it in the bucket.
		b, err := q.Serialize()
		if err != nil {
			return errors.Wrapf(err, "Create: cannot serialize quote from", q.Author)
		}
		err = bucket.Put([]byte(q.Author), b)
		if err != nil {
			return errors.Wrapf(err, "Create: cannot put quote from %s into bucket", q.Author)
		}
		return nil
	})

	if err != nil {
		return errors.Wrapf(err, "Create: cannot create record for (%s|%s|%s)", q.Author, q.Text, q.Source)
	}
	return nil
}

// Get takes an author name and retrieves the corresponding quote from the DB.
func (d *DB) Get(author string) (*Quote, error) {
	q := &Quote{}
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(quoteBucket))
		if bucket == nil {
			return errors.Errorf("Cannot get %s - bucket %s not found", author, quoteBucket)
		}
		v := bucket.Get([]byte(author))
		err := q.Deserialize(v)
		if err != nil {
			return errors.Wrapf(err, "Get: cannot deserialize %s", v)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Get: DB.View() failed")
	}
	return q, nil
}

// List lists all records in the DB.
func (d *DB) List() ([]*Quote, error) {
	// The database returns byte slices that we need to de-serialize
	// into Quote structures.
	structList := []*Quote{}

	// We use a view as we don't update anything.
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(quoteBucket))
		if b == nil {
			// It is valid to attempt listing an empty database.
			// Hence no error is returned in this case.
			return nil
		}
		// ForEach iterates over all elements of a bucket and
		// executes the passed-in function for each element.
		err := b.ForEach(func(k []byte, v []byte) error {
			q := &Quote{}
			err := q.Deserialize(v)
			if err != nil {
				return errors.Wrapf(err, "List: cannot deserialize data of author %s", k)
			}
			// We're inside a closure, so we can access structList
			// that resides in the outer scope.
			structList = append(structList, q)
			return nil
		})
		if err != nil {
			return errors.Wrapf(err, "List: failed iterating over bucket %s", quoteBucket)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "List: view failed")
	}
	return structList, nil
}
