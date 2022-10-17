package gitreport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-git/config"
	"github.com/go-git/models"
	"github.com/go-git/utils"
)

type gitWorker struct {
	config config.Configuration
}

const (
	queryParam = "/pulls?state=all&&sort=updated&&direction=desc&&page="
)

func GitWorker(config config.Configuration) (gw *gitWorker) {
	return &gitWorker{
		config: config,
	}
}

// FetchGitPRSummary will fetch data from git repo for the specified days in config.
func (gw *gitWorker) FetchGitPRSummary() (summary map[string]int, err error) {
	next := true
	page := 1
	summary = make(map[string]int)
	// calculate previous day from the date of current
	previousDays := time.Now().AddDate(0, 0, -1*gw.config.PreviousDays)
	for next {
		GitURL := gw.config.GitURL + queryParam + strconv.Itoa(page)
		response, err := utils.MakeRequest(http.MethodGet, GitURL)
		if err != nil {
			log.Println(err.Error())
			return summary, err
		}
		prDetails := []models.GitPR{}
		err = json.Unmarshal(response.ResponseBody, &prDetails)
		if err != nil {
			log.Println(err.Error())
			return summary, err
		}
		for _, prDetail := range prDetails {
			if prDetail.UpdatedAt.After(previousDays) {
				summary[prDetail.State] += 1
				summary["total"] += 1
				if prDetail.MergedAt != nil {
					summary["merged"] += 1
				}
			} else {
				next = false
				break
			}
		}
		//go to next page
		page++
	}
	return summary, err
}

// SendMail() to send mail to scrum-master
// From the GitURL we can take repository name
// As defined in assignment we are sending mail to the scrum master
func (gw *gitWorker) SendMail(summaryData map[string]int) (err error) {

	repositoryName := strings.Split(gw.config.GitURL, "/")[4:]
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("To: " + gw.config.ReceiverEmail)
	fmt.Println("From: " + gw.config.SenderEmail)
	fmt.Println("Subject: Summary Report of last weeks github PRs for repo: ", strings.Join(repositoryName, "/"))
	fmt.Println(" The summary of pull request is as follows")
	fmt.Println("--------------------------------------")
	fmt.Println("|   State of PR    |       Count      |")
	fmt.Println("--------------------------------------")
	for key, val := range summaryData {
		fmt.Println("| " + key + "    |       " + strconv.Itoa(val) + "         |")
	}
	fmt.Println("------------------------------------------------------------------------")
	return nil
}
