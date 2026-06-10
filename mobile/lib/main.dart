import 'package:flutter/material.dart';
import 'package:qr_flutter/qr_flutter.dart';
import 'screens/contacts_screen.dart';
import 'screens/qr_screen.dart';
import 'screens/scan_screen.dart';


void main() {
  runApp(const Connect6App());
}

class Connect6App extends StatelessWidget {
  const Connect6App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,

      title: 'Connect6',

      theme: ThemeData(
        useMaterial3: true,

        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.blue,
        ),

        scaffoldBackgroundColor: const Color(
          0xFFF4F8FF,
        ),

        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.blue,
          foregroundColor: Colors.white,
          centerTitle: true,
        ),
      ),

      home: ContactsScreen(),
    );
  }
}