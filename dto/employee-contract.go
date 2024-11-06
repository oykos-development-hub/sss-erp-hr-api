package dto

import (
	"encoding/json"
	"log"
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetEmployeeContracts struct {
	Active *bool `json:"active"`
}

type EmployeeContractDTO struct {
	UserProfileID      int        `json:"user_profile_id" validate:"required"`
	ContractTypeID     int        `json:"contract_type_id" validate:"required"`
	OrganizationUnitID int        `json:"organization_unit_id" validate:"required"`
	DepartmentID       *int       `json:"organization_unit_department_id"`
	Abbreviation       *string    `json:"abbreviation"`
	NumberOfConference *string    `json:"number_of_conference"`
	Description        *string    `json:"description"`
	Active             bool       `json:"active"`
	SerialNumber       *string    `json:"serial_number"`
	NetSalary          *string    `json:"net_salary"`
	GrossSalary        *string    `json:"gross_salary"`
	BankAccount        *string    `json:"bank_account"`
	BankName           *string    `json:"bank_name"`
	DateOfSignature    *time.Time `json:"date_of_signature"`
	DateOfEligibility  *time.Time `json:"date_of_eligibility"`
	DateOfStart        *time.Time `json:"date_of_start"`
	DateOfEnd          *time.Time `json:"date_of_end"`
	FileIDs            []int      `json:"file_ids"`
}

func (dto EmployeeContractDTO) ToEmployeeContract() *data.EmployeeContract {
	marshaledFileIds, err := json.Marshal(dto.FileIDs)
	if err != nil {
		log.Println(err)
	}

	return &data.EmployeeContract{
		UserProfileID:      dto.UserProfileID,
		ContractTypeID:     dto.ContractTypeID,
		OrganizationUnitID: dto.OrganizationUnitID,
		DepartmentID:       dto.DepartmentID,
		Abbreviation:       dto.Abbreviation,
		Description:        dto.Description,
		Active:             dto.Active,
		NumberOfConference: dto.NumberOfConference,
		SerialNumber:       dto.SerialNumber,
		NetSalary:          dto.NetSalary,
		GrossSalary:        dto.GrossSalary,
		BankAccount:        dto.BankAccount,
		BankName:           dto.BankName,
		DateOfSignature:    dto.DateOfSignature,
		DateOfEligibility:  dto.DateOfEligibility,
		DateOfStart:        dto.DateOfStart,
		DateOfEnd:          dto.DateOfEnd,
		FileID:             string(marshaledFileIds),
	}
}

type EmployeeContractResponseDTO struct {
	ID                 int        `json:"id"`
	UserProfileID      int        `json:"user_profile_id"`
	ContractTypeID     int        `json:"contract_type_id"`
	OrganizationUnitID int        `json:"organization_unit_id"`
	DepartmentID       *int       `json:"organization_unit_department_id"`
	Abbreviation       *string    `json:"abbreviation"`
	Description        *string    `json:"description"`
	Active             bool       `json:"active"`
	NumberOfConference *string    `json:"number_of_conference"`
	SerialNumber       *string    `json:"serial_number"`
	NetSalary          *string    `json:"net_salary"`
	GrossSalary        *string    `json:"gross_salary"`
	BankAccount        *string    `json:"bank_account"`
	BankName           *string    `json:"bank_name"`
	DateOfSignature    *time.Time `json:"date_of_signature"`
	DateOfEligibility  *time.Time `json:"date_of_eligibility"`
	DateOfStart        *time.Time `json:"date_of_start"`
	DateOfEnd          *time.Time `json:"date_of_end"`
	FileIDs            []int      `json:"file_ids"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

func ToEmployeeContractResponseDTO(data data.EmployeeContract) EmployeeContractResponseDTO {
	var decodedFileIDs []int
	err := json.Unmarshal([]byte(data.FileID), &decodedFileIDs)
	if err != nil {
		log.Println(err)
	}

	return EmployeeContractResponseDTO{
		ID:                 data.ID,
		UserProfileID:      data.UserProfileID,
		ContractTypeID:     data.ContractTypeID,
		OrganizationUnitID: data.OrganizationUnitID,
		DepartmentID:       data.DepartmentID,
		Abbreviation:       data.Abbreviation,
		Description:        data.Description,
		Active:             data.Active,
		NumberOfConference: data.NumberOfConference,
		SerialNumber:       data.SerialNumber,
		NetSalary:          data.NetSalary,
		GrossSalary:        data.GrossSalary,
		BankAccount:        data.BankAccount,
		BankName:           data.BankName,
		DateOfSignature:    data.DateOfSignature,
		DateOfEligibility:  data.DateOfEligibility,
		DateOfStart:        data.DateOfStart,
		DateOfEnd:          data.DateOfEnd,
		FileIDs:            decodedFileIDs,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToEmployeeContractListResponseDTO(employeecontract []*data.EmployeeContract) []EmployeeContractResponseDTO {
	dtoList := make([]EmployeeContractResponseDTO, len(employeecontract))
	for i, x := range employeecontract {
		dtoList[i] = ToEmployeeContractResponseDTO(*x)
	}
	return dtoList
}
