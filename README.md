# DCard 2024 Backend Intern Assignment

This project is about Dcard Backend Intern Assignment.

The APIs are build with Golang and use RESTful to adding or posting ads.
The information will be storage in a SQLite database.

## Requirement

- Golang: 1.22 or higher

## Usage

Build the service, Create SQLite database and create a table in database.

```sh
make build
make create_db
```

Run the service.

```sh
make run
```

## API Detail

Default API URL: `https://<hostname>/ad`.
Using `POST` operate to add a new ad information with a JSon table as input data.
Using `GET` operate with conditions to query eligible ads, the API will return a JSon table.

### `Post`

`POST` operate requires these information with a JSon table:

- title (Require)
  The title of ad
- startAt, endAt (Require)
  The period of time when ad are shown
- conditions: (Optional)
  Conditions for displaying ad, all conditions are also optional
  - ageStart, ageEnd
    The age range of the advertising audience
  - gender
    The gender of the advertising audience (enum: `F`, `M`)
  - country:
    The countries of the advertising audience (enum: `TW`, `JP` or any country code complies with ISO 3166-1)
  - platform:
    The platforms of the advertising audience (enum: `android`, `ios`, `web`)

Example:

```json
{
    "title": "Dcard Intern"
    "startAt": "2024-1-31T03:00:00.000Z"
    "endAt": "2024-4-6T03:00:00.000Z"
    "conditions": {
        "ageStart": 18,
        "ageEnd": 24,
        "gender": ["M", "F"]
        "country": ["TW", "JP", "US"]
        "platform": ["ios", "android", "web"]
    }
}
```

The API will return a string `Data added successfully` if the data is be added into database successfully, or return an error message.

### `GET`

`GET` operate requires these conditions with query parameter:

- offset (Require, default value: 5)
  Which ad to start with.
- limit (Require, default value: 5)
  How many ads to list.
- age (Optional)
  Targe age.
- gender (Optional)
  Targe gender.
- country (Optional)
  Target country.
- platform (Optional)
  Target platform.

Example: `/ad?offset=10&limit=3&age=24&gender=F&country=TW&platform=ios`

The API will return a JSon table of eligible ads, for example:

```json
{
  "item": [
    { "title": "AD 52", "endAt": "2024-12-29T11:34:00Z" },
    { "title": "AD 53", "endAt": "2024-12-30T19:29:00Z" },
    { "title": "AD 54", "endAt": "2024-12-31T13:15:00Z" }
  ]
}
```
