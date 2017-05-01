# aws-env [![Build Status][svg-travis]][travis]

A simple shell utility for exporting a given AWS credentials profile to environment variables. Useful for crossing
machine boundaries with SSH and Vagrant. If you have to deal with having multiple AWS accounts or profiles, then this
is useful.

## Usage

Shamelessly ripped from `aws-env -h`:

```
Usage:
  aws-env [command]

Available Commands:
  export      Export the aws credentials to be read in by environment
  list        List the aws profiles configured

Flags:
      --config string   config file (default is $HOME/.aws-env.yaml)
  -t, --toggle          Help message for toggle

Use "aws-env [command] --help" for more information about a command.
```

If you have a profile named `brangus`, you can extract environment variables like so:

```shell
$ aws-env export brangus
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export AWS_DEFAULT_REGION=...
```

As a shortcut, you can directly source the output of this command to export the variables into your shell session:

```shell
$ $(aws-env export brangus)
```

This will cause your shell to execute the output of `aws-env`, exporting these environment variables.

> **WARNING:** This is _potentially very dangerous_ if you don't trust the script you're executing, so use with care and
> establish trust. All commits and releases here are [PGP signed][keybase] with my key, so if you know me and trust me,
> you should be able to use this, but as always, _read the source code_ and check it before you blindly pipe code into
> your shell session.

## Installation

##### User Install

To install `aws-env` as an ordinary user:

```shell
go get github.com/hhercules/aws-env

or

git clone https://github.com/hhercules/aws-env.git
cd aws-env
go install
```

## Building

You can just use go build commands:

```shell
$ go build
```

## License

Read the file called `LICENSE`, but it's basically MIT. If you want or need a dual-license for some reason that I have
yet to understand, please ask and I can dual-license it as appropriate.

 [travis]: https://travis-ci.org/hhercules/aws-env
 [svg-travis]: https://travis-ci.org/hhercules/aws-env.svg?branch=master
 [releases]: https://github.com/hhercules/aws-env/releases
 [keybase]: https://keybase.io/cxhercules
