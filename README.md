# Shortening GitHub URL Command



## Description

Shortening GitHub URL Command.

Refer to “[Git.io: GitHub URL Shortener](https://github.com/blog/985-git-io-github-url-shortener)”.

## Usage

```shell
$ gitio-shorten --help
usage: gitio-shorten [--version] [--help] <command> [<args>]

Available commands are:
    decode     Decode Git.io URL
    encode     Shorten GitHub URL
    version    Print gitio-shorten version and quit

$ gitio-shorten encode -c t https://github.com/technoweenie
http://git.io/t

$ gitio-shorten decode http://git.io/t
https://github.com/technoweenie
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/spiegel-im-spiegel/gitio-shorten
```

## Contribution

1. Fork ([https://github.com/spiegel-im-spiegel/gitio-shorten/fork](https://github.com/spiegel-im-spiegel/gitio-shorten/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## License

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.ja)
