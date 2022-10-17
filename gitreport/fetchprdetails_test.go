package gitreport

import (
	"testing"

	"github.com/go-git/config"
	"github.com/stretchr/testify/suite"
)

type FetchPRDetailsSuite struct {
	suite.Suite
}

func (suite *FetchPRDetailsSuite) TestFetchGitPRSummary() {
	gw := &gitWorker{
		config: config.Configuration{
			SenderEmail:   "sender@gmail.com",
			ReceiverEmail: "receiver@gmail.com",
			PreviousDays:  3,
			GitURL:        "https://api.github.com/repos/aws/aws-sdk-go",
		},
	}
	_, err := gw.FetchGitPRSummary()
	suite.NoError(err)
}

func (suite *FetchPRDetailsSuite) TestFetchGitPRSummary_Error() {
	gw := &gitWorker{
		config: config.Configuration{
			SenderEmail:   "sender@gmail.com",
			ReceiverEmail: "receiver@gmail.com",
			PreviousDays:  3,
			GitURL:        "https://api.github.com",
		},
	}
	_, err := gw.FetchGitPRSummary()
	suite.Error(err)
}

func (suite *FetchPRDetailsSuite) TestFetchGitPRSummary_ErrorMakeRequest() {
	gw := &gitWorker{
		config: config.Configuration{
			SenderEmail:   "sender@gmail.com",
			ReceiverEmail: "receiver@gmail.com",
			PreviousDays:  3,
			GitURL:        "",
		},
	}
	_, err := gw.FetchGitPRSummary()
	suite.Error(err)
}

func (suite *FetchPRDetailsSuite) TestSendMail() {
	summary := make(map[string]int)
	summary["open"] = 1
	summary["closed"] = 3
	gw := &gitWorker{
		config: config.Configuration{
			SenderEmail:   "sender@gmail.com",
			ReceiverEmail: "receiver@gmail.com",
			PreviousDays:  3,
			GitURL:        "https://api.github.com/repos/aws/aws_gogit",
		},
	}
	err := gw.SendMail(summary)
	suite.NoError(err)
}

func (suite *FetchPRDetailsSuite) TestGitWorker() {
	config := config.Configuration{
		SenderEmail:   "sender@gmail.com",
		ReceiverEmail: "receiver@gmail.com",
		PreviousDays:  3,
		GitURL:        "https://api.github.com/repos/aws/aws_gogit",
	}

	_ = GitWorker(config)
}
func TestFetchGitPRSummary(t *testing.T) {
	suite.Run(t, new(FetchPRDetailsSuite))
}
