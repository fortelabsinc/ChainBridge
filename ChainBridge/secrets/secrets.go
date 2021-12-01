/* Copyright (c) 2021 Forte Labs
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Forte Labs and their suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Forte Labs and their suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Forte Labs.
 */
// [START secrets]

// Sample quickstart is a basic program that uses Secret Manager.
package main

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

type TokenResponse struct {
	Expires int    `json:"expires_in"`
	Token   string `json:"access_token"`
}

func main() {
	// GCP project in which to store secrets in Secret Manager.
	projectID := "your-project-id"

	gcpSATokenUrl := getEnv("GKE_SA_TOKEN_URL", "/some_fallback_url")
	gcpProject := getEnv("GKE_PROJECT", "some_forte_project_id")
	gcpCluster := getEnv("GKE_CLUSTER", "some_forte_cluster_id")
	gcpLocation := getEnv("GKE_LOCATION", "some_location")
	gcpStage := getEnv("DEVELOPMENT", "test")

	scopes := getEnv("SCOPES", "")

	url := gcpSATokenUrl
	if scopes != "" {
		url += "?scopes=" + scopes
	}
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)

	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Create the request to create the secret.
	//createSecretReq := &secretmanagerpb.CreateSecretRequest{
	//	Parent:   fmt.Sprintf("projects/%s", projectID),
	//	SecretId: "my-secret",
	//	Secret: &secretmanagerpb.Secret{
	//		Replication: &secretmanagerpb.Replication{
	//			Replication: &secretmanagerpb.Replication_Automatic_{
	//				Automatic: &secretmanagerpb.Replication_Automatic{},
	//			},
	//		},
	//	},
	//}

	//secret, err := client.CreateSecret(ctx, createSecretReq)
	//if err != nil {
	//	log.Fatalf("failed to create secret: %v", err)
	//}
	//
	//// Declare the payload to store.
	//payload := []byte("my super secret data")
	//
	//// Build the request.
	//addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
	//	Parent: secret.Name,
	//	Payload: &secretmanagerpb.SecretPayload{
	//		Data: payload,
	//	},
	//}
	//
	//// Call the API.
	//version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
	//if err != nil {
	//	log.Fatalf("failed to add secret version: %v", err)
	//}

	// Build the request.
	//accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
	//	Name: version.Name,
	//}

	// Call the API.
	//result, err := client.AccessSecretVersion(ctx, accessRequest)
	//if err != nil {
	//	log.Fatalf("failed to access secret version: %v", err)
	//}

	// Print the secret payload.
	//
	// WARNING: Do not print the secret in a production environment - this
	// snippet is showing how to access the secret material.
	//log.Printf("Plaintext: %s", result.Payload.Data)
}
