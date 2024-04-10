package secretmanager

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type SecretsManager interface {
	// I added NewSecretManager myself to the interface - so I'm not sure it's right
	NewSecretManager(ctx context.Context, projectId string) (*SecretManagerClient, error)
	GetSecret(secretId string) ([]byte, error)
}

type SecretManagerClient struct {
	client    *secretmanager.Client
	projectID string
}

func NewSecretManager(ctx context.Context, projectID string) (*SecretManagerClient, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &SecretManagerClient{client: client, projectID: projectID}, nil
}

func (sm *SecretManagerClient) GetSecret(secretID string) ([]byte, error) {
	// Build name from secretId
	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", sm.projectID, secretID)

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{Name: name}

	// Call the API
	secret, err := sm.client.AccessSecretVersion(context.Background(), accessRequest)
	if err != nil {
		return nil, err
	}
	return secret.Payload.Data, nil
}
