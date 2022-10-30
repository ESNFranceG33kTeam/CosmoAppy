---
title: Modules
---

# Modules

All endpoints are defined here !

Each endpoint has its own documentation as `README.md`.

## New modules

If you want to create a new module, you only need to inspire yourself from already created module :
- `connector.go` __MANDATORY__
- `controller.go` __MANDATORY__
- `model.go` __MANDATORY__
- `router.go` __MANDATORY__
- `swagger.go` __MANDATORY__
- `README.md` __MANDATORY__
---
- `specimen.go` __OPTIONAL__
- `database.go` __OPTIONAL__

> ⚠ Do not forget to add your module to the Launcher at `modules/launcher/launcher.go`, it's him whose controls and manages all modules.
