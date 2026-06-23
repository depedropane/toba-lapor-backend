package usecase

import (
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
)

type AgencyUsecase interface {
	GetAllAgencies() ([]dto.AgencyResponse, error)
	GetAgencyByID(id uint) (*dto.AgencyResponse, error)
	CreateAgency(req dto.CreateAgencyRequest) (*dto.AgencyResponse, error)
	UpdateAgency(id uint, req dto.UpdateAgencyRequest) (*dto.AgencyResponse, error)
	DeleteAgency(id uint) error
}

type agencyUsecase struct {
	agencyRepo repository.AgencyRepository
}

func NewAgencyUsecase(agencyRepo repository.AgencyRepository) AgencyUsecase {
	return &agencyUsecase{agencyRepo}
}

func (u *agencyUsecase) GetAllAgencies() ([]dto.AgencyResponse, error) {
	agencies, err := u.agencyRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.AgencyResponse
	for _, a := range agencies {
		res = append(res, dto.AgencyResponse{
			ID:          a.ID,
			Name:        a.Name,
			Description: a.Description,
		})
	}
	return res, nil
}

func (u *agencyUsecase) GetAgencyByID(id uint) (*dto.AgencyResponse, error) {
	agency, err := u.agencyRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.AgencyResponse{
		ID:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
	}, nil
}

func (u *agencyUsecase) CreateAgency(req dto.CreateAgencyRequest) (*dto.AgencyResponse, error) {
	agency := &model.Agency{
		Name:        req.Name,
		Description: req.Description,
	}
	err := u.agencyRepo.Create(agency)
	if err != nil {
		return nil, err
	}

	return &dto.AgencyResponse{
		ID:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
	}, nil
}

func (u *agencyUsecase) UpdateAgency(id uint, req dto.UpdateAgencyRequest) (*dto.AgencyResponse, error) {
	agency, err := u.agencyRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	agency.Name = req.Name
	agency.Description = req.Description

	err = u.agencyRepo.Update(agency)
	if err != nil {
		return nil, err
	}

	return &dto.AgencyResponse{
		ID:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
	}, nil
}

func (u *agencyUsecase) DeleteAgency(id uint) error {
	_, err := u.agencyRepo.FindByID(id)
	if err != nil {
		return err
	}
	return u.agencyRepo.Delete(id)
}
