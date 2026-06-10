import 'package:flutter/material.dart';

import '../models/peer.dart';
import '../services/api_service.dart';

import 'chat_screen.dart';
import 'profile_screen.dart';
import 'qr_screen.dart';
import 'scan_screen.dart';

class ContactsScreen extends StatefulWidget {

  const ContactsScreen({
    super.key,
  });

  @override
  State<ContactsScreen> createState() =>
      _ContactsScreenState();
}

class _ContactsScreenState
    extends State<ContactsScreen> {

  final ApiService api =
  ApiService();

  List<Peer> peers = [];

  @override
  void initState() {

    super.initState();

    loadPeers();
  }

  Future<void> loadPeers() async {

    final loadedPeers =
    await api.getPeers();

    setState(() {

      peers = loadedPeers;
    });
  }

  @override
  Widget build(BuildContext context) {

    return Scaffold(

      appBar: AppBar(

        title: const Text(
          "Connect6",
        ),

        actions: [

          IconButton(

            tooltip: "Scan QR",

            icon: const Icon(
              Icons.qr_code_scanner,
              size: 22,
            ),

            onPressed: () {

              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (_) =>
                  const ScanScreen(),
                ),
              );
            },
          ),

          PopupMenuButton<String>(

            onSelected: (value) {

              if (value == "qr") {

                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (_) =>
                    const QRScreen(),
                  ),
                );
              }

              if (value == "profile") {

                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (_) =>
                    const ProfileScreen(),
                  ),
                );
              }
            },

            itemBuilder: (context) => [

              const PopupMenuItem(

                value: "qr",

                child: Text(
                  "My QR",
                ),
              ),

              const PopupMenuItem(

                value: "profile",

                child: Text(
                  "Profile",
                ),
              ),
            ],
          ),
        ],
      ),

      body: peers.isEmpty

          ? const Center(
        child:
        CircularProgressIndicator(),
      )

          : ListView.builder(

        itemCount: peers.length,

        itemBuilder:
            (context, index) {

          final peer =
          peers[index];

          return ListTile(

            onTap: () {

              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (_) =>
                      ChatScreen(
                        peer: peer,
                      ),
                ),
              );
            },

            leading: Icon(

              Icons.circle,

              color: peer.online
                  ? Colors.green
                  : Colors.grey,

              size: 14,
            ),

            title: Text(
              peer.name,
            ),

            subtitle: Text(
              peer.online
                  ? "Online"
                  : "Offline",
            ),

            trailing:
            const Icon(
              Icons.chevron_right,
            ),
          );
        },
      ),
    );
  }
}