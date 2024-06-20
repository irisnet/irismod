protoc_install_gopulsar() {
  go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
}

echo "Formatting protobuf files"
find ./ -name "*.proto" -exec clang-format -i {} \;

set -e

home=$PWD

echo "Generating proto code"
proto_dirs=$(find ./ -name 'buf.yaml' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  echo "Generating proto code for $dir"

  cd $dir
  # check if buf.gen.pulsar.yaml exists in the proto directory
  if [ -f "buf.gen.pulsar.yaml" ]; then
    protoc_install_gopulsar
    buf generate --template buf.gen.pulsar.yaml
    # move generated files to the right places
    if [ -d "../irismod" -a "$dir" != "./proto" ]; then
      cp -r ../irismod $home/api
      rm -rf ../irismod
    fi
  fi

  # check if buf.gen.gogo.yaml exists in the proto directory
  if [ -f "buf.gen.gogo.yaml" ]; then
      for file in $(find . -maxdepth 5 -name '*.proto'); do
        # this regex checks if a proto file has its go_package set to irismod.io/api/...
        # gogo proto files SHOULD ONLY be generated if this is false
        # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
        if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*irismod.io/api' "$file" | grep -q ':0$'; then
          buf generate --template buf.gen.gogo.yaml $file
        fi
      done

    # move generated files to the right places
    if [ -d "../irismod.io" ]; then
      cp -r ../irismod.io/* ../../
      rm -rf ../irismod.io
    fi
  fi

  cd $home
done