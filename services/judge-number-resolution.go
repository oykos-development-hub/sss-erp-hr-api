package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type JudgeNumberResolutionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.JudgeNumberResolution
}

func NewJudgeNumberResolutionServiceImpl(app *celeritas.Celeritas, repo data.JudgeNumberResolution) JudgeNumberResolutionService {
	return &JudgeNumberResolutionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *JudgeNumberResolutionServiceImpl) CreateJudgeNumberResolution(ctx context.Context, input dto.JudgeNumberResolutionDTO) (*dto.JudgeNumberResolutionResponseDTO, error) {
	data := input.ToJudgeNumberResolution()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	if data.Active {
		err = h.repo.InactivateOtherResolutions(ctx, id)
		if err != nil {
			return nil, errors.ErrInternalServer
		}
	}

	res := dto.ToJudgeNumberResolutionResponseDTO(*data)

	return &res, nil
}

func (h *JudgeNumberResolutionServiceImpl) UpdateJudgeNumberResolution(ctx context.Context, id int, input dto.JudgeNumberResolutionDTO) (*dto.JudgeNumberResolutionResponseDTO, error) {
	data := input.ToJudgeNumberResolution()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToJudgeNumberResolutionResponseDTO(*data)

	return &response, nil
}

func (h *JudgeNumberResolutionServiceImpl) DeleteJudgeNumberResolution(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *JudgeNumberResolutionServiceImpl) GetJudgeNumberResolution(id int) (*dto.JudgeNumberResolutionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToJudgeNumberResolutionResponseDTO(*data)

	return &response, nil
}

func (h *JudgeNumberResolutionServiceImpl) GetJudgeNumberResolutionList(input dto.GetJudgeNumberResolutionInputDTO) ([]dto.JudgeNumberResolutionResponseDTO, *uint64, error) {
	cond := up.Cond{}

	if input.Active != nil {
		cond["active"] = input.Active
	}

	data, total, err := h.repo.GetAll(input.Page, input.PageSize, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToJudgeNumberResolutionListResponseDTO(data)

	return response, total, nil
}
