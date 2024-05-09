import 'dart:convert';

import 'package:blog_attraction_reviews_flutter/constants.dart';
import 'package:blog_attraction_reviews_flutter/pages/homePage.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

class UserController extends GetxController {
  final TextEditingController usernameController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final isCheck = false.obs;
  final userID = ''.obs;
  final username = ''.obs;
  final avatar = ''.obs;

  @override
  void onClose() {
    usernameController.dispose();
    passwordController.dispose();
    super.onClose();
  }

  Future<void> login() async {
    if (kDebugMode) {
      print('Username: ${usernameController.text}');
      print('Password: ${passwordController.text}');
      print('Remember me: $isCheck');
    }
    if (usernameController.text.isEmpty || passwordController.text.isEmpty) {
      Get.snackbar(
        'Error',
        'Please fill in all fields',
        snackPosition: SnackPosition.BOTTOM,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      return;
    }
    Get.dialog(
      const Center(
        child: CircularProgressIndicator(),
      ),
      barrierDismissible: false,
    );
    final res = await http.post(
      Uri.parse('$baseUrl/v1/auth/login'),
      body: {
        'username': usernameController.text,
        'password': passwordController.text,
      },
    );
    Get.back();
    final data = jsonDecode(res.body);
    if (kDebugMode) {
      print('Data: $data');
    }
    if (res.statusCode == 200) {
      var token = data['token'];

      if (isCheck.value) {
        saveData(token);
      } else {
        clearData();
      }

      Get.offAll(const HomePage());
    } else {
      Get.snackbar('เข้าสู่ระบบไม่สำเร็จ', 'กรุณาตรวจสอบชื่อผู้ใช้หรือรหัสผ่าน',
          snackPosition: SnackPosition.BOTTOM,
          backgroundColor: Colors.red.shade800,
          colorText: Colors.white,
          margin: const EdgeInsets.only(bottom: 20, left: 10, right: 10),
          duration: const Duration(seconds: 3));
    }
  }

  void onCheck(bool? value) {
    isCheck.value = value!;
  }

  void saveData(String token) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString('token', token);
  }

  void clearData() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.remove('token');
  }

  void loadData() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('token');
    if (token != null) {
      final res = await http.post(
        Uri.parse('$baseUrl/v1/auth/verify'),
        headers: {
          'Authorization': 'Bearer $token',
        },
      );
      final data = jsonDecode(res.body);
      if (kDebugMode) {
        print('Data: $data');
      }
      if (res.statusCode == 200) {
        userID.value = data['id'];
        username.value = data['username'];
        avatar.value = data['avatar'];

        Get.offAll(const HomePage());
      } else {
        clearData();
      }
    }
  }
}