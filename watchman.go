/*
 *  Watchman - Simple HTTP Reverse Proxy with authentication
 *  Copyright DesertBit
 *  Author: Roland Singer
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"flag"
	"net/http"
	"net/http/httputil"

	log "github.com/Sirupsen/logrus"
	auth "github.com/abbot/go-http-auth"
)

const (
	configName = "watchman.conf"
)

func main() {
	// Get the config path from the command line arguments.
	// The default config path is just the config name.
	// This will load the config from the current working directory.
	var configPath string
	flag.StringVar(&configPath, "config", configName, "set config file path.")

	// Parse the flags.
	flag.Parse()

	// Load the config.
	err := LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create the authenticator.
	authenticator := auth.NewBasicAuthenticator(Config.Description, auth.HtpasswdFileProvider(Config.PasswdFile))

	// Set the HTTP routes.
	http.HandleFunc("/", authenticator.Wrap(handleReverseProxyFunc))

	// Start the HTTP server.
	log.Fatal(http.ListenAndServe(Config.listenAddress, nil))
}

// handleReverseProxyFunc proxies the traffic to the destination host.
func handleReverseProxyFunc(w http.ResponseWriter, authReq *auth.AuthenticatedRequest) {
	// Get the http Request.
	r := &authReq.Request

	// Get the remote address.
	address, _ := remoteAddress(r)

	// Log
	log.Infof("request from client '%s@%s': %s", authReq.Username, address, r.URL)

	// Create the director.
	director := func(req *http.Request) {
		req = r
		req.URL.Scheme = "http"
		req.URL.Host = Config.destinationAddress
	}

	// Proxy the request.
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
