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

	status, _, err := api.StatusApi.StatusGetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Healthy:", status.IsHealthy, "; Version:", *status.Version)

	token, err := avaclient.NewToken("the-client-id", "super-secret-client-secret")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(token.Token())

	authCtx := context.WithValue(context.Background(), avaclient.ContextOAuth2, token)

	gaeb, _, err := api.AvaConversionApi.AvaConversionConvertToGaeb(authCtx).
		AvaProject(avaclient.ProjectDto{
			ServiceSpecifications: []avaclient.ServiceSpecificationDto{{
				Elements: []avaclient.IElementDto{{
					ElementTypeDiscriminator: "Position",
				}},
			}},
			GaebXmlId: nil,
		}).
		DestinationGaebType(string(avaclient.DESTINATIONGAEBTYPE_GAEB_XML_V3_3_COMMERCE)).
		Execute()
	if err != nil {
		fmt.Println(string(err.(*avaclient.GenericOpenAPIError).Body()))
		fmt.Println(err)
		os.Exit(1)
	}

	fi, err := gaeb.Stat()
	fmt.Println("GAEB Size:", fi.Size())
	var buffer = make([]byte, fi.Size())
	gaeb.Read(buffer)
	gaeb.Close()
	fmt.Println(string(buffer))
}
