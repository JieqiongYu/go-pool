# Hello From PubSub

[Reference](https://qiita.com/shuichiro/items/db2e138dcada8ee97901)

## Cloud Function

### Deploy Function

```bash
gcloud functions deploy HelloFromPubSub --runtime go111 --trigger-topic HELLO_WORLD
```

### Publish Topic

```bash
gcloud pubsub topics publish HELLO_WORLD --message '{"name":"MerJQ", "hobby":"Netflix"}'
```

### Check logs

```bash
gcloud functions log read --limit 50
```