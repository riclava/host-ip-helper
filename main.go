package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/spf13/viper"
)

// Constants
const (
	AppName = "HostIpHelper"
)

// App .
type App struct {
	config *AppConfig
}

// AppConfig .
type AppConfig struct {
	IntervalSeconds int
	Callback        []string
}

// Request of http
type Request struct {
	ID    string
	Addrs []string
}

func httpPost(url string, contentType string, payload string) (int, string, error) {
	resp, err := http.Post(url, contentType, strings.NewReader(payload))
	if err != nil {
		return 0, "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode, string(body), nil
}

func getInterfaceAddrs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []string

	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			result = append(result, ip.String())
		}
	}
	return result, nil
}

func getMachineID() (string, error) {
	id, err := machineid.ProtectedID(AppName)
	if err != nil {
		return "", err
	}
	return id, nil
}

// NewApp create an app instance
func NewApp() (*App, error) {
	viper.SetConfigName("config.yaml")           // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.host-ip-helper") // call multiple times to add many search paths
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	err := viper.ReadInConfig()                  // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	app := &App{
		config: &AppConfig{
			IntervalSeconds: viper.GetInt("interval"),
			Callback:        viper.GetStringSlice("callback"),
		},
	}

	return app, nil
}

func (a *App) work() {

	machineID, err := getMachineID()
	if err != nil {
		log.Fatal(err)
	}

	contentType := "applicaton/json"

	for {
		addrs, err := getInterfaceAddrs()
		if err != nil {
			log.Println(err)
		} else {
			request := &Request{
				ID:    machineID,
				Addrs: addrs,
			}

			payload, _ := json.Marshal(request)

			for _, url := range a.config.Callback {
				code, _, err := httpPost(url, contentType, string(payload))
				if err != nil {
					log.Printf("Error send data to %s due to %s", url, err.Error())
				} else {
					log.Printf("Send data to %s success [code = %d]", url, code)
				}
			}
		}

		time.Sleep(time.Duration(a.config.IntervalSeconds) * time.Second)
	}
}

func jsonMarshalIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "\t")
}

func main() {

	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := jsonMarshalIndent(app.config)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(bytes))
	}

	app.work()
}
