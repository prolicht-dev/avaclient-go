/*
IElementDto is a polymorphic element, OpenAPI generator does not support such elements,
thus a custom IElementDtoHolder struct and interface is used.
*/

package avaclient

// IElementDtoInterface Base interface definition for elements within an ElementContainerBase.
type IElementDtoInterface interface {
	isIElementDto()
}
