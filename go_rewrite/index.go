package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func ArticleList(articleInfo []ArticleInfo) FragmentNode {
	node := Fragment()
	for _, article := range articleInfo {
		if !article.Published {
			continue
		}
		node.AssignChildren(Children{
			P(
				AHref(
					fmt.Sprintf("/%s", article.Path),
					article.Name,
				),
			),
		})
	}
	return node
}

func Index(articleInfo []ArticleInfo) FragmentNode {
	return Fragment(
		Br(),
		Div(
			ClassName("content"),
			P(`Hi. I'm a programmer based in Oxford, England. I work at the <a href="https://www.centreforeffectivealtruism.org/">Centre for Effective Altruism</a>, primarily on the <a href="https://forum.effectivealtruism.org/">EA Forum</a>. Below you can find examples of my open-source work.`),
		),
		Br(),
		Div(
			ClassName("content"),
			H3("Code"),
			Projects(),
		),
		Br(),
		Div(
			ClassName("content"),
			H3("Writing"),
			Div(
				ClassName("content"),
				ArticleList(articleInfo),
			),
		),
		Br(),
		Div(
			ClassName("content"),
			H3("Contact"),
			Div(
				ClassName("content"),
				P(
					ClassName("tall"),
					"Email:&nbsp;",
					Code("['o','l','l','i','e'].join('').concat('@etherington.xyz')"),
				),
			),
		),
	)
}