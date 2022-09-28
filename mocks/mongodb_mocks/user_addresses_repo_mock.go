package mongodbmocks

import (
	"auth-and-db-service/repositories/mongodb"
)

type MockUserAddressesRepo struct{}

func (mua MockUserAddressesRepo) InsertUserAddress(userRepo mongodb.IUsersRepo, userMail, addressName, name, surname, phoneNumber, province, county, address string) error {
	return nil
}

func (mua MockUserAddressesRepo) UpdateUserAdress(addressName, name, surname, phoneNumber, province, county, address string) error {
	return nil
}

func (mua MockUserAddressesRepo) DeleteUserAddress(userMail, addressName string) error {
	return nil
}

func (mua MockUserAddressesRepo) FindUserAddress(addressID string) (*mongodb.Address, error) {
	var address = mongodb.Address{
		Id:              "123",
		Title:           "testTitle",
		Name:            "testName",
		Surname:         "testSurname",
		PhoneNumber:     "5441111111",
		Province:        "testProvince",
		County:          "testCounty",
		DetailedAddress: "testAddress",
	}
	return &address, nil
}

func (mua MockUserAddressesRepo) FindAllUserAddresses(userRepo mongodb.IUsersRepo, userMail string) ([]mongodb.Address, error) {
	addresses := []mongodb.Address{
		{
			Id:              "123",
			Title:           "testTile",
			Name:            "testName",
			Surname:         "testSurname",
			PhoneNumber:     "5441111111",
			Province:        "testProvince",
			County:          "testCounty",
			DetailedAddress: "testAddress",
		},
		{
			Id:              "123",
			Title:           "testTile",
			Name:            "testName",
			Surname:         "testSurname",
			PhoneNumber:     "5441111111",
			Province:        "testProvince",
			County:          "testCounty",
			DetailedAddress: "testAddress",
		},
	}
	return addresses, nil
}
