---
title: Module Event / Event Attendees
---

# Table of Contents

- [Events endpoint](#events-endpoint)
- [Event Attendees endpoint](#event-attendees-endpoint)

# Events endpoint

**Object description**

- Id        int     `json:"id"`
- Name      string  `json:"name"`
- Date      string  `json:"date"`
- Location  string  `json:"location"`
- Type      string  `json:"type"`
- Price     float64 `json:"price"`
- Url       string  `json:"url_facebook"`
- Actif     bool    `json:"actif"`

The parameter `id` is calculated automatically.

## Get

- Get full list of events :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/events" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of events objects

- Get only one event :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/events/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) events object

## Post

- Create a new event :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/events" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Conseil des 4",
        "date": "2023-05-23",
        "location": "Plateau indigo, Johto",
        "type": "sport",
        "price": 0,
        "url_facebook": "https://facebook.com/ligue",
        "actif": true
    }'
```

Output : (`json`) event object

## Put

- To update an event :

```bash
curl -X PUT "https://${MYSERVER}/${SECURE}/events/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Conseil des 4",
        "date": "2023-05-23",
        "location": "Plateau indigo, Kanto",
        "type": "sport",
        "price": 0,
        "url_facebook": "https://facebook.com/ligue",
        "actif": true
    }'
```

Output : (`json`) event object

## Delete

- To delete an event :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/events/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

# Event Attendees endpoint

**Object description**

- Id            int     `json:"id"`
- Id_event      int     `json:"id_event"`
- Id_adherent   int     `json:"id_adherent"`
- Staff         bool    `json:"staff"`

The parameter `id` is calculated automatically.

## Requirements

This part of the module is dependant of the `adherent` module and need it to works.

## Get

- Get full list of attendees :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/event_attendees" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees objects

- Get only the attendees of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/event_attendees/id_adherent/2" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees object

- Get only the attendees of a specific event :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/event_attendees/id_event/1" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of attendees object

## Post

- To create a new attendee :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/event_attendees" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_event": 2,
        "id_adherent": 1,
        "staff": false
    }'
```

Output : (`json`) attendee object

## Put

- To update an attendee :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/event_attendees/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_event": 2,
        "id_adherent": 1,
        "staff": true
    }'
```

Output : (`json`) attendee object

## Delete

- To delete an attendee :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/event_attendees/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

## Data Dependances

In case of suppressing the a `planning` or an `adherent` on who has a link an `attendee`, **this `attendee` will be delete as well**.
