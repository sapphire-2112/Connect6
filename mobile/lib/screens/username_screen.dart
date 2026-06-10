import 'package:flutter/material.dart';

import '../services/storage_service.dart';
import 'contacts_screen.dart';

class UsernameScreen extends StatefulWidget {

  const UsernameScreen({
    super.key,
  });

  @override
  State<UsernameScreen> createState() =>
      _UsernameScreenState();
}

class _UsernameScreenState
    extends State<UsernameScreen> {

  final TextEditingController controller =
  TextEditingController();

  final StorageService storageService =
  StorageService();

  @override
  Widget build(BuildContext context) {

    return Scaffold(

      appBar: AppBar(
        title: const Text(
          "Welcome To Connect6",
        ),
      ),

      body: Padding(

        padding: const EdgeInsets.all(20),

        child: Column(

          children: [

            const SizedBox(height: 30),

            TextField(

              controller: controller,

              decoration:
              const InputDecoration(
                labelText: "Username",
                border:
                OutlineInputBorder(),
              ),
            ),

            const SizedBox(height: 20),

            ElevatedButton(

              onPressed: () async {

                if (controller.text
                    .trim()
                    .isEmpty) {
                  return;
                }

                await storageService
                    .saveName(
                  controller.text.trim(),
                );

                Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(
                    builder: (_) =>
                        ContactsScreen(),
                  ),
                );
              },

              child: const Text(
                "Continue",
              ),
            ),
          ],
        ),
      ),
    );
  }
}