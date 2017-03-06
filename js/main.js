//  This is outside of $(document).ready() to enable lazy loading
$(".carousel").slick({
	dots: true,
	infinite: true,
	fade: true,
	lazyLoad: "progressive",
	slidesToShow: 1,
	slidesToScroll: 1,
	autoplay: true,
	autoplaySpeed: 3500,
});

// $(document).ready(function() {});
