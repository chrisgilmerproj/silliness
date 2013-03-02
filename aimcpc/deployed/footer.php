
      <hr>

      <div class="footer">
        <p><a href="<?php bloginfo('url'); ?>" title="<?php bloginfo('name'); ?>"><?php bloginfo('name'); ?></a> <?php _e('Copyright',''); ?> &#169; <?php print(date(__('Y',''))); ?></p>
        <p>For problems with the website contact the <a href="mailto:webadmin@aimcpc.com">Website Administrator</a></p>
      </div>

    <?php wp_footer(); ?>

    </div> <!-- /container -->

    <script src="<?php echo get_template_directory_uri(); ?>/static/js/jquery-1.9.1.min.js"></script>
    <script src="<?php echo get_template_directory_uri(); ?>/static/js/bootstrap.min.js"></script>
    <script>
      !function ($) {
        $(function(){
          // carousel demo
          $('#myCarousel').carousel()
        })
      }(window.jQuery)
    </script>

</body>
</html>
