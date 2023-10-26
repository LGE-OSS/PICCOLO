LifecycleManagingRule
=====================

The rule to manage lifecycle of Application Pod for Vehicle.

+----------------------+------------+-------+----------------------+
| Attribute            | Value Type | List  | Kind                 |
+======================+============+=======+======================+
| name                 | string     | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| action               | string     | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| target               | VehPod     | true  | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| lifeTime             | string     | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| priority             | int        | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| condition            | object     | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| conditionName        | string     | false | LifecycleManagin     |
|                      |            |       | gRule.condition.attr |
+----------------------+------------+-------+----------------------+
| conditionTopic       | string     | false | LifecycleManagin     |
|                      |            |       | gRule.condition.attr |
+----------------------+------------+-------+----------------------+
| conditionValue       | string     | false | LifecycleManagin     |
|                      |            |       | gRule.condition.attr |
+----------------------+------------+-------+----------------------+
| conditionComparison  | string     | false | LifecycleManagin     |
|                      |            |       | gRule.condition.attr |
+----------------------+------------+-------+----------------------+
| conditions           | object     | false | Lifecy               |
|                      |            |       | cleManagingRule.attr |
+----------------------+------------+-------+----------------------+
| conditionUnits       | object     | true  | LifecycleManaging    |
|                      |            |       | Rule.conditions.attr |
+----------------------+------------+-------+----------------------+
| c                    | string     | false | LifecycleMa          |
| onditionsPreOperator |            |       | nagingRule.condition |
|                      |            |       | s.conditionUnit.attr |
+----------------------+------------+-------+----------------------+
| conditionsValue      | condition  | false | LifecycleMa          |
|                      |            |       | nagingRule.condition |
|                      |            |       | s.conditionUnit.attr |
+----------------------+------------+-------+----------------------+
| co                   | string     | false | LifecycleMa          |
| nditionsPostOperator |            |       | nagingRule.condition |
|                      |            |       | s.conditionUnit.attr |
+----------------------+------------+-------+----------------------+

Attribute Description
---------------------

.. _lifecyclemanagingrule-2:

*LifecycleManagingRule*
~~~~~~~~~~~~~~~~~~~~~~~

--------------

name
^^^^

-  .. rubric:: Value Type
      :name: value-type

   string

-  .. rubric:: List
      :name: list

   False

-  .. rubric:: Kind
      :name: kind

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option

   True

-  .. rubric:: Description
      :name: description

   The name of LifecycleManagingRule.

--------------

action
^^^^^^

-  .. rubric:: Value Type
      :name: value-type-2

   VehPod(VehDeployment)

-  .. rubric:: List
      :name: list-2

   True

-  .. rubric:: Kind
      :name: kind-2

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-2

   True

-  .. rubric:: Description
      :name: description-2

   The action what manager want to do.

   start, update, delete, etc...

--------------

target
^^^^^^

-  .. rubric:: Value Type
      :name: value-type-3

   string

-  .. rubric:: List
      :name: list-3

   True

-  .. rubric:: Kind
      :name: kind-3

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-3

   True

-  .. rubric:: Description
      :name: description-3

   Target VehPod(VehDeployment) spec.

--------------

lifeTime
^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-4

   string

-  .. rubric:: List
      :name: list-4

   False

-  .. rubric:: Kind
      :name: kind-4

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-4

   True

-  .. rubric:: Description
      :name: description-4

   Repeat or not.

   OneTime or sticky.

--------------

priority
^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-5

   int

-  .. rubric:: List
      :name: list-5

   False

-  .. rubric:: Kind
      :name: kind-5

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-5

   False(default is 3)

-  .. rubric:: Description
      :name: description-5

   Processing priority. 1 is highest, 5 is lowest

--------------

condition
^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-6

   object

-  .. rubric:: List
      :name: list-6

   False

-  .. rubric:: Kind
      :name: kind-6

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-6

   False

-  .. rubric:: Description
      :name: description-6

   Execute condition with single conditionTopic.

--------------

conditionName
^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-7

   string

-  .. rubric:: List
      :name: list-7

   False

-  .. rubric:: Kind
      :name: kind-7

   LifecycleManagingRule.condition.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-7

   True(when use condition object)

-  .. rubric:: Description
      :name: description-7

   Condition object name.

--------------

conditionTopic
^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-8

   string

-  .. rubric:: List
      :name: list-8

   False

-  .. rubric:: Kind
      :name: kind-8

   LifecycleManagingRule.condition.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-8

   True(when use condition object)

-  .. rubric:: Description
      :name: description-8

   Vss topic what manager want to set execute condition.

   Lifecycle manager subscribe this topic after rule set.

--------------

conditionValue
^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-9

   string

-  .. rubric:: List
      :name: list-9

   False

-  .. rubric:: Kind
      :name: kind-9

   LifecycleManagingRule.condition.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-9

   True(when use condition object)

-  .. rubric:: Description
      :name: description-9

   The value of vss topic what manager want to set execute condition.

--------------

conditionComparison
^^^^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-10

   string

-  .. rubric:: List
      :name: list-10

   False

-  .. rubric:: Kind
      :name: kind-10

   LifecycleManagingRule.condition.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-10

   True(when use condition object)

-  .. rubric:: Description
      :name: description-10

   The comparison with conditionValue.

   GreaterThan, GreaterThanOrEqualTo, Equal, smaller or Equal,
   SmallerThan.

   example)

   .. code:: 

      condition
        conditionName: Safe_Update
        conditionTopic: Gear_Status
        conditionValue: Park
        conditionComparison: Equal

--------------

conditions
^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-11

   string

-  .. rubric:: List
      :name: list-11

   False

-  .. rubric:: Kind
      :name: kind-11

   LifecycleManagingRule.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-11

   False

-  .. rubric:: Description
      :name: description-11

   The combination of condition.

--------------

conditionUnits
^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-12

   object

-  .. rubric:: List
      :name: list-12

   True

-  .. rubric:: Kind
      :name: kind-12

   LifecycleManagingRule.conditions.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-12

   True(When use conditions)

-  .. rubric:: Description
      :name: description-12

   The unit condition via condition object and operator for combine
   conditions.

--------------

conditionsPreOperator
^^^^^^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-13

   string

-  .. rubric:: List
      :name: list-13

   False

-  .. rubric:: Kind
      :name: kind-13

   LifecycleManagingRule.conditions.conditionUnit.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-13

   False.

-  .. rubric:: Description
      :name: description-13

   PreOperator for single condition object.

   Not or None.(! operator)

--------------

conditionsValue
^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-14

   condition

-  .. rubric:: List
      :name: list-14

   False

-  .. rubric:: Kind
      :name: kind-14

   LifecycleManagingRule.conditions.conditionUnit.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-14

   True(When use conditions)

-  .. rubric:: Description
      :name: description-14

   Condition object name.

   --------------

conditionsPostOperator
^^^^^^^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-15

   string

-  .. rubric:: List
      :name: list-15

   False

-  .. rubric:: Kind
      :name: kind-15

   LifecycleManagingRule.conditions.conditionUnit.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-15

   False

-  .. rubric:: Description
      :name: description-15

   PostOperator for single condition object.

   And or Or.(&& or \|\| operator)

   Xor and Xand is not supported.

   Last object of conditionUnit should be blank.

   example)

   .. code:: 

      conditions
        -conditionUnits:
          conditionsValue: Safe_Update
          conditionsPostOperator: And
        -conditionUnits:
          conditionsValue: ZeroSpeed

   .. code:: 

      conditions
        -conditionUnits:
          conditionsValue: HigherThan20KmPH
          conditionsPostOperator: And
        -conditionUnits:
          conditionsPreOperator: Not
          conditionsValue: HigherThan110KmPH
