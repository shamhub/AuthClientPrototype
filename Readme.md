Prototype Current status:
========================
1) Initial design draft with two layer(`api/` & `../internal/platform/`) design that fits into "model" component of MVC
2) Two layer design because client must inevitably know which "kind" of authentication takes place



Design motivation:
=================

1) Client inevitably must know which "kind" of authentication takes place - So, this leads to two layer design without decoupling layer(middle layer)

2) Grouping types(`type`) by "what they do" but not by "who they are"


Pending items:
=============
Handling response  from `platform` layer -> `api` layer -> client code


Build instructions:
====================

    `$ make install`

    `$ authclient`

