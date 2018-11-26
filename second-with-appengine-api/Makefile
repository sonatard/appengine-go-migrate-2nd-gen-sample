project=$(PROJECT)
service=sample
version=second-with-appengine-api
yaml_dir=./main
url=https://$(version)-dot-$(service)-dot-$(project).appspot.com

deploy:
	time gcloud app deploy --quiet --version $(version) $(yaml_dir)/app.yaml

deploy_legacy:
	# appcfg.py -A $(project) -V $(version) --oauth2_access_token=$(gcloud auth print-access-token) update $(yaml_dir)
	# appcfg.py: error: Error parsing ./src/app.yaml: Unable to assign value 'go111' to attribute 'runtime':
    # Value 'go111' for runtime does not match expression '^(?:contrib-dart|dart|go|php|php55|php72|python|python27|python-compat|java|java7|java8|vm|custom|nodejs|ruby)$'
    #  in "./src/app.yaml", line 1, column 10.
    # make: *** [deploy_legacy] Error 2

logs:
	gcloud app logs tail --service $(service) --version $(version)

browse:
	gcloud app browse --service $(service) --version $(version)

# APIs
index:
	curl $(url)

log:
	curl $(url)/log

cmd:
	curl $(url)/cmd

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

dep:
	dep ensure
