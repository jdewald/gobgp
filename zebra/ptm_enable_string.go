// Code generated by "stringer -type=PTM_ENABLE"; DO NOT EDIT.

package zebra

import "fmt"

const _PTM_ENABLE_name = "PTM_ENABLE_OFFPTM_ENABLE_ONPTM_ENABLE_UNSPEC"

var _PTM_ENABLE_index = [...]uint8{0, 14, 27, 44}

func (i PTM_ENABLE) String() string {
	if i >= PTM_ENABLE(len(_PTM_ENABLE_index)-1) {
		return fmt.Sprintf("PTM_ENABLE(%d)", i)
	}
	return _PTM_ENABLE_name[_PTM_ENABLE_index[i]:_PTM_ENABLE_index[i+1]]
}
