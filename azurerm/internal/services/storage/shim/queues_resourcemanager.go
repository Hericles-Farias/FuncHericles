package shim

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
	"github.com/tombuildsstuff/giovanni/storage/2019-12-12/queue/queues"
)

type ResourceManagerStorageQueueWrapper struct {
	client         *storage.QueueClient
	servicesClient *storage.QueueServicesClient
}

func NewResourceManagerStorageQueueWrapper(client *storage.QueueClient, servicesClient *storage.QueueServicesClient) StorageQueuesWrapper {
	return ResourceManagerStorageQueueWrapper{
		client:         client,
		servicesClient: servicesClient,
	}
}

func (w ResourceManagerStorageQueueWrapper) Create(ctx context.Context, resourceGroup, accountName, queueName string, metaData map[string]string) error {
	rmInput := storage.Queue{
		QueueProperties: &storage.QueueProperties{
			Metadata: w.mapDataPlaneMetaData(metaData),
		},
	}
	_, err := w.client.Create(ctx, resourceGroup, accountName, queueName, rmInput)
	return err
}

func (w ResourceManagerStorageQueueWrapper) Delete(ctx context.Context, resourceGroup, accountName, queueName string) error {
	_, err := w.client.Delete(ctx, resourceGroup, accountName, queueName)
	return err
}

func (w ResourceManagerStorageQueueWrapper) Exists(ctx context.Context, resourceGroup, accountName, queueName string) (*bool, error) {
	existing, err := w.client.Get(ctx, resourceGroup, accountName, queueName)
	if err != nil {
		if utils.ResponseWasNotFound(existing.Response) {
			return utils.Bool(false), nil
		}
		return nil, err
	}

	return utils.Bool(existing.QueueProperties != nil), nil
}

func (w ResourceManagerStorageQueueWrapper) Get(ctx context.Context, resourceGroup, accountName, queueName string) (*StorageQueueProperties, error) {
	existing, err := w.client.Get(ctx, resourceGroup, accountName, queueName)
	if err != nil {
		if utils.ResponseWasNotFound(existing.Response) {
			return nil, nil
		}
		return nil, err
	}

	if existing.QueueProperties == nil {
		return nil, fmt.Errorf("`properties` is null in the API response")
	}

	return &StorageQueueProperties{
		MetaData: w.mapResourceManagerMetaData(existing.QueueProperties.Metadata),
	}, nil
}

func (w ResourceManagerStorageQueueWrapper) GetServiceProperties(ctx context.Context, resourceGroup, accountName string) (*queues.StorageServiceProperties, error) {
	//serviceProps, err := w.servicesClient.GetServiceProperties(ctx, resourceGroup, accountName)
	//if err != nil {
	//	if utils.ResponseWasNotFound(existing.Response) {
	//		return nil, nil
	//	}
	//	return nil, err
	//}
	//
	//return &queues.StorageServiceProperties{
	//	Logging:       nil,
	//	HourMetrics:   nil,
	//	MinuteMetrics: nil,
	//	Cors:          nil,
	//}

	// TODO: work out how to proceed here
	return nil, fmt.Errorf("TODO: API doesn't have all of the available information")
}

func (w ResourceManagerStorageQueueWrapper) UpdateMetaData(ctx context.Context, resourceGroup, accountName, queueName string, metaData map[string]string) error {
	rmInput := storage.Queue{
		QueueProperties: &storage.QueueProperties{
			Metadata: w.mapDataPlaneMetaData(metaData),
		},
	}
	_, err := w.client.Update(ctx, resourceGroup, accountName, queueName, rmInput)
	return err
}

func (w ResourceManagerStorageQueueWrapper) UpdateServiceProperties(ctx context.Context, resourceGroup, accountName string, properties queues.StorageServiceProperties) error {
	// TODO: work out how to proceed here, since the API doesn't contain all of the required fields...
	//w.servicesClient.SetServiceProperties(ctx, resourceGroup, accountName, properties)
	return fmt.Errorf("TODO: API doesn't support all of the available functionality")
}

func (w ResourceManagerStorageQueueWrapper) mapDataPlaneMetaData(input map[string]string) map[string]*string {
	output := make(map[string]*string, len(input))

	for k, v := range input {
		output[k] = &v
	}

	return output
}

func (w ResourceManagerStorageQueueWrapper) mapResourceManagerMetaData(input map[string]*string) map[string]string {
	output := make(map[string]string, len(input))

	for k, v := range input {
		if v == nil {
			continue
		}
		output[k] = *v
	}

	return output
}
