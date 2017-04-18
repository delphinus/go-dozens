# go-dozens

[![CircleCI](https://circleci.com/gh/delphinus/go-dozens/tree/master.svg?style=svg)](https://circleci.com/gh/delphinus/go-dozens/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/delphinus/go-dozens/badge.svg?branch=master)](https://coveralls.io/github/delphinus/go-dozens?branch=master)

[Dozens][] is a DNS service that has a simple Web interface and high
functionality.  It has published [API Reference][api] and can be managed from
any CLI tools.

[Dozens]: http://dozens.jp
[api]: http://help.dozens.jp/categories/apiリファレンス/

This repo **go-dozens** is an implementation for the whole API.  This has been
fully tested and has much reliability.

## Usage

1. Get token

    ```go
    import "github.com/delphinus/go-dozens"

    const (
      DozensKey = "hogehoge"  // API Key from Dozens
      DozensUser = "fugafuga" // Dozens ID
    )

    resp, err := dozens.GetAuthorize(DozensKey, DozensUser)
    if err != nil {
      panic(err)
    }

    token := resp.AuthToken
    ```

2. Call API methods

    ```go
    zone, err := dozens.ZoneList(token)
    if err != nil {
      return panic(err)
    }

    fmt.Printf("%+v", zone) // {Domain:[{ID:12345 Name:example.com}]}
    ```

    That's all!

## Available Methods

The list below is the same as [the official documents for API][api].

### zones

* `ZoneList(token)`
* `ZoneCreate(token, ZoneCreateBody{})`
* `ZoneUpdate(token, zoneID, mailAddress)`
* `ZoneDelete(token, zoneID)`

### records

* `RecordList(token, zoneName)`
* `RecordCreate(token, RecordCreateBody{})`
* `RecordUpdate(token, RecordUpdateBody{})`
* `RecordDelete(token, recordID)`

## Tool

* [godo][] - CLI tool for dozens with this library.

[godo]: https://github.com/delphinus/godo
