package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeeEducationServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeEducation
}

func NewEmployeeEducationServiceImpl(app *celeritas.Celeritas, repo data.EmployeeEducation) EmployeeEducationService {
	return &EmployeeEducationServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeEducationServiceImpl) CreateEmployeeEducation(input dto.EmployeeEducationDTO) (*dto.EmployeeEducationResponseDTO, error) {
	data := input.ToEmployeeEducation()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee education insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee education get")
	}

	res := dto.ToEmployeeEducationResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeEducationServiceImpl) UpdateEmployeeEducation(id int, input dto.EmployeeEducationDTO) (*dto.EmployeeEducationResponseDTO, error) {
	data := input.ToEmployeeEducation()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee education update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee education get")
	}

	response := dto.ToEmployeeEducationResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeEducationServiceImpl) DeleteEmployeeEducation(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee education get")
	}

	return nil
}

func (h *EmployeeEducationServiceImpl) GetEmployeeEducationList(input dto.EducationInput) ([]dto.EmployeeEducationResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": input.UserProfileID,
	}

	if input.SubTypeID != nil {
		cond["sub_type_id"] = *input.SubTypeID
	}
	if input.TypeID != nil {
		cond["type_id"] = *input.TypeID
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee education get all")
	}
	response := dto.ToEmployeeEducationListResponseDTO(data)

	return response, nil
}
