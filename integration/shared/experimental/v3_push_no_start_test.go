package experimental

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("v3-push with --no-start", func() {
	var (
		orgName   string
		spaceName string
		appName   string
		userName  string
	)

	BeforeEach(func() {
		orgName = helpers.NewOrgName()
		spaceName = helpers.NewSpaceName()
		appName = helpers.PrefixedRandomName("app")
		userName, _ = helpers.GetCredentials()
		helpers.TurnOffExperimental()
	})

	AfterEach(func() {
		helpers.TurnOnExperimental()
	})

	When("the environment is set up correctly", func() {
		BeforeEach(func() {
			helpers.SetupCF(orgName, spaceName)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		When("the app exists and is already running", func() {
			var session *Session

			BeforeEach(func() {
				helpers.WithHelloWorldApp(func(appDir string) {
					Eventually(helpers.CustomCF(helpers.CFEnv{WorkingDirectory: appDir}, "v3-push", appName)).Should(Exit(0))
				})

				It("uploads the app, does not restage or restart it, and stops the initial app", func() {
					helpers.WithHelloWorldApp(func(appDir string) {
						session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: appDir}, "v3-push", appName, "--no-start")
						Eventually(session).Should(Say("Updating app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Say(""))
						Eventually(session).Should(Say("Uploading and creating bits package for app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Say(""))
						Eventually(session).Should(Say("Stopping app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Say(""))
						Consistently(session).ShouldNot(Say("Staging package for app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
						Consistently(session).ShouldNot(Say("Starting app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
						Eventually(session).Should(Exit(0))
					})

					session = helpers.CF("v3-app", appName)

					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("requested state:\\s+stopped"))

					Eventually(session).Should(Say("web:1/1"))
					Eventually(session).Should(Say(`state\s+since\s+cpu\s+memory\s+disk`))
					Eventually(session).Should(Say("#0\\s+running\\s+\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2} [AP]M"))

					Eventually(session).Should(Exit(0))
				})
			})
		})

		When("the app does not exist", func() {
			var session *Session

			It("uploads the app, does not restage or restart it", func() {
				helpers.WithHelloWorldApp(func(appDir string) {
					session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: appDir}, "v3-push", appName, "--no-start")
					Eventually(session).Should(Say("Creating app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Say(""))
					Eventually(session).Should(Say("Uploading and creating bits package for app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Say(""))
					Consistently(session).ShouldNot(Say("Staging package for app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))
					Consistently(session).ShouldNot(Say("Starting app %s in org %s / space %s as %s\\.\\.\\.", appName, orgName, spaceName, userName))

					Eventually(session).Should(Exit(0))
				})
			})
		})
	})
})
