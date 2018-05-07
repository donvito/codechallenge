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
	apiRoot := "https://api.github.com/repos/" + repo
	response, err := http.Get(apiRoot)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var _repoMetadata repoMetadata
	err = json.Unmarshal(b, &_repoMetadata)

	//bodyString := string(bodyInBytes)

	//fmt.Printf("Repo name = %s , Clone URL = %s, Commits URL = %s \n", _repoMetadata.Name, _repoMetadata.CloneURL, _repoMetadata.CommitsURL)

	latstCommitDate, author := retrieveRepoCommits(_repoMetadata.CommitsURL)
	fmt.Printf("%s, %s, %s, %s \n", _repoMetadata.Name, _repoMetadata.CloneURL, latstCommitDate, author)
}

type repoCommits struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

func retrieveRepoCommits(commitsURL string) (latestCommitDate, author string) {
	//https://api.github.com/repos/donvito/codechallenge/commits
	latestCommitDate = "06/05/2018" //TODO: implement this
	author = "Melvin Vivas"         //TODO: implement this

	return

}
