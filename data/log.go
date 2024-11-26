package data

import (
	"encoding/json"
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

type LogOperation string
type LogEntity string

var (
	OperationInsert LogOperation = "INSERT"
	OperationUpdate LogOperation = "UPDATE"
	OperationDelete LogOperation = "DELETE"
)

var (
	EntityOrganzationUnit       LogEntity = "organization_units"
	EntityJobPositions          LogEntity = "job_positions"
	EntityUserProfiles          LogEntity = "user_profiles"
	EntitySystematizations      LogEntity = "systematizations"
	EntityContracts             LogEntity = "employee_contracts"
	EntityResolutions           LogEntity = "employee_resolutions"
	EntityAbsents               LogEntity = "employee_absents"
	EntityTenders               LogEntity = "tenders_in_organization_units"
	EntityUserNorm              LogEntity = "user_norm_fulfilments"
	EntitySalary                LogEntity = "salaries"
	EntityJudgeNumberResolution LogEntity = "judge_number_resolutions"
	EntityPlan                  LogEntity = "plans"
	EntityRevision              LogEntity = "revisions"
	EntityRevisionTips          LogEntity = "revision_tips"
)

// Log struct
type Log struct {
	ID        int              `db:"id,omitempty"`
	ChangedAt time.Time        `db:"changed_at"`
	UserID    int              `db:"user_id"`
	ItemID    int              `db:"item_id"`
	Operation LogOperation     `db:"operation"`
	Entity    LogEntity        `db:"entity"`
	OldState  *json.RawMessage `db:"old_state"`
	NewState  *json.RawMessage `db:"new_state"`
}

// Table returns the table name
func (t *Log) Table() string {
	return "logs"
}

// GetAll gets all records from the database, using Upper
func (t *Log) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*Log, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Log
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

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *Log) Get(id int) (*Log, error) {
	var one Log
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Log) Update(m Log) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Log) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Log) Insert(m Log) (int, error) {
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
