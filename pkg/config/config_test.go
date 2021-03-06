// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package config

import (
	"path/filepath"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	cfg, err := FromFile(filepath.Join("testdata", "valid.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Logger == nil {
		t.Fatal("nil Logger")
	}
	if cfg.Logging.Format != "json" {
		t.Errorf("cfg.Logging.Format=%s", cfg.Logging.Format)
	}

	if cfg.ODFI.RoutingNumber != "987654320" {
		t.Errorf("ODFIConfig=%#v", cfg.ODFI)
	}
}

func TestInvalidConfig(t *testing.T) {
	cfg, err := FromFile(filepath.Join("testdata", "invalid.yaml"))
	if err == nil {
		t.Error("expected error")
	}

	if err := cfg.Validate(); err == nil {
		t.Error("expected error")
	}
}

func TestReadConfig(t *testing.T) {
	conf := []byte(`logging:
  format: plain
customers:
  endpoint: "http://localhost:8087"
  accounts:
    decryptor:
      symmetric:
        keyURI: 'MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI='
odfi:
  routing_number: "987654320"
  gateway:
    origin: "CUSTID"
  inbound_path: "/files/inbound/"
  outbound_path: "/files/outbound/"
  return_path: "/files/return/"
  allowed_ips: "10.1.0.1,10.2.0.0/16"
  cutoffs:
    timezone: "America/New_York"
    windows: ["17:00"]
  ftp:
    hostname: sftp.moov.io
    username: moov
    password: secret
  storage:
    keep_remote_files: false
    local:
      directory: "/opt/moov/storage/"
validation:
  micro_deposits:
    source:
      customerID: "user"
      accountID: "acct"
pipeline:
  merging:
    directory: "./storage/"
  stream:
    inmem:
      url: "mem://paygate"
`)
	cfg, err := Read(conf)
	if err != nil {
		t.Fatal(err)
	}
	if err := cfg.Validate(); err != nil {
		t.Error(err)
	}

	if cfg == nil {
		t.Fatal("nil Config")
	}

	if cfg.Customers.Endpoint != "http://localhost:8087" {
		t.Errorf("customers endpoint: %q", cfg.Customers.Endpoint)
	}
	if cfg.Customers.Accounts.Decryptor.Symmetric.KeyURI != "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=" {
		t.Errorf("accounts decryptor KeyURI=%q", cfg.Customers.Accounts.Decryptor.Symmetric.KeyURI)
	}

	if cfg.Validation.MicroDeposits.Source.CustomerID != "user" {
		t.Errorf("unexpected verification: %#v", cfg.Validation)
	}

	if cfg.Pipeline.Stream.InMem.URL != "mem://paygate" {
		t.Errorf("missing pipeline stream config: %#v", cfg.Pipeline.Stream)
	}
}

func TestConfig__FTP(t *testing.T) {
	cfg := Empty().ODFI.FTP
	if cfg != nil {
		t.Fatalf("unexpected %#v", cfg)
	}

	if v := cfg.CAFile(); v != "" {
		t.Errorf("ca_file=%q", v)
	}
	if v := cfg.Timeout(); v != 10*time.Second {
		t.Errorf("dial_timeout=%v", v)
	}
	if v := cfg.DisableEPSV(); v {
		t.Errorf("disabled_epsv=%v", v)
	}
}

func TestConfig__SFTP(t *testing.T) {
	cfg := Empty().ODFI.SFTP
	if cfg != nil {
		t.Fatalf("unexpected %#v", cfg)
	}

	if v := cfg.Timeout(); v != 10*time.Second {
		t.Errorf("dial_timeout=%v", v)
	}
	if v := cfg.MaxConnections(); v != 8 {
		t.Errorf("max_connections_per_file=%d", v)
	}
	if v := cfg.PacketSize(); v != 20480 {
		t.Errorf("max_packet_size=%d", v)
	}
}
