// Code generated by "enumer -type=nationFields"; DO NOT EDIT.

//
package queries

import (
	"fmt"
)

const _nationFieldsName = "N_NATIONKEYN_NAMEN_REGIONKEYN_COMMENT"

var _nationFieldsIndex = [...]uint8{0, 11, 17, 28, 37}

func (i nationFields) String() string {
	if i < 0 || i >= nationFields(len(_nationFieldsIndex)-1) {
		return fmt.Sprintf("nationFields(%d)", i)
	}
	return _nationFieldsName[_nationFieldsIndex[i]:_nationFieldsIndex[i+1]]
}

var _nationFieldsValues = []nationFields{0, 1, 2, 3}

var _nationFieldsNameToValueMap = map[string]nationFields{
	_nationFieldsName[0:11]:  0,
	_nationFieldsName[11:17]: 1,
	_nationFieldsName[17:28]: 2,
	_nationFieldsName[28:37]: 3,
}

// nationFieldsString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func nationFieldsString(s string) (nationFields, error) {
	if val, ok := _nationFieldsNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to nationFields values", s)
}

// nationFieldsValues returns all values of the enum
func nationFieldsValues() []nationFields {
	return _nationFieldsValues
}

// IsAnationFields returns "true" if the value is listed in the enum definition. "false" otherwise
func (i nationFields) IsAnationFields() bool {
	for _, v := range _nationFieldsValues {
		if i == v {
			return true
		}
	}
	return false
}