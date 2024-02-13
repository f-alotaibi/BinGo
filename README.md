# BinGo
One paste frontend to rule them all, now more performant 😎, and less code size 🚀

Running instance is at https://pastebin-go.onrender.com

# implemented pastes
- Pastebin
- Github gists
- Gitlab snippets
- Rentry.co

# building
make sure you have `templ` installed
```
go install github.com/a-h/templ/cmd/templ@latest
templ generate .
go build .
```