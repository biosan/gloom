# Gloom

> **G**o B**loom** Filter

Gloom is a simple project written to learn how to build something useful in Go.

This repo contains a library and a simple REST API to use that library with HTTP requests.

Gloom is nothing more than a toy project, so do not use it for anything serious without at least reading the code.

## Why

I was always fascinated by Bloom Filters, it's a datastucture that combines many things I love: hashing functions, bit manipulations, simplicity, efficiency, etc.

I also thought that building a "Bloom Filter as a Service" would be something cool and not so common.

So it was the perfect project to learn Go.

## How to use it

### Install

```bash
go get github.com/biosan/gloom
```

### Run as API

```bash
go run github.com/biosan/gloom/api
```

> NOTE:
> Remember to configure the server port. Default is `8888`

#### Example

- Create a new BloomFilter
```JSON
POST - /create
request = {
    "number_of_hash_functions" : "3",
    "size_in_bits" : "512"
}
response : 201 
```

- Add an item
```JSON
POST - /add
request  : {
    "data": "WHATEVER STRING YOU WANT"
}
response : 201
```

- Query
```JSON
POST - /query
request  : {
    "data": "WHATEVER STRING YOU WANT"
}
response : {
    "probably_in": "true"|"false"
}
```

### Use it in your project

```Go
import "github.com/biosan/gloom"
```

> NOTE:
> Data added/queryed to the filter must be `[]byte`


#### Example

```Go
package main

import (
    "fmt"
    "github.com/biosan/gloom"
)

func main() {
	// Create a new bloom filter with
	// 3 hash functions and a size of 1024-bit
	bf := gloom.New(3, 1024)
	// Convert string to byte slice
	hi := []byte("hi")
	hello := []byte("hello")
	// Insert "hi" inside the filter
	bf.Add(hi)
	// Query the filter
	query1 := bf.Query(hi)
	query2 := bf.Query(hello)

	fmt.Printf("%v\n", query1)    // Prints "true"
	fmt.Printf("%v\n", query2)    // Prints "false"
}
```