package api

var Pastebin = PasteService{ServiceName: "Pastebin", Endpoint: "https://pastebin.com/raw/%s"}
var GithubGists = PasteService{ServiceName: "Github", Endpoint: "https://gist.githubusercontent.com/%s/raw"}
var GitlabSnippets = PasteService{ServiceName: "Gitlab", Endpoint: "https://gitlab.com/snippets/%s/raw"}
var RentryCo = PasteService{ServiceName: "Rentry", Endpoint: "https://rentry.co/%s/raw"}
