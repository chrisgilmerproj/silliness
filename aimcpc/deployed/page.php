<?php get_header(); ?>

      <div class"row-fluid">

   <?php if(have_posts()) : ?>
      <?php while(have_posts()) : the_post(); ?>

         <div class="post" id="post-<?php the_ID(); ?>">

               <h2 class="page-header muted"><?php the_title(); ?></h2>

               <?php the_content(); ?>
               <?php edit_post_link('Edit', '<p>', '</p>'); ?>

         </div>

      <?php endwhile; ?>

   <?php else : ?>

      <div class="post">
         <h2><?php _e('Not Found'); ?></h2>
      </div>

   <?php endif; ?>
      </div>

<?php get_footer(); ?>
