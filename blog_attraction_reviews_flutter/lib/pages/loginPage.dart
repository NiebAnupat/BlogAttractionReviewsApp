import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

import '../getx/user/logic.dart';

class LoginPage extends StatefulWidget {
  LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final userController = Get.put(UserController());
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Container(
        color: Colors.white,
        padding: const EdgeInsets.all(30),
        child: Column(
          // mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const SizedBox(
              height: 100,
            ),
            Center(
              child: ClipRRect(
                borderRadius: BorderRadius.circular(100),
                child: Image.network(
                  "https://cdn.dribbble.com/users/470545/screenshots/3120636/trip-icons-animation.gif",
                  height: 200,
                ),
              ),
            ),
            const SizedBox(
              height: 30,
            ),
            Center(
              child: Text(
                "เที่ยวไหนดี?",
                style: GoogleFonts.sriracha(
                  textStyle: const TextStyle(
                    fontSize: 30,
                    // fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ),
            const SizedBox(
              height: 25,
            ),
            Padding(
              padding: const EdgeInsets.only(left: 15),
              child: Text(
                "เข้าสู่ระบบ",
                textAlign: TextAlign.start,
                style: GoogleFonts.notoSansThaiLooped(
                  textStyle: const TextStyle(fontSize: 20
                      // fontWeight: FontWeight.bold,
                      ),
                ),
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            Form(
                key: _formKey,
                child: Column(
                  children: [
                    TextFormField(
                      controller: userController.usernameController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'กรุณากรอกชื่อผู้ใช้';
                        }
                        return null;
                      },
                      style: GoogleFonts.rubik(
                        textStyle: const TextStyle(
                          fontSize: 15,
                          fontWeight: FontWeight.normal,
                        ),
                      ),
                      decoration: InputDecoration(
                        labelText: 'ชื่อผู้ใช้',
                        labelStyle: GoogleFonts.notoSansThaiLooped(
                          textStyle: TextStyle(
                            fontSize: 14,
                            color: Colors.red.shade900,
                          ),
                        ),
                        contentPadding: const EdgeInsets.only(
                            left: 15, top: 8, right: 15, bottom: 0),
                        border: const OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(30)),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius:
                              const BorderRadius.all(Radius.circular(30)),
                          borderSide: BorderSide(color: Colors.red.shade900),
                        ),
                      ),
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    TextFormField(
                      keyboardType: TextInputType.visiblePassword,
                      obscureText: true,
                      controller: userController.passwordController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'กรุณากรอกรหัสผ่าน';
                        }
                        return null;
                      },
                      style: GoogleFonts.rubik(
                        textStyle: const TextStyle(
                          fontSize: 15,
                          fontWeight: FontWeight.normal,
                        ),
                      ),
                      decoration: InputDecoration(
                        labelText: 'รหัสผ่าน',
                        labelStyle: GoogleFonts.notoSansThaiLooped(
                          textStyle: TextStyle(
                            fontSize: 14,
                            color: Colors.red.shade900,
                          ),
                        ),
                        contentPadding: const EdgeInsets.only(
                            left: 15, top: 8, right: 15, bottom: 0),
                        border: const OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(30)),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius:
                              const BorderRadius.all(Radius.circular(30)),
                          borderSide: BorderSide(color: Colors.red.shade900),
                        ),
                      ),
                    ),
                    const SizedBox(
                      height: 5,
                    ),
                    Row(
                      children: [
                        Obx(() => Checkbox(
                              activeColor: Colors.red.shade900,
                              value: userController.isCheck.value,
                              onChanged: (value) =>
                                  userController.onCheck(value),
                            )),
                        Text(
                          'จดจำรหัสผ่าน',
                          style: GoogleFonts.notoSansThaiLooped(
                            textStyle: const TextStyle(
                              fontSize: 15,
                              fontWeight: FontWeight.normal,
                            ),
                          ),
                        ),
                      ],
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 15),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        crossAxisAlignment: CrossAxisAlignment.end,
                        children: [
                          TextButton(
                              style: TextButton.styleFrom(
                                minimumSize: Size.zero,
                                // padding: EdgeInsets.zero,
                                padding: const EdgeInsets.fromLTRB(
                                    10.0, 3.0, 10.0, 3.0),
                                tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                              ),
                              onPressed: () {},
                              child: Text(
                                "ลืมรหัสผ่าน",
                                style: GoogleFonts.notoSansThaiLooped(
                                  textStyle: TextStyle(
                                      fontSize: 13,
                                      fontWeight: FontWeight.bold,
                                      color:
                                          Colors.red.shade900.withOpacity(0.7)),
                                ),
                              )),
                          TextButton(
                              onPressed: () {
                                if (_formKey.currentState!.validate()) {
                                  userController.login();
                                }
                              },
                              style: TextButton.styleFrom(
                                backgroundColor: Colors.red.shade900,
                                // padding: const EdgeInsets.symmetric(
                                //     horizontal: 20, vertical: 10),
                                padding: const EdgeInsets.fromLTRB(
                                    15.0, 3.0, 15.0, 3.0),
                              ),
                              child: Text(
                                "เข้าสู่ระบบ",
                                style: GoogleFonts.notoSansThaiLooped(
                                  textStyle: const TextStyle(
                                      fontSize: 15,
                                      fontWeight: FontWeight.bold,
                                      color: Colors.white),
                                ),
                              )),
                        ],
                      ),
                    ),
                    const SizedBox(
                      height: 50,
                    ),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          "ยังไม่มีบัญชีผู้ใช้?",
                          style: GoogleFonts.notoSansThaiLooped(
                            textStyle: const TextStyle(
                              fontSize: 15,
                              fontWeight: FontWeight.normal,
                            ),
                          ),
                        ),
                        TextButton(
                            onPressed: () {},
                            style: TextButton.styleFrom(
                              minimumSize: Size.zero,
                              padding: const EdgeInsets.symmetric(
                                  horizontal: 10, vertical: 0),
                              tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                            ),
                            child: Text(
                              "สมัครสมาชิก",
                              style: GoogleFonts.notoSansThaiLooped(
                                textStyle: TextStyle(
                                    fontSize: 15,
                                    fontWeight: FontWeight.bold,
                                    color: Colors.red.shade900),
                              ),
                            )),
                      ],
                    ),
                  ],
                )),
          ],
        ));
  }
}
