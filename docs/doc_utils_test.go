package docs

import "testing"

func TestDocs(t *testing.T) {

	err := GenerateDocs("../.tmp", "../reference")
	if err != nil {
		t.Error(err)
	}

}
