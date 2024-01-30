/*
Team API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package team

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func Test_team_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsGet(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsPost(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdDelete(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdGet(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdParentPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdParentPut(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdPatch(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}