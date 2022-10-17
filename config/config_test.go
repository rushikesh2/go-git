package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestReadConfigValues_Success() {
	expected := Configuration{
		SenderEmail:   "senr@gmail.com",
		ReceiverEmail: "mmaster@gmail.com",
		PreviousDays:  2,
		GitURL:        "https://apis.gitl.com/repos/aws",
	}
	conf, err := ReadConfigValues("test.json")
	suite.NoError(err)
	suite.Equal(expected, conf)
}

func (suite *ConfigTestSuite) TestReadConfigValues_Error() {
	_, err := ReadConfigValues("")
	suite.Error(err)
	suite.Equal("open : The system cannot find the file specified.", err.Error())
}

func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
