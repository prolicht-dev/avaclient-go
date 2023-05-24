/*
IElementDto is a polymorphic element, OpenAPI generator does not support such elements,
thus a custom IElementDtoHolder struct and interface is used.
*/

package avaclient

import (
	"encoding/json"
	"fmt"
)

type IElementDtoHolder struct {
	Holder IElementDtoInterface
}

func (t *IElementDtoHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Holder)
}

func (t *IElementDtoHolder) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	// We detect the correct struct type by looking at the elementTypeDiscriminator field.
	jsonType, ok := objMap["elementTypeDiscriminator"]
	if !ok {
		return fmt.Errorf("no Type specification")
	}
	var goType string
	err = json.Unmarshal(*jsonType, &goType)
	if err != nil {
		return fmt.Errorf("error getting type: %s", err.Error())
	}

	switch goType {
	case "Position", "PositionDto":
		var elem PositionDto
		err = json.Unmarshal(b, &elem)
		if err != nil {
			return err
		}
		t.Holder = &elem
	case "ExecutionDescription", "ExecutionDescriptionDto":
		var elem ExecutionDescriptionDto
		err = json.Unmarshal(b, &elem)
		if err != nil {
			return err
		}
		t.Holder = &elem
	case "NoteText", "NoteTextDto":
		var elem NoteTextDto
		err = json.Unmarshal(b, &elem)
		if err != nil {
			return err
		}
		t.Holder = &elem
	case "ServiceSpecificationGroup", "ServiceSpecificationGroupDto":
		var elem ServiceSpecificationGroupDto
		err = json.Unmarshal(b, &elem)
		if err != nil {
			return err
		}
		t.Holder = &elem
	default:
		return fmt.Errorf("unknown IElementDTO type %s", goType)
	}

	return nil
}

func (p PositionDto) isIElementDto() {
}

func (e ExecutionDescriptionDto) isIElementDto() {
}

func (n NoteTextDto) isIElementDto() {
}

func (s ServiceSpecificationGroupDto) isIElementDto() {
}
