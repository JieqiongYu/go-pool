# Hello World Cloud Function 

### Commands about GCP project

```bash
# Check current project
gcloud config get-value core/project
# Set project
gcloud config set project <PROJECT_ID>
```

### Deploy

```bash
gcloud functions deploy HelloWorld --runtime go111 --trigger-http
```

### Test

```bash
curl https://us-central1-connect-cf.cloudfunctions.net/HelloWorld
```

### Clean up

```bash
gcloud functions delete HelloWorld
```