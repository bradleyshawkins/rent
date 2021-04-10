if [ -z $GOOS ]
then
	GOOS=linux
fi

echo $GOOS

CGO_ENABLED=0 \
GOOS=${GOOS} \
GOARCH=amd64 \
go build -v -i