import 'package:blog_attraction_reviews_flutter/model/BlogContent.dart';

class BlogPost {
  String id;
  String title;
  String description;
  String thumbnail;
  DateTime createAt;
  String authorID;
  int like;
  int favorite;
  List<BlogContent?>? contents;

  BlogPost(
      {required this.id,
      required this.title,
      required this.description,
      required this.thumbnail,
      required this.createAt,
      required this.authorID,
      required this.like,
      required this.favorite,
      required this.contents});

  factory BlogPost.fromJson(Map<String, dynamic> json) {
    return BlogPost(
      id: json['ID'],
      title: json['Title'],
      description: json['Description'],
      thumbnail: json['Thumbnail'],
      createAt: DateTime.parse(json['CreateAt']),
      authorID: json['AuthorID'],
      like: json['Likes'],
      favorite: json['Favorites'],

      // if contents is not null, map the contents to BlogContent object
      contents: json['Contents'] != null
          ? (json['Contents'] as List)
              .map((content) => BlogContent.fromJson(content))
              .toList()
          : null,
    );
  }
}
