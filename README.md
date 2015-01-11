# ts4tosqs

Send messages from [ts4](https://github.com/tsileo/ts4) to an AWS SQS queue.

## Quick start

	$ go get github.com/tsileo/ts4tosqs
	$ export AWS_SECRET_ACCESS_KEY=yourkey
	$ export AWS_ACCESS_KEY_ID=yourkey
	$ go run ts4tosqs.go -debug -start=2015-01-10T11 -end=2015-01-10T11:30 -queue=yourkey

## License

MIT
