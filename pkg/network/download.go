package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func DownloadFile(log *logrus.Entry, url string) ([]byte, error) {
	log.Infof("download file [%s]", url)
	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		log.Errorf("failed to download: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		log.Errorf("status code [%d]", res.StatusCode)
		return nil, fmt.Errorf("status code [%d]", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("failed to read response body: %s", err)
		return nil, err
	}
	return body, nil
}
