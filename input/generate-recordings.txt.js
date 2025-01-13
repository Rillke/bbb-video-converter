// execute in browser context (wordpress bbb plugin)
//
//
let a = jQuery('#bb_recordings').find('tbody tr').map(function(i, tr) {
    let $tds = jQuery(tr).find('td');
    let location = $tds.eq(0).text();
    let start = $tds.eq(1).text();
    let duration = $tds.eq(2).text();
    let file_namme = `${location}_${start}_${duration}.webm`;

    let id = $tds.eq(5).find('.dashicons-format-video').parent().attr('href').replace(/.+\//, '');

    return id + '|' +  file_namme;
});

jQuery('<pre>').text(a.toArray().join('\n')).appendTo('body');

