package task_2

import (
	"github.com/google/go-cmp/cmp"
	"slices"
	"testing"
)

func Test_download(t *testing.T) {
	tests := []struct {
		name          string
		timemoutLimit int
		urls          []string
		want          []string
		wantErr       bool
	}{
		{name: "from the task",
			timemoutLimit: 1000,
			urls: []string{
				"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
				"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
				"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
				"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
				"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
			},
			want: []string{
				"downloaded data from https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml\n",
				"downloaded data from https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml\n",
				"downloaded data from https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml\n",
				"downloaded data from https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml\n",
				"downloaded data from https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml\n",
			},
			wantErr: false,
		},
		{name: "timeoutLimit is 1",
			timemoutLimit: 1,
			urls: []string{
				"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
				"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
				"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
				"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
				"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
			},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		timeoutLimit = tt.timemoutLimit
		t.Run(tt.name, func(t *testing.T) {
			got, err := download(tt.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			slices.Sort(got)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
