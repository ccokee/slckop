make generate
make manifests

make bundle IMG="dockier.io/ccokee/slckop:latest"
make bundle-build bundle-push BUNDLE_IMG="docker.io/ccokee/slckop-bundle:latest"