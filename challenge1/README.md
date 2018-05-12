[![Go Report Card](https://goreportcard.com/badge/github.com/donvito/codechallenge)](https://goreportcard.com/report/github.com/donvito/codechallenge)

This is an application which queries Github stats using the api.github.com. Input is a text file containing repo names. Output prints out the name, clone URL, date of latest commit and name of latest author of each repo.

# Build Docker image

```docker build --tag 'challenge1:1.0' .```

# Updated repos.txt with the ff. info
```
​$orgname/$repo
```

Sample
```
donvito/learngo
geocine/golem
```

# Run the application

```docker run -it --mount type=bind,source="$(pwd)",target=/tmp,readonly challenge1:1.0 /usr/bin/dumb-init ./app -f /tmp/<file where repos are listed>```

Example

```docker run -it --mount type=bind,source="$(pwd)",target=/tmp,readonly challenge1:1.0 /usr/bin/dumb-init ./app -f /tmp/repos.txt```

# Output prints out the name, clone URL, date of latest commit and name of latest author for each repo
```
$ docker run -it --mount type=bind,source="$(pwd)",target=/tmp,readonly challenge1:1.0 /usr/bin/dumb-init ./app -f /tmp/repos.txt
Repo Name, Clone URL, Last Commit Date, Author
learngo, https://github.com/donvito/learngo.git, 2018-05-01T01:51:39Z, Melvin Vivas
golem, https://github.com/geocine/golem.git, 2018-03-30T03:52:12Z, Aivan Monceller
```