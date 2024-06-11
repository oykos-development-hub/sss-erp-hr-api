package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
)

// EmployeeContract struct
type EmployeeContract struct {
	ID                 int        `db:"id,omitempty"`
	UserProfileID      int        `db:"user_profile_id"`
	ContractTypeID     int        `db:"contract_type_id"`
	OrganizationUnitID int        `db:"organization_unit_id"`
	DepartmentID       *int       `db:"organization_unit_department_id"`
	NumberOfConference *string    `db:"number_of_conference"`
	Abbreviation       *string    `db:"abbreviation"`
	Description        *string    `db:"description"`
	Active             bool       `db:"active"`
	SerialNumber       *string    `db:"serial_number"`
	NetSalary          *string    `db:"net_salary"`
	GrossSalary        *string    `db:"gross_salary"`
	BankAccount        *string    `db:"bank_account"`
	BankName           *string    `db:"bank_name"`
	DateOfSignature    *time.Time `db:"date_of_signature"`
	DateOfEligibility  *time.Time `db:"date_of_eligibility"`
	DateOfStart        *time.Time `db:"date_of_start"`
	DateOfEnd          *time.Time `db:"date_of_end"`
	FileID             *int       `db:"file_id"`
	CreatedAt          time.Time  `db:"created_at,omitempty"`
	UpdatedAt          time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeContract) Table() string {
	return "employee_contracts"
}

// Delete deletes a record from the database by id, using Upper
func (t *EmployeeContract) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (t *EmployeeContract) GetByUserProfileId(userProfileId int, condition *up.Cond) ([]*EmployeeContract, error) {
	var employeeContracts []*EmployeeContract

	collection := Upper.Collection(t.Table())
	res := collection.Find(up.Cond{"user_profile_id =": userProfileId})

	if condition != nil {
		res = res.And(*condition)
	}

	err := res.OrderBy("created_at desc").All(&employeeContracts)
	if err != nil {
		fmt.Println(err)
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	return employeeContracts, nil
}

// Delete deletes a records from the database for user id, using Upper
func (t *EmployeeContract) DeleteForUser(ctx context.Context, userID int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(up.Cond{"user_profile_id": userID})
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *EmployeeContract) Insert(ctx context.Context, m EmployeeContract) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates a record in the database, using Upper
func (t *EmployeeContract) Update(ctx context.Context, m EmployeeContract) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Get gets one record from the database, by id, using Upper
func (t *EmployeeContract) Get(id int) (*EmployeeContract, error) {
	var one EmployeeContract
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}

	return &one, nil
}
