/**
* Errors tests
* @author  Diana Lucia Serna Higuita
* Test all the functions: $ go test ./...
 */
package errors

import (
	"testing"
)

/**
*  Test errors package
 */
func TestErrors(t *testing.T) {
	/** Test function getErrorIdDescription
	* This function returns the description of the error
	 */
	errorIdDescription := NoErrorId.getErrorIdDescription()
	if len(errorIdDescription) == 0 {
		t.Errorf("Error by the function TestErrors")
	}

	/** Test function getErrorTypeName
	* This function returns the error type
	 */
	errorTypeName := Error.getErrorTypeName()
	if len(errorTypeName) == 0 {
		t.Errorf("Error by the function getErrorTypeName")
	}
}
