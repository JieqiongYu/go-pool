# Connect to Firebase from Cluod Function 

## SetUp

### Change account for Google Cloud SDK

```bash
gcloud auth login
```

### Commands about GCP project

```bash
# Check current project
gcloud config get-value core/project
# Set project
gcloud config set project <PROJECT_ID>
```

### Install Firebase Admin SDK

```bash
go get -u firebase.google.com/go
```

### Deploy

This function would be triggered by Firestore Event.

```bash
# cd /path/to/Sunrise-Sunset
gcloud functions deploy ScoreReview --runtime go111 \
--trigger-event providers/cloud.firestore/eventTypes/document.create \
--trigger-resource "projects/connect-cf-241923/databases/(default)/documents/movie_reviews/{pushId}"
## projects/<PROJECT_ID>/databases/(default)/documents/movie_reviews/(pushId)
```

### Test

Add one record in cloud firestore of `movie_reviews`

### Delete

```bash
gcloud functions delete ScoreReview
```

## Reference 
* https://medium.com/google-cloud/firebase-developing-serverless-functions-in-go-963cb011265d
* https://firebase.google.com/docs/firestore/quickstart
* https://cloud.google.com/functions/docs/quickstart