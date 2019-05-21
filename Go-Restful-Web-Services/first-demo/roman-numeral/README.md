# Roman Numeral

This service takes the number range (1-10) from the client and returns its Roman string. 

## API Design Doc

|HTTP Verb|PATH|Action|Resource|
|-|-|-|-|
|GET|/roman_number/2|show|roman_number|

## Source Code Tree

```bash
├── README.md
├── romannumerals
│   └── data.go
└── romanserver
    └── server.go
```

## data.go
* It works as a data service. 

## Test 

* Valid request

```bash
curl -X GET "http://localhost:8000/roman_number/5"
```

* Resource out of range

```bash
curl -X GET "http://localhost:8000/roman_number/12"
```

* Invalid resource 

```bash
curl -X GET "http://localhost:8000/random_resource/3" 
```