# Cloud Function to Big Query

[Reference](https://qiita.com/shuichiro/items/0383f2cf07b9114f9cdf)

## Pre Setup

### Big Query Table

1. Table definition: `names.json`
2. `bq` command to create the table

```bash
cd /PATH/TO/{CF-BQ}
bq mk go_dataset
bq mk --table go_dataset.names names.json
```

## Hint

`cloud.google.com/go/bigquery`

## Main Function

* function-to-bq.go


## Prepare for `go.mod`

* go.mod

## Deploy

```bash
gcloud functions deploy ToBigQuery --runtime go111 --trigger-http
```

## Test

```bash
curl https://us-central1-kouzoh-p-jieqiong-yu.cloudfunctions.net/ToBigQuery?name=MerJQ
```

Check the item 

```bash
bq query "SELECT * FROM go_dataset.names limit 10"
```