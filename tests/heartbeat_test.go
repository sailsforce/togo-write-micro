package tests

import (
	"log"
	"net/http/httptest"
	"os"

	"github.com/go-chi/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	kit_models "github.com/sailsforce/gomicro-kit/models"
	"github.com/sailsforce/togo-write-micro/internal/config"
	"github.com/sailsforce/togo-write-micro/internal/rest"
)

var _ = Describe("Heartbeat Controller", func() {

	router := rest.HeartbeatRotues()

	BeforeEach(func() {
		// load custom env variables
		os.Setenv("DATABASE_URL", "postgres://postgres:admin@localhost:5432/postgres")
		os.Setenv("APP_NAME", "test-app-name")
		os.Setenv("RELEASE_VERSION", "test-version")
		os.Setenv("RELEASE_SLUG", "test-slug")
		// create test config
		err := config.NewServiceConfig(true)
		if err != nil {
			log.Fatalf("error setting up test config: %v", err)
		}
	})

	AfterEach(func() {
		os.Clearenv()
	})

	Describe("calling heartbeat", func() {
		It("should equal all the env variables", func() {
			var respBody kit_models.Heartbeat
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			router.ServeHTTP(rr, req)
			// parse resp body
			if err := render.DecodeJSON(rr.Body, &respBody); err != nil {
				log.Fatalf("error decoding json: %v", err)
			}
			// validate
			Expect(rr.Code).To(Equal(200))
			Expect(respBody.AppName).To(Equal("test-app-name"))
			Expect(respBody.ReleaseVersion).To(Equal("test-version"))
			Expect(respBody.Slug).To(Equal("test-slug"))
			Expect(respBody.Message).To(Equal("togo write service"))
		})
	})
})
