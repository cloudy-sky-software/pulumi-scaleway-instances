// Copyright 2022, Cloudy Sky Software LLC.

package gen

import (
	"bytes"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	openapigen "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/provider/pkg/gen/examples"
)

const packageName = "scaleway-instances"

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openapiDoc openapi3.T) (pschema.PackageSpec, openapigen.ProviderMetadata) {
	pkg := pschema.PackageSpec{
		Name:        packageName,
		Description: "A Pulumi package for creating and managing Scaleway Instances resources.",
		DisplayName: "Scaleway Instances",
		License:     "Apache-2.0",
		Keywords: []string{
			"pulumi",
			packageName,
			"category/cloud",
			"kind/native",
		},
		Homepage:   "https://cloudysky.software",
		Publisher:  "Cloudy Sky Software",
		Repository: "https://github.com/cloudy-sky-software/pulumi-scaleway-instances",

		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"apiKey": {
					Description: "The API key",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
			},
		},

		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the Scaleway Instances package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"apiKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"SCALEWAY_INSTANCES_APIKEY",
						},
					},
					Description: "The Scaleway Instances API key.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
			},
		},

		PluginDownloadURL: "github://api.github.com/cloudy-sky-software/pulumi-scaleway-instances",
		Types:             map[string]pschema.ComplexTypeSpec{},
		Resources:         map[string]pschema.ResourceSpec{},
		Functions:         map[string]pschema.FunctionSpec{},
		Language:          map[string]pschema.RawMessage{},
	}

	csharpNamespaces := map[string]string{
		"scaleway-instances": "ScalewayInstances",
		// TODO: Is this needed?
		"": "Provider",
	}

	openAPICtx := &openapigen.OpenAPIContext{
		Doc:                       openapiDoc,
		Pkg:                       &pkg,
		ExcludedPaths:             []string{},
		UseParentResourceAsModule: true,
	}

	providerMetadata, err := openAPICtx.GatherResourcesFromAPI(csharpNamespaces)
	if err != nil {
		contract.Failf("generating resources from OpenAPI spec: %v", err)
	}

	// Add examples to resources
	for k, v := range examples.ResourceExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"rootNamespace": "Pulumi",
		"packageReferences": map[string]string{
			"Pulumi": "3.*",
		},
		"namespaces": csharpNamespaces,
		// TODO: What does this enable?
		// "dictionaryConstructors": true,
	})

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst",
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"packageName": "@cloudyskysoftware/pulumi-scaleway-instances",
		"dependencies": map[string]string{
			"@pulumi/pulumi":    "^3.0.0",
			"shell-quote":       "^1.6.1",
			"tmp":               "^0.0.33",
			"@types/tmp":        "^0.0.33",
			"glob":              "^7.1.2",
			"@types/glob":       "^5.0.35",
			"node-fetch":        "^2.3.0",
			"@types/node-fetch": "^2.1.4",
		},
		"devDependencies": map[string]string{
			"mocha":              "^5.2.0",
			"@types/mocha":       "^5.2.5",
			"@types/shell-quote": "^1.6.0",
		},
		"moduleToPackage": map[string]string{
			// Snapshot export module conflicts with Node's reserved keyword,
			// so an override is necessary.
			"export": "snapshot_export",
		},
	})
	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"packageName": "pulumi_scaleway_instances",
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
	})

	metadata := openapigen.ProviderMetadata{
		ResourceCRUDMap: providerMetadata.ResourceCRUDMap,
		AutoNameMap:     providerMetadata.AutoNameMap,
	}
	return pkg, metadata
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
