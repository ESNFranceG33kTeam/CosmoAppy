# Volunteers endpoint

## Get

- Get full list of volunteers :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/volunteers" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of volunteers objects

- Get only the volunteer status of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/volunteers/id_adherent/1" \
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
        "id_adherent": 2,
        "actif": true,
        "bureau": false
    }'
```

Output : (`json`) volunteer object

## Put

- To update a volunteer :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/volunteers/id_adherent/2" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_adherent": 2,
        "actif": false,
        "bureau": false
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