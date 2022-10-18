import Table from "@material-ui/core/Table"
import TableBody from "@material-ui/core/TableBody"
import TableCell from "@material-ui/core/TableCell"
import TableContainer from "@material-ui/core/TableContainer"
import TableHead from "@material-ui/core/TableHead"
import TableRow from "@material-ui/core/TableRow"
import {
  Badges,
  DisabledBadge,
  EnabledBadge,
} from "components/DeploySettingsLayout/Badges"
import { useDeploySettings } from "components/DeploySettingsLayout/DeploySettingsLayout"
import { Header } from "components/DeploySettingsLayout/Header"
import {
  OptionDescription,
  OptionName,
  OptionValue,
} from "components/DeploySettingsLayout/Option"
import { Stack } from "components/Stack/Stack"
import React from "react"

const AuthSettingsPage: React.FC = () => {
  const { deploymentConfig } = useDeploySettings()

  return (
    <>
      <Stack direction="column" spacing={6}>
        <div>
          <Header
            title="Login with OpenID Connect"
            secondary
            description="Set up authentication to login with OpenID Connect."
            docsHref="https://coder.com/docs/coder-oss/latest/admin/auth#openid-connect-with-google"
          />

          <Badges>
            {deploymentConfig.oidc_client_id.value ? (
              <EnabledBadge />
            ) : (
              <DisabledBadge />
            )}
          </Badges>

          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell width="50%">Option</TableCell>
                  <TableCell width="50%">Value</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oidc_client_id.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_client_id.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oidc_client_id.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oidc_client_secret.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_client_secret.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oidc_client_secret.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oidc_allow_signups.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_allow_signups.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oidc_allow_signups.value.toString()}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oidc_email_domain.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_email_domain.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oidc_email_domain.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oidc_issuer_url.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_issuer_url.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oidc_issuer_url.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>{deploymentConfig.oidc_scopes.name}</OptionName>
                    <OptionDescription>
                      {deploymentConfig.oidc_scopes.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      <ul>
                        {deploymentConfig.oidc_scopes.value.map((scope) => (
                          <li key={scope}>{scope}</li>
                        ))}
                      </ul>
                    </OptionValue>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
        </div>

        <div>
          <Header
            title="Login with GitHub"
            secondary
            description="Set up authentication to login with GitHub."
            docsHref="https://coder.com/docs/coder-oss/latest/admin/auth#github"
          />

          <Badges>
            {deploymentConfig.oauth2_github_client_id.value ? (
              <EnabledBadge />
            ) : (
              <DisabledBadge />
            )}
          </Badges>

          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell width="50%">Option</TableCell>
                  <TableCell width="50%">Value</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_client_id.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oauth2_github_client_id.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oauth2_github_client_id.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_client_secret.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oauth2_github_client_secret.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oauth2_github_client_secret.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_allow_signups.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oauth2_github_allow_signups.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oauth2_github_allow_signups.value.toString()}
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_allowed_organizations.name}
                    </OptionName>
                    <OptionDescription>
                      {
                        deploymentConfig.oauth2_github_allowed_organizations
                          .description
                      }
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      <ul>
                        {deploymentConfig.oauth2_github_allowed_organizations.value.map(
                          (org) => (
                            <li key={org}>{org}</li>
                          ),
                        )}
                      </ul>
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_allowed_teams.name}
                    </OptionName>
                    <OptionDescription>
                      {deploymentConfig.oauth2_github_allowed_teams.description}
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      <ul>
                        {deploymentConfig.oauth2_github_allowed_teams.value.map(
                          (team) => (
                            <li key={team}>{team}</li>
                          ),
                        )}
                      </ul>
                    </OptionValue>
                  </TableCell>
                </TableRow>

                <TableRow>
                  <TableCell>
                    <OptionName>
                      {deploymentConfig.oauth2_github_enterprise_base_url.name}
                    </OptionName>
                    <OptionDescription>
                      {
                        deploymentConfig.oauth2_github_enterprise_base_url
                          .description
                      }
                    </OptionDescription>
                  </TableCell>

                  <TableCell>
                    <OptionValue>
                      {deploymentConfig.oauth2_github_enterprise_base_url.value}
                    </OptionValue>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
        </div>
      </Stack>
    </>
  )
}

export default AuthSettingsPage
