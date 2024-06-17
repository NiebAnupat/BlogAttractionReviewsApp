import 'package:blog_attraction_reviews_flutter/constants.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogPost.dart';
import 'package:blog_attraction_reviews_flutter/pages/%E0%B8%B4viewBlog.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:http/http.dart' as http;

class BlogPostPreviewCard extends StatefulWidget {
  final BlogPost blogPost;

  const BlogPostPreviewCard({super.key, required this.blogPost});

  @override
  State<BlogPostPreviewCard> createState() => _BlogPostPreviewCardState();
}

class _BlogPostPreviewCardState extends State<BlogPostPreviewCard> {
  @override
  void initState() {
    super.initState();
  }

  Future<void> loadThumbnail() async {
    final res = await http
        .get(Uri.parse('$baseUrl/v1/fileStorage/${widget.blogPost.thumbnail}'));
    if (res.statusCode == 200) {
      print(res.body);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        InkWell(
          onTap: () {
            Get.to(() => ViewBlogPage(blogPost: widget.blogPost));
          },
          child: Card(
            color: Colors.white,
            child: Column(
              children: [
                ClipRRect(
                    borderRadius: const BorderRadius.only(
                      topLeft: Radius.circular(10),
                      topRight: Radius.circular(10),
                    ),
                    child: Image.network(
                        '$baseUrl/v1/fileStorage/${widget.blogPost.thumbnail}')),
                SizedBox(
                  width: double.infinity,
                  // color: Colors.red,
                  child: Padding(
                    padding: const EdgeInsets.symmetric(
                        vertical: 10, horizontal: 20),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      // mainAxisAlignment: MainAxisAlignment.start,
                      children: [
                        Text(
                          widget.blogPost.title,
                          style: GoogleFonts.notoSansThaiLooped(
                            textStyle: const TextStyle(
                              fontSize: 25,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                        const SizedBox(height: 5),
                        Text(
                          widget.blogPost.description,
                          style: const TextStyle(
                            fontSize: 16,
                            // fontWeight: FontWeight.w300,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
        const SizedBox(height: 10),
      ],
    );
  }
}
