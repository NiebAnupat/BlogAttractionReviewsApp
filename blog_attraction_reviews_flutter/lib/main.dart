import 'package:blog_attraction_reviews_flutter/pages/loginPage.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

void main() {
  runApp(const GetMaterialApp(home: MainApp()));
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: LoginPage(),
      ),
    );
  }
}
