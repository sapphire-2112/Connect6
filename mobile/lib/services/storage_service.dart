import 'package:shared_preferences/shared_preferences.dart';

class StorageService {

  Future<void> saveName(
      String name,
      ) async {

    final prefs =
    await SharedPreferences.getInstance();

    await prefs.setString(
      "username",
      name,
    );
  }

  Future<String?> loadName() async {

    final prefs =
    await SharedPreferences.getInstance();

    return prefs.getString(
      "username",
    );
  }
}