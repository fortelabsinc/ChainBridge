// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

const DefaultConfigPath = "./config.json"
const DefaultKeystorePath = "./keys"
const DefaultBlockTimeout = int64(180) // 3 minutes
const EnvironmentConfigFlag = "FORTE_RELAYER_CONFIG_FROM_ENV"

type Config struct {
	Chains       []RawChainConfig `json:"chains"`
	KeystorePath string           `json:"keystorePath,omitempty"`
}

// RawChainConfig is parsed directly from the config file and should be using to construct the core.ChainConfig
type RawChainConfig struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Id       string            `json:"id"`       // ChainID
	Endpoint string            `json:"endpoint"` // url for rpc endpoint
	From     string            `json:"from"`     // address of key to use
	Opts     map[string]string `json:"opts"`
}

func NewConfig() *Config {
	return &Config{
		Chains: []RawChainConfig{},
	}
}

func (c *Config) ToJSON(file string) *os.File {
	var (
		newFile *os.File
		err     error
	)

	var raw []byte
	if raw, err = json.Marshal(*c); err != nil {
		log.Warn("error marshalling json", "err", err)
		os.Exit(1)
	}

	newFile, err = os.Create(file)
	if err != nil {
		log.Warn("error creating config file", "err", err)
	}
	_, err = newFile.Write(raw)
	if err != nil {
		log.Warn("error writing to config file", "err", err)
	}

	if err := newFile.Close(); err != nil {
		log.Warn("error closing file", "err", err)
	}
	return newFile
}

func (c *Config) validate() error {
	for _, chain := range c.Chains {
		if chain.Type == "" {
			return fmt.Errorf("required field chain.Type empty for chain %s", chain.Id)
		}
		if chain.Endpoint == "" {
			return fmt.Errorf("required field chain.Endpoint empty for chain %s", chain.Id)
		}
		if chain.Name == "" {
			return fmt.Errorf("required field chain.Name empty for chain %s", chain.Id)
		}
		if chain.Id == "" {
			return fmt.Errorf("required field chain.Id empty for chain %s", chain.Id)
		}
		if chain.From == "" {
			return fmt.Errorf("required field chain.From empty for chain %s", chain.Id)
		}
	}
	return nil
}

func GetConfig(ctx *cli.Context) (*Config, error) {
	var fig Config
	path := DefaultConfigPath
	if file := ctx.String(ConfigFileFlag.Name); file != "" {
		path = file
	}
	_, configFromEnv := os.LookupEnv(EnvironmentConfigFlag)
	if configFromEnv {
		err := loadEnvConfig(&fig)
		if err != nil {
			log.Warn("err loading json file", "err", err.Error())
			return &fig, err
		}
	} else {
		err := loadConfig(path, &fig)
		if err != nil {
			log.Warn("err loading json file", "err", err.Error())
			return &fig, err
		}
	}

	if ksPath := ctx.String(KeystorePathFlag.Name); ksPath != "" {
		fig.KeystorePath = ksPath
	}
	log.Debug("Loaded config", "path", path)
	err := fig.validate()
	if err != nil {
		return nil, err
	}
	return &fig, nil
}

func loadConfig(file string, config *Config) error {
	ext := filepath.Ext(file)
	fp, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	log.Debug("Loading configuration", "path", filepath.Clean(fp))

	f, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return err
	}

	if ext == ".json" {
		if err = json.NewDecoder(f).Decode(&config); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unrecognized extention: %s", ext)
	}

	return nil
}

func makeSecretsId(projectId string, clusterId string, stageId string, key string) string {
	return projectId + "-" + clusterId + "-" + stageId + "-" + key
}

func makeSecretVersionName(projectId string, secretsId string, versionsId string) string {
	return "/v1/projects/" + projectId + "/secrets/" + secretsId + "/versions/" + versionsId + ":access"
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getSecret(client *secretmanager.Client, key string, fallbackEnvKey ...string) string {
	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: key,
	}

	ctx := context.Background()
	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Debug("failed to access secret version", "err", err)
		envKey := ""
		if len(fallbackEnvKey) > 0 {
			envKey = fallbackEnvKey[0]
		}
		return getEnv(envKey, "")
	}
	return string(result.Payload.Data)
}

func getAuthToken(gcpSATokenUrl string, scopes string) (*oauth2.Token, error) {
	url := gcpSATokenUrl
	if scopes != "" {
		url += "?scopes=" + scopes
	}
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var token *oauth2.Token
	if err := json.Unmarshal(body, &token); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}
	return token, nil
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func loadEnvConfig(config *Config) error {

	// Load GKE Specific Environment Variables
	gcpSATokenUrl := getEnv("GKE_SA_TOKEN_URL", "/some_fallback_url")
	// GCP project in which to store secrets in Secret Manager.
	gcpProject := getEnv("GKE_PROJECT", "some_forte_project_id")
	gcpCluster := getEnv("GKE_CLUSTER", "some_forte_cluster_id")
	// Only needed to set secrets
	//gcpLocation := getEnv("GKE_LOCATION", "some_location")
	gcpStage := getEnv("DEVELOPMENT", "test")

	scopes := getEnv("SCOPES", "")

	//threashold := getEnvAsInt("THRESHOLD", 1)
	//
	//randomKeystorePassword, err := GenerateRandomString(32)
	//if err != nil {
	//	randomKeystorePassword = "spacebacon"
	//}
	//keystorePassword := getEnv("KEYSTORE_PASSWORD", randomKeystorePassword)

	// Default Aleo Chain Options
	aleoChainName := getEnv("ALEO_CHAIN_NAME", "ALEO")
	aleoChainId := getEnv("ALEO_CHAIN_ID", "0")
	aleoChainType := getEnv("ALEO_CHAIN_TYPE", "aleo")
	aleoHost := getEnv("ALEO_HOST", "127.0.0.1")
	aleoPort := getEnv("ALEO_PORT", "4000")
	aleoHttp := getEnv("ALEO_HTTP", "true")

	// Aleo Chain Secrets
	aleoPrivateKey := getEnv("ALEO_PRIVATE_KEY", "APrivateKey1yftg7T4nExo3QiFi84a1zmnYnQKq9988giRCdHrmxsJAvZe")
	aleoViewKey := getEnv("ALEO_VIEW_KEY", "AViewKey1gYQaJfSq2iGvN5ertmzWh1PEuVdseTPukGW9W9dpCkB9")
	aleoAddress := getEnv("ALEO_ADDRESS", "aleo10cu0dcvtqd5m3pfrl4hdpee3upzc4dfjly5png835kz9kzhvpvgsp40npd")

	// Default ETH Chain Options
	ethChainName := getEnv("ETH_CHAIN_NAME", "ETH")
	ethChainId := getEnv("ETH_CHAIN_ID", "1")
	ethChainType := getEnv("ETH_CHAIN_TYPE", "ethereum")
	ethHost := getEnv("ETH_HOST", "127.0.0.1")
	//ethPort := getEnv("ETH_PORT", "8545")
	ethWSPort := getEnv("ETH_WS_PORT", "7545")
	//ethGateway := "http://" + ethHost + ":" + ethPort
	ethEndpoint := "ws://" + ethHost + ":" + ethWSPort
	//ethNetworkId := getEnv("ETH_NETWORK_ID", "196285")
	//ethHttp := getEnv("ETH_HTTP", "false")
	ethBridgeAddress := getEnv("ETH_BRIDGE_ADDRESS", "0x498504b4691dF9dd2C58E73EA1D212df297a7b9F")
	ethERC721Handler := getEnv("ETH_ERC721_HANDLER", "0xA24Fb7a54dD0C906ecF63654384a8fAdE91801b1")
	ethGasLimit := getEnv("ETH_GAS_LIMIT", "1000000")
	ethMaxGasPrice := getEnv("ETH_MAX_GAS_PRICE", "10000000000")

	// ETH Chain Secrets
	ethPrivateKey := getEnv("ETH_PRIVATE_KEY", "2dd574512132480f1a81d46157806bbee5dc712a526185dabf925112e54f2aaa")
	ethAddress := getEnv("ETH_ADDRESS", "0xced59731d42bf1079e08f04a703a9b5a379d70b4")

	// Get Auth Token From Google secrets
	token, err := getAuthToken(gcpSATokenUrl, scopes)

	if token != nil && err == nil {
		// Create secrets manager client.
		ctx := context.Background()
		tokenSource := oauth2.StaticTokenSource(token)
		clientOptions := []option.ClientOption{option.WithTokenSource(tokenSource)}

		client, err := secretmanager.NewClient(ctx, clientOptions...)

		if client != nil && err == nil {
			defer client.Close()
			aleoPrivateKeySecretId := makeSecretsId(gcpProject, gcpCluster, gcpStage, "aleo-private-key")
			aleoPrivateKeySecretVersion := makeSecretVersionName(gcpProject, aleoPrivateKeySecretId, "latest")

			aleoPrivateKey = getSecret(client, aleoPrivateKeySecretVersion, aleoPrivateKey)

			aleoViewKeySecretId := makeSecretsId(gcpProject, gcpCluster, gcpStage, "aleo-view-key")
			aleoViewKeySecretVersion := makeSecretVersionName(gcpProject, aleoViewKeySecretId, "latest")

			aleoViewKey = getSecret(client, aleoViewKeySecretVersion, aleoViewKey)

			aleoAddressSecretId := makeSecretsId(gcpProject, gcpCluster, gcpStage, "aleo-address")
			aleoAddressSecretVersion := makeSecretVersionName(gcpProject, aleoAddressSecretId, "latest")

			aleoAddress = getSecret(client, aleoAddressSecretVersion, aleoAddress)

			ethAddressSecretId := makeSecretsId(gcpProject, gcpCluster, gcpStage, "eth-address")
			ethAddressSecretVersion := makeSecretVersionName(gcpProject, ethAddressSecretId, "latest")

			ethAddress = getSecret(client, ethAddressSecretVersion, ethAddress)

			ethPrivateKeySecretId := makeSecretsId(gcpProject, gcpCluster, gcpStage, "eth-private-key")
			ethPrivateKeySecretVersion := makeSecretVersionName(gcpProject, ethPrivateKeySecretId, "latest")

			ethPrivateKey = getSecret(client, ethPrivateKeySecretVersion, ethPrivateKey)
		}
	}

	//_, err = importPrivKey(ctx, keytype, dHandler.datadir, ethPrivateKey, keystorePassword)

	ethChainConfig := RawChainConfig{
		Name:     ethChainName,
		Type:     ethChainType,
		Id:       ethChainId,
		Endpoint: ethEndpoint,
		From:     ethAddress,
		Opts: map[string]string{
			"bridge":        ethBridgeAddress,
			"erc721Handler": ethERC721Handler,
			"gasLimit":      ethGasLimit,
			"maxGasPrice":   ethMaxGasPrice,
		},
	}

	aleoChainConfig := RawChainConfig{
		Name:     aleoChainName,
		Type:     aleoChainType,
		Id:       aleoChainId,
		Endpoint: "http://" + aleoHost + ":" + aleoPort,
		From:     aleoAddress,
		Opts: map[string]string{
			"http": aleoHttp,
		},
	}

	log.Info("Aleo Chain Config", aleoChainConfig)
	log.Info("ETH Chain Config", ethChainConfig)

	config.Chains = []RawChainConfig{}

	return nil
}
