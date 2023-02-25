# www.etherington.io

This repo contains the source and content for www.etherington.io. The site
is written using [Smetana](https://github.com/oetherington/smetana) and is
compiled into the static `/build/` directory to be served by Github Pages.

Articles are written in Markdown and a sitemap is automatically generated using
the git history to get the last modified date for each page.

The only dependencies are `go` and `git`.

### Commands

 - Build the site: `make run`
 - Run the linter: `make lint`
 - Check the code formatting: `make check-fmt`
 - Run the dev server on port 8000: `make serve`
