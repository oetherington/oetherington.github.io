// Avoid `console` errors in browsers that lack a console.
(function() {
    var method;
    var noop = function () {};
    var methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeline', 'timelineEnd', 'timeStamp', 'trace', 'warn'
    ];
    var length = methods.length;
    var console = (window.console = window.console || {});
    while (length--) {
        method = methods[length];
        if (!console[method]) console[method] = noop;
    }
}());

// Setup lazyloading for the carousel
$(function() {
	return $(".carousel.lazy").on("slide.bs.carousel", function(ev) {
		var lazy;
		lazy = $(ev.relatedTarget).find("img[data-src]");
		lazy.attr("src", lazy.data("src"));
		lazy.removeAttr("data-src");
	});
});

(function($) {
    "use strict";

    // Page scrolling on click
    $(document).on("click", "a.page-scroll", function(event) {
        var $anchor = $(this);
        $("html, body").stop().animate({
            scrollTop: ($($anchor.attr("href")).offset().top - 50)
        }, 1250, "easeInOutExpo");
        event.preventDefault();
    });

    // Highlight the top nav on scroll
	$("body").scrollspy({
		target: ".navbar-fixed-top",
		offset: 100
	});

    // Close the responsive menu on menu item click
    $(".navbar-collapse ul li a").click(function() {
        $(".navbar-toggle:visible").click();
    });

    // Offset for main nav
	$("#mainNav").affix({
		offset: {
			top: 50
		}
	});

	// Lazyload media
	new LazyLoad({
		parent: "#media",
		childs: "iframe",
		offset: 50
	}).init();

	new LazyLoad({
		parent: "#chrd",
		childs: "img",
		offset: 50
	}).init();

	// Init email API for contact form
	emailjs.init("user_lcLubOKg8r0jVcldtyn7v");

})(jQuery);

function recaptcha_cb() {
	$("#submit-button").css({ "display": "inline-block" });
	$(".g-recaptcha").css({ "display": "none" });
}

$(document).ready(function() {
	$('#contact_form')
		.bootstrapValidator({
			feedbackIcons: {
				valid: "glyphicon glyphicon-ok",
				invalid: "glyphicon glyphicon-remove",
				validating: "glyphicon glyphicon-refresh"
			},
			fields: {
				name: {
					validators: {
						stringLength: {
							min: 2,
						},
						notEmpty: {
							message: "Please supply your name"
						}
					}
				},
				email: {
					validators: {
						notEmpty: {
							message: "Please supply your email address"
						},
						emailAddress: {
							message: "Please supply a valid email address"
						}
					}
				},
				message: {
					validators: {
						notEmpty: {
							message: "Please write a message"
						}
					}
				}
			}
		})
		.submit(function(e) {
			e.preventDefault();

			var v = grecaptcha.getResponse();
			if(v.length == 0) {
				document.getElementById('captcha').innerHTML =
					"You can't leave Captcha Code empty";
				$('#error_message').slideDown({ opacity: "show" }, "slow");
				return;
			}

			$("#success_message").slideDown({ opacity: "show" }, "slow");

			var msg = "Name='" + escape($("input[name=name]").val()) +
				"', Email='" + escape($("input[name=email]").val()) +
				"', Phone='" + escape($("input[name=phone]").val()) +
				"', Message='" + escape($("textarea[name=message]").val()) +
				"'";

			emailjs.send("gmail", "template_D8nLf7kF", {
				to_name: "Ollie",
				from_name: escape($("input[name=name]").val()),
				message_html: msg
			});

			$("input[name=name]").val("");
			$("input[name=email]").val("");
			$("input[name=phone]").val("");
			$("textarea[name=message]").val("");
		});
});
