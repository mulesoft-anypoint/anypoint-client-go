/*
Deployment v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiGetAllDeploymentsRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	xAnypntEnvId *string
	xAnypntOrgId *string
}

// The environment id where the applications are deployed
func (r DefaultApiGetAllDeploymentsRequest) XAnypntEnvId(xAnypntEnvId string) DefaultApiGetAllDeploymentsRequest {
	r.xAnypntEnvId = &xAnypntEnvId
	return r
}

// The org id where the applications are deployed
func (r DefaultApiGetAllDeploymentsRequest) XAnypntOrgId(xAnypntOrgId string) DefaultApiGetAllDeploymentsRequest {
	r.xAnypntOrgId = &xAnypntOrgId
	return r
}

func (r DefaultApiGetAllDeploymentsRequest) Execute() (*GetAllDeploymentsResponse, *http.Response, error) {
	return r.ApiService.GetAllDeploymentsExecute(r)
}

/*
GetAllDeployments List deployments

Retrieves a list of deployments for the specified query. If the query returns no results, then an empty list is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DefaultApiGetAllDeploymentsRequest
*/
func (a *DefaultApiService) GetAllDeployments(ctx context.Context) DefaultApiGetAllDeploymentsRequest {
	return DefaultApiGetAllDeploymentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAllDeploymentsResponse
func (a *DefaultApiService) GetAllDeploymentsExecute(r DefaultApiGetAllDeploymentsRequest) (*GetAllDeploymentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAllDeploymentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAllDeployments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAnypntEnvId == nil {
		return localVarReturnValue, nil, reportError("xAnypntEnvId is required and must be specified")
	}
	if r.xAnypntOrgId == nil {
		return localVarReturnValue, nil, reportError("xAnypntOrgId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anypnt-env-id", r.xAnypntEnvId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anypnt-org-id", r.xAnypntOrgId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
