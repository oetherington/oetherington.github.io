import * as fs from "fs-extra";
import pug from "pug";
import simpleIcons from "simple-icons";
import { marked } from "marked";
import cheerio from "cheerio";
import hljs from "highlight.js";
import gitDateExtractor from "git-date-extractor";

const ROOT = "https://www.etherington.io/";
const ARTICLES = "articles/manifest.json";
const PUBLIC = "public";
const OUT = "docs";

const articles = JSON
	.parse(fs.default.readFileSync(ARTICLES, "utf8"))
	.filter((article) => !article.ignored);

async function compileFile(input, output, params = {}) {
	try {
		const res = pug.renderFile(input, {
			icons: simpleIcons,
			...params,
		});

		const $ = cheerio.load(res);
		$("code").replaceWith(function() {
			const code = $(this).html();
			const language = $(this).attr("lang");
			const hl = language
				? hljs.highlight(code, { language })
				: hljs.highlightAuto(code);
			return `<code>${hl.value}</code>`;
		});

		const html = $.html().replaceAll("&amp;", "&");
		await fs.outputFile(`${OUT}/${output}`, html);
	} catch (e) {
		console.log(e);
		process.exit(1);
	}
}

function makeMarkdownRenderer() {
	const r = new marked.Renderer();
	r.toc = [];
	r.heading = function(text, level, raw) {
		const anchor = this.options.headerPrefix + raw.toLowerCase().replace(/[^\w]+/g, "-");
		r.toc.push({ anchor, level, text, });
		return `<h${level} id="${anchor}"><a href="#${anchor}">&gt; ${text}</a></h${level}>\n`;
	}
	return r;
}

async function build() {
	console.log("Removing old build");
	await fs.remove(OUT);

	console.log("Copying public files");
	await fs.copy(PUBLIC, OUT);

	console.log("Compiling index");
	await compileFile("pages/index.pug", "index.html", {
		title: "Ollie Etherington",
		articles,
	});

	console.log("Compiling articles");
	for (let article of articles) {
		const renderer = makeMarkdownRenderer();
		const md = await fs.default.readFile(`articles/${article.path}.md`);
		const content = marked(md.toString(), {
			renderer,
			gfm: true,
			smartLists: true,
		});
		await compileFile("pages/article.pug", `${article.path}.html`, {
			title: `${article.name} | Ollie Etherington`,
			toc: renderer.toc,
			article,
			content,
		});
	}

	console.log("Getting timestamps");
	const stamps = await gitDateExtractor.getStamps({});

	console.log("Compiling sitemap");
	let locations = articles.map(a => { return {
		url: `${ROOT}${a.path}`,
		modified: stamps[`articles/${a.path}.md`].modified,
	}});
	locations.unshift({
		url: ROOT,
		modified: stamps["pages/index.pug"].modified,
	});
	await compileFile("pages/sitemap.pug", "sitemap.xml", { locations });

	console.log("Done!");
}

build();
