package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/donvito/gopkg/githubstats"
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

	m := map[string]string{"RepoName": "Repo Name", "CloneURL": "Clone URL", "LastCommitDate": "Last Commit Date", "Author": "Author"}
	repoInfos = append(repoInfos, m)

	for _, repo := range repos {
		m := githubstats.RetrieveRepoMetadata(repo)
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
