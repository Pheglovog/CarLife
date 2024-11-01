package gateway

import (
	"carlife-chaincode-go/chaincode"
	"crypto/x509"
	"encoding/json"
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
)

var GatewayMap = make(map[string]*client.Gateway)
var channel = "carchannel"
var SmartContract = "cartrace"

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

	orgMap := make(map[string]OrgConfig)
	for _, org := range orgConfigs.Organizations {
		orgMap[org.MSPID] = org
	}
	return orgMap, nil
}

func InitGateways(configFilePath string) {
	orgConfigs, err := loadOrgConfigs(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load org configs: %v", err)
	}

	for mspID, config := range orgConfigs {
		GatewayMap[mspID] = newGateway(config)
	}
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
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
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
	case chaincode.ComponentSupplier:
		mspID = "ComponentSupplierMSP"
	case chaincode.Manufacturer:
		mspID = "ManufacturerMSP"
	case chaincode.Store:
		mspID = "StoreMSP"
	case chaincode.Insurer:
		mspID = "InsurerMSP"
	case chaincode.Maintenancer:
		mspID = "MaintenancerMSP"
	case chaincode.Consumer:
		mspID = "StoreMSP"
	default:
		return "", fmt.Errorf("user type %s is not supported", userType)
	}
	gw := GatewayMap[mspID]
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
	gw := GatewayMap["StoreMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.EvaluateTransaction("GetUser", userID)
	return string(result), err
}

func GetCar(carID string) (string, error) {
	gw := GatewayMap["StoreMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.EvaluateTransaction("GetCar", carID)
	return string(result), err
}

func SetCarTires(userID string, carID string, width float32, radius float32, workshop string) (string, error) {
	gw := GatewayMap["ComponentSupplierMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarTires",
		userID,
		carID,
		strconv.FormatFloat(float64(width), 'f', 2, 32),
		strconv.FormatFloat(float64(radius), 'f', 2, 32),
		workshop,
	)
	return string(result), err
}

func SetCarBody(userID string, carID string, material string,
	weitght float32, color string, workshop string) (string, error) {
	gw := GatewayMap["ComponentSupplierMSP"]
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
	)
	return string(result), err
}

func SetCarInterior(userID string, carID string, material string,
	weitght float32, color string, workshop string) (string, error) {
	gw := GatewayMap["ComponentSupplierMSP"]
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
	)
	return string(result), err
}

func SetCarManu(userID string, carID string, workshop string) (string, error) {
	gw := GatewayMap["ManufacturerMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarManu",
		userID,
		carID,
		workshop,
	)
	return string(result), err
}

func SetCarStore(userID string, carID string, store string,
	cost float32, ownerID string) (string, error) {
	gw := GatewayMap["StoreMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarStore",
		userID,
		carID,
		store,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		ownerID,
	)
	return string(result), err
}

func SetCarInsure(userID string, carID string, name string,
	cost float32, years int) (string, error) {
	gw := GatewayMap["InsurerMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarInsure",
		userID,
		carID,
		name,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
		strconv.Itoa(years),
	)
	return string(result), err
}

func SetCarMaint(userID string, carID string, part string,
	entent string, cost float32) (string, error) {
	gw := GatewayMap["MaintenancerMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"SetCarMaint",
		userID,
		carID,
		part,
		entent,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
	)
	return string(result), err
}

func TransferCar(userID string, carID string, newUserID string, cost float32) (string, error) {
	gw := GatewayMap["StoreMSP"]
	net := gw.GetNetwork(channel)
	contract := net.GetContract(SmartContract)
	result, err := contract.SubmitTransaction(
		"TransferCar",
		userID,
		carID,
		newUserID,
		strconv.FormatFloat(float64(cost), 'f', 2, 32),
	)
	return string(result), err
}
