# appengine-go-migrate-2nd-gen-sample

Reference is [Migrating your App Engine app from Go 1.9 to Go 1.11](https://cloud.google.com/appengine/docs/standard/go111/go-differences).

### first
- Go 1.9
- App Engine first generation
- Use App Engine API
- dep


### second-with-appengine-api
- Go 1.11
- App Engine second generation
- Use App Engine API
- Use appengine.Main Method
- dep

### second-without-appengine-api
- Go 1.11
- App Engine second generation
- Not use [App Engine API](https://cloud.google.com/appengine/docs/standard/go/reference)
  - App Engine URL Fetch -> [net/http package](https://golang.org/pkg/net/http/)
  - App Engine Socket -> [net package](https://golang.org/pkg/net/)
  - App Engine Blobstore -> [Cloud Storage](https://cloud.google.com/storage/) [storage package](https://godoc.org/cloud.google.com/go/storage)
  - App Engine datastore -> [Cloud datastore](https://cloud.google.com/datastore/), [datastore package](https://godoc.org/cloud.google.com/go/datastore)
  - App Engine taskqueue -> [Cloud Tasks](https://cloud.google.com/tasks/), [cloudtasks package](https://godoc.org/google.golang.org/api/cloudtasks/v2beta3)
  - App Engine cron -> [Cloud Scheduler](https://cloud.google.com/scheduler/)
  - App Engine Modules -> [Environment variables](https://cloud.google.com/appengine/docs/standard/go111/runtime#environment_variables)
- Not Use appengine.Main Method, Use Normal http.ListenAndServe Method
- go mod


## TODO

- Use Cloud Scheduler
- socket sample
- URL Fetch Sample
- Log Sample with request structured logging
- Mail sample
- Image sample
- Blobstore sample
- Cloud SQL sample
- test for Spanner 


## Usage

```
export GO111MODULE=on
export PROJECT=your_project_id
cd [first, second-with-appengine-api or second-without-appengine-api]
make dep
make deploy
make index
```

