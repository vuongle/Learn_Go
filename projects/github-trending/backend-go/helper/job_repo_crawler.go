package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"github-trending-api/constants"
	app_errors "github-trending-api/errors"
	"github-trending-api/models"
	"github-trending-api/repositories"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/labstack/gommon/log"
)

func CrawlTrendingRepos(githubRepo repositories.GithubRepo) {

	// Create a collector(crawler)
	c := colly.NewCollector()

	// Create a new slice storing crawled trendind repos
	repos := make([]models.Github, 0, 30)

	c.OnHTML(`article[class=Box-row]`, func(e *colly.HTMLElement) {
		var githubRepo models.Github

		// repo name is "a" tag inside h2
		githubRepo.Name = e.ChildText("h2 > a")
		n := strings.Replace(e.ChildText("h2 > a"), "\n", "", -1)
		githubRepo.Name = strings.Replace(n, " ", "", -1)

		// crawl description
		githubRepo.Description = e.ChildText("p.col-9")

		// color
		bgColor := e.ChildAttr(".repo-language-color", "style")
		re := regexp.MustCompile("#[a-zA-Z0-9_]+")
		match := re.FindStringSubmatch(bgColor)
		if len(match) > 0 {
			githubRepo.Color = match[0]
		}

		// url
		githubRepo.Url = e.ChildAttr("h1.h2 > a", "href")

		// lang
		githubRepo.Lang = e.ChildText("span[itemprop=programmingLanguage]")

		// stars and forks
		e.ForEach(".mt-2 a", func(index int, el *colly.HTMLElement) {
			if index == 0 {
				githubRepo.Stars = strings.TrimSpace(el.Text)
			} else if index == 1 {
				githubRepo.Fork = strings.TrimSpace(el.Text)
			}
		})
		e.ForEach(".mt-2 .float-sm-right", func(index int, el *colly.HTMLElement) {
			githubRepo.StarsToday = strings.TrimSpace(el.Text)
		})

		// build by
		var buildBy []string
		e.ForEach(".mt-2 span a img", func(index int, el *colly.HTMLElement) {
			avatarContributor := el.Attr("src")
			buildBy = append(buildBy, avatarContributor)
		})
		githubRepo.BuildBy = strings.Join(buildBy, ",")

		// add to slice
		repos = append(repos, githubRepo)
	})

	// Register a callback when the "OnHTML()" done
	c.OnScraped(func(r *colly.Response) {
		queue := NewJobQueue(runtime.NumCPU())
		queue.Start()
		defer queue.Stop()

		// cache the data in redis
		encodedRepos, err := json.Marshal(repos)
		if err == nil {
			SetCache(context.Background(), constants.REDIS_TRENDING_REPO_KEY, string(encodedRepos))
		}

		// save or update the data in db
		for _, repo := range repos {
			queue.Submit(&RepoProcess{
				repo:       repo,
				githubRepo: githubRepo,
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://github.com/trending")
}

type RepoProcess struct {
	repo       models.Github
	githubRepo repositories.GithubRepo
}

func (rp *RepoProcess) Process() {
	// select repo by name
	cacheRepo, err := rp.githubRepo.SelectRepoByName(context.Background(), rp.repo.Name)
	if err == app_errors.ErrRepoNotFound {
		// khong tim thay repo - insert repo to database
		fmt.Println("Add: ", rp.repo.Name)
		_, err = rp.githubRepo.SaveRepo(context.Background(), rp.repo)
		if err != nil {
			log.Error(err)
		}
		return
	}

	// Neu tim thấy thì update
	if rp.repo.Stars != cacheRepo.Stars ||
		rp.repo.StarsToday != cacheRepo.StarsToday ||
		rp.repo.Fork != cacheRepo.Fork {
		fmt.Println("Updated: ", rp.repo.Name)
		rp.repo.UpdatedAt = time.Now()
		_, err = rp.githubRepo.UpdateRepo(context.Background(), rp.repo)
		if err != nil {
			log.Error(err)
		}
	}
}
