# gseo

[![build-test](https://github.com/xiexianbin/gseo/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/gseo/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/gseo?status.svg)](https://pkg.go.dev/github.com/xiexianbin/gseo)
[![Go Report Card](https://goreportcard.com/badge/github.com/xiexianbin/go-actions-demo)](https://goreportcard.com/report/github.com/xiexianbin/go-actions-demo)

a golang client to optimize [hugo](https://www.xiexianbin.cn/tags/hugo/) seo by Google Search Console. read [gseo spec](./docs/specification.md) for more information.

## Install

- source

```
go install github.com/xiexianbin/gseo
```

- bin

```
curl -Lfs -o gseo https://github.com/xiexianbin/gseo/releases/latest/download/gseo-{linux|darwin|windows}
chmod +x gseo
./gseo
```

## Usage

### show help

- root help

```
$ gseo -h
 golang client to optimize [hugo](https://www.xiexianbin.cn/tags/hugo/) seo by Google Search Console.

Usage:
  gseo [flags]
  gseo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        init gseo configure
  keyword     show site keywords
  render      render hugo post markdown files
  sites       site list
  version     Print the version number of gseo.

Flags:
  -h, --help      help for gseo
  -v, --verbose   debug info.

Use "gseo [command] --help" for more information about a command.
```

- keywords help


```
$ gseo help keyword
download hugo post keywords from Google Search Console API, and cache it in `./.gseo/` dir

Usage:
  gseo keyword [flags]

Flags:
  -h, --help          help for keyword
  -l, --last int      last days (default 90)
  -s, --site string   site url
```

- render help

```
$ gseo help render
render hugo post markdown files.
default args is:
  gseo render --content PATH_OF_HUGO_CONTENT --position 10 --ctr 0 --impressions 100 --clicks 0.3 --max 8 --dryrun

Usage:
  gseo render [flags]

Flags:
  -k, --clicks float        >=clicks to render seo.
      --content string      hugo content path
  -c, --ctr float           ctr = clicks / impressions to render seo, and 0 <= ctr <= 1. (default 0.3)
  -r, --dryrun              dry run mode.
  -h, --help                help for render
  -i, --impressions float   >=impressions to render seo. (default 100)
  -m, --max int             max seo items, -1 is un-limit. (default 8)
  -p, --position float      >=position to render seo. (default 10)
  -s, --skip-err            skip error. (default true)
```

### how to get Google Search Console API token

- step-1: config google auth secret in file `~/.gseo/client_secret.json` (can i build it as constant in go pkg?)

```
$ cat ~/.gseo/client_secret.json
{"installed":{"client_id":"1017408311257-hq3j99vk9ludpoff862mnp52v36nv4gc.apps.googleusercontent.com","project_id":"adept-button-344010","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]}}
```

- step-2: init google token

```
$ gseo init
Please enter Google API client_secret.json path (default is /Users/xiexianbin/.gseo/client_secret.json):
init config success!
Go to the following link in your browser then type the authorization code:
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=1017408311257-hq3j99vk9ludpoff862mnp52v36nv4gc.apps.googleusercontent.com&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fwebmasters.readonly&state=state-token
---> <Visit the URL above, and paste the (authorization code/授权代码) over here> <---
Saving credential file to: /Users/xiexianbin/.gseo/token.json
init Google API OAuth2.0 token success!
```

### Demo

- list sites

```
$ gseo sites
PermissionLevel  SiteUrl
siteOwner        sc-domain:xiexianbin.cn
siteOwner        https://docs.xiexianbin.cn/
siteOwner        https://www.xiexianbin.cn/
```

- show special site keywords, while write keywords to cache file `~/.gseo/cache-<date>.json`

```
$ gseo keyword -s "https://www.xiexianbin.cn/"
==> get 1000 lines Results
==> get 1000 lines Results
Write to file /Users/xiexianbin/.gseo/cache-2022-06-03.json success, bytes 276427
```

- render demo, read keywords from cache file `~/.gseo/cache-<date>.json`

```
$ export HUGO_CONTENT="/Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content"

# default args
$ gseo render --content ${HUGO_CONTENT} --position 10 --ctr 0 --impressions 10 --clicks 0.1 --max 8 --dryrun true

# my args 1 **
$ gseo render --content ${HUGO_CONTENT} --position 10 --ctr 0 --impressions 100 --clicks 0.3 --max 8

# my args 2 *****
$ gseo render --content ${HUGO_CONTENT} --position 0 --ctr 0 --impressions 1 --clicks 0 --max 8 --dryrun true
```

now, you can find seo keywords has auto update/render to `${HUGO_CONTENT}`.

## Action

- keyword can not contain `:`, occur golang yaml parse err

## arch

[Hugo 博客 SEO 优化 - 自动调用 Google Search Console API](https://www.xiexianbin.cn/open-sources/google-search-console-render-hugo-post/index.html)

## dev

```
# run some TestCase
go test -v -run TestClient google/client_test.go

# build
$ make all
```

## ref

- https://developers.google.com/webmaster-tools/about
- https://developers.google.com/webmaster-tools/v1/api_reference_index
- [Quickstart: Run a Search Console App in Python](https://developers.google.com/webmaster-tools/v1/quickstart/quickstart-python)
  - [Search Console APP Dashboard](https://console.cloud.google.com/apis/api/cloudsearch.googleapis.com/overview)
- [Search Console Testing Tools API (Experimental)](https://developers.google.com/webmaster-tools/search-console-api)
