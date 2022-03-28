package servicebus

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type ServiceBus struct {
	*azservicebus.Client
}

type ServiceBusOptions struct {
	*azservicebus.ClientOptions
}

func Connect(options *ServiceBusOptions) (*ServiceBus, error) {
	var client *ServiceBus
	var err error
	client.Client, err = azservicebus.NewClientFromConnectionString(
		os.Getenv("AZURE_SERVICEBUS_CONNECTION_URI"),
		options.ClientOptions,
	)
	if err != nil {
		return nil, err
	}

	return client, nil

}
