<?php get_header(); ?>

      <div class="row-fluid">
        <div class="span12">
          <div id="myCarousel" class="carousel slide">
            <ol class="carousel-indicators">
              <li data-target="#myCarousel" data-slide-to="0" class="active"></li>
              <li data-target="#myCarousel" data-slide-to="1"></li>
              <li data-target="#myCarousel" data-slide-to="2"></li>
              <li data-target="#myCarousel" data-slide-to="3"></li>
            </ol>
            <!-- Carousel items -->
            <div class="carousel-inner">
              <div class="active item">
                  <img src="<?php echo get_template_directory_uri(); ?>/static/img/aimcpc/office1.jpg" alt="">
              </div>
              <div class="item">
                  <img src="<?php echo get_template_directory_uri(); ?>/static/img/aimcpc/office2.jpg" alt="">
              </div>
              <div class="item">
                  <img src="<?php echo get_template_directory_uri(); ?>/static/img/aimcpc/office3.jpg" alt="">
              </div>
              <div class="item">
                  <img src="<?php echo get_template_directory_uri(); ?>/static/img/aimcpc/office4.jpg" alt="">
              </div>
            </div>
            <!-- Carousel nav -->
            <a class="carousel-control left" href="#myCarousel" data-slide="prev">&lsaquo;</a>
            <a class="carousel-control right" href="#myCarousel" data-slide="next">&rsaquo;</a>
          </div>
        </div>
      </div>

      <hr>

      <div class="row-fluid">
        <?php
        $args = array( 'numberposts' => 3 );
        $lastposts = get_posts( $args );
        foreach($lastposts as $post) : setup_postdata($post); ?>
          <div class="span4">
            <h2><?php the_title(); ?></h2>
            <?php the_content(); ?>
          </div>
        <?php endforeach; ?>
      </div>

<?php get_footer(); ?>
