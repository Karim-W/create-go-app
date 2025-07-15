//
//  assert.go
//  newest_template
//
//  Created by karim-w on 12/07/2025.
//

package assert

import "github.com/karim-w/glose"

// NoErr panics if the provided error is not nil.
// This function is useful for ensuring that errors are handled immediately,
func NoErr(err error) {
	if err != nil {
		glose.Panik(err)
	}
}
