package main

import (
	"fmt"
	"net/http"
	"strings"
	"os/exec"
	"testing"
    "github.com/stretchr/testify/assert"

)

var (

	subscriptionID = "< >"
	resourceGroupName = "< >"
	appGatewayName = "< >"
	actualAppGateway = true
	
)



func TestAppGatewayMK(t *testing.T) {

	accessToken, err := accessToken()
	if err != nil {
		fmt.Println("Error generating access token:", err)
		return
	}

	expectedAppGateway, err := checkApplicationGatewayExists(subscriptionID, resourceGroupName, appGatewayName, accessToken)
	
	fmt.Print(expectedAppGateway)

	t.Run("Existence", func(t *testing.T) {
	assert.Equal(t, actualAppGateway, expectedAppGateway, "Not exisits")
	})
}




func accessTokenMK() (string, error) {
	cmd := exec.Command("az", "account", "get-access-token", "--query", "accessToken", "--output", "tsv")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}




func checkGatewayExistsMK(subscriptionID, resourceGroupName, applicationGatewayName, accessToken string) (bool, error) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/applicationGateways/%s?api-version=2019-09-01",
		subscriptionID, resourceGroupName, applicationGatewayName)

	req, err := http.NewRequest("GET", url, nil)
	fmt.Print(req)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

