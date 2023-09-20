package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// EmployeeExperience struct
type EmployeeExperience struct {
	ID                        int       `db:"id,omitempty"`
	UserProfileID             int       `db:"user_profile_id"`
	Relevant                  bool      `db:"relevant"`
	OrganizationUnit          *string   `db:"organization_unit"`
	OrganizationUnitID        *int      `db:"organization_unit_id"`
	AmountOfExperience        *int      `db:"amount_of_experience"`
	AmountOfInsuredExperience *int      `db:"amount_of_insured_experience"`
	DateOfSignature           time.Time `db:"date_of_signature"`
	DateOfStart               time.Time `db:"date_of_start"`
	DateOfEnd                 time.Time `db:"date_of_end"`
	FileID                    int       `db:"file_id"`
	CreatedAt                 time.Time `db:"created_at,omitempty"`
	UpdatedAt                 time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeExperience) Table() string {
	return "employee_experiences"
}

// GetAll gets all records from the database, using upper
func (t *EmployeeExperience) GetAll(condition *up.Cond) ([]*EmployeeExperience, error) {
	collection := upper.Collection(t.Table())
	var all []*EmployeeExperience
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("updated_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *EmployeeExperience) Get(id int) (*EmployeeExperience, error) {
	var one EmployeeExperience
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *EmployeeExperience) Update(m EmployeeExperience) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *EmployeeExperience) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeeExperience) Insert(m EmployeeExperience) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
