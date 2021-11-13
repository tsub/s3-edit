# s3-edit

Edit directly a file on Amazon S3 in CLI.

[![GitHub release](https://img.shields.io/github/release/tsub/s3-edit.svg?style=flat-square)](https://github.com/tsub/s3-edit/releases)
[![CircleCI branch](https://img.shields.io/circleci/project/github/tsub/s3-edit/master.svg?style=flat-square)](https://circleci.com/gh/tsub/s3-edit/tree/master)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/tsub/s3-edit)
[![MIT license](https://img.shields.io/github/license/tsub/s3-edit.svg?style=flat-square)](https://github.com/tsub/s3-edit/blob/master/LICENSE)

## Installation

### Use go get

```
$ go get -u github.com/tsub/s3-edit
```

### Install with Homebrew

For macOS and Linux

```
$ brew install tsub/s3-edit/s3-edit
```

### Get binary from GitHub releases

Download latest binary from https://github.com/tsub/s3-edit/releases

### Using [Nix](https://nixos.org/)

```
$ nix-env -if default.nix
```

## Requirements

* AWS credentials
* Upload files to S3 in advance

For examples, use aws-cli

```
$ aws configure --profile myaccount
$ export AWS_PROFILE=myaccount
```

Other methods,

```
$ export AWS_ACCESS_KEY_ID=xxxx
$ export AWS_SECRET_ACCESS_KEY=xxxx
$ export AWS_REGION=ap-northeast-1
```

## Usage

Upload the file to S3 in advance.

```
$ echo "This is a test file." > myfile.txt
$ aws s3 cp test.txt s3://mybucket/myfile.txt
```

To directly edit a file on S3, use `edit` subcommand.

```
$ s3-edit edit s3://mybucket/myfile.txt
```

Then, open a file with the default editor specified by `$EDITOR`.

[![https://gyazo.com/96c9225da700f91e7b44c04f439fdd23](https://i.gyazo.com/96c9225da700f91e7b44c04f439fdd23.png)](https://gyazo.com/96c9225da700f91e7b44c04f439fdd23)

When you close the editor after edit, a file is automatically re-uploaded to S3.

```
$ aws s3 cp s3://mybucket/myfile.txt -
This is a test file.
Edited with s3-edit.
```

## Development

### Requirements

* Golang >= 1.11

### How to setup

```
$ git clone git@github.com:tsub/s3-edit.git
$ export GO111MODULE=on
```

### Using [Nix](https://nixos.org/)

From the project root directory, this will enter you into a shell environment with s3-edit executable available, built from local source files.

```
$ nix-shell
```
