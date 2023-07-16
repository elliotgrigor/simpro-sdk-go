# simPRO Software API Go SDK

**Disclaimer**: This project is not affiliated with simPRO.

An unofficial SDK for the [simPRO Software API](https://developer.simprogroup.com/apidoc/).

## Install

```sh
go get github.com/elliotgrigor/simpro-sdk-go
```

## Usage example

```go
package main

import (
    "log"

    "github.com/elliotgrigor/simpro-sdk-go/simpro"
)

func main() {
    sdk, err := simpro.NewSimPROSDK(
        "my-organisation.simprocloud.com",          // simPRO instance's FQDN
        "0a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t", // API access token
        0,                                          // Optional: company ID
    )
    if err != nil {
        log.Fatal(err)
    }

    // Retrieve list of company IDs
    cl, err := sdk.GetCompanies()
    if err != nil {
        log.Fatal(err)
    }

    // `cl` (dereferenced, pretty-printed)
    // [
    //     {
    //         ID: 0
    //         Name: My Organisation
    //         Phone: 0123 456 7890
    //         Email: contact@my.org
    //         Address: {
    //             Line1: 123 Fake Street
    //             Line2: Glasgow
    //         }
    //         Country: United Kingdom
    //         Timezone: Europe/London
    //         Currency: GBP
    //     }
    //     {
    //         ID: 1
    //         Name: My Other Organisation
    //         Phone: 0123 456 7891
    //         Email: contact@my-other.org
    //         Address: {
    //             Line1: 321 Made-up Avenue
    //             Line2: Manchester
    //         }
    //         Country: United Kingdom
    //         Timezone: Europe/London
    //         Currency: GBP
    //     }
    // ]

    // Company ID can be set after initialisation
    sdk.SetCompany(1)

    // Uses company ID set in NewSimPROSDK or SetCompany
    ci, err := sdk.GetCompanyInfo()
    if err != nil {
        log.Fatal(err)
    }

    // `ci` (dereferenced, pretty-printed)
    // {
    //     ID: 1
    //     Name: My Other Organisation
    //     Phone: 0123 456 7891
    //     Email: contact@my-other.org
    //     Address: {
    //         Line1: 321 Made-up Avenue
    //         Line2: Manchester
    //     }
    //     Country: United Kingdom
    //     Timezone: Europe/London
    //     Currency: GBP
    // }
}
```
