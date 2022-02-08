# Forked GoDotEnv (https://goreportcard.com/report/github.com/joho/godotenv)

A Go (golang) environment loader (which loads env vars from a .env file)

## Installation

```shell
go get github.com/bendt-indonesia/env
```

## Usage

Add your application configuration to your `.env` file in the root of your project:

```shell
DB_HOST=localhost
DB_USERNAME=root
DB_PASSWORD=
JWT_SECRET=YOURSECRETKEYGOESHERE
```

Then in your Go app you can do something like

```go
package main

import (
    "github.com/bendt-indonesia/env"
    "log"
    "os"
)

func main() {
    //Default will read .env files on root folder
    err := env.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    s3Bucket := os.Getenv("S3_BUCKET")
    secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
```

## Who?

The original library [env](https://github.com/bendt-indonesia/env) was written by [John Barton](https://johnbarton.co/)based off the tests/fixtures in the original library.