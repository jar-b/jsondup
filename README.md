# jsondup
[![build](https://github.com/jar-b/jsondup/actions/workflows/build.yml/badge.svg)](https://github.com/jar-b/jsondup/actions/workflows/build.yml)

Detect duplicate keys in a JSON object.


## Installation

```console
go install github.com/jar-b/jsondup/cmd/jsondup@latest
```

## Usage

```console
$ jsondup -h
Detect duplicate keys in a JSON object.

Usage: jsondup [filename]
```

## Examples

### Simple

Given the following content in `dup.json`:

```json
{
  "a": "foo",
  "a": "bar"
}
```

```console
$ jsondup dup.json
2023/09/12 15:53:29 duplicate key "a"
```

### Nested

A more complex example might be an AWS IAM policy document with duplicated `Condition` keys:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:PutObject",
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "s3:prefix": ["one/", "two/"]
        },
        "StringEquals": {
          "s3:versionid": "abc123"
        }
      }
    }
  ]
}
```

```console
$ jsondup iam.json
2023/09/12 15:57:15 duplicate key "Statement.0.Condition.StringEquals"
```
