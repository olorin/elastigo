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
	"fmt"
	//"github.com/araddon/gou"
	"github.com/bmizerany/assert"
	"testing"
)

func TestFilters(t *testing.T) {
	// search for docs that are missing repository.name
	qry := Search("github").Filter(
		Filter().Exists("repository.name"),
	)
	out, err := qry.Result()
	assert.T(t, err == nil, t, "should not have error")
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 7695, fmt.Sprintf("Should have 7695 total= %v", out.Hits.Total))

	qry = Search("github").Filter(
		Filter().Missing("repository.name"),
	)
	out, _ = qry.Result()
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 389, fmt.Sprintf("Should have 389 total= %v", out.Hits.Total))

	//actor_attributes: {type: "User",
	qry = Search("github").Filter(
		Filter().Terms("actor_attributes.location", "portland"),
	)
	out, _ = qry.Result()
	//gou.Debug(out)
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 71, fmt.Sprintf("Should have 71 total= %v", out.Hits.Total))

	/*
		Should this be an AND by default?
	*/
	qry = Search("github").Filter(
		Filter().Terms("actor_attributes.location", "portland"),
		Filter().Terms("repository.has_wiki", true),
	)
	out, err = qry.Result()
	//gou.Debug(out)
	assert.T(t, err == nil, t, "should not have error")
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 44, fmt.Sprintf("Should have 44 total= %v", out.Hits.Total))

	// NOW, lets try with two query calls instead of one
	qry = Search("github").Filter(
		Filter().Terms("actor_attributes.location", "portland"),
	)
	qry.Filter(
		Filter().Terms("repository.has_wiki", true),
	)
	out, err = qry.Result()
	//gou.Debug(out)
	assert.T(t, err == nil, t, "should not have error")
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 44, fmt.Sprintf("Should have 44 total= %v", out.Hits.Total))

	qry = Search("github").Filter(
		"or",
		Filter().Terms("actor_attributes.location", "portland"),
		Filter().Terms("repository.has_wiki", true),
	)
	out, err = qry.Result()
	assert.T(t, err == nil, t, "should not have error")
	assert.T(t, out.Hits.Len() == 10, fmt.Sprintf("Should have 10 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 6676, fmt.Sprintf("Should have 6676 total= %v", out.Hits.Total))
}

func TestFilterRange(t *testing.T) {

	// now lets filter range for repositories with more than 100 forks
	out, _ := Search("github").Size("25").Filter(
		Range().Field("repository.forks").From("100"),
	).Result()
	if out == nil || &out.Hits == nil {
		t.Fail()
		return
	}

	assert.T(t, out.Hits.Len() == 25, fmt.Sprintf("Should have 25 docs %v", out.Hits.Len()))
	assert.T(t, out.Hits.Total == 725, fmt.Sprintf("Should have total=725 but was %v", out.Hits.Total))
}
