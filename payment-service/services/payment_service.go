package services

//
//import (
//	"bytes"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"encoding/base64"
//	"encoding/json"
//	"encoding/pem"
//	"fmt"
//	"io/ioutil"
//	"mime/multipart"
//	"net/http"
//)
//
//type TokenResponse struct {
//	AccessToken  string `json:"access_token"`
//	ExpiresIn    string `json:"expires_in"`
//	RefreshToken string `json:"refresh_token"`
//	Scope        string `json:"scope"`
//	TokenType    string `json:"token_type"`
//}
//
//type CryptogramResponse struct {
//	Hpan       string `json:"hpan"`
//	ExpDate    string `json:"expDate"`
//	Cvc        string `json:"cvc"`
//	TerminalId string `json:"terminalId"`
//}
//
//type TestData struct {
//	url          string
//	email        string
//	password     string
//	clientID     string
//	clientSecret string
//	terminalID   string
//}
//
//var testData = TestData{url: "https://test-epay.epayment.kz",
//	email:        "epay@halykbank.kz",
//	password:     "XZG1E@Mm",
//	clientID:     "test",
//	clientSecret: "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG",
//	terminalID:   "67e34d63-102f-4bd1-898e-370781d0074d"}
//
//func GetToken() (string, error) {
//	tokenURL := "https://test-epay-oauth.epayment.kz/oauth2/token"
//
//	// Создаем буфер для тела запроса и writer для multipart формы
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//
//	// Добавляем поля формы
//	err := writer.WriteField("grant_type", "client_credentials")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("scope", "webapi usermanagement email_send verification statement statistics payment")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("client_id", "test")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("client_secret", "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("invoiceId", "000000001")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("amount", "100")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("currency", "KZT")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("terminalId", "67e34d63-102f-4bd1-898e-370781d0074d")
//	if err != nil {
//		return "", err
//	}
//	err = writer.WriteField("failurePostLink", "")
//	if err != nil {
//		return "", err
//	}
//
//	// Закрываем writer чтобы отправить все данные
//	err = writer.Close()
//	if err != nil {
//		return "", err
//	}
//
//	req, err := http.NewRequest("POST", tokenURL, body)
//	if err != nil {
//		return "", err
//	}
//
//	// Устанавливаем заголовок Content-Type включая boundary
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("status: %s, body: %s ", resp.Status, respBody)
//	}
//
//	var token TokenResponse
//	err = json.Unmarshal(respBody, &token)
//	if err != nil {
//		return "", err
//	}
//	fmt.Println(token)
//	return token.AccessToken, nil
//}
//
//func GetPublicKey() (*rsa.PublicKey, error) {
//	publicKeyURL := "https://test-epay-api.epayment.kz/public.rsa"
//	resp, err := http.Get(publicKeyURL)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//	block, _ := pem.Decode(body)
//	if block == nil {
//		return nil, fmt.Errorf("failed to decode PEM block containing public key")
//	}
//	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
//	if !ok {
//		return nil, fmt.Errorf("failed to parse RSA public key")
//	}
//	return rsaPublicKey, nil
//}
//
//func encryptData() (string, error) {
//	publicKey, err := GetPublicKey()
//	if err != nil {
//		return "", err
//	}
//	data := map[string]string{
//		"hpan":       "4405639704015096",
//		"expDate":    "0125",
//		"cvc":        "815",
//		"terminalId": "67e34d63-102f-4bd1-898e-370781d0074d",
//	}
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		return "", err
//	}
//	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonData)
//	if err != nil {
//		return "", err
//	}
//	return base64.StdEncoding.EncodeToString(encryptedData), nil
//}
//
//type PaymentResponse struct {
//	Status    string  `json:"status"`
//	Message   string  `json:"message"`
//	PaymentID string  `json:"payment_id"`
//	Amount    float64 `json:"amount"`
//	Currency  string  `json:"currency"`
//	InvoiceID string  `json:"invoice_id"`
//}
//
//func MakePayment() (*PaymentResponse, error) {
//	paymentUrl := "https://test-epay-api.epayment.kz/payment/cryptopay"
//	token, err := GetToken()
//	if err != nil {
//		return nil, fmt.Errorf("failed to get token: %v", err)
//	}
//
//	encryptedData, err := encryptData()
//	if err != nil {
//		return nil, fmt.Errorf("failed to encrypt data: %v", err)
//	}
//
//	body := map[string]interface{}{
//		"amount":      100,
//		"currency":    "KZT",
//		"name":        "JON JONSON",
//		"cryptogram":  encryptedData,
//		"invoiceId":   "000001",
//		"description": "test payment",
//		"cardSave":    true,
//		"postLink":    "https://testmerchant/order/1123",
//	}
//	jsonBody, err := json.Marshal(body)
//	if err != nil {
//		return nil, fmt.Errorf("failed to marshal body: %v", err)
//	}
//
//	req, err := http.NewRequest("POST", paymentUrl, bytes.NewBuffer(jsonBody))
//	if err != nil {
//		return nil, fmt.Errorf("failed to create request: %v", err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+token)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf("failed to perform request: %v", err)
//	}
//	defer resp.Body.Close()
//	var paymentResponse PaymentResponse
//
//	if err = json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
//		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
//	}
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, paymentResponse)
//	}
//	fmt.Println(paymentResponse)
//	return &paymentResponse, nil
//}
