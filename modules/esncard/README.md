---
title: Module ESNcard
---

# ESNcards endpoint

## Requirements

This module is dependant of the `adherent` module and need it to works.

## Get

- Get full list of esncards :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/esncards" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of esncards objects

- Get only the esncard of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/esncards/id_adherent/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

- Get only the esncard with the specific esncard code :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/esncards/esncard/my-esncard-code" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

## Post

- To create a new esncard :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/esncards" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_adherent": 2,
        "esncard": "aVeryTooLongCode"
    }'
```

Output : (`json`) esncard object

## Delete

- To delete an esncard :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/esncards/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
