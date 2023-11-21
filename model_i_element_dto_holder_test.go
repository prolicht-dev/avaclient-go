package avaclient

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestMarshall(t *testing.T) {
	avaProject := ProjectDto{
		ServiceSpecifications: []ServiceSpecificationDto{
			{
				Elements: []IElementDtoHolder{
					{
						PositionDto{
							IElementDto: IElementDto{
								ElementTypeDiscriminator: "Position",
							},
							UnitPriceOverride: PtrFloat32(2),
							QuantityOverride:  PtrFloat32(5),
							ShortText:         PtrString("test text"),
						},
					},
				},
			},
		},
	}

	jsonBytes, err := json.MarshalIndent(avaProject, "", "  ")
	ok(t, err)

	expJson := `{
  "forceStrictTotals": false,
  "id": "",
  "priceAccuracy": 0,
  "priceRoundingMode": "",
  "serviceSpecifications": [
    {
      "deductionFactor": 0,
      "elements": [
        {
          "additionType": "",
          "amountToBeEnteredByBidder": false,
          "comissionStatus": "",
          "complemented": false,
          "deductionFactor": 0,
          "elementTypeDiscriminator": "Position",
          "gaebComplementingType": "",
          "hierarchyLevel": 0,
          "id": "",
          "ignoreProjectCataloguePropagation": false,
          "isComplementingPosition": false,
          "isLumpSum": false,
          "notOffered": false,
          "positionType": "",
          "priceCompositionRequired": false,
          "priceType": "",
          "quantityOverride": 5,
          "serviceType": "",
          "shortText": "test text",
          "taxRate": 0,
          "unitPriceOverride": 2,
          "useDifferentTaxRate": false
        }
      ],
      "exchangePhase": "",
      "executionGuaranteePercentage": 0,
      "id": "",
      "ignoreChildPriceUpdates": false,
      "ignoreDuplicateElementIds": false,
      "ignoreDuplicateItemNumbers": false,
      "ignoreProjectCataloguePropagation": false,
      "origin": "",
      "priceType": "",
      "projectTaxRate": 0,
      "warrantyBondPercentage": 0
    }
  ]
}`
	got := string(jsonBytes)
	equals(t, expJson, got)
}

func TestUnMarshall(t *testing.T) {
	jsonStr := `{
  "forceStrictTotals": false,
  "id": "",
  "priceAccuracy": 0,
  "priceRoundingMode": "Normal",
  "serviceSpecifications": [
    {
      "deductionFactor": 0,
      "elements": [
        {
          "additionType": "None",
          "amountToBeEnteredByBidder": false,
          "comissionStatus": "Undefined",
          "complemented": false,
          "deductionFactor": 0,
          "elementTypeDiscriminator": "Position",
          "gaebComplementingType": "Undefined",
          "hierarchyLevel": 0,
          "id": "",
          "isComplementingPosition": false,
          "isLumpSum": false,
          "notOffered": false,
          "positionType": "Regular",
          "priceCompositionRequired": false,
          "priceType": "WithTotal",
          "quantityOverride": 5,
          "serviceType": "Regular",
          "shortText": "test text",
          "taxRate": 0,
          "unitPriceOverride": 2,
          "useDifferentTaxRate": false
        }
      ],
      "exchangePhase": "Base",
      "executionGuaranteePercentage": 0,
      "id": "",
      "ignoreChildPriceUpdates": false,
      "ignoreDuplicateElementIds": false,
      "ignoreDuplicateItemNumbers": false,
      "origin": "GaebXml",
      "priceType": "WithTotal",
      "projectTaxRate": 0,
      "warrantyBondPercentage": 0
    }
  ]
}`

	var decodedProject ProjectDto
	err := json.Unmarshal([]byte(jsonStr), &decodedProject)
	ok(t, err)

	avaProject := ProjectDto{
		PriceRoundingMode: PRICEROUNDINGMODEDTO_NORMAL,
		ServiceSpecifications: []ServiceSpecificationDto{
			{
				Elements: []IElementDtoHolder{
					{
						&PositionDto{
							AdditionType:          ADDITIONTYPEDTO_NONE,
							ComissionStatus:       COMISSIONSTATUSDTO_UNDEFINED,
							GaebComplementingType: POSITIONCOMPLEMENTINGTYPEDTO_UNDEFINED,
							PositionType:          POSITIONTYPEDTO_REGULAR,
							PriceType:             PRICETYPEDTO_WITH_TOTAL,
							ServiceType:           SERVICETYPEDTO_REGULAR,
							IElementDto: IElementDto{
								ElementTypeDiscriminator: "Position",
							},
							UnitPriceOverride: PtrFloat32(2),
							QuantityOverride:  PtrFloat32(5),
							ShortText:         PtrString("test text"),
						},
					},
				},
				ExchangePhase: EXCHANGEPHASEDTO_BASE,
				Origin:        ORIGINDTO_GAEB_XML,
				PriceType:     PRICETYPEDTO_WITH_TOTAL,
			},
		},
	}

	equals(t, avaProject, decodedProject)
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
