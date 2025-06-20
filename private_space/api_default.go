/*
Private Space API

Description of the Private Space API

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package private_space

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiCreatePrivateSpaceRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	privateSpacePostBody *PrivateSpacePostBody
}

func (r DefaultApiCreatePrivateSpaceRequest) PrivateSpacePostBody(privateSpacePostBody PrivateSpacePostBody) DefaultApiCreatePrivateSpaceRequest {
	r.privateSpacePostBody = &privateSpacePostBody
	return r
}

func (r DefaultApiCreatePrivateSpaceRequest) Execute() (*PrivateSpace, *http.Response, error) {
	return r.ApiService.CreatePrivateSpaceExecute(r)
}

/*
CreatePrivateSpace Method for CreatePrivateSpace

creates a private space for the given organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @return DefaultApiCreatePrivateSpaceRequest
*/
func (a *DefaultApiService) CreatePrivateSpace(ctx context.Context, orgId string) DefaultApiCreatePrivateSpaceRequest {
	return DefaultApiCreatePrivateSpaceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return PrivateSpace
func (a *DefaultApiService) CreatePrivateSpaceExecute(r DefaultApiCreatePrivateSpaceRequest) (*PrivateSpace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateSpace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreatePrivateSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/privatespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.privateSpacePostBody
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

type DefaultApiDeletePrivateSpaceRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	privateSpaceId string
}

func (r DefaultApiDeletePrivateSpaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePrivateSpaceExecute(r)
}

/*
DeletePrivateSpace Method for DeletePrivateSpace

deletes a private space for the given organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param privateSpaceId The ID of the private space in GUID format
 @return DefaultApiDeletePrivateSpaceRequest
*/
func (a *DefaultApiService) DeletePrivateSpace(ctx context.Context, orgId string, privateSpaceId string) DefaultApiDeletePrivateSpaceRequest {
	return DefaultApiDeletePrivateSpaceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		privateSpaceId: privateSpaceId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeletePrivateSpaceExecute(r DefaultApiDeletePrivateSpaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeletePrivateSpace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/privatespaces/{privateSpaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"privateSpaceId"+"}", url.PathEscape(parameterValueToString(r.privateSpaceId, "privateSpaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiGetPrivateSpaceRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	privateSpaceId string
}

func (r DefaultApiGetPrivateSpaceRequest) Execute() (*PrivateSpace, *http.Response, error) {
	return r.ApiService.GetPrivateSpaceExecute(r)
}

/*
GetPrivateSpace Method for GetPrivateSpace

retrieves a private space for the given organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param privateSpaceId The ID of the private space in GUID format
 @return DefaultApiGetPrivateSpaceRequest
*/
func (a *DefaultApiService) GetPrivateSpace(ctx context.Context, orgId string, privateSpaceId string) DefaultApiGetPrivateSpaceRequest {
	return DefaultApiGetPrivateSpaceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		privateSpaceId: privateSpaceId,
	}
}

// Execute executes the request
//  @return PrivateSpace
func (a *DefaultApiService) GetPrivateSpaceExecute(r DefaultApiGetPrivateSpaceRequest) (*PrivateSpace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateSpace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPrivateSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/privatespaces/{privateSpaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"privateSpaceId"+"}", url.PathEscape(parameterValueToString(r.privateSpaceId, "privateSpaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DefaultApiGetPrivateSpacesRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	offset *int32
	limit *int32
	sort *string
	ascending *string
}

// The offset specifies the offset of the first row to return
func (r DefaultApiGetPrivateSpacesRequest) Offset(offset int32) DefaultApiGetPrivateSpacesRequest {
	r.offset = &offset
	return r
}

// Amount of objects retrieved in the response
func (r DefaultApiGetPrivateSpacesRequest) Limit(limit int32) DefaultApiGetPrivateSpacesRequest {
	r.limit = &limit
	return r
}

// Property to sort by
func (r DefaultApiGetPrivateSpacesRequest) Sort(sort string) DefaultApiGetPrivateSpacesRequest {
	r.sort = &sort
	return r
}

// Order for sorting
func (r DefaultApiGetPrivateSpacesRequest) Ascending(ascending string) DefaultApiGetPrivateSpacesRequest {
	r.ascending = &ascending
	return r
}

func (r DefaultApiGetPrivateSpacesRequest) Execute() (*PrivateSpaceSummary, *http.Response, error) {
	return r.ApiService.GetPrivateSpacesExecute(r)
}

/*
GetPrivateSpaces Method for GetPrivateSpaces

retrieves private spaces for the given organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @return DefaultApiGetPrivateSpacesRequest
*/
func (a *DefaultApiService) GetPrivateSpaces(ctx context.Context, orgId string) DefaultApiGetPrivateSpacesRequest {
	return DefaultApiGetPrivateSpacesRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return PrivateSpaceSummary
func (a *DefaultApiService) GetPrivateSpacesExecute(r DefaultApiGetPrivateSpacesRequest) (*PrivateSpaceSummary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateSpaceSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPrivateSpaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/privatespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.ascending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ascending", r.ascending, "")
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

type DefaultApiUpdatePrivateSpaceRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	privateSpaceId string
	privateSpacePatchBody *PrivateSpacePatchBody
}

func (r DefaultApiUpdatePrivateSpaceRequest) PrivateSpacePatchBody(privateSpacePatchBody PrivateSpacePatchBody) DefaultApiUpdatePrivateSpaceRequest {
	r.privateSpacePatchBody = &privateSpacePatchBody
	return r
}

func (r DefaultApiUpdatePrivateSpaceRequest) Execute() (*PrivateSpace, *http.Response, error) {
	return r.ApiService.UpdatePrivateSpaceExecute(r)
}

/*
UpdatePrivateSpace Method for UpdatePrivateSpace

updates a private space for the given organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param privateSpaceId The ID of the private space in GUID format
 @return DefaultApiUpdatePrivateSpaceRequest
*/
func (a *DefaultApiService) UpdatePrivateSpace(ctx context.Context, orgId string, privateSpaceId string) DefaultApiUpdatePrivateSpaceRequest {
	return DefaultApiUpdatePrivateSpaceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		privateSpaceId: privateSpaceId,
	}
}

// Execute executes the request
//  @return PrivateSpace
func (a *DefaultApiService) UpdatePrivateSpaceExecute(r DefaultApiUpdatePrivateSpaceRequest) (*PrivateSpace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateSpace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdatePrivateSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/privatespaces/{privateSpaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"privateSpaceId"+"}", url.PathEscape(parameterValueToString(r.privateSpaceId, "privateSpaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.privateSpacePatchBody
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
