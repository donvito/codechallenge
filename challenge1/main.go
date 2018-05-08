package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

		//fmt.Print(reposFilename, "\n")
		repos = readReposFromFile(reposFilename)
		//fmt.Println(repos)
	} else {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	for _, repo := range repos {
		//fmt.Printf("repo = %s \n", repo)
		retrieveRepoMetadata(repo)
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

func retrieveRepoMetadata(repo string) {
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

	_repoMetadata.CommitsURL = fmt.Sprintf("%s%s%s", "https://api.github.com/repos/donvito/", _repoMetadata.Name, "/commits") // TODO: need to take this from commit URL in API response
	latstCommitDate, author := retrieveRepoCommits(_repoMetadata.CommitsURL)
	fmt.Printf("%s, %s, %s, %s \n", _repoMetadata.Name, _repoMetadata.CloneURL, latstCommitDate, author)
}

type RepoCommits struct {
	Sha    string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func retrieveRepoCommits(commitsURL string) (latestCommitDate, author string) {

	//println(commitsURL)

	response, err := http.Get(commitsURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	// Unmarshal string into structs.
	var repos []RepoCommits
	json.Unmarshal(body, &repos)

	// Loop over structs and display them.
	for l := range repos {
		latestCommitDate = repos[l].Commit.Author.Date
		author = repos[l].Commit.Author.Name
		break // just get latest so break after first iteration, need to improve this
	}

	return

}
