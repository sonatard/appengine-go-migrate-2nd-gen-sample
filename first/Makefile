project=projectid
service=sample
version=first
yaml_dir=./yaml
url=https://$(version)-dot-$(service)-dot-$(project).appspot.com

deploy:
	# gcloud app deploy --quiet --version $(version) $(yaml_dir)/app.yaml
	# ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
    # 2018/11/23 22:06:31 go-app-builder: Failed parsing input: package "github.com/sonatard/appengine-go-sample/first/vendor/google.golang.org/appengine/log" cannot import internal package "google.golang.org/appengine/internal/log"
    # issue :

deploy_legacy:
	#appcfg.py -A $(project) -V $(version) --oauth2_access_token=$(gcloud auth print-access-token) update $(yaml_dir)
	appcfg.py -A $(project) -V first3 --oauth2_access_token=$(gcloud auth print-access-token) update $(yaml_dir)

logs:
	gcloud app logs tail --service $(service) --version $(version)

browse:
	gcloud app browse --service $(service) --version $(version)

# APIs
log:
	curl $(url)/log

appenginelog:
	curl $(url)/appenginelog

panic:
	curl $(url)/panic

taskqueue:
	curl $(url)/taskqueue

memcache:
	curl $(url)/memcache

memcachedelete:
	curl $(url)/memcachedelete

datastore:
	curl $(url)/datastore

datastoredelete:
	curl $(url)/datastoredelete

search:
	curl $(url)/search

searchdelete:
	curl $(url)/searchdelete

auth:
	curl $(url)/auth

