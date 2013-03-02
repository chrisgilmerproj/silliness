<div class="nav">

   <?php if(function_exists('wp_page_menu')) : ?>

      <?php wp_page_menu('show_home=1&depth=1&sort_column=menu_order&title_li='); ?>

   <?php endif; ?>

</div>