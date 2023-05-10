/*
AVACloud API 1.41.4

Testing OenormConversionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package avaclient

import (
	"context"
	openapiclient "github.com/prolicht-dev/avaclient-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_avaclient_OenormConversionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OenormConversionApiService OenormConversionConvertToAva", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OenormConversionApi.OenormConversionConvertToAva(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OenormConversionApiService OenormConversionConvertToExcel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OenormConversionApi.OenormConversionConvertToExcel(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OenormConversionApiService OenormConversionConvertToGaeb", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OenormConversionApi.OenormConversionConvertToGaeb(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OenormConversionApiService OenormConversionConvertToOenorm", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OenormConversionApi.OenormConversionConvertToOenorm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
