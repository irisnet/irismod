package farm

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/irisnet/irismod/modules/farm/client/testutil"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(testutil.IntegrationTestSuite))
}
