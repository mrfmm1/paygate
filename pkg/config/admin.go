// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package config

type Admin struct {
	BindAddress           string `yaml:"bind_address" json:"bind_address"`
	DisableConfigEndpoint bool   `yaml:"disable_config_endpoint" json:"disable_config_endpoint"`
}
