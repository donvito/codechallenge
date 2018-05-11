package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	errMsg := "Please indicate filename: ./app -f repos.txt or ./app --filaname repos.txt"
	if len(os.Args) != 3 {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	repos := []string{}
	if os.Args[1] == "--filename" || os.Args[1] == "-f" {
		reposFilename := os.Args[2]
		repos = readReposFromFile(reposFilename)
	} else {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	var repoInfos = make([]map[string]string, 0)

	m := make(map[string]string)
	m["RepoName"] = "RepoName"
	m["CloneURL"] = "CloneURL"
	m["LastCommitDate"] = "LastCommitDate"
	m["Author"] = "Author"
	repoInfos = append(repoInfos, m)

	for _, repo := range repos {
		m := retrieveRepoMetadata(repo)
		repoInfos = append(repoInfos, m)
	}

	for _, m := range repoInfos {
		fmt.Printf("%s, %s, %s, %s \n", m["RepoName"], m["CloneURL"], m["LastCommitDate"], m["Author"])
	}

}

func readReposFromFile(filepath string) (repos []string) {
	file, _ := os.Open(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		repos = append(repos, scanner.Text()) //add repo name to slice
	}

	return

}

type repoMetadata struct {
	Name       string `json:"name"`
	CloneURL   string `json:"clone_url"`
	CommitsURL string `json:"commits_url"`
}

func retrieveRepoMetadata(repo string) (m map[string]string) {
	//https://api.github.com/users/donvito/repos
	apiRoot := fmt.Sprintf("%s%s", "https://api.github.com/repos/", repo)
	response, err := http.Get(apiRoot)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var _repoMetadata repoMetadata
	err = json.Unmarshal(body, &_repoMetadata)

	parsedCommitsURL := parseCommitsURL(_repoMetadata.CommitsURL)
	lastCommitDate, author := retrieveRepoCommits(parsedCommitsURL)

	m = map[string]string{"RepoName": _repoMetadata.Name, "CloneURL": _repoMetadata.CloneURL, "LastCommitDate": lastCommitDate, "Author": author}

	return

}

func parseCommitsURL(commitsURL string) (parsedCommitsURL string) {

	i := strings.Index(commitsURL, "{/sha}")
	parsedCommitsURL = commitsURL[:i]

	return

}

type repoCommits struct {
	Sha    string `json:"sha"`
	Commit commit `json:"commit"`
}

type commit struct {
	Author author `json:"author"`
}

type author struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func retrieveRepoCommits(commitsURL string) (latestCommitDate, author string) {

	response, err := http.Get(commitsURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	// Unmarshal string into structs.
	var repos []repoCommits
	json.Unmarshal(body, &repos)

	//get first element of slice only
	if len(repos) > 0 {
		latestCommitDate = repos[0].Commit.Author.Date
		author = repos[0].Commit.Author.Name
	}

	return

}
