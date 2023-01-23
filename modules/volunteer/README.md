---
title: Module Volunteer
---

# Volunteers endpoint

**Object description**

- Id                int     `json:"id"`             - Id of the volunteer
- Firstname         string  `json:"firstname"`      - Firstname of the volunteer
- Lastname          string  `json:"lastname"`       - Lastname of the volunteer
- Email             string  `json:"email"`          - Email of the volunteer
- Discord           string  `json:"discord"`        - Discord pseudo
- Phone             string  `json:"phone"`          - Phone number
- University        string  `json:"university"`     - University name
- Postal address    string  `json:"postal_address"` - Postal address
- Actif             bool    `json:"actif"`          - Status of the volunteer
- Bureau            bool    `json:"bureau"`         - Does the volunteer has a bureau role
- Employee          bool    `json:"employee"`       - Does the volunteer is an employee
- Started date      string  `json:"started_date"`   - Date of volunteering started

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
        "discord": "Alola",
        "phone": "0123456789",
        "university": "Indigo",
        "postal_address": "01 bis house, 00001 Pallet Town, Kanto",
        "actif": true,
        "bureau": false,
        "employee": false,
        "started_date": "1997-04-01"
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
        "email": "ash.ketchum@indego.com",
        "discord": "Alola",
        "phone": "0123456789",
        "university": "Indigo",
        "postal_address": "01 bis house, 00001 Pallet Town, Kanto",
        "actif": false,
        "bureau": false,
        "employee": false,
        "started_date": "1997-04-01"
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
