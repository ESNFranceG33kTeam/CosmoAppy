---
title: Module Money
---

# Moneys endpoint

## Get

- Get full list of money operations :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/moneys" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of moneys objects

- Get only the money operations of a specific label :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/moneys/label/Event" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of moneys object

## Post

- To create a new money operation :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/moneys" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "label": Event,
        "price": 5
    }'
```

Output : (`json`) money object
