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
	"fmt"

	"github.com/BurntSushi/toml"
)

const (
	maxPort = 65535
)

var (
	Config = config{
		ListenPort: 80,

		DestinationHost: "127.0.0.1",
		DestinationPort: 8080,

		Description: "Secured Area",
		PasswdFile:  "watchman.passwd",
	}
)

type config struct {
	listenAddress string
	ListenHost    string
	ListenPort    int

	destinationAddress string
	DestinationHost    string
	DestinationPort    int

	Description string
	PasswdFile  string
}

// Init initializes the config values.
func (c *config) Init() error {
	// Validate the ports
	if c.ListenPort <= 0 || c.ListenPort > maxPort {
		return fmt.Errorf("invalid listen port: %v", c.ListenPort)
	}
	if c.DestinationPort <= 0 || c.DestinationPort > maxPort {
		return fmt.Errorf("invalid destination port: %v", c.DestinationPort)
	}

	// Create the listen address and destination address.
	c.listenAddress = fmt.Sprintf("%s:%v", c.ListenHost, c.ListenPort)
	c.destinationAddress = fmt.Sprintf("%s:%v", c.DestinationHost, c.DestinationPort)

	// Check if the passwd file exists.
	e, err := exists(c.PasswdFile)
	if err != nil {
		return err
	} else if !e {
		return fmt.Errorf("watchman passwd file is missing!")
	}

	return nil
}

// LoadConfig loads the config.
func LoadConfig(configPath string) error {
	// Load and decode the file.
	_, err := toml.DecodeFile(configPath, &Config)
	if err != nil {
		return fmt.Errorf("failed to load config file '%s': %v", configPath, err)
	}

	// Initialize the config.
	return Config.Init()
}
