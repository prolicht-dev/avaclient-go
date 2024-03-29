/*
AVACloud API 1.41.4

Testing ExcelConversionApiService

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

func Test_avaclient_ExcelConversionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExcelConversionApiService ExcelConversionConvertToAva", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExcelConversionApi.ExcelConversionConvertToAva(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExcelConversionApiService ExcelConversionConvertToExcel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExcelConversionApi.ExcelConversionConvertToExcel(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExcelConversionApiService ExcelConversionConvertToGaeb", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExcelConversionApi.ExcelConversionConvertToGaeb(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExcelConversionApiService ExcelConversionConvertToOenorm", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExcelConversionApi.ExcelConversionConvertToOenorm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
