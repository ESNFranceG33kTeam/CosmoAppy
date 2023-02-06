---
title: Module Planning / Planning Attendees
---

# Table of Contents

- [Plannings endpoint](#plannings-endpoint)
- [Planning Attendees endpoint](#planning-attendees-endpoint)

# Plannings endpoint

**Object description**

- Id            int     `json:"id"`             - Id of the planning
- Name          string  `json:"name"`           - Name of the planning
- Type          string  `json:"type"`           - Type of the planning
- Location      string  `json:"location"`       - Location of the planning
- Date_begins   string  `json:"date_begins"`    - Start date of the planning
- Date_end      string  `json:"date_end"`       - End date of the planning
- Hour_begins   string  `json:"hour_begins"`    - Start hour of the planning
- Hour_end      string  `json:"hour_end"`       - End hour of the planning

The parameters `id` is calculated automatically.

## Get

- Get full list of plannings :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/plannings" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of plannings objects

- Get only one planning :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/plannings/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) plannings object

## Post

- Create a new planning :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/plannings" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Permanence",
        "type": "Permanence",
        "location": "Labo Prof Chen",
        "date_begins": "2023-05-23",
        "date_end": "2023-05-23",
        "hour_begins": "9:00:00",
        "hour_end": "18:00:00"
    }'
```

Output : (`json`) planning object

## Put

- To update an planning :

```bash
curl -X PUT "https://${MYSERVER}/${SECURE}/plannings/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Permanence",
        "type": "Permanence",
        "location": "Labo Prof Chen",
        "date_begins": "2023-05-23",
        "date_end": "2023-05-23",
        "hour_begins": "9:00:00",
        "hour_end": "12:00:00"
    }'
```

Output : (`json`) planning object

## Delete

- To delete an planning :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/plannings/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

# Planning Attendees endpoint

**Object description**

- Id            int     `json:"id"`             - Id of the attendee
- Id_planning   int     `json:"id_planning"`    - Id of the planning
- Id_volunteer  int     `json:"id_volunteer"`   - Id of the volunteer
- Job           string  `json:"job"`            - Job of the shift
- Date          string  `json:"date"`           - Date of the shift
- Hour_begins   string  `json:"hour_begins"`    - Start hour of the shift
- Hour_end      string  `json:"hour_end"`       - End hour of the shift

The parameters `id` is calculated automatically.

## Requirements

This part of the module is dependant of the `volunteer` module and need it to works.

## Get

- Get full list of attendees :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/planning_attendees" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees objects

- Get only the attendees of a specific volunteer :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/planning_attendees/id_volunteer/2" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees object

- Get only the attendees of a specific planning :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/planning_attendees/id_planning/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees object

## Post

- To create a new attendee :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/planning_attendees" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_planning": 2,
        "id_volunteer": 1,
        "job": "Staff",
        "date": "2023-05-23",
        "hour_begins": "10:00:00",
        "hour_end": "12:00:00"
    }'
```

Output : (`json`) attendee object

## Put

- To update an attendee :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/planning_attendees/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_planning": 2,
        "id_volunteer": 1,
        "job": "Staff",
        "date": "2023-05-23",
        "hour_begins": "12:00:00",
        "hour_end": "14:00:00"
    }'
```

Output : (`json`) attendee object

## Delete

- To delete an attendee :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/planning_attendees/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

## Data Dependances

In case of suppressing the a `planning` or an `volunteer` on who has a link an `attendee`, **this `attendee` will be delete as well**.
