package insights

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
	"net/http"
)

// WorkbooksClient is the composite Swagger for Application Insights Management Client
type WorkbooksClient struct {
	BaseClient
}

// NewWorkbooksClient creates an instance of the WorkbooksClient client.
func NewWorkbooksClient(subscriptionID string) WorkbooksClient {
	return NewWorkbooksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkbooksClientWithBaseURI creates an instance of the WorkbooksClient client.
func NewWorkbooksClientWithBaseURI(baseURI string, subscriptionID string) WorkbooksClient {
	return WorkbooksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByResourceGroup get all Workbooks defined within a specified resource group and category.
//
// resourceGroupName is the name of the resource group. location is the name of location where workbook is stored.
// category is category of workbook to return. tags is tags presents on each workbook returned. canFetchContent is
// flag indicating whether or not to return the full content for each applicable workbook. If false, only return
// summary content for workbooks.
func (client WorkbooksClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, location string, category CategoryType, tags []string, canFetchContent *bool) (result Workbooks, err error) {
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, location, category, tags, canFetchContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkbooksClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.WorkbooksClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkbooksClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkbooksClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, location string, category CategoryType, tags []string, canFetchContent *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"category":    autorest.Encode("query", category),
		"location":    autorest.Encode("query", location),
	}
	if tags != nil && len(tags) > 0 {
		queryParameters["tags"] = autorest.Encode("query", tags, ",")
	}
	if canFetchContent != nil {
		queryParameters["canFetchContent"] = autorest.Encode("query", *canFetchContent)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkbooksClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkbooksClient) ListByResourceGroupResponder(resp *http.Response) (result Workbooks, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
