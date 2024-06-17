import 'package:flutter/material.dart';

class MainPadding extends StatelessWidget {
  final Widget child;
  const MainPadding({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      padding: const EdgeInsets.only(bottom: 110),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 30),
        child: ClipRRect(
            borderRadius: const BorderRadius.all(Radius.circular(20)),
            child: SingleChildScrollView(child: child)),
      ),
    );
  }
}
