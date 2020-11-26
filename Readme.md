Prototype Current status:
========================
1) Initial design draft with two layer design(`cmd/` + `platform/`) that fits into "model" layer of MVC
2) Two layer design should be final because client must inevitably know which "kind" of authentication takes place



Design motivation:
=================

1) Client inevitably must know which "kind" of authentication takes place - So, this leads to two layer design without decoupling layer(middle)

2) Grouping types(`type`) by "what they do" but not by "who they are"



Pending items:
=============
0) Hande the return values from platform layer to public layer to client
1) Refactor cmd.Authenticate() api
2) Refactoring current single threaded version based on review
3) Converting it to multi-threaded version at platform layer


Remarks:
========

> Customised error handling or customised logging will not be taken care as part of this prototype
