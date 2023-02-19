package main

import (
	. "github.com/oetherington/smetana"
)

func Index() FragmentNode {
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
				P("TODO"),
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
