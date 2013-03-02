<!DOCTYPE html>
<html lang="en">
<head>
	<meta http-equiv="Content-Type" content="<?php bloginfo('html_type'); ?>; charset=<?php bloginfo('charset'); ?>" />	
    <meta charset="utf-8">
    <title><?php bloginfo('name'); ?><?php wp_title(); ?></title>

    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

	<?php if ( is_singular() ) wp_enqueue_script( 'comment-reply' ); ?>
    <?php wp_head(); ?>

    <!-- Le styles -->
    <link href="<?php echo get_template_directory_uri(); ?>/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="<?php echo get_template_directory_uri(); ?>/static/css/bootstrap-responsive.min.css" rel="stylesheet">
    <link href="<?php echo get_template_directory_uri(); ?>/static/css/style.css" rel="stylesheet">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="<?php echo get_template_directory_uri(); ?>/static/js/html5shiv.js"></script>
    <![endif]-->

    <!-- Fav and touch icons -->
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="http://twitter.github.com/bootstrap/assets/ico/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="http://twitter.github.com/bootstrap/assets/ico/apple-touch-icon-114-precomposed.png">
      <link rel="apple-touch-icon-precomposed" sizes="72x72" href="http://twitter.github.com/bootstrap/assets/ico/apple-touch-icon-72-precomposed.png">
                    <link rel="apple-touch-icon-precomposed" href="http://twitter.github.com/bootstrap/assets/ico/apple-touch-icon-57-precomposed.png">
                                   <link rel="shortcut icon" href="http://twitter.github.com/bootstrap/assets/ico/favicon.png">
  </head>

  <body data-spy="scroll" data-target=".navbar">

    <div id="home" class="container">

      <div class="row-fluid">
        <div class="span12">
          <dl class="dl-horizontal pull-right muted">
              <dt>Office</dt><dd>303.805.1800</dd>
              <dt>Fax</dt><dd>303.805.9323</dd>
              <dt><abbr title="For emergencies call 911 immediately">Emergency</abbr></dt><dd><code>911</code></dd>
          </dl>
          <a href="/">
            <img src="<?php echo get_template_directory_uri(); ?>/static/img/aimcpc/logo.jpg">
          </a>
          <h3 class="text-info muted">
          <small class="text-info lead"><?php bloginfo('description'); ?></small>
          </h3>
        </div>
      </div>

      <div class="row-fluid">
        <div class="span12 navbar">
          <div class="navbar-inner">
            <?php if(function_exists('wp_page_menu')) : ?>

               <?php wp_page_menu('show_home=1&depth=1&sort_column=menu_order&title_li=&menu_class=container nav'); ?>

            <?php endif; ?>
          </div>
        </div><!-- /.navbar -->
      </div>

