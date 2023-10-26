Monitoring
==========

Vehicle Pod with Monitor structure of containers.

================ ========== ===== ===============
Attribute        Value Type List  Kind
================ ========== ===== ===============
primaryContainer object     true  Monitoring.attr
watcherContainer object     false Monitoring.attr
================ ========== ===== ===============

Attribute Description
---------------------

.. _monitoring-2:

*Monitoring*
~~~~~~~~~~~~

--------------

primaryContainer
^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type

   object

-  .. rubric:: List
      :name: list

   true

-  .. rubric:: Kind
      :name: kind

   Monitoring.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option

   true

-  .. rubric:: Description
      :name: description

   The container responsible for the main functionality of the
   application.

--------------

watcherContainer
^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-2

   object

-  .. rubric:: List
      :name: list-2

   false

-  .. rubric:: Kind
      :name: kind-2

   Monitoring.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-2

   true

-  .. rubric:: Description
      :name: description-2

   The container that monitors and logs the primary containers, and
   handles lifecycle and application status-related requests to the
   system.
