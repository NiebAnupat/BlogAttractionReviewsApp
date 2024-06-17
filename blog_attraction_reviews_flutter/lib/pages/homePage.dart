import 'package:blog_attraction_reviews_flutter/components/main_padding.dart';
import 'package:blog_attraction_reviews_flutter/pages/%E0%B8%B4viewBlog.dart';
import 'package:blog_attraction_reviews_flutter/pages/addBlog.dart';
import 'package:blog_attraction_reviews_flutter/pages/mainPage.dart';
import 'package:blog_attraction_reviews_flutter/pages/profilePage.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:animated_notch_bottom_bar/animated_notch_bottom_bar/animated_notch_bottom_bar.dart';

class HomePage extends StatefulWidget {
  HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final _pageController = PageController(initialPage: 0);
  final NotchBottomBarController _controller =
      NotchBottomBarController(index: 0);

  int maxCount = 3;

  @override
  void dispose() {
    _pageController.dispose();

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final List<Widget> bottomBarPages = [
      MainPage(controller: (_controller)),
      AddBlogPage(),
      const ProfilePage(),
    ];
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
        body: PageView(
          controller: _pageController,
          physics: const NeverScrollableScrollPhysics(),
          children: List.generate(
              bottomBarPages.length, (index) => bottomBarPages[index]),
        ),
        extendBody: true,
        bottomNavigationBar: (bottomBarPages.length <= maxCount)
            ? AnimatedNotchBottomBar(
                notchBottomBarController: _controller,
                color: Colors.red.shade900,
                showLabel: true,
                textOverflow: TextOverflow.visible,
                maxLine: 1,
                shadowElevation: 5,
                kBottomRadius: 28.0,
                notchColor: Colors.grey.shade800,
                removeMargins: false,
                bottomBarWidth: 500,
                showShadow: false,
                durationInMilliSeconds: 300,
                itemLabelStyle: GoogleFonts.notoSansThai(
                  textStyle: const TextStyle(
                    fontSize: 12,
                    color: Colors.white,
                  ),
                ),
                elevation: 1,
                bottomBarItems: const [
                  BottomBarItem(
                      inActiveItem: Icon(
                        Icons.home_filled,
                        color: Colors.white,
                      ),
                      activeItem: Icon(
                        Icons.home_filled,
                        color: Colors.white,
                      ),
                      itemLabel: 'หน้าหลัก'),
                  BottomBarItem(
                    inActiveItem: Icon(
                      Icons.add_circle_sharp,
                      color: Colors.white,
                    ),
                    activeItem: Icon(
                      Icons.add,
                      color: Colors.white,
                    ),
                  ),
                  BottomBarItem(
                      inActiveItem: Icon(
                        Icons.home_filled,
                        color: Colors.white,
                      ),
                      activeItem: Icon(
                        Icons.home_filled,
                        color: Colors.white,
                      ),
                      itemLabel: 'โปรไฟล์'),
                ],
                onTap: (index) {
                  // log('current selected index $index');
                  _pageController.jumpToPage(index);
                },
                kIconSize: 24.0,
              )
            : null);
  }
}
