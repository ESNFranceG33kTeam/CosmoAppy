---
title: Module Report
---

# Table of Contents

- [Reports endpoint](#reports-endpoint)
- [Statistics endpoint](#statistics-endpoint)

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
- NbHours           float  `json:"nb_hours"`                - Number of hours of the event/planning/custom
- NbStaffsVlt          float  `json:"nb_staffs_vlt"`               - Number of volunteer staffs of the event/planning/custom
- NbStaffsEmp          float  `json:"nb_staffs_emp"`               - Number of employee staffs of the event/planning/custom
- NbStaffsScv          float  `json:"nb_staffs_scv"`               - Number of civic service staffs of the event/planning/custom
- TauxValorisationVlt  float  `json:"taux_valorisation_vlt"`       - Taux volunteer Valorisation of the event/planning/custom
- TauxValorisationEmp  float  `json:"taux_valorisation_emp"`       - Taux employee Valorisation of the event/planning/custom
- TauxValorisationScv  float  `json:"taux_valorisation_scv"`       - Taux civic service Valorisation of the event/planning/custom
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
        "nb_staffs_vlt": 3,
        "nb_staffs_emp": 0,
        "nb_staffs_scv": 0,
        "taux_valorisation_vlt": 18,
        "taux_valorisation_emp": 18,
        "taux_valorisation_scv": 18,
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
        "nb_staffs_vlt": 3,
        "nb_staffs_emp": 0,
        "nb_staffs_scv": 0,
        "taux_valorisation_vlt": 18,
        "taux_valorisation_emp": 18,
        "taux_valorisation_scv": 18,
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

# Statistics endpoint

**Object description**

- **Id**                int             `json:"id"`                     - Id of the stat
- **ArchiveDate**       string          `json:"archive_date"`           - Date of the stat
- **NbPerCodePublic**   json.RawMessage `json:"nb_per_codepublic"`      - Number of report by code public
- **NbPerCodeProject**  json.RawMessage `json:"nb_per_codeproject"`     - Number of report by code project
- **NbPerType**         json.RawMessage `json:"nb_per_type"`            - Number of report by type
- **NbTotal**           int             `json:"nb_total"`               - Number of report total
- **HoursVltCodes**     json.RawMessage `json:"hours_vlt_per_codes"`    - Number of hours for a volunteer by code
- **HoursEmpCodes**     json.RawMessage `json:"hours_emp_per_codes"`    - Number of hours for an employee by code
- **HoursScvCodes**     json.RawMessage `json:"hours_scv_per_codes"`    - Number of hours for a scv by code
- **ValoVltCodes**      json.RawMessage `json:"valo_vlt_per_codes"`     - Valorisation for a volunteer by code
- **ValoEmpCodes**      json.RawMessage `json:"valo_emp_per_codes"`     - Valorisation for an employee by code
- **ValoScvCodes**      json.RawMessage `json:"valo_scv_per_codes"`     - Valorisation for a scv by code
- **CreatedAt**         time.Time       `json:"created_at"`             - Created date of the stat
- **UpdatedAt**         time.Time       `json:"updated_at"`             - Updated date of the stat

The parameters `id`, `created_at` and `updated_at` are calculated automatically.

## Get

- Get full list of stats :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/reports/stats/monthly" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of stats objects

- Get only stat of a certain date :

```bash
curl -X GET \
    "https://${MYSERVER}/${SECURE}/reports/stats/monthly/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) stat object

## Post

- Generate stat of the last month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/reports/stats/monthly/create" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

- Generate stat of a certain past month :

```bash
curl -X POST \
    "https://${MYSERVER}/${SECURE}/reports/stats/monthly/force/2022-04" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output
