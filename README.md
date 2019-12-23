# PGL - Pretty Git Log
### Monitors your Pull Requests and commits on git (current version only supports Github)

#### Prerequisites
1. Install Golang v1.12.5 <br/>
    -- verify by `go version` command in console.
2. Set GOPATH <br/>
    -- Run `export GOPATH=/path/desired/here` command 
    
### For Contributors
3. Fork and clone repo -> Send PR

### For Users
3. run `go get github.com/Vhndaree/preety-git-log` <br/><br/>
4. open repo and run `cp .example.env .env` <br/>
  -- change `GITHUB_TOKEN` with your token and remove # of the line<br/>
    generate token [here](https://github.com/settings/tokens/new?scopes=&description=pgl) 
5. Now You are ready to go, run `go run main.go` <br/> <br/>


&copy; [Ram Bhandari](https://github.com/Vhndaree )
