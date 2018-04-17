// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// taken from http://golang.org/src/pkg/net/ipraw_test.go
//
// notes: in MAC system, use root user to run.
package ping

import (	
	"net/http"
	"net/url"
	"net"
	"time"
)

func Ping(address string, timeoutSecond int) (alive bool, err error, timedelay time.Duration) {
	t := time.Now()
	httcode,err := Pinger(address, timeoutSecond)
	return httcode == 200, err, time.Now().Sub(t)
}

func Pinger(domain string, timeoutSecond int) (int, error) {
	proxyUrl, err := url.Parse(domain)
    if err != nil {
       return 0, err
    }	
	client := &http.Client{
	  Transport: &http.Transport{
		 Dial: (&net.Dialer{Timeout: time.Duration(timeoutSecond) * time.Second}).Dial,
		 TLSHandshakeTimeout: time.Duration(timeoutSecond) * time.Second,
		 Proxy: http.ProxyURL(proxyUrl),
	  },
	}	
    url := "http://www.google.com"
    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
       return 0, err
    }
    resp, err := client.Do(req)
    if err != nil {
       return 0, err
    }
    resp.Body.Close()
    return resp.StatusCode, nil
}
