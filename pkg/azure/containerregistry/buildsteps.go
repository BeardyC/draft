package containerregistry

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
	"net/http"

	"github.com/BeardyC/go-autorest/autorest"
	"github.com/BeardyC/go-autorest/autorest/azure"
	"github.com/BeardyC/go-autorest/autorest/validation"
)

// BuildStepsClient is the client for the BuildSteps methods of the Containerregistry service.
type BuildStepsClient struct {
	BaseClient
}

// NewBuildStepsClient creates an instance of the BuildStepsClient client.
func NewBuildStepsClient(subscriptionID string) BuildStepsClient {
	return NewBuildStepsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBuildStepsClientWithBaseURI creates an instance of the BuildStepsClient client.
func NewBuildStepsClientWithBaseURI(baseURI string, subscriptionID string) BuildStepsClient {
	return BuildStepsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task. stepName is the
// name of a build step for a container registry build task. buildStepCreateParameters is the parameters for
// creating a build step.
func (client BuildStepsClient) Create(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string, buildStepCreateParameters BuildStep) (result BuildStepsCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: stepName,
			Constraints: []validation.Constraint{{Target: "stepName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "stepName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "stepName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "Create")
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, buildTaskName, stepName, buildStepCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BuildStepsClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string, buildStepCreateParameters BuildStep) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps/{stepName}", pathParameters),
		autorest.WithJSON(buildStepCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) CreateSender(req *http.Request) (future BuildStepsCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) CreateResponder(resp *http.Response) (result BuildStep, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task. stepName is the
// name of a build step for a container registry build task.
func (client BuildStepsClient) Delete(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (result BuildStepsDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: stepName,
			Constraints: []validation.Constraint{{Target: "stepName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "stepName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "stepName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "Delete")
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, buildTaskName, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BuildStepsClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) DeleteSender(req *http.Request) (future BuildStepsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task. stepName is the
// name of a build step for a container registry build task.
func (client BuildStepsClient) Get(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (result BuildStep, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: stepName,
			Constraints: []validation.Constraint{{Target: "stepName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "stepName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "stepName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "Get")
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, buildTaskName, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BuildStepsClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) GetResponder(resp *http.Response) (result BuildStep, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task.
func (client BuildStepsClient) List(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (result BuildStepListPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "List")
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName, buildTaskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.bsl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "List", resp, "Failure sending request")
		return
	}

	result.bsl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BuildStepsClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) ListResponder(resp *http.Response) (result BuildStepList, err error) {
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
func (client BuildStepsClient) listNextResults(lastResults BuildStepList) (result BuildStepList, err error) {
	req, err := lastResults.buildStepListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BuildStepsClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (result BuildStepListIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, registryName, buildTaskName)
	return
}

// ListBuildArguments sends the list build arguments request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task. stepName is the
// name of a build step for a container registry build task.
func (client BuildStepsClient) ListBuildArguments(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (result BuildArgumentListPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: stepName,
			Constraints: []validation.Constraint{{Target: "stepName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "stepName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "stepName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "ListBuildArguments")
	}

	result.fn = client.listBuildArgumentsNextResults
	req, err := client.ListBuildArgumentsPreparer(ctx, resourceGroupName, registryName, buildTaskName, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "ListBuildArguments", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuildArgumentsSender(req)
	if err != nil {
		result.bal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "ListBuildArguments", resp, "Failure sending request")
		return
	}

	result.bal, err = client.ListBuildArgumentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "ListBuildArguments", resp, "Failure responding to request")
	}

	return
}

// ListBuildArgumentsPreparer prepares the ListBuildArguments request.
func (client BuildStepsClient) ListBuildArgumentsPreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps/{stepName}/listBuildArguments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBuildArgumentsSender sends the ListBuildArguments request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) ListBuildArgumentsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBuildArgumentsResponder handles the response to the ListBuildArguments request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) ListBuildArgumentsResponder(resp *http.Response) (result BuildArgumentList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBuildArgumentsNextResults retrieves the next set of results, if any.
func (client BuildStepsClient) listBuildArgumentsNextResults(lastResults BuildArgumentList) (result BuildArgumentList, err error) {
	req, err := lastResults.buildArgumentListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listBuildArgumentsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBuildArgumentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listBuildArgumentsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBuildArgumentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "listBuildArgumentsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBuildArgumentsComplete enumerates all values, automatically crossing page boundaries as required.
func (client BuildStepsClient) ListBuildArgumentsComplete(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string) (result BuildArgumentListIterator, err error) {
	result.page, err = client.ListBuildArguments(ctx, resourceGroupName, registryName, buildTaskName, stepName)
	return
}

// Update sends the update request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. buildTaskName is the name of the container registry build task. stepName is the
// name of a build step for a container registry build task. buildStepUpdateParameters is the parameters for
// updating a build step.
func (client BuildStepsClient) Update(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string, buildStepUpdateParameters BuildStepUpdateParameters) (result BuildStepsUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: stepName,
			Constraints: []validation.Constraint{{Target: "stepName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "stepName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "stepName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildStepsClient", "Update")
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, buildTaskName, stepName, buildStepUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildStepsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BuildStepsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, stepName string, buildStepUpdateParameters BuildStepUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/steps/{stepName}", pathParameters),
		autorest.WithJSON(buildStepUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BuildStepsClient) UpdateSender(req *http.Request) (future BuildStepsUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BuildStepsClient) UpdateResponder(resp *http.Response) (result BuildStep, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
