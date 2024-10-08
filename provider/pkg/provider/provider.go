package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	fwCallback "github.com/cloudy-sky-software/pulumi-provider-framework/callback"
	fwRest "github.com/cloudy-sky-software/pulumi-provider-framework/rest"
)

type scalewayInstancesProvider struct {
	name    string
	version string

	apiKey string
}

var (
	handler  *fwRest.Provider
	callback fwCallback.ProviderCallback
)

func makeProvider(host *provider.HostClient, name, version string, pulumiSchemaBytes, openapiDocBytes, metadataBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	p := &scalewayInstancesProvider{
		name:    name,
		version: version,
	}

	callback = p
	rp, err := fwRest.MakeProvider(host, name, version, pulumiSchemaBytes, openapiDocBytes, metadataBytes, callback)

	handler = rp.(*fwRest.Provider)

	return rp, err
}

func extractOutput(outputs interface{}) (map[string]interface{}, error) {
	// All create/read/update endpoints return the object under a top-level object.
	// For example, if creating a `Server` resource, then the response
	// is under a `server` property.
	for _, v := range outputs.(map[string]interface{}) {
		return v.(map[string]interface{}), nil
	}

	return nil, errors.New("outputs did not have a top-level object")
}

func (p *scalewayInstancesProvider) GetAuthorizationHeader() string {
	return p.apiKey
}

func (p *scalewayInstancesProvider) OnPreInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, _ *http.Request) error {
	return nil
}

func (p *scalewayInstancesProvider) OnPostInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, outputs interface{}) (map[string]interface{}, error) {
	return outputs.(map[string]interface{}), nil
}

// OnConfigure is called by the provider framework when Pulumi calls Configure on
// the resource provider server.
func (p *scalewayInstancesProvider) OnConfigure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	apiKey, ok := req.GetVariables()["scaleway-instances:config:apiKey"]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["apiKey"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				apiKey = v
			}
		}

		// Return an error if the API key is still empty.
		if apiKey == "" {
			return nil, errors.New("api key is required")
		}
	}

	logging.V(3).Info("Configuring Scaleway Instances API key")
	p.apiKey = apiKey

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// OnDiff checks what impacts a hypothetical update will have on the resource's properties.
func (p *scalewayInstancesProvider) OnDiff(_ context.Context, _ *pulumirpc.DiffRequest, _ string, _ *resource.ObjectDiff, _ *openapi3.MediaType) (*pulumirpc.DiffResponse, error) {
	return nil, nil
}

func (p *scalewayInstancesProvider) OnPreCreate(_ context.Context, _ *pulumirpc.CreateRequest, _ *http.Request) error {
	return nil
}

// OnPostCreate allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *scalewayInstancesProvider) OnPostCreate(_ context.Context, _ *pulumirpc.CreateRequest, outputs interface{}) (map[string]interface{}, error) {
	return extractOutput(outputs)
}

func (p *scalewayInstancesProvider) OnPreRead(_ context.Context, _ *pulumirpc.ReadRequest, _ *http.Request) error {
	return nil
}

func (p *scalewayInstancesProvider) OnPostRead(_ context.Context, _ *pulumirpc.ReadRequest, outputs interface{}) (map[string]interface{}, error) {
	return extractOutput(outputs)
}

func (p *scalewayInstancesProvider) OnPreUpdate(_ context.Context, _ *pulumirpc.UpdateRequest, _ *http.Request) error {
	return nil
}

func (p *scalewayInstancesProvider) OnPostUpdate(_ context.Context, _ *pulumirpc.UpdateRequest, _ http.Request, outputs interface{}) (map[string]interface{}, error) {
	return extractOutput(outputs)
}

func (p *scalewayInstancesProvider) OnPreDelete(_ context.Context, _ *pulumirpc.DeleteRequest, _ *http.Request) error {
	return nil
}

func (p *scalewayInstancesProvider) OnPostDelete(_ context.Context, _ *pulumirpc.DeleteRequest) error {
	return nil
}
