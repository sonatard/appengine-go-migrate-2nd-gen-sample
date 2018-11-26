# Google App Engine second generation with appengine package sample code


## Support Function

**Tested** means that it can be called now.
**Deprecated** means that it is deprecated by Google.

|Function Name|Tested|Deprecated|
|-------------|-------|
|App Engine User API|OK|in the future|
|App Engine log|OK|in the future|
|App Engine Storage|OK|in the future|
|App Engine taskqueue|OK|in the future|
|App Engine cron|OK|in the future|
|App Engine Memcache|OK|in the future|
|App Engine Datastore|OK|in the future|
|App Engine Search API|OK|in the future|
|App Engine Image API|OK|in the future|
|vendor|OK|No|
|dep|OK|No|
|go mod|OK|OK|
|appcfg.py|OK|in the future|
|gcloud command|OK|No|


- https://cloud.google.com/appengine/docs/standard/go111/go-differences#changes_to_the_app_engine_go_111_runtime

> App Engine no longer modifies the Go toolchain to include the appengine package.
> We strongly recommend using the Google Cloud client library or third party libraries instead of the App Engine-specific APIs.
> For more information, see [Migrating from the App Engine Go SDK](https://cloud.google.com/appengine/docs/standard/go111/go-differences#migrating-appengine-sdk).

