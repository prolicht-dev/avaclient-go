package main

import (
	"context"
	"fmt"
	"os"

	"github.com/prolicht-dev/avaclient-go"
)

// A simple test application that demonstrates the use of the API client.
func main() {
	cfg := avaclient.NewConfiguration()
	api := avaclient.NewAPIClient(cfg)
	status, _, err := api.StatusApi.StatusGetStatus(context.Background())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Healthy:", status.IsHealthy, "; Version:", status.Version)

	token, err := avaclient.NewToken("sample-client-id", "super secret!")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(token)

	authCtx := context.WithValue(context.Background(), avaclient.ContextAccessToken, token.AccessToken)

	avaProject := avaclient.ProjectDto{
		ServiceSpecifications: []avaclient.ServiceSpecificationDto{
			{
				Elements: []avaclient.IElementDtoHolder{
					{
						avaclient.PositionDto{
							ElementTypeDiscriminator: "Position",
							UnitPriceOverride:        2,
							Quantity:                 5,
							ShortText:                "test text",
						},
					},
				},
			},
		},
	}
	gaeb, _, err := api.AvaConversionApi.AvaConversionConvertToGaeb(authCtx, avaProject, nil)
	if err != nil {
		fmt.Println(string(err.(avaclient.GenericOpenAPIError).Body()))
		fmt.Println(err)
		os.Exit(1)
	}

	fi, err := gaeb.Stat()
	fmt.Println(fi.Size())
	var buffer = make([]byte, fi.Size())
	gaeb.Read(buffer)
	gaeb.Close()
	fmt.Println(string(buffer))
}
