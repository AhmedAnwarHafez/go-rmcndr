# go-rmcndr

A social app for sharing music recommendations.

## Installation

Assuming you have configured `tailwindcss` globally using `npm install -g tailwindcss`

```
git clone <URL>
cd go-rmcndr
make dev #tailwind and the webserver should both start in parallel
```

## Tech stack

- Go
- Sqlite3
- Tailwind

## Roadmap

- [ ] When the user searches for new users, the returned results should display the nickname and the top 3 genres in which by counting the number of songs for each genre.
      If there are more genres, [x and more] should be appended
      For example, [3 Trance] [4 Ambient] [5 Dance] and 3 more

- [ ] When the user adds a new song and selects a genre, a Typeahead will be displayed to select an existing genre.
- [ ] Handle authentication
- [ ] add more tables
- [ ] Add friends
