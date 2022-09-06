package coinswap

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/irisnet/irismod/modules/coinswap/client/testutil"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(testutil.IntegrationTestSuite))
}
