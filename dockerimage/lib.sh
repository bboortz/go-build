
set -o errexit
set -o nounset
set -o pipefail


if [ -z "${ARCH}" ]; then
    error "Variable ARCH is not set"
fi


error() {
    error "$1" >&2
    exit 1
}


#export CGO_ENABLED=0 
export GOOS=linux
export GOARCH="${ARCH}" 

