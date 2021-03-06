// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	//"encoding/json"
	"fmt"
	"github.com/araddon/gou"
	"github.com/bmizerany/assert"
	"testing"
)

func TestFacetRegex(t *testing.T) {
	// This is a possible solution for auto-complete
	out, _ := Search("github").Size("0").Facet(
		Facet().Regex("repository.name", "no.*").Size("8"),
	).Result()
	if out == nil || &out.Hits == nil {
		t.Fail()
		return
	}
	//Debug(string(out.Facets))
	fh := gou.NewJsonHelper([]byte(out.Facets))
	facets := fh.Helpers("/repository.name/terms")
	assert.T(t, len(facets) == 8, fmt.Sprintf("Should have 8? but was %v", len(facets)))
	// for _, f := range facets {
	// 	Debug(f)
	// }
}
