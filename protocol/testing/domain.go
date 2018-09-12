// Code generated by cdpgen. DO NOT EDIT.

// Package testing implements the Testing domain. Testing domain is a dumping
// ground for the capabilities requires for browser or app testing that do not
// fit other domains.
package testing

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Testing domain. Testing domain is a
// dumping ground for the capabilities requires for browser or app testing that
// do not fit other domains.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Testing domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GenerateTestReport invokes the Testing method. Generates a report for
// testing.
func (d *domainClient) GenerateTestReport(ctx context.Context, args *GenerateTestReportArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Testing.generateTestReport", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Testing.generateTestReport", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Testing", Op: "GenerateTestReport", Err: err}
	}
	return
}