---
title: Module Event / Event Attendees
---

# Table of Contents

- [Events endpoint](#events-endpoint)
- [Event Attendees endpoint](#event-attendees-endpoint)

# Events endpoint

**Object description**

- Id            int     `json:"id"`             - Id of the event
- Name          string  `json:"name"`           - Name of the event
- Date          string  `json:"date"`           - Date of the event
- Location      string  `json:"location"`       - Location of the event
- NbSpotsMax    int     `json:"nb_spots_max"`   - Number Spots Max of the event
- NbSpotsTaken  int     `json:"nb_spots_taken"` - Number Spots Taken of the event
- Type          string  `json:"type"`           - Type of the event
- Price         float64 `json:"price"`          - Price of the event
- Url           string  `json:"url_facebook"`   - Url of the event
- Actif         bool    `json:"actif"`          - Status of the event

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
        "nb_spots_max": 30,
        "nb_spots_taken": 30,
        "type": "sport",
        "price": 0,
        "url_facebook": "https://facebook.com/ligue",
        "actif": true
    }'
```

Output : (`json`) event object

## Put

- To take a spot :

```bash
curl -X PUT "https://${MYSERVER}/${SECURE}/events/take_spot/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) event object

- To update an event :

```bash
curl -X PUT "https://${MYSERVER}/${SECURE}/events/update/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Conseil des 4",
        "date": "2023-05-23",
        "location": "Plateau indigo, Kanto",
        "nb_spots_max": 30,
        "nb_spots_taken": 30,
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

- Id            int     `json:"id"`             - Id of the attendee
- Id_event      int     `json:"id_event"`       - Id of the event
- Id_adherent   int     `json:"id_adherent"`    - Id of the adherent
- Staff         bool    `json:"staff"`          - Status of the event

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
