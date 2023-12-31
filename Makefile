# How to run:
# make dev
dev:
	go run cmd/smart-schedule/main.go -env development

# How to run:
# make prod
prod:
	go run cmd/smart-schedule/main.go -env production

# How to run:
# make image
image-prod:
	docker build --platform linux/x86_64 -f deployments/Dockerfile -t tdg-prod/smart-schedule .

# make deploy-prod
deploy-prod:
	make image-prod
	docker tag $(docker images -q tdg-prod/smart-schedule) registry-intl.ap-southeast-1.aliyuncs.com/lotuss/cronjob:latest
	docker push registry-intl.ap-southeast-1.aliyuncs.com/lotuss/cronjob:latest