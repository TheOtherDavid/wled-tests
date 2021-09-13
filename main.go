package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	c "github.com/TheOtherDavid/wled-tests/config"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")

	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variables...")

	fmt.Println("WLED URL: ", configuration.Wled.Url)
	baseUrl := configuration.Wled.Url

	sleepTime := 3000 * time.Millisecond

	fmt.Println("Sending Level 0")
	sendWledCommand(baseUrl, configuration.Wled.LevelZeroBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 1")
	sendWledCommand(baseUrl, configuration.Wled.LevelOneBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 2")
	sendWledCommand(baseUrl, configuration.Wled.LevelTwoBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 3")
	sendWledCommand(baseUrl, configuration.Wled.LevelThreeBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 4")
	sendWledCommand(baseUrl, configuration.Wled.LevelFourBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 5")
	sendWledCommand(baseUrl, configuration.Wled.LevelFiveBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 4")
	sendWledCommand(baseUrl, configuration.Wled.LevelFourBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 3")
	sendWledCommand(baseUrl, configuration.Wled.LevelThreeBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 2")
	sendWledCommand(baseUrl, configuration.Wled.LevelTwoBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 1")
	sendWledCommand(baseUrl, configuration.Wled.LevelOneBody)

	time.Sleep(sleepTime)

	fmt.Println("Sending Level 0")
	sendWledCommand(baseUrl, configuration.Wled.LevelZeroBody)

	fmt.Println("Complete")
}

func sendWledCommand(baseUrl string, body string) error {
	jsonBody := []byte(body)

	fullUrl := "http://" + baseUrl + "/json/state"

	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Timeout = time.Second * 5

	resp, err := client.Do(req)
	if resp != nil {
		println("Response code " + strconv.Itoa(resp.StatusCode))
	} else {
		println("Error code " + err.Error())
	}
	return err
}
