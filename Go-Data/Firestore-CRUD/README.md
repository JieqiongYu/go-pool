# Firestore CRUD

## Setup

### Install Go Dependencies

```bash
go get cloud.google.com/go/firestore
go get firebase.google.com/go
```

### Create a private key
Go to https://console.firebase.google.com/u/0/project/< PROJECT ID >/settings/serviceaccounts/adminsdk to generate the 
private key and save it somewhere.


## Reference 

* https://medium.com/eureka-engineering/go-firebase-example-2879d32de678
* [Firebaseの新しいDatabase「Cloud Firestore」をGoで使ってみる](https://log.shinofara.xyz/2018/02/firebase%E3%81%AE%E6%96%B0%E3%81%97%E3%81%84databasecloud-firestore%E3%82%92go%E3%81%A7%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B/)
* https://godoc.org/cloud.google.com/go/firestore
* https://firebase.google.com/docs/firestore/quickstart