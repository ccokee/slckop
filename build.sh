make generate
make manifests
IMAGE_TAG_BASE="docker.io/ccokee/slckop" VERSION="latest" make docker-build docker-push