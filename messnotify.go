package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type JParamsStruct struct {
	Dockermode int    `json:"Dockermode"`
	SlackURL   string `json:"SlackURL"`
}

func main() {
	var JParam JParamsStruct
	const JsonFileName = "settings.json"

	// Открываем файл с настройками
	jSettingsFile, err := os.Open(JsonFileName)
	// Проверяем на ошибки
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	defer jSettingsFile.Close()

	FData, err := ioutil.ReadAll(jSettingsFile)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	json.Unmarshal(FData, &JParam)
	if JParam.Dockermode == 1 {
		JParam.SlackURL = os.Getenv("SLACK_URL")
	}

	http.HandleFunc("/slack", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if r.Method == http.MethodOptions {
				return
			}

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}

			qstring2, _ := url.QueryUnescape(string(body))

			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}

			timeout := time.Duration(10 * time.Second)
			client := &http.Client{Transport: tr, Timeout: timeout}

			req, _ := http.NewRequest("POST", JParam.SlackURL, strings.NewReader(qstring2))
			resp, dreq := client.Do(req)
			if dreq != nil {
				w.WriteHeader(http.StatusGatewayTimeout)
				return
			}
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, string(body))
			return
		}
	})

	http.ListenAndServe("0.0.0.0:8788", nil)
}
