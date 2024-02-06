/*
Team Members API

Description of the Team Members API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_members

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	membershipType *string
	identityType *string
	memberIds *[]string
	search *string
	limit *int32
	offset *int32
	sort *string
	ascending *bool
}

// Include the group access mappings that grant the provided membership type By default, all group access mappings are returned
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) MembershipType(membershipType string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.membershipType = &membershipType
	return r
}

// A search string to use for case-insensitive partial matches on external group name
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) IdentityType(identityType string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.identityType = &identityType
	return r
}

// Include the members of the team that have ids in this list
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) MemberIds(memberIds []string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.memberIds = &memberIds
	return r
}

// A search string to use for case-insensitive partial matches on team name
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Search(search string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.search = &search
	return r
}

// Maximum number of rolegroups to retrieve per request.
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Limit(limit int32) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.limit = &limit
	return r
}

// The number of records to omit from the response.
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Offset(offset int32) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.offset = &offset
	return r
}

// The field to sort on. default identity_type
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Sort(sort string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.sort = &sort
	return r
}

// Whether to sort ascending or descending. Default true
func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Ascending(ascending bool) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	r.ascending = &ascending
	return r
}

func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) Execute() (*TeamMemberCollection, *http.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdMembersGetExecute(r)
}

/*
OrganizationsOrgIdTeamsTeamIdMembersGet Method for OrganizationsOrgIdTeamsTeamIdMembersGet

retrieves team members

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param teamId The ID of the team in GUID format
 @return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersGet(ctx context.Context, orgId string, teamId string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest {
	return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return TeamMemberCollection
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersGetExecute(r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersGetRequest) (*TeamMemberCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamMemberCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdMembersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.membershipType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "membership_type", r.membershipType, "")
	}
	if r.identityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identity_type", r.identityType, "")
	}
	if r.memberIds != nil {
		t := *r.memberIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "member_ids", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "member_ids", t, "multi")
		}
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	userId string
}

func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteExecute(r)
}

/*
OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete Method for OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete

delete the given user from the given team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param teamId The ID of the team in GUID format
 @param userId The ID of the user in GUID format
 @return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete(ctx context.Context, orgId string, teamId string, userId string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest {
	return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteExecute(r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/members/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	userId string
	teamMemberPutBody *TeamMemberPutBody
}

func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest) TeamMemberPutBody(teamMemberPutBody TeamMemberPutBody) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest {
	r.teamMemberPutBody = &teamMemberPutBody
	return r
}

func (r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdMembersUserIdPutExecute(r)
}

/*
OrganizationsOrgIdTeamsTeamIdMembersUserIdPut Method for OrganizationsOrgIdTeamsTeamIdMembersUserIdPut

adds the given user to the given team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param teamId The ID of the team in GUID format
 @param userId The ID of the user in GUID format
 @return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersUserIdPut(ctx context.Context, orgId string, teamId string, userId string) DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest {
	return DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdMembersUserIdPutExecute(r DefaultApiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdMembersUserIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/members/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.teamMemberPutBody
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
