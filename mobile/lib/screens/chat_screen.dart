import 'package:flutter/material.dart';

import '../models/peer.dart';
import '../models/message.dart';

import '../services/api_service.dart';

import '../widgets/message_bubble.dart';

class ChatScreen extends StatefulWidget {

  final Peer peer;

  const ChatScreen({
    super.key,
    required this.peer,
  });

  @override
  State<ChatScreen> createState() =>
      _ChatScreenState();
}

class _ChatScreenState
    extends State<ChatScreen> {

  final ApiService api =
  ApiService();

  List<Message> messages = [];

  @override
  void initState() {

    super.initState();

    loadMessages();
  }

  Future<void> loadMessages() async {

    final loadedMessages =
    await api.getMessages(
      widget.peer.id,
    );

    setState(() {

      messages = loadedMessages;
    });
  }

  @override
  Widget build(BuildContext context) {

    return Scaffold(

      appBar: AppBar(
        title: Text(
          widget.peer.name,
        ),
      ),

      body: Column(

        children: [

          Expanded(

            child: ListView.builder(

              itemCount: messages.length,

              itemBuilder:
                  (context, index) {

                final message =
                messages[index];

                return MessageBubble(

                  sender:
                  message.sender,

                  text:
                  message.text,
                );
              },
            ),
          ),

          Padding(

            padding:
            const EdgeInsets.all(8),

            child: Row(

              children: [

                Expanded(

                  child: TextField(

                    decoration:
                    const InputDecoration(

                      hintText:
                      "Type Message",

                      border:
                      OutlineInputBorder(),
                    ),
                  ),
                ),

                const SizedBox(
                  width: 8,
                ),

                IconButton(

                  onPressed: () {},

                  icon: const Icon(
                    Icons.send,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}