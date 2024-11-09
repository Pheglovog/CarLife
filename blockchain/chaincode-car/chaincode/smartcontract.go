package chaincode

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, password string) error {
	user := User{
		UserID:   userID,
		UserType: userType,
		Password: password,
		CarList:  []string{},
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	v, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return err
	}
	if v != nil {
		return errors.New("User already exists")
	}
	return ctx.GetStub().PutState(userID, userJson)
}

func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userJson, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SmartContract) GetCar(ctx contractapi.TransactionContextInterface, carID string) (*Car, error) {
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return nil, err
	}

	var car Car
	err = json.Unmarshal(carJson, &car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *SmartContract) SetCarTires(ctx contractapi.TransactionContextInterface, userID string,
	carID string, width float32, radius float32, workshop string) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != ComponentSupplier {
		return "", errors.New("only component supplier can set car tires")
	}
	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	Cartires := CarTires{
		Time:     time.Now(),
		Width:    width,
		Radius:   radius,
		Workshop: workshop,
		TxID:     txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		car.CarID = carID
	} else {
		err = json.Unmarshal(carJson, &car)
		if err != nil {
			return "", err
		}
	}
	car.Tires = Cartires
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarBody(ctx contractapi.TransactionContextInterface, userID string,
	carID string, material string, weitght float32, color string, workshop string) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != ComponentSupplier {
		return "", errors.New("only componentSupplier can set car body")
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	Carbody := CarBody{
		Time:     time.Now(),
		Material: material,
		Weitght:  weitght,
		Color:    color,
		Workshop: workshop,
		TxID:     txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		car.CarID = carID
	} else {
		err = json.Unmarshal(carJson, &car)
		if err != nil {
			return "", err
		}
	}
	car.Body = Carbody
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarInterior(ctx contractapi.TransactionContextInterface, userID string,
	carID string, material string, weitght float32, color string, workshop string) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != ComponentSupplier {
		return "", errors.New("only componentSupplier can set car interior")
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	Carinterior := CarInterior{
		Time:     time.Now(),
		Material: material,
		Weitght:  weitght,
		Color:    color,
		Workshop: workshop,
		TxID:     txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		car.CarID = carID
	} else {
		err = json.Unmarshal(carJson, &car)
		if err != nil {
			return "", err
		}
	}
	car.Interior = Carinterior
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarManu(ctx contractapi.TransactionContextInterface, userID string,
	carID string, workshop string) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != Manufacturer {
		return "", errors.New("only manufacturer can set car manu")
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	Carmanu := CarManu{
		Time:     time.Now(),
		Workshop: workshop,
		TxID:     txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		return "", errors.New("car not exist")
	}
	car.Manu = Carmanu
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarStore(ctx contractapi.TransactionContextInterface, userID string,
	carID string, store string, cost float32, ownerID string) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != Store {
		return "", errors.New("only store can set car store")
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	Carstore := CarStore{
		Time:  time.Now(),
		Store: store,
		Cost:  cost,
		Owner: ownerID,
		TxID:  txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		return "", errors.New("car not exist")
	}
	car.Store = Carstore
	car.Owner = ownerID
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarInsure(ctx contractapi.TransactionContextInterface,
	userID string, carID string, name string,
	cost float32, years int) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != Insurer {
		return "", errors.New("only insurer can set car insure")
	}
	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	insure := Insure{
		Name:      name,
		Cost:      cost,
		BeginTime: time.Now(),
		EndTime:   time.Now().AddDate(years, 0, 0),
		TxID:      txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		return "", errors.New("car not exist")
	}
	car.Insure.Insures = append(car.Insure.Insures, insure)
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) SetCarMaint(ctx contractapi.TransactionContextInterface, userID string,
	carID string, part string, extent string, cost float32) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != Maintenancer {
		return "", errors.New("only maintenancer can fix car")
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	maint := Maint{
		Time:   time.Now(),
		Part:   part,
		Extent: extent,
		Cost:   cost,
		TxID:   txID,
	}
	var car Car
	carJson, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return "", err
	}
	if carJson == nil {
		return "", errors.New("car not exist")
	}
	car.Maint.Maints = append(car.Maint.Maints, maint)
	carJson, err = json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}

func (s *SmartContract) TransferCar(ctx contractapi.TransactionContextInterface, userID string,
	carID string, newUserID string, cost float32) (string, error) {
	//check permission
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}
	if user.UserType != Consumer {
		return "", errors.New("only consumer can transfer car")
	}
	car, err := s.GetCar(ctx, carID)
	if err != nil {
		return "", err
	}
	if car == nil {
		return "", errors.New("car not exist")
	}
	if car.Owner != userID {
		return "", errors.New("only owner can transfer car")
	}

	//change car information
	txID := ctx.GetStub().GetTxID()
	record := Record{
		OldUser: userID,
		NewUser: newUserID,
		Time:    time.Now(),
		Cost:    cost,
		TxID:    txID,
	}
	car.Owner = newUserID
	car.Record.Records = append(car.Record.Records, record)
	carJson, err := json.Marshal(car)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(carID, carJson)
	if err != nil {
		return "", err
	}

	//change user information
	user.CarList = append(user.CarList, carID)
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(userID, userJson)
	if err != nil {
		return "", err
	}
	return txID, nil
}
