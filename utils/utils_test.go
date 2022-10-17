package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilTestSuite struct {
	suite.Suite
}

func (suite *UtilTestSuite) TestMakeRequest() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`OK`))
	}))
	defer testServer.Close()
	_, err := MakeRequest("GET", testServer.URL)
	suite.NoError(err)
}

func (suite *UtilTestSuite) TestMakeRequest_EmptyUrl() {
	_, err := MakeRequest("GET", "")
	suite.Equal(ErrInvalidURL, err)
}

func (suite *UtilTestSuite) TestMakeRequest_Error() {
	_, err := MakeRequest("GET", "xyz.com")
	suite.Equal("Get \"xyz.com\": unsupported protocol scheme \"\"", err.Error())
}
func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, new(UtilTestSuite))
}
