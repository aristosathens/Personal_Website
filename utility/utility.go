package utility

import ()

//
// This contains miscellaneous functions for use in other packages in the Web project
//

// ------------------------------------------- Public ------------------------------------------- //

// Given a map src, copies all of its elements into dest
func CopyStringMap(src map[string]string, dest *map[string]string) *map[string]string {
	for key, element := range src {
		(*dest)[key] = element
	}
	return dest
}
