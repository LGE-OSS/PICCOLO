# DeploymentPolicy

Set the structure of containers included on Vehicle Pod.

| Attribute         | Value Type | List  | Kind                  |
| ----------------- | ---------- | ----- | --------------------- |
| monitoring        | string     | false | DeploymentPolicyValue |
| activeRedundant   | string     | false | DeploymentPolicyValue |
| passiveRedundant  | string     | false | DeploymentPolicyValue |
| nModularRedundant | string     | false | DeploymentPolicyValue |

### Attribute Description

#### *DeploymentPolicy*

------

##### monitoring 

- ###### Value Type

  string

- ###### List

  false

- ###### Kind

  DeploymentPolicyValue

- ###### Mandatory option

  none

- ###### Description

  Monitoring type.

------

##### activeRedundant

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  DeploymentPolicyValue

- ###### Mandatory option

  none

- ###### Description

  Active-Redundant type.

------

#####  passiveRedundant

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  DeploymentPolicyValue

- ###### Mandatory option

  none

- ###### Description

  Passive-Redundant type.

------

##### nModularRedundant

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  DeploymentPolicyValue

- ###### Mandatory option

  none

- ###### Description

  N-Modular-Redundant type.
