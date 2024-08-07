# www.etherington.xyz

This repo contains the source and content for
[etherington.xyz](https://www.etherington.xyz). The site is written using
[Smetana](https://github.com/oetherington/smetana) and is compiled into the
static `/build/` directory to be served by Github Pages.

Articles are written in Markdown and a sitemap is automatically generated using
the git history to get the last modified date for each page.

The only dependencies are `go` and `git`, and optionally `caddy` for running
the dev server.

### Commands

 - Build the site: `make run`
 - Run the linter: `make lint`
 - Check the code formatting: `make check-fmt`
 - Run the dev server on port 8000: `make serve`
