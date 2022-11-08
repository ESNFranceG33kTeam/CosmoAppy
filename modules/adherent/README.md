---
title: Module Adherent
---

# Adherents endpoint

**Object description**

- **Id**            int         `json:"id"`             - Id of the adherent
- **Firstname**     string      `json:"firstname"`      - Firstname of the adherent
- **Lastname**      string      `json:"lastname"`       - Lastname of the adherent
- **Email**         string      `json:"email"`          - Email of the adherent
- **Dateofbirth**   string      `json:"dateofbirth"`    - Date of birth of the adherent
- **Student**       bool        `json:"student"`        - Student status of the adherent
- **University**    string      `json:"university"`     - University of the adherent
- **Homeland**      string      `json:"homeland"`       - Homeland of the adherent
- **Speakabout**    string      `json:"speakabout"`     - How the adherent learn about the association
- **Newsletter**    bool        `json:"newsletter"`     - Newsletter status of the adherent
- **CreatedAt**     time.Time   `json:"created_at"`     - Created date of the adherent
- **UpdatedAt**     time.Time   `json:"updated_at"`     - Updated date of the adherent

The parameters `id`, `created_at` and `updated_at` are calculated automatically.


## Get

- Get full list of adherents :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/adherents" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of adherents objects

- Get only one adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/adherents/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) adherent object

## Post

- Create a new adherent :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/adherents" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "dresseur@indigo.com",
        "dateofbirth": "1987-05-22",
        "student": true,
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false
    }'
```

Output : (`json`) adherent object

## Put

- To update an adherent :

```bash
curl -X PUT "https://${MYSERVER}/${SECURE}/adherents/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "dresseur@indigo.com",
        "dateofbirth": "1987-05-22",
        "student": false,
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false
    }'
```

Output : (`json`) adherent object

## Delete

- To delete an adherent :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/adherents/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
