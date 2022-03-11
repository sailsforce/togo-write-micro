package tests

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sailsforce/togo-write-micro/internal"
	"github.com/sailsforce/togo-write-micro/internal/config"
)

var _ = Describe("Read Micro Router", func() {

	rotuer := internal.Routes()

	BeforeEach(func() {
		if err := config.NewServiceConfig(false); err != nil {
			log.Fatalf("error setting up test config: %v", err)
		}
	})

	AfterEach(func() {
		os.Clearenv()
	})

	Describe("creating routes", func() {
		Context("calling routes()", func() {
			It("should return chi.Mux", func() {
				// validate
				Expect(rotuer).To(Not(BeNil()))
			})
		})
	})

})
