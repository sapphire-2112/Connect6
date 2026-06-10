import '../models/peer.dart';
import '../models/profile.dart';
import '../models/message.dart';

class ApiService {

  Future<List<Peer>> getPeers() async {

    return [

      Peer(
        id: "con6-123",
        name: "Alice",
        online: true,
      ),

      Peer(
        id: "con6-456",
        name: "Bob",
        online: false,
      ),

      Peer(
        id: "con6-789",
        name: "Roy",
        online: true,
      ),
    ];
  }

  Future<Profile> getProfile() async {

    return Profile(
      id: "con6-123",
      name: "Parkhi",
    );
  }
  Future<List<Message>> getMessages(
      String peerId,
      ) async {

    return [

      Message(
        sender: "Alice",
        text: "Hi",
      ),

      Message(
        sender: "You",
        text: "Hello",
      ),

      Message(
        sender: "Alice",
        text: "Working?",
      ),

      Message(
        sender: "You",
        text: "Fine",
      ),
    ];
  }

}