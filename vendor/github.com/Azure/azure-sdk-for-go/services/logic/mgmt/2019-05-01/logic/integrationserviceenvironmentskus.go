package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IntegrationServiceEnvironmentSkusClient is the REST API for Azure Logic Apps.
type IntegrationServiceEnvironmentSkusClient struct {
	BaseClient
}

// NewIntegrationServiceEnvironmentSkusClient creates an instance of the IntegrationServiceEnvironmentSkusClient
// client.
func NewIntegrationServiceEnvironmentSkusClient(subscriptionID string) IntegrationServiceEnvironmentSkusClient {
	return NewIntegrationServiceEnvironmentSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationServiceEnvironmentSkusClientWithBaseURI creates an instance of the
// IntegrationServiceEnvironmentSkusClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIntegrationServiceEnvironmentSkusClientWithBaseURI(baseURI string, subscriptionID string) IntegrationServiceEnvironmentSkusClient {
	return IntegrationServiceEnvironmentSkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of integration service environment Skus.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
func (client IntegrationServiceEnvironmentSkusClient) List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result IntegrationServiceEnvironmentSkuListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentSkusClient.List")
		defer func() {
			sc := -1
			if result.isesl.Response.Response != nil {
				sc = result.isesl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup, integrationServiceEnvironmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.isesl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "List", resp, "Failure sending request")
		return
	}

	result.isesl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationServiceEnvironmentSkusClient) ListPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentSkusClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentSkusClient) ListResponder(resp *http.Response) (result IntegrationServiceEnvironmentSkuList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationServiceEnvironmentSkusClient) listNextResults(ctx context.Context, lastResults IntegrationServiceEnvironmentSkuList) (result IntegrationServiceEnvironmentSkuList, err error) {
	req, err := lastResults.integrationServiceEnvironmentSkuListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentSkusClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationServiceEnvironmentSkusClient) ListComplete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result IntegrationServiceEnvironmentSkuListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentSkusClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup, integrationServiceEnvironmentName)
	return
}
