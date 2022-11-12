// Copyright 2022, Cloudy Sky Software LLC.

package provider

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudy-sky-software/pulumi-provider-framework/openapi"
	"github.com/cloudy-sky-software/pulumi-provider-framework/state"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

const testCreateJSONPayload = `{
	"name":"scw-reverent-chatterjee",
	"dynamic_ip_required":false,
	"commercial_type":"PLAY2-PICO",
	"boot_type":"local",
	"enable_ipv6":true,
	"image":"1cfcd1d3-82aa-46d5-b9cd-1682edf62f67",
	"volumes": {
		"0": {
			"size":10000000000,
			"volume_type":"b_ssd"
		}
	},
	"project":"65978379-535c-45a1-87be-92e4e3926db5",
	"zone": "fr-par-1"
}
`

func readFileFromProviderResourceDir(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("..", "..", "cmd", "pulumi-resource-scaleway-instances", filename))
	if err != nil {
		t.Fatalf("Failed reading openapi.yml: %v", err)
	}

	return b
}

func makeTestProvider(ctx context.Context, t *testing.T) pulumirpc.ResourceProviderServer {
	t.Helper()

	openapiBytes := readFileFromProviderResourceDir(t, "openapi_generated.yml")
	d := openapi.GetOpenAPISpec(openapiBytes)
	d.Servers[0].URL = "http://localhost:8080"
	openapiBytes, _ = d.MarshalJSON()

	pSchemaBytes := readFileFromProviderResourceDir(t, "schema.json")
	metadataBytes := readFileFromProviderResourceDir(t, "metadata.json")

	p, err := makeProvider(nil, "", "", pSchemaBytes, openapiBytes, metadataBytes)

	if err != nil {
		t.Fatalf("Could not create a provider instance: %v", err)
	}

	_, err = p.Configure(ctx, &pulumirpc.ConfigureRequest{
		Variables: map[string]string{"scaleway-instances:config:apiKey": "fakeapikey"},
	})

	if err != nil {
		t.Fatalf("Error configuring the provider: %v", err)
	}

	return p
}

func TestDiff(t *testing.T) {
	ctx := context.Background()

	p := makeTestProvider(ctx, t)

	outputs := make(map[string]interface{})
	outputs["name"] = "Test"
	oldsStruct, _ := plugin.MarshalProperties(state.GetResourceState(outputs, resource.NewPropertyMapFromMap(outputs)), state.DefaultMarshalOpts)

	news := make(map[string]interface{})
	news["name"] = "Test2"
	newsStruct, _ := plugin.MarshalProperties(resource.NewPropertyMapFromMap(news), state.DefaultMarshalOpts)

	resp, err := p.Diff(ctx, &pulumirpc.DiffRequest{Id: "", Urn: "urn:pulumi:some-stack::some-project::scaleway-instances:servers:Server::someResourceName", Olds: oldsStruct, News: newsStruct})
	assert.Nil(t, err)
	assert.Equal(t, pulumirpc.DiffResponse_DIFF_SOME, resp.Changes)
	assert.NotEmpty(t, resp.Diffs)
	assert.Len(t, resp.Diffs, 1)
	assert.Empty(t, resp.Replaces)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()

	var inputs map[string]interface{}
	if err := json.Unmarshal([]byte(testCreateJSONPayload), &inputs); err != nil {
		t.Fatal("Failed to unmarshal test payload")
	}

	p := makeTestProvider(ctx, t)

	inputProperties, _ := plugin.MarshalProperties(resource.NewPropertyMapFromMap(inputs), state.DefaultMarshalOpts)

	_, err := p.Create(ctx, &pulumirpc.CreateRequest{
		Urn:        "urn:pulumi:dev::scaleway-instances-ts::scaleway-instances:servers:Server::testServer",
		Properties: inputProperties,
	})

	assert.NotNil(t, err)
	// For now just assume that if we got to the point of making the request, we are good to go.
	assert.Contains(t, err.Error(), "connect: connection refused")
}
