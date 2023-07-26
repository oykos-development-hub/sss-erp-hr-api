package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type EmployeeResolutionDTO struct {
	UserProfileID     int       `json:"user_profile_id" validate:"required"`
	ResolutionTypeID  int       `json:"resolution_type_id" validate:"required"`
	ResolutionPurpose *string   `json:"resolution_purpose"`
	DateOfStart       time.Time `json:"date_of_start" validate:"required,datetime"`
	DateOfEnd         time.Time `json:"date_of_end" validate:"required,datetime"`
	FileId            *int      `json:"file_id"`
}

type EmployeeResolutionResponseDTO struct {
	ID                int       `json:"id"`
	UserProfileID     int       `json:"user_profile_id"`
	ResolutionTypeID  int       `json:"resolution_type_id"`
	ResolutionPurpose *string   `json:"resolution_purpose"`
	DateOfStart       time.Time `json:"date_of_start"`
	DateOfEnd         time.Time `json:"date_of_end"`
	FileID            *int      `json:"file_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (dto EmployeeResolutionDTO) ToEmployeeResolution() *data.EmployeeResolution {
	return &data.EmployeeResolution{
		UserProfileID:     dto.UserProfileID,
		ResolutionTypeID:  dto.ResolutionTypeID,
		ResolutionPurpose: dto.ResolutionPurpose,
		DateOfStart:       dto.DateOfStart,
		DateOfEnd:         dto.DateOfEnd,
		FileID:            dto.FileId,
	}
}

func ToEmployeeResolutionResponseDTO(data data.EmployeeResolution) EmployeeResolutionResponseDTO {
	return EmployeeResolutionResponseDTO{
		ID:                data.ID,
		UserProfileID:     data.UserProfileID,
		ResolutionTypeID:  data.ResolutionTypeID,
		ResolutionPurpose: data.ResolutionPurpose,
		DateOfStart:       data.DateOfStart,
		DateOfEnd:         data.DateOfEnd,
		FileID:            data.FileID,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func ToEmployeeResolutionListResponseDTO(employeeresolutions []*data.EmployeeResolution) []EmployeeResolutionResponseDTO {
	dtoList := make([]EmployeeResolutionResponseDTO, len(employeeresolutions))
	for i, x := range employeeresolutions {
		dtoList[i] = ToEmployeeResolutionResponseDTO(*x)
	}
	return dtoList
}
