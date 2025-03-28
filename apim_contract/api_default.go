/*
API Manager Contract API

API Manager Contract API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_contract

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

type DefaultApiCreateApiContractRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	postApiContract *PostApiContract
}

func (r DefaultApiCreateApiContractRequest) PostApiContract(postApiContract PostApiContract) DefaultApiCreateApiContractRequest {
	r.postApiContract = &postApiContract
	return r
}

func (r DefaultApiCreateApiContractRequest) Execute() (*ContractDetails, *http.Response, error) {
	return r.ApiService.CreateApiContractExecute(r)
}

/*
CreateApiContract Creates a new contract for a given API in a given organization and environment

Creates a new contract for a given API in a given organization and environment Connected Apps require the following scopes: - Manage Contracts - Exchange Viewer


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @return DefaultApiCreateApiContractRequest
*/
func (a *DefaultApiService) CreateApiContract(ctx context.Context, orgId string, envId string, apiId string) DefaultApiCreateApiContractRequest {
	return DefaultApiCreateApiContractRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
	}
}

// Execute executes the request
//  @return ContractDetails
func (a *DefaultApiService) CreateApiContractExecute(r DefaultApiCreateApiContractRequest) (*ContractDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateApiContract")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postApiContract == nil {
		return localVarReturnValue, nil, reportError("postApiContract is required and must be specified")
	}

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
	localVarPostBody = r.postApiContract
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

type DefaultApiDeleteApiContractRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	contractId string
}

func (r DefaultApiDeleteApiContractRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApiContractExecute(r)
}

/*
DeleteApiContract Deletes a contract for a given API in a given organization and environment

Deletes a contract for a given API in a given organization and environment Connected Apps require the following scopes: - Manage Contracts


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @param contractId The contract id
 @return DefaultApiDeleteApiContractRequest
*/
func (a *DefaultApiService) DeleteApiContract(ctx context.Context, orgId string, envId string, apiId string, contractId string) DefaultApiDeleteApiContractRequest {
	return DefaultApiDeleteApiContractRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
		contractId: contractId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteApiContractExecute(r DefaultApiDeleteApiContractRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteApiContract")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contractId"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)

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

type DefaultApiGetApiContractRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	contractId string
}

func (r DefaultApiGetApiContractRequest) Execute() (*ContractDetails, *http.Response, error) {
	return r.ApiService.GetApiContractExecute(r)
}

/*
GetApiContract Retrieves a contract for a given API in a given organization and environment

Retrieves a contract for a given API in a given organization and environment Connected Apps require the following scopes: - View Contracts


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @param contractId The contract id
 @return DefaultApiGetApiContractRequest
*/
func (a *DefaultApiService) GetApiContract(ctx context.Context, orgId string, envId string, apiId string, contractId string) DefaultApiGetApiContractRequest {
	return DefaultApiGetApiContractRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
		contractId: contractId,
	}
}

// Execute executes the request
//  @return ContractDetails
func (a *DefaultApiService) GetApiContractExecute(r DefaultApiGetApiContractRequest) (*ContractDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetApiContract")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contractId"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)

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

type DefaultApiGetApiContractsRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	limit *int32
	offset *int32
	sort *string
	order *string
	ascending *bool
	status *string
	query *string
	includeExtraApplicationData *bool
}

// The maximum number of contracts to return
func (r DefaultApiGetApiContractsRequest) Limit(limit int32) DefaultApiGetApiContractsRequest {
	r.limit = &limit
	return r
}

// The offset of the first contract to return
func (r DefaultApiGetApiContractsRequest) Offset(offset int32) DefaultApiGetApiContractsRequest {
	r.offset = &offset
	return r
}

// The field to sort by
func (r DefaultApiGetApiContractsRequest) Sort(sort string) DefaultApiGetApiContractsRequest {
	r.sort = &sort
	return r
}

// The order to sort by
func (r DefaultApiGetApiContractsRequest) Order(order string) DefaultApiGetApiContractsRequest {
	r.order = &order
	return r
}

// The order to sort by
func (r DefaultApiGetApiContractsRequest) Ascending(ascending bool) DefaultApiGetApiContractsRequest {
	r.ascending = &ascending
	return r
}

// The status of the contract
func (r DefaultApiGetApiContractsRequest) Status(status string) DefaultApiGetApiContractsRequest {
	r.status = &status
	return r
}

// A string that will be checked for a partial or similar matches of the name, description, label and tags
func (r DefaultApiGetApiContractsRequest) Query(query string) DefaultApiGetApiContractsRequest {
	r.query = &query
	return r
}

// Whether to include extra application data
func (r DefaultApiGetApiContractsRequest) IncludeExtraApplicationData(includeExtraApplicationData bool) DefaultApiGetApiContractsRequest {
	r.includeExtraApplicationData = &includeExtraApplicationData
	return r
}

func (r DefaultApiGetApiContractsRequest) Execute() (*GetContractsResponse, *http.Response, error) {
	return r.ApiService.GetApiContractsExecute(r)
}

/*
GetApiContracts Retrieves all contracts for a given API in a given organization and environment

Retrieves all contracts for a given API in a given organization and environment Connected Apps require the following scopes: - View Contracts


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @return DefaultApiGetApiContractsRequest
*/
func (a *DefaultApiService) GetApiContracts(ctx context.Context, orgId string, envId string, apiId string) DefaultApiGetApiContractsRequest {
	return DefaultApiGetApiContractsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
	}
}

// Execute executes the request
//  @return GetContractsResponse
func (a *DefaultApiService) GetApiContractsExecute(r DefaultApiGetApiContractsRequest) (*GetContractsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetContractsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetApiContracts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.ascending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ascending", r.ascending, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	}
	if r.includeExtraApplicationData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeExtraApplicationData", r.includeExtraApplicationData, "")
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

type DefaultApiUpdateApiContractRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	contractId string
	patchApiContract *PatchApiContract
}

func (r DefaultApiUpdateApiContractRequest) PatchApiContract(patchApiContract PatchApiContract) DefaultApiUpdateApiContractRequest {
	r.patchApiContract = &patchApiContract
	return r
}

func (r DefaultApiUpdateApiContractRequest) Execute() (*ContractDetails, *http.Response, error) {
	return r.ApiService.UpdateApiContractExecute(r)
}

/*
UpdateApiContract Updates a contract for a given API in a given organization and environment

Updates a contract for a given API in a given organization and environment Connected Apps require the following scopes: - Manage Contracts - Exchange Viewer


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @param contractId The contract id
 @return DefaultApiUpdateApiContractRequest
*/
func (a *DefaultApiService) UpdateApiContract(ctx context.Context, orgId string, envId string, apiId string, contractId string) DefaultApiUpdateApiContractRequest {
	return DefaultApiUpdateApiContractRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
		contractId: contractId,
	}
}

// Execute executes the request
//  @return ContractDetails
func (a *DefaultApiService) UpdateApiContractExecute(r DefaultApiUpdateApiContractRequest) (*ContractDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateApiContract")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contractId"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchApiContract == nil {
		return localVarReturnValue, nil, reportError("patchApiContract is required and must be specified")
	}

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
	localVarPostBody = r.patchApiContract
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

type DefaultApiUpdateApiContractStatusRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiId string
	contractId string
	action string
}

func (r DefaultApiUpdateApiContractStatusRequest) Execute() (*ContractDetails, *http.Response, error) {
	return r.ApiService.UpdateApiContractStatusExecute(r)
}

/*
UpdateApiContractStatus Performs an action on a contract for a given API in a given organization and environment

Performs an action on a contract for a given API in a given organization and environment Connected Apps require the following scopes: - Manage Contracts


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param apiId The api manager instance id for a given environment
 @param contractId The contract id
 @param action The action to be performed on the contract
 @return DefaultApiUpdateApiContractStatusRequest
*/
func (a *DefaultApiService) UpdateApiContractStatus(ctx context.Context, orgId string, envId string, apiId string, contractId string, action string) DefaultApiUpdateApiContractStatusRequest {
	return DefaultApiUpdateApiContractStatusRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiId: apiId,
		contractId: contractId,
		action: action,
	}
}

// Execute executes the request
//  @return ContractDetails
func (a *DefaultApiService) UpdateApiContractStatusExecute(r DefaultApiUpdateApiContractStatusRequest) (*ContractDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateApiContractStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/xapi/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId}/{action}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiId"+"}", url.PathEscape(parameterValueToString(r.apiId, "apiId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contractId"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"action"+"}", url.PathEscape(parameterValueToString(r.action, "action")), -1)

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
