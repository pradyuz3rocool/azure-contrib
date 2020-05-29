package iothub

import (
	// "bytes"
	// "crypto/hmac"
	// "crypto/sha256"
	// "encoding/base64"
	// "fmt"
	// "html/template"
	// "io"
	// "io/ioutil"
	// "net/http"
	// "net/url"
	// "strconv"
	// "strings"
	// "time"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

// const {
// 	maxIdleConnections int    = 100
// 	requestTimeout     int    = 10
// 	tokenValidSecs     int    = 3600
// 	apiVersion         string = "2016-11-14"
// }

// type sharedAccessKey = string
// type sharedAccessKeyName = string
// type hostName = string
// type deviceID = string

// IotHubHTTPClient is a simple client to connect to Azure IoT Hub
// type IotHubHTTPClient struct {
// 	sharedAccessKeyName sharedAccessKeyName
// 	sharedAccessKey     sharedAccessKey
// 	hostName            hostName
// 	deviceID            deviceID
// 	client              *http.Client
// }

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings *Settings
	logger   log.Logger
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Debugf("DeviceId: %s", input.DeviceId)

	//connectionString, methodType := a.settings.AZURE_IOTHUB_CONNECTION_STRING, a.settings.Operation
	//deviceID := a.input.DeviceId
	
	// output := &Output{}
	
	// connectionString, methodType := a.settings.AZURE_IOTHUB_CONNECTION_STRING, a.settings.Operation
	// deviceID := input.DeviceId

	// client, err := NewIotHubHTTPClientFromConnectionString(connectionString)
	// if err != nil {
	// 	log.Error("Error creating http client from connection string", err)
	// }

	// switch methodType {
	// case "Add Device":
	// 	resp, status := client.CreateDeviceID(deviceID)
	// 	ctx.SetOutputObject(output, resp)
	// 	//ctx.SetOutputObject(ovStatus, status)
	// case "Delete Device":
	// 	resp, status := client.DeleteDeviceID(deviceID)
	// 	ctx.SetOutputObject(output, resp)
	// 	//ctx.SetOutputObject(ovStatus, status)
	// case "Purge Device":
	// 	resp, status := client.PurgeCommandsForDeviceID(deviceID)
	// 	ctx.SetOutputObject(output, resp)
	// 	//ctx.SetOutputObject(ovStatus, status)
	// case "Get Twin Details":
	// 	resp, status := client.GetDeviceTwin(deviceID)
	// 	ctx.SetOutputObject(output, resp)
	// 	//ctx.SetOutputObject(ovStatus, status)
	// }

	return true, nil
}

// func parseConnectionString(connString string) (hostName, sharedAccessKey, sharedAccessKeyName, deviceID, error) {
// 	url, err := url.ParseQuery(connString)
// 	if err != nil {
// 		return "", "", "", "", err
// 	}

// 	h := tryGetKeyByName(url, "HostName")
// 	kn := tryGetKeyByName(url, "SharedAccessKeyName")
// 	k := tryGetKeyByName(url, "SharedAccessKey")
// 	d := tryGetKeyByName(url, "DeviceId")

// 	return hostName(h), sharedAccessKey(k), sharedAccessKeyName(kn), deviceID(d), nil
// }

// func tryGetKeyByName(v url.Values, key string) string {
// 	if len(v[key]) == 0 {
// 		return ""
// 	}

// 	return strings.Replace(v[key][0], " ", "+", -1)
// }

// // NewIotHubHTTPClient is a constructor of IutHubClient
// func NewIotHubHTTPClient(hostName string, sharedAccessKeyName string, sharedAccessKey string, deviceID string) *IotHubHTTPClient {
// 	return &IotHubHTTPClient{
// 		sharedAccessKeyName: sharedAccessKeyName,
// 		sharedAccessKey:     sharedAccessKey,
// 		hostName:            hostName,
// 		deviceID:            deviceID,
// 		client: &http.Client{
// 			Transport: &http.Transport{
// 				MaxIdleConnsPerHost: maxIdleConnections,
// 			},
// 			Timeout: time.Duration(requestTimeout) * time.Second,
// 		},
// 	}
// }

// // NewIotHubHTTPClientFromConnectionString creates new client from connection string
// func NewIotHubHTTPClientFromConnectionString(connectionString string) (*IotHubHTTPClient, error) {
// 	h, k, kn, d, err := parseConnectionString(connectionString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return NewIotHubHTTPClient(h, kn, k, d), nil
// }

// // IsDevice tell either device id was specified when client created.
// // If device id was specified in connection string this will enabled device scoped requests.
// func (c *IotHubHTTPClient) IsDevice() bool {
// 	return c.deviceID != ""
// }

// // Service API

// // CreateDeviceID creates record for for given device identifier in Azure IoT Hub
// func (c *IotHubHTTPClient) CreateDeviceID(deviceID string) (string, string) {
// 	url := fmt.Sprintf("%s/devices/%s?api-version=%s", c.hostName, deviceID, apiVersion)
// 	data := fmt.Sprintf(`{"deviceId":"%s"}`, deviceID)
// 	return c.performRequest("PUT", url, data)
// }

// // GetDeviceID retrieves device by id
// func (c *IotHubHTTPClient) GetDeviceID(deviceID string) (string, string) {
// 	url := fmt.Sprintf("%s/devices/%s?api-version=%s", c.hostName, deviceID, apiVersion)
// 	return c.performRequest("GET", url, "")
// }

// // DeleteDeviceID deletes device by id
// func (c *IotHubHTTPClient) DeleteDeviceID(deviceID string) (string, string) {
// 	url := fmt.Sprintf("%s/devices/%s?api-version=%s", c.hostName, deviceID, apiVersion)
// 	return c.performRequest("DELETE", url, "")
// }

// // PurgeCommandsForDeviceID removed commands for specified device
// func (c *IotHubHTTPClient) PurgeCommandsForDeviceID(deviceID string) (string, string) {
// 	url := fmt.Sprintf("%s/devices/%s/commands?api-version=%s", c.hostName, deviceID, apiVersion)
// 	return c.performRequest("DELETE", url, "")
// }

// // ListDeviceIDs list all device ids
// func (c *IotHubHTTPClient) ListDeviceIDs() (string, string) {
// 	url := fmt.Sprintf("%s/devices/query?api-version=%s", c.hostName, apiVersion)
// 	return c.performRequest("GET", url, "")
// }

// // Get Device Twin
// func (c *IotHubHTTPClient) GetDeviceTwin(deviceID string) (string, string) {
// 	url := fmt.Sprintf("%s/twins/%s?api-version=2018-06-30", c.hostName, deviceID)
// 	return c.performRequest("GET", url, "")
// }

// // // TODO: SendMessageToDevice as soon as that endpoint is exposed via HTTP

// // // Device API

// // // SendMessage from a logged in device
// // func (c *IotHubHTTPClient) SendMessage(message string) (string, string) {
// // 	url := fmt.Sprintf("%s/devices/%s/messages/events?api-version=%s", c.hostName, c.deviceID, apiVersion)
// // 	return c.performRequest("POST", url, message)
// // }

// // // ReceiveMessage to a logged in device
// // func (c *IotHubHTTPClient) ReceiveMessage() (string, string) {
// // 	url := fmt.Sprintf("%s/devices/%s/messages/deviceBound?api-version=%s", c.hostName, c.deviceID, apiVersion)
// // 	return c.performRequest("GET", url, "")
// // }

// func (c *IotHubHTTPClient) buildSasToken(uri string) string {
// 	timestamp := time.Now().Unix() + int64(3600)
// 	encodedURI := template.URLQueryEscaper(uri)

// 	toSign := encodedURI + "\n" + strconv.FormatInt(timestamp, 10)

// 	binKey, _ := base64.StdEncoding.DecodeString(c.sharedAccessKey)
// 	mac := hmac.New(sha256.New, []byte(binKey))
// 	mac.Write([]byte(toSign))

// 	encodedSignature := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

// 	if c.sharedAccessKeyName != "" {
// 		return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&skn=%s&sr=%s", encodedSignature, timestamp, c.sharedAccessKeyName, encodedURI)
// 	}

// 	return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&sr=%s", encodedSignature, timestamp, encodedURI)
// }

// func (c *IotHubHTTPClient) performRequest(method string, uri string, data string) (string, string) {
// 	token := c.buildSasToken(uri)
// 	log.Info("The SaStoken is [%s]", token)
// 	//log.("%s https://%s\n", method, uri)
// 	//log.Printf(data)
// 	req, _ := http.NewRequest(method, "https://"+uri, bytes.NewBufferString(data))

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("User-Agent", "golang-iot-client")
// 	req.Header.Set("Authorization", token)

// 	//log.Println("Authorization:", token)

// 	if method == "DELETE" {
// 		req.Header.Set("If-Match", "*")
// 	}

// 	resp, err := c.client.Do(req)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	// read the entire reply to ensure connection re-use
// 	text, _ := ioutil.ReadAll(resp.Body)

// 	io.Copy(ioutil.Discard, resp.Body)
// 	defer resp.Body.Close()

// 	return string(text), resp.Status
// }