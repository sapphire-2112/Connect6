import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:qr_flutter/qr_flutter.dart';
import '../models/invite.dart';

class QRScreen extends StatelessWidget {
  const QRScreen({super.key});

  @override
  Widget build(BuildContext context) {

    final invite = Invite(
      type: "connect6",
      id: "con6-demo",
      name: "Parkhi",
    );

    final qrData = jsonEncode({
      "type": invite.type,
      "id": invite.id,
      "name": invite.name,
    });

    return Scaffold(
      appBar: AppBar(
        title: const Text("My Connect6 QR"),
      ),

      body: Center(
        child: QrImageView(
          data: qrData,
          size: 300,
        ),
      ),
    );
  }
}