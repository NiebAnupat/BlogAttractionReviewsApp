import 'package:blog_attraction_reviews_flutter/constants.dart';
import 'package:blog_attraction_reviews_flutter/getx/blog/logic.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogContent.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class EditBlogPage extends StatelessWidget {
  late String bid;
  late String title;
  late String description;
  late String thumbnail;
  List<BlogContent>? contents;
  EditBlogPage(
      {super.key,
      required this.bid,
      required this.title,
      required this.description,
      required this.thumbnail});

  final blogController = Get.put(BlogController());
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        leading: null,
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
                    Image.network('$baseUrl/v1/fileStorage/$thumbnail'),
                    const SizedBox(height: 10),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 10),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(title,
                              style: GoogleFonts.notoSansThaiLooped(
                                textStyle: const TextStyle(
                                  fontSize: 35,
                                  fontWeight: FontWeight.bold,
                                ),
                              )),
                          const SizedBox(height: 10),
                          Text(description,
                              style: GoogleFonts.notoSansThaiLooped(
                                textStyle: const TextStyle(
                                  fontSize: 16,
                                ),
                              )),
                          const SizedBox(height: 10),

                          // display currentEditContent
                          Obx(() => Column(children: [
                                ListView.builder(
                                  shrinkWrap: true,
                                  itemCount:
                                      blogController.currentEditContent.length,
                                  itemBuilder: (context, index) {
                                    if (blogController
                                            .currentEditContent[index].type ==
                                        0) {
                                      return Column(
                                        crossAxisAlignment:
                                            CrossAxisAlignment.start,
                                        children: [
                                          Text(
                                              blogController
                                                  .currentEditContent[index]
                                                  .content as String,
                                              style: GoogleFonts
                                                  .notoSansThaiLooped(
                                                textStyle: const TextStyle(
                                                  fontSize: 15,
                                                ),
                                              )),
                                          const SizedBox(height: 15),
                                        ],
                                      );
                                    } else if (blogController
                                            .currentEditContent[index].type ==
                                        1) {
                                      return Column(
                                        crossAxisAlignment:
                                            CrossAxisAlignment.start,
                                        children: [
                                          ClipRRect(
                                            borderRadius:
                                                const BorderRadius.all(
                                                    Radius.circular(20)),
                                            child: Image.network(
                                                '$baseUrl/v1/fileStorage/${blogController.currentEditContent[index].content}'),
                                          ),
                                          const SizedBox(height: 15),
                                        ],
                                      );
                                    }
                                  },
                                )
                              ])),

                          const SizedBox(height: 10),

                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              ButtonTheme(
                                minWidth: 100.0,
                                child: ElevatedButton(
                                  onPressed: () {
                                    blogController.typeContent.value = 0;

                                    // show text input dialog
                                    showDialog(
                                        context: context,
                                        builder: (context) {
                                          return AlertDialog(
                                            title: Text(
                                              'เพิ่มข้อความ',
                                              style: GoogleFonts
                                                  .notoSansThaiLooped(
                                                      textStyle:
                                                          const TextStyle(
                                                fontSize: 20,
                                                fontWeight: FontWeight.bold,
                                              )),
                                            ),
                                            content: TextField(
                                                maxLines: 5,
                                                controller: blogController
                                                    .contentTextController,
                                                decoration: const InputDecoration(
                                                    hintText: 'กรอกข้อความ',
                                                    border:
                                                        OutlineInputBorder()),
                                                style: GoogleFonts
                                                    .notoSansThaiLooped(
                                                        textStyle:
                                                            const TextStyle(
                                                  fontSize: 16,
                                                  fontWeight: FontWeight.bold,
                                                ))),
                                            actions: [
                                              ElevatedButton(
                                                  onPressed: () {
                                                    blogController
                                                        .addContent(bid);
                                                    Navigator.of(context).pop();
                                                  },
                                                  child: Text(
                                                    'เพิ่ม',
                                                    style: GoogleFonts
                                                        .notoSansThai(),
                                                  ))
                                            ],
                                          );
                                        });
                                  },
                                  style: ElevatedButton.styleFrom(
                                      backgroundColor: Colors.red.shade900,
                                      foregroundColor: Colors.white,
                                      shape: const RoundedRectangleBorder(
                                          borderRadius: BorderRadius.only(
                                              topLeft: Radius.circular(10),
                                              bottomLeft:
                                                  Radius.circular(10)))),
                                  child: Text(
                                    'เพิ่มข้อความ',
                                    style: GoogleFonts.notoSansThaiLooped(
                                        textStyle: const TextStyle(
                                      fontSize: 12,
                                      fontWeight: FontWeight.bold,
                                    )),
                                  ),
                                ),
                              ),
                              ButtonTheme(
                                minWidth: 100.0,
                                child: ElevatedButton(
                                  onPressed: () {
                                    blogController.typeContent.value = 1;
                                  },
                                  style: ElevatedButton.styleFrom(
                                      backgroundColor: Colors.red.shade900,
                                      foregroundColor: Colors.white,
                                      shape: const RoundedRectangleBorder(
                                          borderRadius: BorderRadius.only(
                                              topRight: Radius.circular(10),
                                              bottomRight:
                                                  Radius.circular(10)))),
                                  child: Text('เพิ่มรูปภาพ',
                                      style: GoogleFonts.notoSansThaiLooped(
                                          textStyle: const TextStyle(
                                              fontSize: 12,
                                              fontWeight: FontWeight.bold))),
                                ),
                              )
                            ],
                          )
                        ],
                      ),
                    ),
                  ])))),
    );
  }
}
