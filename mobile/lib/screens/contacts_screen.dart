import 'package:flutter/material.dart';
import '../models/peer.dart';
import 'chat_screen.dart';
import 'qr_screen.dart';
import 'scan_screen.dart';

class ContactsScreen extends StatelessWidget {
  ContactsScreen({super.key});

  final List<Peer> peers = [
    Peer(id: "con6-123",name: "ALice", online: true),
    Peer(id: "con6-456",name: "Bob",online: false),
    Peer(id: "con6-789",name: "Roy", online: true),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Connect6"),

        actions: [

          IconButton(
            tooltip: "My QR",
            icon: const Icon(
              Icons.qr_code,
              size: 22,
            ),

            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (_) => const QRScreen(),
                ),
              );
            },
          ),

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
                  builder: (_) => const ScanScreen(),
                ),
              );
            },
          ),

        ],
      ),
      body: ListView.builder(
        itemCount: peers.length,
        itemBuilder: (context, index) {
          final peer = peers[index];

          return ListTile(
            onTap: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => ChatScreen(
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
            title: Text(peer.name),
            subtitle: Text(
              peer.online
                  ? "Online"
                  : "Offline",
            ),
          );
        },
      ),
    );
  }
}