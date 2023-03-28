---
title: Module Adherent
---

# Table of Contents

- [Adherents endpoint](#adherents-endpoint)
- [Statistics endpoint](#statistics-endpoint)

# Adherents endpoint

**Object description**

- **Id**            int         `json:"id"`             - Id of the adherent
- **Firstname**     string      `json:"firstname"`      - Firstname of the adherent
- **Lastname**      string      `json:"lastname"`       - Lastname of the adherent
- **Email**         string      `json:"email"`          - Email of the adherent
- **Dateofbirth**   string      `json:"dateofbirth"`    - Date of birth of the adherent
- **Situation**     string      `json:"situation"`      - Situation status of the adherent
- **University**    string      `json:"university"`     - University of the adherent
- **Homeland**      string      `json:"homeland"`       - Homeland of the adherent
- **Speakabout**    string      `json:"speakabout"`     - How the adherent learn about the association
- **Newsletter**    bool        `json:"newsletter"`     - Newsletter status of the adherent
- **AdhesionDate**  string      `jsjon:"adhesion_date"` - Date of the adhesion
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
        "situation": "student",
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false,
        "adhesion_date": "1990-05-22"
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
        "situation": "worker",
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false,
        "adhesion_date": "1990-05-22"
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

# Statistics endpoint

**Object description**

- **Id**                int             `json:"id"`                 - Id of the stat
- **ArchiveDate**       string          `json:"archive_date"`       - Date of the stat
- **NbPerCountry**      json.RawMessage `json:"nb_per_country"`     - Number of adherent by country
- **NbPerUniversity**   json.RawMessage `json:"nb_per_university"`  - Number of adherent by university
- **NbPerSituation**    json.RawMessage `json:"nb_per_situation"`   - Number of adherent by situation
- **AboutusPerType**    json.RawMessage `json:"aboutus_per_type"`   - Number of how adherent heared about us per type
- **NbTotal**           int             `json:"nb_total"`           - Number of adherent total
- **CreatedAt**         time.Time       `json:"created_at"`         - Created date of the stat
- **UpdatedAt**         time.Time       `json:"updated_at"`         - Updated date of the stat

The parameters `id`, `created_at` and `updated_at` are calculated automatically.

## Get

- Get full list of stats :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/adherents/stats/monthly" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of stats objects

- Get only stat of a certain date :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/adherents/stats/monthly/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) stat object

## Post

- Generate stat of the last month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/adherents/stats/monthly/create" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

- Generate stat of a certain past month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/adherents/stats/monthly/force/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
