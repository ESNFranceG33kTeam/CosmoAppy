---
title: Module Report
---

# Reports endpoint

**Object description**

- Id                int     `json:"id"`                 - Id of the report
- Name              string  `json:"name"`               - Name of the report
- Date              string  `json:"date"`               - Date of the event/planning/custom
- Comment           string  `json:"comment"`            - Comment of the report
- NbHours           string  `json:"nb_hours"`           - Number of hours of the event/planning/custom
- NbStaffs          string  `json:"nb_staffs"`          - Number of staffs of the event/planning/custom
- TauxValorisation  string  `json:"taux_valorisation"`  - Taux Valorisation of the event/planning/custom

The parameters `id` is calculated automatically.

## Get

- Get full list of reports :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/reports" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of reports objects

- Get only the report status :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/reports/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) report object

## Post

- To create a new report :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/reports" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Report NEM",
        "date": "2021-04-14",
        "comment": "It was amazing !"
        "nb_hours": 12,
        "nb_staffs": 10,
        "taux_valorisation": 18
    }'
```

Output : (`json`) report object

## Put

- To update a report :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/reports/2" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Report NEM",
        "date": "2021-04-14",
        "comment": "It was amaziiiiing !"
        "nb_hours": 12,
        "nb_staffs": 10,
        "taux_valorisation": 14
    }'
```

Output : (`json`) report object

## Delete

- To delete an report :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/reports/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
