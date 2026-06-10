import 'package:flutter/material.dart';

class MessageBubble extends StatelessWidget {
  final String sender;
  final String text;

  const MessageBubble({
    super.key,
    required this.sender,
    required this.text,
  });

  @override
  Widget build(BuildContext context) {

    bool isMe = sender == "You";

    return Align(
      alignment:
      isMe
          ? Alignment.centerRight
          : Alignment.centerLeft,

      child: Container(
        margin: const EdgeInsets.symmetric(
          vertical: 4,
          horizontal: 8,
        ),

        padding: const EdgeInsets.all(12),

        decoration: BoxDecoration(
          color: isMe
              ? Colors.blue
              : Colors.green.shade100,

          borderRadius:
          BorderRadius.circular(12),
        ),

        child: Column(
          crossAxisAlignment:
          CrossAxisAlignment.start,

          children: [

            Text(
              sender,
              style: const TextStyle(
                fontWeight: FontWeight.bold,
              ),
            ),

            const SizedBox(height: 4),

            Text(text),
          ],
        ),
      ),
    );
  }
}