import 'package:blog_attraction_reviews_flutter/components/main_padding.dart';
import 'package:blog_attraction_reviews_flutter/getx/blog/logic.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class AddBlogPage extends StatefulWidget {
  AddBlogPage({super.key});
  final _formKey = GlobalKey<FormState>();
  @override
  State<AddBlogPage> createState() => _AddBlogPageState();
}

class _AddBlogPageState extends State<AddBlogPage> {
  final blogController = Get.put(BlogController());

  @override
  Widget build(BuildContext context) {
    return MainPadding(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 20),
          Column(children: [
            Text(
              "เพิ่มบทความใหม่",
              style: GoogleFonts.notoSansThaiLooped(
                textStyle: const TextStyle(
                  fontSize: 30,
                  fontWeight: FontWeight.bold,
                ),
              ),
            ),
            const SizedBox(height: 10),
            Form(
              key: widget._formKey,
              child: Column(children: [
                TextFormField(
                  controller: blogController.titleController,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'กรุณากรอกข้อมูล';
                    }
                    return null;
                  },
                  decoration: const InputDecoration(
                    labelText: 'หัวข้อ',
                    prefixIcon: Icon(Icons.subject),
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.all(Radius.circular(10)),
                    ),
                  ),
                ),
                const SizedBox(height: 10),
                TextFormField(
                  controller: blogController.descriptionController,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'กรุณากรอกข้อมูล';
                    }
                    return null;
                  },
                  keyboardType: TextInputType.multiline,
                  maxLines: 5,
                  decoration: const InputDecoration(
                    labelText: 'รายละเอียด',
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 10),
                TextFormField(
                  controller: blogController.thumbnailController,
                  keyboardType: TextInputType.none,
                  readOnly: true,
                  onTap: () {
                    blogController.selectImage();
                  },
                  decoration: const InputDecoration(
                    prefixIcon: Icon(
                      Icons.image,
                    ),
                    labelText: 'รูปภาพพรีวิว',
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 10),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    ElevatedButton(
                      onPressed: () {
                        if (widget._formKey.currentState!.validate()) {
                          blogController.createBlog();
                        }
                      },
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.red.shade400,
                        foregroundColor: Colors.white,

                        // padding: const EdgeInsets.symmetric(horizontal: 50),
                        // shape: RoundedRectangleBorder(
                        //   borderRadius: BorderRadius.circular(10),
                        // ),
                      ),
                      child: Text(
                        'เพิ่มบทความ',
                        style: GoogleFonts.notoSansThaiLooped(
                            textStyle: const TextStyle(
                          fontSize: 12,
                          fontWeight: FontWeight.bold,
                        )),
                      ),
                    ),
                  ],
                ),
              ]),
            ),
          ]),
        ],
      ),
    );
  }
}
