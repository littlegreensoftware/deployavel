package resources_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/littlegreensoftware/deployavel/mocks"
	"github.com/littlegreensoftware/deployavel/resources"
)

var siteData *os.File

var params = resources.SiteRequest{
	Site: resources.SiteInput{
		Domain:      "test.com",
		ProjectType: "php",
		Directory:   "/test",
	},
}

func createSiteTest(t *testing.T) {
	var site *resources.Site

	r := mocks.MakeMockAPIRequest(siteData, http.StatusOK)

	site, err := resources.SiteCreate(r, 1, params)

	if err != nil {
		t.Errorf("Should not have an error: %v", err)
	}

	if site == nil {
		t.Error("Should have a site")
	}
}

func createSiteWithErrorsTest(t *testing.T) {
	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		server, err := resources.SiteCreate(r, 1, params)

		if err == nil {
			t.Errorf("Should have an error")
		}

		if server != nil {
			t.Error("Should not have a site")
		}
	}
}

func TestSiteSetup(t *testing.T) {
	siteData = JsonData("site.json")
	defer siteData.Close()

	t.Run("should create a server", createSiteTest)
	t.Run("should handle errors", createSiteWithErrorsTest)
}
