# s3-edit

Edit directly a file on Amazon S3.

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

For only macOS

```
$ brew install tsub/s3-edit/s3-edit
```

### Get binary from GitHub releases

For Linux

```
$ VERSION=0.0.6
$ curl -fSL -o s3-edit.tar.gz "https://github.com/tsub/s3-edit/releases/download/v${VERSION}/s3-edit_v${VERSION}_linux_amd64.tar.gz"
$ tar -zxvf s3-edit.tar.gz -C /usr/local/bin
$ rm s3-edit.tar.gz
```

## Requirements

* AWS credentials
* Upload files to S3 in advance

For examples, use aws-cli

```
$ aws configure --profile myaccount
$ export AWS_PROFILE=myaccount
$ export AWS_REGION=ap-northeast-1
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

## Future

* [x] Provide the Homebrew formula
* [ ] Add server-side encryption option

## Development

### Requirements

* [dep](https://github.com/golang/dep)

### How to setup

```
$ go get -u github.com/tsub/s3-edit
$ cd $GOPATH/src/github.com/tsub/s3-edit
$ dep ensure
```
