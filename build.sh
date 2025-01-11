#!/bin/bash

set -e

package_name='sshm'

platforms=(
"darwin/amd64" # macOS x86_64
"darwin/arm64" # macOS ARM64 (Apple Silicon)
"linux/amd64"  # Linux x86_64
"linux/arm"    # Linux ARM32
"linux/arm64"  # Linux ARM64
)

rm -rf './dist'

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='./dist/'$package_name'-'$GOOS'-'$GOARCH

	echo 'Building '$package_name'-'$GOOS'-'$GOARCH
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	echo 'Archiving'
	zip -m $output_name'.zip' $output_name
done
