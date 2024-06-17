import 'package:blog_attraction_reviews_flutter/components/main_padding.dart';
import 'package:blog_attraction_reviews_flutter/constants.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogPost.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:get/get_connect/http/src/utils/utils.dart';
import 'package:google_fonts/google_fonts.dart';

class ViewBlogPage extends StatelessWidget {
  final BlogPost blogPost;
  const ViewBlogPage({super.key, required this.blogPost});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        surfaceTintColor: Colors.transparent,
        shape: const RoundedRectangleBorder(
          borderRadius: BorderRadius.vertical(
            bottom: Radius.circular(15),
          ),
        ),
        elevation: 50,
        title: Row(
          children: [
            Image.asset(
              'assets/attraction_icon.png',
              width: 50,
            ),
            const SizedBox(width: 10),
            Text(
              'เที่ยวไหนดี ?',
              style: GoogleFonts.sriracha(
                textStyle: const TextStyle(
                  fontSize: 30,
                  // fontWeight: FontWeight.bold,
                ),
              ),
            ),
          ],
        ),
      ),
      extendBody: true,
      body: Container(
          color: Colors.white,
          child: ClipRRect(
            borderRadius: const BorderRadius.only(
                topLeft: Radius.circular(20), topRight: Radius.circular(20)),
            child: SingleChildScrollView(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Image.network(
                      '$baseUrl/v1/fileStorage/${blogPost.thumbnail}'),
                  const SizedBox(height: 10),
                  Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 10),
                    child: Column(
                      children: [
                        Text(blogPost.title,
                            style: GoogleFonts.notoSansThaiLooped(
                              textStyle: const TextStyle(
                                fontSize: 35,
                                fontWeight: FontWeight.bold,
                              ),
                            )),
                        const SizedBox(height: 10),
                        Text(blogPost.description,
                            style: GoogleFonts.notoSansThaiLooped(
                              textStyle: const TextStyle(
                                fontSize: 16,
                              ),
                            )),
                      ],
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 20),
                    child: Column(
                      children: [
                        // check if blogPost has content then create listbuilder
                        if (blogPost.contents != null)
                          ListView.builder(
                            shrinkWrap: true,
                            physics: const NeverScrollableScrollPhysics(),
                            itemCount: blogPost.contents!.length,
                            itemBuilder: (context, index) {
                              if (blogPost.contents![index]!.type == 0) {
                                return Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    Text(
                                        blogPost.contents![index]!.content
                                            as String,
                                        style: GoogleFonts.notoSansThaiLooped(
                                          textStyle: const TextStyle(
                                            fontSize: 15,
                                          ),
                                        )),
                                    const SizedBox(height: 15),
                                  ],
                                );
                              } else if (blogPost.contents![index]!.type == 1) {
                                return Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    ClipRRect(
                                      borderRadius: const BorderRadius.all(
                                          Radius.circular(20)),
                                      child: Image.network(
                                          '$baseUrl/v1/fileStorage/${blogPost.contents![index]!.content}'),
                                    ),
                                    const SizedBox(height: 15),
                                  ],
                                );
                              }
                            },
                          ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          )),
    );
  }
}
