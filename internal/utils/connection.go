package utils

/**
* This file contains all the utility functions related to network connections.
* @Author: Yinfeng Lu
* @Date: 2024-08-20
 */

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// connectionTester is a http client that is used to test the connection.
// It does not keep the connection alive and has a timeout of 2 seconds.
var connectionTester = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// CanConnect tests the connection to the given URL.
// If the given URL is invalid
// (i.e. connection request times out after 2 seconds, or the status code is not 200),
// it returns false. Otherwise, it returns true.
// Note that redirecting URLs are not accepted.
func CanConnect(url string) bool {
	resp, err := connectionTester.Get(url)
	if err != nil {
		logx.Errorw("connection test failed", logx.LogField{Key: "error", Value: err.Error()})
		return false
	}
	if resp.StatusCode != http.StatusOK {
		logx.Errorw("connection test failed", logx.LogField{Key: "status_code", Value: resp.StatusCode})
		return false
	}
	return true
}
