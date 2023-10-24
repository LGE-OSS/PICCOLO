# LifecycleManagingRule

The rule to manage lifecycle of Application Pod for Vehicle.

| Attribute              | Value Type | List  | Kind                                                |
| ---------------------- | ---------- | ----- | --------------------------------------------------- |
| name                   | string     | false | LifecycleManagingRule.attr                          |
| action                 | string     | false | LifecycleManagingRule.attr                          |
| target                 | VehPod     | true  | LifecycleManagingRule.attr                          |
| lifeTime               | string     | false | LifecycleManagingRule.attr                          |
| priority               | int        | false | LifecycleManagingRule.attr                          |
| condition              | object     | false | LifecycleManagingRule.attr                          |
| conditionName          | string     | false | LifecycleManagingRule.condition.attr                |
| conditionTopic         | string     | false | LifecycleManagingRule.condition.attr                |
| conditionValue         | string     | false | LifecycleManagingRule.condition.attr                |
| conditionComparison    | string     | false | LifecycleManagingRule.condition.attr                |
| conditions             | object     | false | LifecycleManagingRule.attr                          |
| conditionUnits         | object     | true  | LifecycleManagingRule.conditions.attr               |
| conditionsPreOperator  | string     | false | LifecycleManagingRule.conditions.conditionUnit.attr |
| conditionsValue        | condition  | false | LifecycleManagingRule.conditions.conditionUnit.attr |
| conditionsPostOperator | string     | false | LifecycleManagingRule.conditions.conditionUnit.attr |

### Attribute Description

#### *LifecycleManagingRule*

------

##### name

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  True

- ###### Description

  The name of LifecycleManagingRule.



------

##### action

- ######  Value Type

  VehPod(VehDeployment)

- ###### List

  True

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  True

- ###### Description

  The action what manager want to do.
  
  start, update, delete, etc...

------
##### target

- ######  Value Type

  string

- ###### List

  True

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  True

- ###### Description

  Target VehPod(VehDeployment) spec.

------
##### lifeTime

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  True

- ###### Description

  Repeat or not.
  
  OneTime or sticky.

------

##### priority

- ######  Value Type

  int

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  False(default is 3)

- ###### Description

  Processing priority. 1 is highest, 5 is lowest

------

##### condition

- ######  Value Type

  object

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  False

- ###### Description

  Execute condition with single conditionTopic.

------

##### conditionName

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.condition.attr

- ###### Mandatory option

  True(when use condition object)

- ###### Description

  Condition object name.

------

##### conditionTopic

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.condition.attr

- ###### Mandatory option

  True(when use condition object)

- ###### Description

  Vss topic what manager want to set execute condition.
  
  Lifecycle manager subscribe this topic after rule set.

------

##### conditionValue

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.condition.attr

- ###### Mandatory option

  True(when use condition object)

- ###### Description

  The value of vss topic what manager want to set execute condition.

------

##### conditionComparison

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.condition.attr

- ###### Mandatory option

  True(when use condition object)

- ###### Description

  The comparison with conditionValue.

  GreaterThan, GreaterThanOrEqualTo, Equal, smaller or Equal, SmallerThan.

  

  example)

  ```
  condition
    conditionName: Safe_Update
    conditionTopic: Gear_Status
    conditionValue: Park
    conditionComparison: Equal
  ```
------
##### conditions

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.attr

- ###### Mandatory option

  False

- ###### Description

  The combination of condition.

------

##### conditionUnits

- ######  Value Type

  object

- ###### List

  True

- ###### Kind

  LifecycleManagingRule.conditions.attr

- ###### Mandatory option

  True(When use conditions)

- ###### Description

  The unit condition via condition object and operator for combine conditions.

------

##### conditionsPreOperator

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.conditions.conditionUnit.attr

- ###### Mandatory option

  False.

- ###### Description

  PreOperator for single condition object.

  Not or None.(! operator)

------

##### conditionsValue

- ######  Value Type

  condition

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.conditions.conditionUnit.attr

- ###### Mandatory option

  True(When use conditions)

- ###### Description

  Condition object name.

  ------

##### conditionsPostOperator

- ######  Value Type

  string

- ###### List

  False

- ###### Kind

  LifecycleManagingRule.conditions.conditionUnit.attr

- ###### Mandatory option

  False

- ###### Description

  PostOperator for single condition object.
  
  And or Or.(&& or || operator)
  
  Xor and Xand is not supported.
  
  Last object of conditionUnit should be blank.

  

  example)
  
  ```
  conditions
    -conditionUnits:
      conditionsValue: Safe_Update
      conditionsPostOperator: And
    -conditionUnits:
      conditionsValue: ZeroSpeed
  ```

  ```
  conditions
    -conditionUnits:
      conditionsValue: HigherThan20KmPH
      conditionsPostOperator: And
    -conditionUnits:
      conditionsPreOperator: Not
      conditionsValue: HigherThan110KmPH
  ```