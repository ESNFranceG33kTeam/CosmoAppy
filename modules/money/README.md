---
title: Module Money
---

# Moneys endpoint

**Object description**

- Id            int         `json:"id"`             - Id of the entry
- Label         string      `json:"label"`          - Label of the operation
- Price         float64     `json:"price"`          - Price of the operation
- PaymentDate   string      `json:"payment_date"`   - Payment date of the operation
- CreatedAt     time.Time   `json:"created_at"`     - Created date of the operation

The parameters `id` and `created_at` are calculated automatically.

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
