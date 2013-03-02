<?php
if ( function_exists('register_sidebar') )
    register_sidebar(array(
        'before_widget' => '<div id="%1$s" class="widget %2$s">',
        'after_widget' => '</div>',
        'before_title' => '<div class="title">',
        'after_title' => '</div>',
    ));

// Add ID and CLASS attributes to the first <ul> occurence in wp_page_menu
function add_menuclass($ulclass) {
return preg_replace('/<ul>/', '<ul class="nav">', $ulclass, 1);
}
add_filter('wp_page_menu','add_menuclass');
?>
