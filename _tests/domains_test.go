package _tests_test

import (
    "fmt"
	. "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Domains", func() {
	Context("with a deployed app", func() {

		It("can add, list, and remove domains", func() {
            //setup
            appName := "app-test"
			execute("deis apps:create %s", appName)
            fmt.Println("hey man")

            //add domain
            addCmd := "deis domains:add %s --app=%s"
            domain := "domainsample"
            _, err := execute(addCmd, domain, appName)
            Expect(err).NotTo(HaveOccurred())

            //list domains
            listCmd := "deis domains:list --app=%s"
            _, err = execute(listCmd, appName)
            Expect(err).NotTo(HaveOccurred())
            // @todo: Expect(err).to include domain

			// curl app at both root and custom domain

            //remove domain
            removeCmd := "deis domains:remove %s --app=%s"
            _, err = execute(removeCmd, domain, appName)
            Expect(err).NotTo(HaveOccurred())
		})

		XIt("can add, list, and remove certs", func() {
			// "deis domains:add %s --app=%s", domain, app
			// "deis certs:list", app
			// "deis certs:add %s %s", certPath, keyPath
			// wait for 60 seconds until cert generation is done?
			// curl the custom SSL endpoint
			// "deis certs:remove %s", domain
			// "deis certs:list", app
			// curl the custom SSL endpoint, should fail
			// curl app at both root and custom domain, custom should fail
			// "deis domains:remove %s --app=%s", domain, app
		})
	})
})
