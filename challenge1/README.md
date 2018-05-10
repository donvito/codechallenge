[![Go Report Card](https://goreportcard.com/badge/github.com/donvito/codechallenge)](https://goreportcard.com/report/github.com/donvito/codechallenge)

# Build Docker image

```docker build --tag 'challenge1:1.0' .```

# Run the application

```docker run -it --mount type=bind,source="$(pwd)",target=/tmp,readonly challenge1:1.0 /usr/bin/dumb-init ./app -f /tmp/<file where repos are listed>```

Example

```docker run -it --mount type=bind,source="$(pwd)",target=/tmp,readonly challenge1:1.0 /usr/bin/dumb-init ./app -f /tmp/repos.txt```