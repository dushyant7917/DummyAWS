// +build aws

package s3

type S3Client struct {
}

func NewS3Client() *S3Client {
	return &S3Client{}
}

func (c *S3Client) GetBucketName() string {
	return "foo"
}

func (c *S3Client) GetLatestFile(key string) string {
	return key + "/" + "foo"
}