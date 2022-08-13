package woo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/ssm"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
)

func GetOrder(ctx context.Context, orderID string) (*Order, error) {
	log := logger.FromContext(ctx)
	log.Infof("get order [%s]", orderID)

	host, err := ssm.GetStageParameter(ctx, "vs2", "/woo/host", false)
	if err != nil {
		log.Errorf("failed to read woo host: %s", err)
		return nil, err
	}

	username, err := ssm.GetStageParameter(ctx, "vs2", "/woo/username", false)
	if err != nil {
		log.Errorf("failed to read woo username: %s", err)
		return nil, err
	}

	password, err := ssm.GetStageParameter(ctx, "vs2", "/woo/password", true)
	if err != nil {
		log.Errorf("failed to read woo password: %s", err)
		return nil, err
	}

	client := &http.Client{Timeout: 30 * time.Second}
	url := fmt.Sprintf("%s/wp-json/wc/v3/orders/%s", host, orderID)
	log.Infof("get order from: %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("failed to create HTTP request to get woo order: %s", err)
		return nil, err
	}
	req.SetBasicAuth(username, password)
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to send HTTP request to get woo order: %s", err)
		return nil, err
	}
	log.Infof("response status: %s", res.Status)
	if res.StatusCode >= http.StatusBadRequest {
		log.Errorf("response code [%s] from woo", res.Status)
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("failed to read response body to get woo order: %s", err)
		return nil, err
	}
	order := new(Order)
	err = json.Unmarshal(bodyText, order)
	if err != nil {
		log.Errorf("failed to parse response body to get woo order: %s", err)
		log.Errorf("%s", string(bodyText))
		return nil, err
	}
	log.Infof("order: %+v", *order)
	return order, nil
}
