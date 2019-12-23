# PGL - Pretty Git Log
### Monitors your Pull Requests and commits on git (current version only supports Github)

#### Prerequisites
1. Install Golang v-a.12.5 <enter>
    -- verify by `go version` command in console.
2. Set GOPATH <enter>
    -- Run `export GOPATH=/path/desired/here` command 
3. create `github.com/vhndaree` by using `mkdir github.com && github.com/Vhndaree` command 
4. clone pgl inside **Vhndaree** directory 
    -- HTTPS<enter>
      `git clone https://github.com/Vhndaree/pretty-git-log.git` <enter>
    -- SSH<enter>
      `git@github.com:Vhndaree/pretty-git-log.git`
5. open repo and run `cp .example.env` <enter>
  -- change `GITHUB_TOKEN` with your token <enter>
    generate token [here](https://github.com/settings/tokens/new?scopes=&description=pgl)
6. Now Your are ready to go, run `go run main.go` <enter> <enter>


&copy [Ram Bhandari](https://github.com/Vhndaree )