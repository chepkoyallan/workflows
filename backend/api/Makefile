IMAGE_TAG=workflow-api-generator
VERSION=0.0.1

# Generate clients starting by building api-generator image locally.
# Note, this should only be used for local development purposes. Once any change is made to the Dockerfile,
# we should push the new image remotely to ensure everyone is using the same tools.
.PHONY: generate-from-scratch
generate-from-scratch: .image-built hack/generator.sh *.proto
	docker run  --interactive --rm \
	    --user $$(id -u):$$(id -g) \
		--mount type=bind,source="$$(pwd)/../..",target=/go/src/github.com/chepkoyallan/workflows \
		$(IMAGE_TAG) /go/src/github.com/chepkoyallan/workflows/backend/api/hack/generator.sh

.PHONY: generate-from-scratch-sdk
generate-from-scratch-sdk: .image-built hack/sdk.sh sdk_protos/*.proto
	docker run  --interactive --rm \
	    --user $$(id -u):$$(id -g) \
		--mount type=bind,source="$$(pwd)/../..",target=/go/src/github.com/chepkoyallan/workflows \
		$(IMAGE_TAG) /go/src/github.com/chepkoyallan/workflows/backend/api/hack/sdk.sh

documentation:
	docker run --rm \
	  -v $$(pwd)../:/backend\
      -v $$(pwd)/docs:/out \
      -v $$(pwd)/sdk_protos/:/protos \
      pseudomuto/protoc-gen-doc \

# Build a local api-generator image.
.PHONY: image
image: .image-built

# Push api-generator image remotely.
.PHONY: push
push: image
	docker tag $(IMAGE_TAG) $(REMOTE_IMAGE)
	docker push $(REMOTE_IMAGE)

# .image-built is a local token file to help Make determine the latest successful build.
.image-built: Dockerfile
	docker build ../.. -t $(IMAGE_TAG) -f Dockerfile
	touch .image-built


