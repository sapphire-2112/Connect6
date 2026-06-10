import 'package:flutter/material.dart';
import '../services/storage_service.dart';

class ProfileScreen extends StatefulWidget {

  const ProfileScreen({
    super.key,
  });

  @override
  State<ProfileScreen> createState() =>
      _ProfileScreenState();
}

class _ProfileScreenState
    extends State<ProfileScreen> {

  final StorageService storageService =
  StorageService();

  String username = "";

  @override
  void initState() {

    super.initState();

    loadUsername();
  }

  Future<void> loadUsername() async {

    final savedName =
    await storageService.loadName();

    setState(() {

      username =
          savedName ?? "Unknown";
    });
  }
  @override
  Widget build(BuildContext context) {

    return Scaffold(

      appBar: AppBar(
        title: const Text(
          "Profile",
        ),
      ),

      body: Padding(

        padding:
        const EdgeInsets.all(20),

        child: Column(

          crossAxisAlignment:
          CrossAxisAlignment.start,

          children: [

            const SizedBox(height: 20),

            const Text(
              "Username",
              style: TextStyle(
                fontSize: 18,
                fontWeight:
                FontWeight.bold,
              ),
            ),

            const SizedBox(height: 10),

            Text(
              username,
              style: const TextStyle(
                fontSize: 16,
              ),
            ),

            const SizedBox(height: 10),

            ElevatedButton(

              onPressed: () {

                final controller =
                TextEditingController(
                  text: username,
                );

                showDialog(

                  context: context,

                  builder: (context) {

                    return AlertDialog(

                      title: const Text(
                        "Edit Username",
                      ),

                      content: TextField(
                        controller: controller,
                      ),

                      actions: [

                        TextButton(

                          onPressed: () {
                            Navigator.pop(context);
                          },

                          child: const Text(
                            "Cancel",
                          ),
                        ),

                        TextButton(

                          onPressed: () async {

                            await storageService
                                .saveName(
                              controller.text.trim(),
                            );

                            setState(() {

                              username =
                                  controller.text.trim();
                            });

                            Navigator.pop(context);
                          },

                          child: const Text(
                            "Save",
                          ),
                        ),
                      ],
                    );
                  },
                );
              },

              child: const Text(
                "Edit Username",
              ),
            ),

            const SizedBox(height: 30),

            const Text(
              "Node ID",
              style: TextStyle(
                fontSize: 18,
                fontWeight:
                FontWeight.bold,
              ),
            ),

            const SizedBox(height: 10),

            const Text(
              "Not Connected Yet",
            ),
          ],
        ),
      ),
    );
  }
}