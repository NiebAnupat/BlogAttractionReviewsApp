import 'package:animated_notch_bottom_bar/animated_notch_bottom_bar/animated_notch_bottom_bar.dart';
import 'package:blog_attraction_reviews_flutter/components/blogpostPreviewCard.dart';
import 'package:blog_attraction_reviews_flutter/components/main_padding.dart';
import 'package:blog_attraction_reviews_flutter/getx/blog/logic.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogPost.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class MainPage extends StatelessWidget {
  final NotchBottomBarController? controller;
  MainPage({super.key, this.controller});

  final blogController = Get.put(BlogController());
  @override
  Widget build(BuildContext context) {
    return MainPadding(
      child: Column(
        children: [
          const SizedBox(height: 20),
          Obx(() {
            if (blogController.isLoading.value) {
              return Center(
                child: CircularProgressIndicator(),
              );
            }
            return Column(
              children: [
                ListView.builder(
                  shrinkWrap: true,
                  physics: const NeverScrollableScrollPhysics(),
                  itemCount: blogController.blogPosts.length,
                  itemBuilder: (context, index) {
                    BlogPost blogPost = blogController.blogPosts[index];
                    return BlogPostPreviewCard(blogPost: blogPost);
                  },
                ),
              ],
            );
          }),
        ],
      ),
    );
  }
}
