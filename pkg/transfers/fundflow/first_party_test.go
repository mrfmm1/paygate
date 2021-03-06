// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fundflow

import (
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	customers "github.com/moov-io/customers/client"
	"github.com/moov-io/paygate/pkg/client"
	"github.com/moov-io/paygate/pkg/config"
)

func TestOriginate__DebitCheck(t *testing.T) {
	cfg := config.Empty()
	cfg.ODFI.RoutingNumber = "987654320"

	fp := NewFirstPerson(cfg.Logger, cfg.ODFI)

	companyID := "MOOV"
	xfer := &client.Transfer{}
	src := Source{
		Customer: customers.Customer{
			Status: customers.RECEIVE_ONLY,
		},
	}
	dest := Destination{
		Account: customers.Account{
			RoutingNumber: "987654320",
		},
	}
	if _, err := fp.Originate(companyID, xfer, src, dest); err == nil {
		t.Error("expected error")
	} else {
		if !strings.Contains(err.Error(), "does not support debit") {
			t.Error(err)
		}
	}
}

func TestOriginate__RoutingNumberErr(t *testing.T) {
	cfg := config.Empty() // leave off RoutingNumber for first test
	fp := NewFirstPerson(log.NewNopLogger(), cfg.ODFI)

	src := Source{
		Account: customers.Account{
			RoutingNumber: "987654320",
		},
	}
	dst := Destination{
		Account: customers.Account{
			RoutingNumber: "987654320",
		},
	}

	if file, err := fp.Originate("companyID", nil, src, dst); file != nil || err == nil {
		t.Fatal("expected nil File and error")
	} else {
		if !strings.Contains(err.Error(), "rejecting transfer between two accounts within") {
			t.Error(err)
		}
	}

	cfg.ODFI.RoutingNumber = "987654320"
	src.Account.RoutingNumber = "123"
	dst.Account.RoutingNumber = "987"

	if file, err := fp.Originate("companyID", nil, src, dst); file != nil || err == nil {
		t.Fatal("expected nil File and error")
	} else {
		if !strings.Contains(err.Error(), "rejecting third-party transfer between FI's we don't represent") {
			t.Error(err)
		}
	}
}
