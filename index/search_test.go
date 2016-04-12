//
// search_test.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"reflect"
	"testing"

	"fmt"
	"net/http"
)

func TestSearchService_Repositories(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"q": "ubuntu",
		})

		fmt.Fprintf(w, `{"query":"ubuntu","num_results":1,"results":[{"star_count":0,"is_official":true,"name":"official/ubuntu-upstart","is_trusted":true,"description":""}]}`)
	})

	result, _, err := client.Search.Repositories("ubuntu", nil)

	if err != nil {
		t.Errorf("Search.Repositories returned error: %v", err)
	}

	want := &RepositoriesSearchResult{
		NumberResults: Int(1),
		QueryString:   String("ubuntu"),
		Repositories: []Repository{
			{
				Description: String(""),
				IsOfficial:  Bool(true),
				IsTrusted:   Bool(true),
				Name:        String("official/ubuntu-upstart"),
				StarCount:   Int(0),
			},
		},
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Search.Repositories returned %+v, want %+v", result, want)
	}
}
