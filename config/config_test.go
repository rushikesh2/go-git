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
		BaseURL:       "https://apis.gitl.com/repos/",
		Repository:    "abc/abc",
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
