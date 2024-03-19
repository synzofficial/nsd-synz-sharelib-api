# nsd-synz-sharelib-api

## API ID LIST
- BFF and ShareLib -> 00
- AUTH -> 01
- MASTER_DATA -> 02
- User -> 03
- Account -> 04
- Media -> 05

## How to install? (go get this private repo)
1. Set go env GOPRIVATE
```
$ go env GOPRIVATE | grep "github.com/synzofficial" || go env -w GOPRIVATE="$(go env GOPRIVATE),github.com/synzofficial"
```

2. Check go env GOPRIVATE
```
$ go env GOPRIVATE

// github.com/synzofficial
```

## setup for support downloading lib from github using ssh.

1. set .gitconfig
```
vi ~/.gitconfig
```

2. in ~/.gitconfig file **ADD**
```
[url "ssh://git@github.com/"]
        insteadOf = https://github.com/
```