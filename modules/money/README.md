---
title: Module Money
---

# Table of Contents

- [Moneys endpoint](#moneys-endpoint)
- [Statistics endpoint](#statistics-endpoint)

# Moneys endpoint

**Object description**

- Id            int         `json:"id"`             - Id of the entry
- Label         string      `json:"label"`          - Label of the operation
- Price         float64     `json:"price"`          - Price of the operation
- PaymentType   string      `json:"payment_type"`   - Payment type of the operation
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
        "price": 5,
        "payment_type": "cash",
        "payment_date": "2023-01-23"
    }'
```

Output : (`json`) money object

# Statistics endpoint

**Object description**

- **Id**                int             `json:"id"`                 - Id of the stat
- **ArchiveDate**       string          `json:"archive_date"`       - Date of the stat
- **NbPerLabel**        json.RawMessage `json:"nb_per_label"`       - Number of each label
- **AvgPrices**         json.RawMessage `json:"avg_prices"`         - Average prices per type of label
- **MinPrices**         json.RawMessage `json:"min_prices"`         - Min prices per type of label
- **MaxPrices**         json.RawMessage `json:"max_prices"`         - Max prices per type of label
- **NbPaymentTypes**    json.RawMessage `json:"nb_payments_type"`   - Number of Payments type
- **CreatedAt**         time.Time       `json:"created_at"`         - Created date of the stat
- **UpdatedAt**         time.Time       `json:"updated_at"`         - Updated date of the stat

The parameters `id`, `created_at` and `updated_at` are calculated automatically.

## Get

- Get full list of stats :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/moneys/stats/monthly" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of stats objects

- Get only stat of a certain date :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/moneys/stats/monthly/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) stat object

## Post

- Generate stat of the last month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/moneys/stats/monthly/create" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

- Generate stat of a certain past month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/moneys/stats/monthly/force/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
