# Adherents endpoint

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
