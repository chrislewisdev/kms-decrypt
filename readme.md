# KMS Decrypt Tool

It's a sane way to decrypt values that you encrypted with KMS, via the command-line.

## Usage
1. Install Go
2. Run `go get github.com/chrislewisdev/kms-decrypt`
3. Ensure your environment has your AWS credentials set up as per the [AWS CLI documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
4. Run `kms-decrypt -region <region> -value <value>`. Region defaults to ap-southeast-2 if not provided.

## Updating
If you already have this tool but want to update to the latest version, run `go get -u github.com/chrislewisdev/kms-decrypt`.