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
	"net/http"
	"os"
	"strings"
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// remoteAddress returns the IP address of the request.
// If the X-Forwarded-For or X-Real-Ip http headers are set, then
// they are used to obtain the remote address.
// The boolean is true, if the remote address is obtained using the
// request RemoteAddr() method.
func remoteAddress(r *http.Request) (string, bool) {
	hdr := r.Header

	// Try to obtain the ip from the X-Forwarded-For header
	ip := hdr.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(ip, ",")
		if len(parts) > 0 {
			ip = strings.TrimSpace(parts[0])

			if ip != "" {
				return ip, false
			}
		}
	}

	// Try to obtain the ip from the X-Real-Ip header
	ip = strings.TrimSpace(hdr.Get("X-Real-Ip"))
	if ip != "" {
		return ip, false
	}

	// Fallback to the request remote address
	return removePortFromRemoteAddr(r.RemoteAddr), true
}

// removePortFromRemoteAddr removes the port if present from the remote address.
func removePortFromRemoteAddr(remoteAddr string) string {
	pos := strings.LastIndex(remoteAddr, ":")
	if pos < 0 {
		return remoteAddr
	}

	return remoteAddr[:pos]
}
