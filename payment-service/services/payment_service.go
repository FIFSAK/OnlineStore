package services

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type CryptogramResponse struct {
	Hpan       string `json:"hpan"`
	ExpDate    string `json:"expDate"`
	Cvc        string `json:"cvc"`
	TerminalId string `json:"terminalId"`
}

func GetToken() (string, error) {
	tokenURL := "https://testoauth.homebank.kz/epay2/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "webapi usermanagement email_send verification statement statistics payment")
	data.Set("client_id", "test")
	data.Set("client_secret", "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}
	var token TokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {
	publicKeyURL := "https://testepay.homebank.kz/api/public.rsa"
	resp, err := http.Get(publicKeyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(body)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse RSA public key")
	}
	return rsaPublicKey, nil
}

func encryptData() (string, error) {
	publicKey, err := GetPublicKey()
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"hpan":       "4405639704015096",
		"expDate":    "0125",
		"cvc":        "815",
		"terminalId": "67e34d63-102f-4bd1-898e-370781d0074d",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func CreatePayment() (string, error) {
	paymentUrl := "https://testepay.homebank.kz/api/payment/cryptopay"
	token, err := GetToken()
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	encryptedData, err := encryptData()
	if err != nil {
		return "", fmt.Errorf("failed to encrypt data: %v", err)
	}

	body := map[string]interface{}{
		"amount":          100,
		"currency":        "KZT",
		"name":            "JON JONSON",
		"cryptogram":      encryptedData,
		"invoiceId":       "000001",
		"invoiceIdAlt":    "8564546",
		"description":     "test payment",
		"accountId":       "uuid000001",
		"email":           "jj@example.com",
		"phone":           "77777777777",
		"cardSave":        true,
		"data":            `{"statement":{"name":"Arman Ali","invoiceID":"80000016"}}`,
		"postLink":        "https://testmerchant/order/1123",
		"failurePostLink": "https://testmerchant/order/1123/fail",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal body: %v", err)
	}

	req, err := http.NewRequest("POST", paymentUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %s body: %s", resp.Status, respBody)
	}

	return string(respBody), nil
}
