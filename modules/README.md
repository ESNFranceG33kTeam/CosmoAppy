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
- `vars.go` __MANDATORY__
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

## Router access

2 subrouters are availables :
- 1 subrouter for public access who can be accessible by anyone, example :

```bash
MYSERVER='127.0.0.1:8080'

curl -X GET \
    "https://${MYSERVER}/health`" \
    -H "accept: application/json"
```

- 1 subrouter `SECURE` for private access requiring a token and can be accessible from `auth` after the server name, example :

```bash
MYSERVER='127.0.0.1:8080'
SECURE='auth'
MYTOKEN='MyVeryLongToken'

curl -X GET \
    "https://${MYSERVER}/${SECURE}/status`" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```