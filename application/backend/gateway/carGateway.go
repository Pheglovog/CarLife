package gateway

import (
	"carlife-backend/model"
	"context"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// var GatewayMap = make(map[string]*client.Gateway)
var channel = "carchannel"
var SmartContract = "carlife-chaincode-go"

var orgMap = make(map[string]OrgConfig)

type OrgConfig struct {
	MSPID        string `json:"mspID"`
	CryptoPath   string `json:"cryptoPath"`
	CertPath     string `json:"certPath"`
	KeyPath      string `json:"keyPath"`
	TlsCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
}

func loadOrgConfigs(filePath string) (map[string]OrgConfig, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var orgConfigs struct {
		Organizations []OrgConfig `json:"organizations"`
	}
	if err := json.Unmarshal(file, &orgConfigs); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	for _, org := range orgConfigs.Organizations {
		orgMap[org.MSPID] = org
	}
	return orgMap, nil
}

func InitGateways(configFilePath string) {
	_, err := loadOrgConfigs(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load org configs: %v", err)
	}

	// for mspID, config := range orgConfigs {
	// 	GatewayMap[mspID] = newGateway(config)
	// }
}

func readFirstFile(dirPath string) ([]byte, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := dir.Readdirnames(1)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(dirPath, fileNames[0]))
}

func newIdentity(orgConfig OrgConfig) *identity.X509Identity {
	certificatePEM, err := readFirstFile(orgConfig.CertPath)
	if err != nil {
		panic(fmt.Errorf("failed to read certificate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(orgConfig.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func newSign(orgConfig OrgConfig) identity.Sign {
	privateKeyPEM, err := readFirstFile(orgConfig.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func newGrpcConnection(orgConfig OrgConfig) *grpc.ClientConn {
	certificatePEM, err := os.ReadFile(orgConfig.TlsCertPath)
	if err != nil {
		panic(fmt.Errorf("failed to read TLS certificate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, orgConfig.GatewayPeer)

	connection, err := grpc.NewClient(orgConfig.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func newGateway(orgConfig OrgConfig) *client.Gateway {
	clientConnection := newGrpcConnection(orgConfig)
	id := newIdentity(orgConfig)
	sign := newSign(orgConfig)

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(1*time.Minute),
		client.WithEndorseTimeout(1*time.Minute),
		client.WithSubmitTimeout(1*time.Minute),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}

	return gw
}

func RegisterUser(userID string, userType string, password string) (string, error) {
	var mspID string
	switch userType {
	case model.ComponentSupplier:
		mspID = "ComponentSupplierMSP"
	case model.Manufacturer:
		mspID = "ManufacturerMSP"
	case model.Store:
		mspID = "StoreMSP"
	case model.Insurer:
		mspID = "InsurerMSP"
	case model.Maintenancer:
		mspID = "MaintenancerMSP"
	case model.Consumer:
		mspID = "StoreMSP"
	default:
		return "", fmt.Errorf("user type %s is not supported", userType)
	}
	// gw := GatewayMap[mspID]
	gw := newGateway(orgMap[mspID])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	_, commit, err := contract.SubmitAsync("RegisterUser", client.WithArguments(userID, userType, password))
	if err != nil {
		return "", fmt.Errorf("failed to submit register transaction asynchronously: %w", err)
	}
	if commitStatus, err := commit.Status(); err != nil {
		return "", fmt.Errorf("failed to get commit status: %w", err)
	} else if !commitStatus.Successful {
		return "", fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code))
	}
	return commit.TransactionID(), err
}

func GetUser(userID string) (string, error) {
	gw := newGateway(orgMap["StoreMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.EvaluateTransaction("GetUser", userID)
	return string(result), CheckErr(err)
}

func GetCar(carID string) (string, error) {
	gw := newGateway(orgMap["StoreMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.EvaluateTransaction("GetCar", carID)
	return string(result), CheckErr(err)
}

func SetCarTires(userID string, carID string, width float32, radius float32, workshop string) (string, error) {
	gw := newGateway(orgMap["ComponentSupplierMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	fmt.Println("SetCarTires:", userID, carID, width, radius, workshop)
	result, err := contract.SubmitTransaction(
		"SetCarTires",
		userID,
		carID,
		strconv.FormatFloat(float64(width), 'f', 2, 32),
		strconv.FormatFloat(float64(radius), 'f', 2, 32),
		workshop,
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarBody(userID string, carID string, material string,
	weitght float32, color string, workshop string) (string, error) {
	gw := newGateway(orgMap["ComponentSupplierMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarBody",
		userID,
		carID,
		material,
		strconv.FormatFloat(float64(weitght), 'f', 2, 32),
		color,
		workshop,
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarInterior(userID string, carID string, material string,
	weitght float32, color string, workshop string) (string, error) {
	gw := newGateway(orgMap["ComponentSupplierMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarInterior",
		userID,
		carID,
		material,
		strconv.FormatFloat(float64(weitght), 'f', 2, 32),
		color,
		workshop,
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarManu(userID string, carID string, workshop string) (string, error) {
	gw := newGateway(orgMap["ManufacturerMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarManu",
		userID,
		carID,
		workshop,
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarStore(userID string, carID string, store string,
	cost float32, ownerID string) (string, error) {
	gw := newGateway(orgMap["StoreMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarStore",
		userID,
		carID,
		store,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		ownerID,
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarInsure(userID string, carID string, name string,
	cost float32, years int) (string, error) {
	gw := newGateway(orgMap["InsurerMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarInsure",
		userID,
		carID,
		name,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		strconv.Itoa(years),
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func SetCarMaint(userID string, carID string, part string,
	entent string, cost float32) (string, error) {
	gw := newGateway(orgMap["MaintenancerMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarMaint",
		userID,
		carID,
		part,
		entent,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func TransferCar(userID string, carID string, newUserID string, cost float32) (string, error) {
	gw := newGateway(orgMap["StoreMSP"])
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"TransferCar",
		userID,
		carID,
		newUserID,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		time.Now().Format(time.RFC3339),
	)
	return string(result), CheckErr(err)
}

func CheckErr(err error) error {
	var newErr error
	if err != nil {
		var endorseErr *client.EndorseError
		var submitErr *client.SubmitError
		var commitStatusErr *client.CommitStatusError
		var commitErr *client.CommitError

		if errors.As(err, &endorseErr) {
			newErr = fmt.Errorf("Endorse error for transaction %s with gRPC status %v: %s\n", endorseErr.TransactionID, status.Code(endorseErr), endorseErr)
		} else if errors.As(err, &submitErr) {
			newErr = fmt.Errorf("Submit error for transaction %s with gRPC status %v: %s\n", submitErr.TransactionID, status.Code(submitErr), submitErr)
		} else if errors.As(err, &commitStatusErr) {
			if errors.Is(err, context.DeadlineExceeded) {
				newErr = fmt.Errorf("Timeout waiting for transaction %s commit status: %s", commitStatusErr.TransactionID, commitStatusErr)
			} else {
				newErr = fmt.Errorf("Error obtaining commit status for transaction %s with gRPC status %v: %s\n", commitStatusErr.TransactionID, status.Code(commitStatusErr), commitStatusErr)
			}
		} else if errors.As(err, &commitErr) {
			newErr = fmt.Errorf("Transaction %s failed to commit with status %d: %s\n", commitErr.TransactionID, int32(commitErr.Code), err)
		} else {
			newErr = fmt.Errorf("unexpected error type %T: %w", err, err)
		}

		// Any error that originates from a peer or orderer node external to the gateway will have its details
		// embedded within the gRPC status error. The following code shows how to extract that.
		statusErr := status.Convert(err)

		details := statusErr.Details()
		if len(details) > 0 {
			fmt.Println("Error Details:%v", details)
		}
	}
	return newErr
}
