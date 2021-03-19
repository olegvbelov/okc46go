$(document).ready(function() {

    "use strict";

    $('.city').click(function (e) {
        if ( !$(this).hasClass('selected') ) {
            var oldAddressId = $('a.selected').attr("href");
            var newAddressId = $(this).attr("href");
            var oldMapId = "#map" + oldAddressId.substring(1);
            var newMapId = "#map" + newAddressId.substring(1);
            $(oldAddressId).css('display', 'none');
            $('a.selected').removeClass('selected');
            $(this).addClass('selected');
            $(newAddressId).css('display', '');
            $(oldMapId).css('display', 'none');
            $(newMapId).css('display', '');
        }
    });
});