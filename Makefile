all: build-binary build-image push

TAG ?= v0.0.5
REGISTRY ?= rokibulhasan114
APP_NAME ?= book-api
RELEASE_NAME ?= book-server


build-binary:
	go build -o ${APP_NAME} .


build-image: build-binary
	docker build -t ${REGISTRY}/${APP_NAME}:${TAG} .


push: build-image
	docker push ${REGISTRY}/${APP_NAME}:${TAG}


helm-install:
	helm install ${RELEASE_NAME} chart

helm-uninstall:
	helm uninstall ${RELEASE_NAME}

clean:
	rm -f ${APP_NAME}