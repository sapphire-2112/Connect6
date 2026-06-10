import 'package:flutter/material.dart';
import '../models/peer.dart';
import '../widgets/message_bubble.dart';


class ChatScreen extends StatelessWidget {
  final Peer peer;

  const ChatScreen({
    super.key,
    required this.peer,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(peer.name),
      ),
      body: Column(
        children: [

          Expanded(
            child: ListView(
              children: const [

                MessageBubble(
                  sender: "Alice",
                  text: "Hi",
                ),

                MessageBubble(
                  sender: "You",
                  text: "Hello",
                ),

                MessageBubble(
                  sender: "Alice",
                  text: "Working?",
                ),

                MessageBubble(
                  sender: "You",
                  text: "Fine",
                ),

              ],
            ),
          ),

          Padding(
            padding: EdgeInsets.all(8),

            child: Row(
              children: [

                Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: "Type Message",
                      border: OutlineInputBorder(),
                    ),
                  ),
                ),

                SizedBox(width: 8),

                IconButton(
                  onPressed: () {},
                  icon: Icon(Icons.send),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}