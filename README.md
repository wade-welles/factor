# Factor

Factor is a tool for building SPA web applications in Go using Web Assembly.

Factor is heavily inspired by [Svelte and Sapper](https://sapper.svelte.technology/guide#getting-started), but doesn't attempt to be a straight port, but rather take those ideas and present them in a way natural for a Go developer.

## Application Structure

```
├ app
│ ├ App.html # Main/initial view
│ └ template.html # application template
├ assets
│ ├ # your files here
├ components 
│ ├ # your files here
├ routes
│ ├ # your routes here - these are pages.
│ ├ _error.html
│ └ index.html
```