package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionRevisorServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RevisionRevisor
}

func NewRevisionRevisorServiceImpl(app *celeritas.Celeritas, repo data.RevisionRevisor) RevisionRevisorService {
	return &RevisionRevisorServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionRevisorServiceImpl) CreateRevisionRevisor(input dto.RevisionRevisorDTO) (*dto.RevisionRevisorResponseDTO, error) {
	data := input.ToRevisionRevisor()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRevisionRevisorResponseDTO(*data)

	return &res, nil
}

func (h *RevisionRevisorServiceImpl) UpdateRevisionRevisor(id int, input dto.RevisionRevisorDTO) (*dto.RevisionRevisorResponseDTO, error) {
	data := input.ToRevisionRevisor()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRevisionRevisorResponseDTO(*data)

	return &response, nil
}

func (h *RevisionRevisorServiceImpl) DeleteRevisionRevisor(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RevisionRevisorServiceImpl) GetRevisionRevisor(id int) (*dto.RevisionRevisorResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRevisionRevisorResponseDTO(*data)

	return &response, nil
}

func (h *RevisionRevisorServiceImpl) GetRevisionRevisorList(input dto.RevisionRevisorFilter) ([]dto.RevisionRevisorResponseDTO, error) {
	cond := up.Cond{}
	if input.RevisionID != nil {
		cond["revision_id"] = input.RevisionID
	}
	if input.RevisorID != nil {
		cond["revisor_id"] = input.RevisorID
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToRevisionRevisorListResponseDTO(data)

	return response, nil
}