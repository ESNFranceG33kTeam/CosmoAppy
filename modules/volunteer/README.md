---
title: Module Volunteer
---

# Table of Contents

- [Volunteers endpoint](#volunteers-endpoint)
- [Statistics endpoint](#statistics-endpoint)

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
- HrStatus          string  `json:"hr_status"`      - Hr status of the volunteer
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
        "hr_status": "volunteer",
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
        "hr_status": "volunteer",
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

# Statistics endpoint

**Object description**

- **Id**                    int             `json:"id"`                         - Id of the stat
- **ArchiveDate**           string          `json:"archive_date"`               - Date of the stat
- **NbActivePerUniversity** json.RawMessage `json:"nb_active_per_university"`   - Number of active volunteer by university
- **NbActivePerHrstatus**   json.RawMessage `json:"nb_active_per_hrstatus"`     - Number of active volunteer by hr status
- **NbNewVlt**              int             `json:"nb_new_vlt_this_month"`      - Number of new volunteer this month
- **NbAlumnus**             int             `json:"nb_alumnus"`                 - Number of alumnus
- **NbBureau**              int             `json:"nb_bureau"`                  - Number of volunteer in the bureau
- **CreatedAt**             time.Time       `json:"created_at"`                 - Created date of the stat
- **UpdatedAt**             time.Time       `json:"updated_at"`                 - Updated date of the stat

The parameters `id`, `created_at` and `updated_at` are calculated automatically.

## Get

- Get full list of stats :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/volunteers/stats/monthly" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of stats objects

- Get only stat of a certain date :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/volunteers/stats/monthly/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) stat object

## Post

- Generate stat of the last month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/volunteers/stats/monthly/create" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
