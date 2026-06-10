import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:mobile_scanner/mobile_scanner.dart';

class ScanScreen extends StatefulWidget {
  const ScanScreen({super.key});

  @override
  State<ScanScreen> createState() => _ScanScreenState();
}

class _ScanScreenState extends State<ScanScreen> {

  bool scanned = false;

  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        title: const Text("Scan Connect6 QR"),
      ),

      body: MobileScanner(

        onDetect: (capture) {

          if (scanned) return;

          scanned = true;

          final barcode =
              capture.barcodes.first.rawValue;

          if (barcode == null) {
            return;
          }

          try {

            final data =
            jsonDecode(barcode);

            if (data["type"] == "connect6") {

              showDialog(
                context: context,
                builder: (_) => AlertDialog(
                  title: const Text(
                    "Connected",
                  ),
                  content: Text(
                    "Connected to ${data["name"]}",
                  ),
                ),
              );

            } else {

              showDialog(
                context: context,
                builder: (_) => const AlertDialog(
                  title: Text("Invalid QR"),
                  content: Text(
                    "Not a Connect6 QR",
                  ),
                ),
              );
            }

          } catch (e) {

            showDialog(
              context: context,
              builder: (_) => const AlertDialog(
                title: Text("Invalid QR"),
                content: Text(
                  "QR format not recognized",
                ),
              ),
            );
          }
        },
      ),
    );
  }
}