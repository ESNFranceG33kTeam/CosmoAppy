---
title: Module Planning / Planning Attendees
---

# Table of Contents

- [Plannings endpoint](#plannings-endpoint)
- [Planning Attendees endpoint](#planning-attendees-endpoint)

# Plannings endpoint

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
        "location": "Labo Prof Chen",
        "date_begins": "2023-05-23",
        "date_end": "2023-05-23",
        "hour_begins": "9-00",
        "hour_end": "18-00"
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
        "location": "Labo Prof Chen",
        "date_begins": "2023-05-23",
        "date_end": "2023-05-23",
        "hour_begins": "9-00",
        "hour_end": "12-00"
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

## Requirements

This part of the module is dependant of the `adherent` module and need it to works.

## Get

- Get full list of attendees :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/planning_attendees" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees objects

- Get only the attendees of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/planning_attendees/id_adherent/2" \
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
        "id_adherent": 1,
        "date": "2023-05-23",
        "hour_begins": "10-00",
        "hour_end": "12-00"
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
        "id_adherent": 1,
        "date": "2023-05-23",
        "hour_begins": "12-00",
        "hour_end": "14-00"
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

In case of suppressing the a `planning` or an `adherent` on who has a link an `attendee`, **this `attendee` will be delete as well**.