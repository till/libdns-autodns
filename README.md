
\<autodns\> for [`libdns`](https://github.com/libdns/libdns)
=======================

[![Go Reference](https://pkg.go.dev/badge/test.svg)](https://pkg.go.dev/github.com/till/libdns-autodns)

This package implements the [libdns interfaces](https://github.com/libdns/libdns) for \<autodns\>, allowing you to manage DNS records.

Example:

```
package main

import (
	"context"
	"os"
	"log"

	autodns "github.com/till/libdns-autodns"
)

func main() {
	provider := autodns.NewWithDefaults(os.Getenv("AUTODNS_USERNAME"), os.Getenv("AUTODNS_PASSWORD"))

	records, err := provider.GetRecords(context.TODO(), "zone.example.org")
	if err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	fmt.Printf("%#v", records)
}
```

As an alternative, configure the provider struct with the following:

| Field      | Description                | Required |
|------------|----------------------------|----------|
| Username   | username                   | yes      |
| Password   | password                   | yes      |
| Endpoint   | https://api.autodns.com/v1 | no       |
| Context    | 4                          | no       |
| httpClient | `&http.Client{}`           | no       |
