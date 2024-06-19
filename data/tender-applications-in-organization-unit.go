package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

// TenderApplicationsInOrganizationUnit struct
type TenderApplicationsInOrganizationUnit struct {
	ID                             int        `db:"id,omitempty"`
	JobTenderID                    int        `db:"job_tender_id"`
	UserProfileID                  *int       `db:"user_profile_id"`
	Active                         bool       `db:"active"`
	IsInternal                     bool       `db:"is_internal"`
	FirstName                      *string    `db:"first_name"`
	LastName                       *string    `db:"last_name"`
	Nationality                    *string    `db:"nationality"`
	DateOfBirth                    *time.Time `db:"date_of_birth"`
	DateOfApplication              *time.Time `db:"date_of_application"`
	OfficialPersonalDocumentNumber *string    `db:"official_personal_document_number"`
	Evaluation                     *int       `db:"evaluation"`
	Status                         string     `db:"status"`
	FileID                         *int       `db:"file_id"`
	CreatedAt                      time.Time  `db:"created_at,omitempty"`
	UpdatedAt                      time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *TenderApplicationsInOrganizationUnit) Table() string {
	return "tender_applications_in_organization_units"
}

// GetAll gets all records from the database, using Upper
func (t *TenderApplicationsInOrganizationUnit) GetAll(page *int, pageSize *int, condition *up.AndExpr) ([]*TenderApplicationsInOrganizationUnit, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*TenderApplicationsInOrganizationUnit
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order by")
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *TenderApplicationsInOrganizationUnit) Get(id int) (*TenderApplicationsInOrganizationUnit, error) {
	var one TenderApplicationsInOrganizationUnit
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *TenderApplicationsInOrganizationUnit) Update(m TenderApplicationsInOrganizationUnit) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *TenderApplicationsInOrganizationUnit) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *TenderApplicationsInOrganizationUnit) Insert(m TenderApplicationsInOrganizationUnit) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
