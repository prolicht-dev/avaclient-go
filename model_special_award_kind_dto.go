/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// SpecialAwardKindDto This enumeration describes awards for project that are not just a regular procurement. For example, it can be used to describe recurring maintenance or an outline contract (German: Rahmenvertrag) which just specifies services and prices but may be requested on demand when necessary
type SpecialAwardKindDto string

// List of SpecialAwardKindDto
const (
	SPECIALAWARDKINDDTO_UNSPECIFIED                      SpecialAwardKindDto = "Unspecified"
	SPECIALAWARDKINDDTO_MAINTENANCE_CONTRACT             SpecialAwardKindDto = "MaintenanceContract"
	SPECIALAWARDKINDDTO_OUTLINE_CONTRACT                 SpecialAwardKindDto = "OutlineContract"
	SPECIALAWARDKINDDTO_OUTLINE_FOR_MAINTENANCE_CONTRACT SpecialAwardKindDto = "OutlineForMaintenanceContract"
)
