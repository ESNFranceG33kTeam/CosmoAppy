---
title: Modules
---

# Modules

All endpoints are defined here !

Each endpoint has its own documentation as `README.md`.

## New modules

If you want to create a new module, you only need to inspire yourself from already created module and follow those requirements :
- `connector.go` __MANDATORY__
- `controller.go` __MANDATORY__
- `model.go` __MANDATORY__
- `router.go` __MANDATORY__
- `swagger.go` __MANDATORY__
- `README.md` __MANDATORY__
---
- `specimen.go` __OPTIONAL__
- `database.go` __OPTIONAL__

> ⚠ Do not forget to add your module to the Launcher at `launcher/launcher.go`, it's him whose controls and manages all modules.

## List of existing modules

- [health](../README.md#healthcheck)
- [adherent](adherent/README.md)
- [esncard](esncard/README.md)
- [volunteer](volunteer/README.md)
- [event](event/README.md)
- [planning](planning/README.md)
- [money](money/README.md)
- [report](report/README.md)
- [stock](stock/README.md)
