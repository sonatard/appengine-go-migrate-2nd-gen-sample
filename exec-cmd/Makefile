project=$(PROJECT)
service=sample
version=exec-cmd
yaml_dir=./main
url=https://$(version)-dot-$(service)-dot-$(project).appspot.com

deploy:
	time gcloud app deploy --quiet --version $(version) $(yaml_dir)/app.yaml

dep:
	go get

logs:
	gcloud app logs tail --service $(service) --version $(version)

browse:
	gcloud app browse --service $(service) --version $(version)

# APIs
index:
	curl $(url)
