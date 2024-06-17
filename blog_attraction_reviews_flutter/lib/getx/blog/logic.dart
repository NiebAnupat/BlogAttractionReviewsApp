import 'dart:convert';

import 'package:blog_attraction_reviews_flutter/constants.dart';
import 'package:blog_attraction_reviews_flutter/getx/user/logic.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogContent.dart';
import 'package:blog_attraction_reviews_flutter/model/BlogPost.dart';
import 'package:blog_attraction_reviews_flutter/pages/editBlog.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_state_manager/src/simple/get_controllers.dart';
import 'package:http/http.dart' as http;
import 'package:image_picker/image_picker.dart';

class BlogController extends GetxController {
  final RxList<BlogPost> blogPosts = <BlogPost>[].obs;
  final RxBool isLoading = true.obs;

  final RxList<BlogContent> currentEditContent = <BlogContent>[].obs;

  final titleController = TextEditingController();
  final descriptionController = TextEditingController();
  final thumbnailController = TextEditingController();
  final contentTextController = TextEditingController();
  final typeContent = 0.obs;

  final ImagePicker _picker = ImagePicker();
  final Rx<XFile> selectThumbnail = XFile('').obs;

  @override
  void onInit() {
    super.onInit();

    loadData();
  }

  Future<void> loadData() async {
    isLoading.value = true;
    final response = await http.get(Uri.parse('$baseUrl/v1/blog'));
    if (response.statusCode == 200) {
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      blogPosts.value =
          data['blogs'].map<BlogPost>((e) => BlogPost.fromJson(e)).toList();
    }
    isLoading.value = false;
  }

  void selectImage() {
    _picker.pickImage(source: ImageSource.gallery).then((value) {
      if (value != null) {
        selectThumbnail.value = value;
        contentTextController.text = value.name;
      }
    });
  }

  Future<void> createBlog() async {
    Get.dialog(
      const Center(
        child: CircularProgressIndicator(),
      ),
      barrierDismissible: false,
    );
    final res = await http.post(
      Uri.parse('$baseUrl/v1/blog'),
      body: jsonEncode({
        'title': titleController.text,
        'description': descriptionController.text,
      }),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${Get.find<UserController>().token.value}',
      },
    );
    Get.back();
    if (res.statusCode == 201) {
      Get.snackbar(
        'เพิ่มบทความสำเร็จ',
        'บทความของคุณถูกเพิ่มเรียบร้อยแล้ว',
        snackPosition: SnackPosition.BOTTOM,
        backgroundColor: Colors.green,
        colorText: Colors.white,
        margin: const EdgeInsets.only(bottom: 10),
        duration: const Duration(seconds: 3),
      );
      titleController.clear();
      descriptionController.clear();
      await loadData();

      final newBlog = blogPosts.last;
      Get.to(EditBlogPage(
          bid: newBlog.id,
          title: newBlog.title,
          description: newBlog.description,
          thumbnail: newBlog.thumbnail));
    } else {
      Get.snackbar(
        'Error',
        'Failed to create blog',
        snackPosition: SnackPosition.BOTTOM,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
    }
  }

  Future<void> addContent(String blogID) async {
    Get.dialog(
      const Center(
        child: CircularProgressIndicator(),
      ),
      barrierDismissible: false,
    );

    final Map<String, String> body = {
      'content': contentTextController.text,
      'blogID': blogID,
      'order': 1.toString(),
      'type': typeContent.value.toString(),
      'value':
          typeContent.value == 0 ? contentTextController.text : 'image.jpg',
    };

    final res = await http.post(
      Uri.parse('$baseUrl/v1/blog/content'),
      body: body,
      headers: {
        'Authorization': 'Bearer ${Get.find<UserController>().token.value}',
      },
    );
    Get.back();
    if (res.statusCode == 201) {
      await loadData();
      // update currentEditContent
      final blog = blogPosts.firstWhere((element) => element.id == blogID);
      currentEditContent.value = blog.contents! as List<BlogContent>;
    } else {
      Get.snackbar(
        'Error',
        'Failed to add content',
        snackPosition: SnackPosition.BOTTOM,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
    }
    contentTextController.clear();
  }
}
