---
title: Module Volunteer
---

# Volunteers endpoint

**Object description**

- Id            int     `json:"id"`             - Id of the volunteer
- Firstname     string  `json:"firstname"`      - Firstname of the volunteer
- Lastname      string  `json:"lastname"`       - Lastname of the volunteer
- Email         string  `json:"email"`          - Email of the volunteer
- Actif         bool    `json:"actif"`          - Status of the volunteer
- Bureau        bool    `json:"bureau"`         - Does the volunteer has a bureau role

The parameters `id` is calculated automatically.

## Get

- Get full list of volunteers :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/volunteers" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of volunteers objects

- Get only the volunteer status :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/volunteers/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) volunteer object

## Post

- To create a new volunteer :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/volunteers" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "ash.ketchum@indego.com"
        "actif": true,
        "bureau": false
    }'
```

Output : (`json`) volunteer object

## Put

- To update a volunteer :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/volunteers/2" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "ash.ketchum@indego.com"
        "actif": false,
        "bureau": false
    }'
```

Output : (`json`) volunteer object

## Delete

- To delete an volunteer :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/volunteers/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
