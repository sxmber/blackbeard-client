# blackbeard-client
Client program that will automate torrenting and move files to media servers

Clone the repo

```$ git clone https://github.com/sxmber/blackbeard-client```

Install golang and set global variables in ~/.profile (or equivalent)

https://go.dev/doc/install

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

```

cd and install using go

```$ cd blackbeard-client && go install```

# Usage

blackbeard only has 1 command, "send" and requires 2 flags, -l and -m

These are the torrent/magnet link and the media type respectively

```$ blackbeard send -l somelegaltorrent.torrent -m Movies```
