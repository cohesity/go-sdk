/*
Cohesity REST API

Testing CloudRetrieveTaskAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/cohesity/go-sdk"
)

func Test_v2_CloudRetrieveTaskAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudRetrieveTaskAPIService CreateCloudRetrieveTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudRetrieveTaskAPI.CreateCloudRetrieveTask(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudRetrieveTaskAPIService GetCloudRetrieveTaskByJobId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId int64

		resp, httpRes, err := apiClient.CloudRetrieveTaskAPI.GetCloudRetrieveTaskByJobId(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudRetrieveTaskAPIService GetCloudRetrieveTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudRetrieveTaskAPI.GetCloudRetrieveTasks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
