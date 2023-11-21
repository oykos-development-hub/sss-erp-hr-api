package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionsInOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RevisionsInOrganizationUnit
}

func NewRevisionsInOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.RevisionsInOrganizationUnit) RevisionsInOrganizationUnitService {
	return &RevisionsInOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionsInOrganizationUnitServiceImpl) CreateRevisionsInOrganizationUnit(input dto.RevisionsInOrganizationUnitDTO) (*dto.RevisionsInOrganizationUnitResponseDTO, error) {
	data := input.ToRevisionsInOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRevisionsInOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *RevisionsInOrganizationUnitServiceImpl) UpdateRevisionsInOrganizationUnit(id int, input dto.RevisionsInOrganizationUnitDTO) (*dto.RevisionsInOrganizationUnitResponseDTO, error) {
	data := input.ToRevisionsInOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRevisionsInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *RevisionsInOrganizationUnitServiceImpl) DeleteRevisionsInOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RevisionsInOrganizationUnitServiceImpl) GetRevisionsInOrganizationUnit(id int) (*dto.RevisionsInOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRevisionsInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *RevisionsInOrganizationUnitServiceImpl) GetRevisionsInOrganizationUnitList(input dto.RevisionOrgUnitFilter) ([]dto.RevisionsInOrganizationUnitResponseDTO, error) {
	cond := up.Cond{}
	if input.RevisionID != nil {
		cond["revision_id"] = input.RevisionID
	}
	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = input.OrganizationUnitID
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToRevisionsInOrganizationUnitListResponseDTO(data)

	return response, nil
}