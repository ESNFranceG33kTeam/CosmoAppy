---
title: Module coCAS
---

# Module coCAS

This module aim for the CAS connexion to ESN Galaxy.

This allow to retrieve roles and give access through galaxy.

**Object description**

- **Username**      string      `json:"username"`       - Galaxy username of the user
- **First**         string      `json:"firstname"`      - Firstname of the user
- **Last**          string      `json:"lastname"`       - Lastname of the user
- **Email**         string      `json:"email"`          - Email of the user
- **Country**       string      `json:"country"`        - ESN Country of the user
- **Nationality**   string      `json:"nationality"`    - Nationality of the user
- **PictureURL**    string      `json:"picture_url"`    - Url profile picture of the user
- **Birthdate**     string      `json:"dateofbirth"`    - Date of birth of the user
- **Gender**        string      `json:"gender"`         - Gender of the user
- **Phone**         string      `json:"phone"`          - Phone number of the user
- **Sc**            string      `json:"section_code"`   - Code of the primary section
- **Section**       string      `json:"association"`    - Primary section
- **Roles**         []string    `json:"roles"`          - Roles match with primary country and section
- **ExtendedRoles** []string    `json:"extended_roles"` - Extended roles match with specific section code
- **Admin**         bool        `json:"admin"`          - Overlay for grant admin privileges

## Integration

In order to use this module, you need to first be connect with the galaxy to retrived information from those endpoints.

In result you will get a json with every information of the user, you can filter any role or section you want in order to connect him. 

But please keep in mind that if the field `admin` is at `true` you should keep him as administrator and give him full permissions. That in order to the g33kteam be able to give you full support.

## Endpoints

This API is callable without any authentication with `GET` call :

- Login

The default behiaviour is login, so you don't need to specify any endpoint :

```bash
curl -X GET "https://${MYSERVER}" \
    -H "accept: application/json"
```

Output : (`json`) profile object

- Logout

To logout, just call the endpoint `/logout` :

```bash
curl -X GET "https://${MYSERVER}/logout" \
    -H "accept: application/json"
```

Output : no output

## Permissions

By default any G33kTeam member and people with IT roles on the Galaxy are grant admin permission [see here](https://github.com/ESNFranceG33kTeam/.github/blob/main/members/galaxy.yaml) for further information and details.

You can also add some extra peoples for admin grant by filling the variable `extra_admin` :

```
# Galaxy username for Extra Admin
extra_admin:
  - "P1chu"
  - "M3wtw0"
  - "Cynth1@"
```
