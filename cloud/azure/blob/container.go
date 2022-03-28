package blob

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Container provides a connection to an azure storage account container.
type Container struct {
	azblob.ContainerClient
}

/// ContainerOptions provides options to be passed to the client connection method.
type ContainerOptions struct {
	*azblob.ClientOptions
}

// Connect connects to the container at the given uri using the credentials held in AZURE_BLOB_ACCOUNT_NAME and AZURE_BLOB_ACCOUNT_KEY.
func Connect(uri string, options *ContainerOptions) (*Container, error) {
	cred, err := azblob.NewSharedKeyCredential(
		os.Getenv("AZURE_BLOB_ACCOUNT_NAME"),
		os.Getenv("AZURE_BLOB_ACCOUNT_KEY"),
	)
	if err != nil {
		return nil, err
	}

	var client *Container
	client.ContainerClient, err = azblob.NewContainerClientWithSharedKey(
		uri,
		cred,
		options.ClientOptions,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetContent retrieves content from the container with the given name.
func (container *Container) GetContent(ctx context.Context, options *azblob.DownloadBlobOptions, name string) (string, error) {
	blob := container.NewBlockBlobClient(name)
	resp, err := blob.Download(ctx, options)
	if err != nil {
		return "", err
	}

	buffer := new(strings.Builder)
	reader := resp.Body(&azblob.RetryReaderOptions{})
	io.Copy(buffer, reader)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil

}

// CreateContent uploads content to the container under the given name and verifies the uploaded content against the container content.
func (container *Container) CreateContent(ctx context.Context, options *azblob.UploadBlockBlobOptions, name, content string) error {
	blob := container.NewBlockBlobClient(name)
	_, err := blob.Upload(ctx, streaming.NopCloser(io.ReadSeeker(strings.NewReader(content))), options)
	if err != nil {
		return err
	}

	uploaded, err := container.GetContent(ctx, nil, name)
	if err != nil {
		return err
	}

	if uploaded != content {
		container.DeleteContent(ctx, nil, name)
		return fmt.Errorf("Uploaded content corrupt. Try again.")
	}

	return nil

}

// DeleteContent removes the content with the given name from the container.
func (container *Container) DeleteContent(ctx context.Context, options *azblob.DeleteBlobOptions, name string) error {
	blob := container.NewBlockBlobClient(name)
	_, err := blob.Delete(ctx, options)
	if err != nil {
		return err
	}
	return nil
}
