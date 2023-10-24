# Monitoring

Vehicle Pod with Monitor structure of containers.

| Attribute        | Value Type | List  | Kind            |
| ---------------- | ---------- | ----- | --------------- |
| primaryContainer | object     | true  | Monitoring.attr |
| watcherContainer | object     | false | Monitoring.attr |

### Attribute Description

#### *Monitoring*

------

##### primaryContainer

- ######  Value Type

  object

- ###### List

  true

- ###### Kind

  Monitoring.attr

- ###### Mandatory option

  true

- ###### Description

  The container responsible for the main functionality of the application.

------

##### watcherContainer

- ######  Value Type

  object

- ###### List

  false

- ###### Kind

  Monitoring.attr

- ###### Mandatory option

  true

- ###### Description

  The container that monitors and logs the primary containers, and handles lifecycle and application status-related requests to the system.

