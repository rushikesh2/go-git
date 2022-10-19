package gitreport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/go-git/config"
	"github.com/go-git/models"
	"github.com/go-git/utils"
)

type gitWorker struct {
	config config.Configuration
}

func GitWorker(config config.Configuration) (gw *gitWorker) {
	return &gitWorker{
		config: config,
	}
}

// FetchGitPRSummary will fetch data from git repository for the specified days in config.
// for more info: https://docs.github.com/en/rest/pulls/pulls#about-the-pulls-api
func (gw *gitWorker) FetchGitPRSummary() (summary map[string]int, err error) {
	next := true
	page := 1
	summary = make(map[string]int)
	// calculate previous days from the current date
	previousDays := time.Now().AddDate(0, 0, -1*gw.config.PreviousDays)
	for next {
		// formating the URL
		gitPRUrl := fmt.Sprintf(gw.config.BaseURL, gw.config.Repository, strconv.Itoa(page))
		response, err := utils.MakeRequest(http.MethodGet, gitPRUrl)
		if err != nil {
			return summary, err
		}
		prDetails := []models.GitPR{}
		err = json.Unmarshal(response.ResponseBody, &prDetails)
		if err != nil {
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
	fmt.Println("------------------------------------------------------------------------")
	fmt.Print("\n")
	fmt.Println("To: " + gw.config.ReceiverEmail)
	fmt.Println("From: " + gw.config.SenderEmail)
	fmt.Println("Subject: Summary Report of last weeks github PRs for repository: " + gw.config.Repository)
	fmt.Print("\n")
	fmt.Println("Hello Scrum-Master,")
	fmt.Print("\n")
	fmt.Println("Please find the summary report of github repository below.")
	fmt.Print("\n")

	// used tabwriter library for more details: https://pkg.go.dev/text/tabwriter
	w := tabwriter.NewWriter(os.Stdout, 5, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "%s\t%s\t\n", "---------------", "-----------")
	fmt.Fprintf(w, "%s\t%s\t\n", "State of the PR", "Count")
	fmt.Fprintf(w, "%s\t%s\t\n", "---------------", "-----------")
	for key, val := range summaryData {
		fmt.Fprintf(w, "%s\t%d\t\n", key, val)
	}
	w.Flush()
	fmt.Println("------------------------------------------------------------------------")
	return nil
}
