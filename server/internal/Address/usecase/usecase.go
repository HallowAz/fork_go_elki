package usecase

import (
	addressRep "server/internal/Address/repository"
	sessionRep "server/internal/Session/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"fmt"
)

type UsecaseI interface {
	CreateAddress(reqAddress *dto.ReqCreateAddress) error
	DeleteAddress(id uint, sessionToken string) error
	//GetAddresses(userId uint) (*dto.RespGetAddresses, error)
	SetAddress(id uint, sessionToken string) error 
}

type addressUsecase struct {
	addressRepo addressRep.AddressRepositoryI
	sessionRepo  sessionRep.SessionRepositoryI
}

func NewAddressUsecase(addressRepI addressRep.AddressRepositoryI, sessionRepI sessionRep.SessionRepositoryI) *addressUsecase {
	return &addressUsecase{
		addressRepo: addressRepI,
		sessionRepo:  sessionRepI,
	}
}

func (ad *addressUsecase) CreateAddress(reqAddress *dto.ReqCreateAddress) error {
	address := dto.ToEntityCreateAddress(reqAddress)

	if len(address.City) == 0 || len(address.Street) == 0 || len(address.House) == 0 {
		return entity.ErrBadRequest
	}

	cookie, err := ad.sessionRepo.Check(reqAddress.Cookie)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	addresses, err := ad.addressRepo.GetAddresses(cookie.UserID)
	if err != nil {
		fmt.Println("get", err)
		return err
	}
	fmt.Println("")
	if addresses != nil {
		for _, checkAd := range addresses.Addresses {
			if checkAd.City == address.City && checkAd.Street == address.Street && checkAd.House == address.House && checkAd.Flat == address.Flat {
				return entity.ErrAddressAlreadyExist
			} 
		}
	}
	return ad.addressRepo.CreateAddress(dto.ToDBCreateAddress(address, cookie.UserID)) 
}

func (ad *addressUsecase) DeleteAddress(id uint, sessionToken string) error {
	cookie, err := ad.sessionRepo.Check(sessionToken)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	address := &dto.DBReqDeleteUserAddress{
		UserId: cookie.UserID,
		AddressId: id,
	}

	return ad.addressRepo.DeleteAddress(address)
}

func (ad *addressUsecase) SetAddress(id uint, sessionToken string) error  {
	
	cookie, err := ad.sessionRepo.Check(sessionToken)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	address := &dto.DBReqUpdateUserAddress{
		UserId: cookie.UserID,
		AddressId: id,
	}

	return ad.addressRepo.SetAddress(address)
}

