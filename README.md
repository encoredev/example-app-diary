<img width="200px" src="https://encore.dev/assets/branding/logo/logo.svg" alt="Encore - Backend development engine" />

# Diary Example App

This is a lightweight Go application that allows you to easily log and record a journal or diary.
It contains two endpoints: one for listing the journal for a single day, and one for creating a new entry in the journal.

This was built in 30 minutes, with tests included, using [Encore](https://encore.dev). It took 2 minutes to deploy, including with the database.

## Endpoints

Listing all my entries for a specific date:

```bash
curl -X GET "http://localhost:4000/logbook/entries/2022-09-08" | jq
```

```json
{
  "Entries": [
    {
      "Id": 1,
      "Time": "2022-09-08T16:25:21.409544Z",
      "Text": "I cooked fried eggs this morning. I forgot how much I loved them!"
    }
  ]
}
```

Create new diary entry:

```bash
curl -X POST --data '{"Text":"I cooked fried eggs this morning. I forgot how much I loved them!"}' "http://localhost:4000/logbook/entries"
```

```json
{
  "Id": 6,
  "Time": "2022-09-08T16:25:21.409544Z",
  "Text": "I cooked fried eggs this morning. I forgot how much I loved them!"
}
```

## Install

```bash
brew install go@1.19.1
brew install encoredev/tap/encore
git clone git@github.com:encorecommunity/example-app-diary.git
encore auth signup # or encore auth login (if you already have an account)
encore run
```

## Deployment

```bash
encore app create my-diary-app-name
git push origin main
```

Then head over to <https://app.encore.dev> to find out your production URL, and off you go into the clouds!

## Testing

```bash
encore test
```

## Contributing

All contributions are welcome! All we ask is that you adhere to the [Code of Conduct](https://github.com/encoredev/encore/blob/main/CODE_OF_CONDUCT.md)

- [Quick Start with Encore](https://encore.dev/docs/quick-start)
- [Create an Account with Encore]()
- [Go Cheatsheet](https://encore.dev/guide/go.mod)
