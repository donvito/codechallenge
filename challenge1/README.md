#Build Docker image
```docker build .```

#Run the application
```docker run -it <container id> /usr/bin/dumb-init ./app -f repos.txt```