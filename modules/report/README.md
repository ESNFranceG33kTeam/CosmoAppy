---
title: Module Report
---

# Reports endpoint

**Object description**

- Id                int     `json:"id"`                     - Id of the report
- Type              string  `json:"type"`                   - Type of the report
- RefExt            int     `json:"ref_ext"`                - External reference of the report
- Name              string  `json:"name"`                   - Name of the report
- Date              string  `json:"date"`                   - Date of the event/planning/custom
- Comment           string  `json:"comment"`                - Comment of the report
- NbReelAtt         int     `json:"nb_reel_attendees"`      - Number reel of attendees of the event/planning/custom 
- NbSubsAtt         int     `json:"nb_subscribe_attendees"` - Number subscribes of attendees of the event/planning/custom
- StaffsList        string  `json:"staffs_list"`            - List of staffs of the event/planning/custom
- NbHoursPrepa      float   `json:"nb_hours_prepa"`         - Number of hours of preparation
- NbHours           string  `json:"nb_hours"`               - Number of hours of the event/planning/custom
- NbStaffs          string  `json:"nb_staffs"`              - Number of staffs of the event/planning/custom
- TauxValorisation  string  `json:"taux_valorisation"`      - Taux Valorisation of the event/planning/custom
- CodePublic        string  `json:"code_public"`            - Public code of the report
- CodeProject       string  `json:"code_project"`           - Project code of the report

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
        "type": "event",
        "ref_ext": 1,
        "name": "Report NEM",
        "date": "2021-04-14",
        "comment": "It was amazing !",
        "nb_reel_attendees": 20,
        "nb_subscribe_attendees": 20,
        "staffs_list": "John Smith, Toto titi, Ash Ketchum",
        "nb_hours_prepa": 4,
        "nb_hours": 12,
        "nb_staffs": 3,
        "taux_valorisation": 18,
        "code_public": "ALL",
        "code_project": "PKM"
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
        "type": "event",
        "ref_ext": 1,
        "name": "Report NEM",
        "date": "2021-04-14",
        "comment": "It was amaziiiiing !",
        "nb_reel_attendees": 20,
        "nb_subscribe_attendees": 20,
        "staffs_list": "John Smith, Toto titi, Ash Ketchum",
        "nb_hours_prepa": 4,
        "nb_hours": 12,
        "nb_staffs": 10,
        "taux_valorisation": 14,
        "code_public": "ALL",
        "code_project": "PKM"
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
