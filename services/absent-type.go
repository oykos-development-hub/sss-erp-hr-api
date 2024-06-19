package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"

	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type AbsentTypeServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.AbsentType
}

func NewAbsentTypeServiceImpl(app *celeritas.Celeritas, repo data.AbsentType) AbsentTypeService {
	return &AbsentTypeServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *AbsentTypeServiceImpl) CreateAbsentType(input dto.AbsentTypeDTO) (*dto.AbsentTypeResponseDTO, error) {
	data := input.ToAbsentType()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo absent type insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo absent type get")
	}

	res := dto.ToAbsentTypeResponseDTO(*data)

	return &res, nil
}

func (h *AbsentTypeServiceImpl) UpdateAbsentType(id int, input dto.AbsentTypeDTO) (*dto.AbsentTypeResponseDTO, error) {
	data := input.ToAbsentType()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo absent type update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo absent type get")
	}

	response := dto.ToAbsentTypeResponseDTO(*data)

	return &response, nil
}

func (h *AbsentTypeServiceImpl) DeleteAbsentType(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo absent type delete")
	}

	return nil
}

func (h *AbsentTypeServiceImpl) GetAbsentType(id int) (*dto.AbsentTypeResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo absent type get")
	}
	response := dto.ToAbsentTypeResponseDTO(*data)

	return &response, nil
}

func (h *AbsentTypeServiceImpl) GetAbsentTypeList(data dto.GetAbesntTypeDTO) ([]dto.AbsentTypeResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if data.ParentID != nil {
		cond["parent_id"] = *data.ParentID
	}

	res, total, err := h.repo.GetAll(data.Page, data.Size, &cond)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo absent type get all")
	}
	response := dto.ToAbsentTypeListResponseDTO(res)

	return response, total, nil
}
