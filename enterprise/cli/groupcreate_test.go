package cli_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DanielRondonGarcia/pretty"

	"github.com/DanielRondonGarcia/coder/v2/cli/clitest"
	"github.com/DanielRondonGarcia/coder/v2/cli/cliui"
	"github.com/DanielRondonGarcia/coder/v2/coderd/coderdtest"
	"github.com/DanielRondonGarcia/coder/v2/coderd/rbac"
	"github.com/DanielRondonGarcia/coder/v2/codersdk"
	"github.com/DanielRondonGarcia/coder/v2/enterprise/coderd/coderdenttest"
	"github.com/DanielRondonGarcia/coder/v2/enterprise/coderd/license"
	"github.com/DanielRondonGarcia/coder/v2/pty/ptytest"
)

func TestCreateGroup(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		client, admin := coderdenttest.New(t, &coderdenttest.Options{LicenseOptions: &coderdenttest.LicenseOptions{
			Features: license.Features{
				codersdk.FeatureTemplateRBAC: 1,
			},
		}})
		anotherClient, _ := coderdtest.CreateAnotherUser(t, client, admin.OrganizationID, rbac.RoleUserAdmin())

		var (
			groupName = "test"
			avatarURL = "https://example.com"
		)

		inv, conf := newCLI(t, "groups",
			"create", groupName,
			"--avatar-url", avatarURL,
		)

		pty := ptytest.New(t)
		inv.Stdout = pty.Output()
		clitest.SetupConfig(t, anotherClient, conf)

		err := inv.Run()
		require.NoError(t, err)

		pty.ExpectMatch(fmt.Sprintf("Successfully created group %s!", pretty.Sprint(cliui.DefaultStyles.Keyword, groupName)))
	})
}
