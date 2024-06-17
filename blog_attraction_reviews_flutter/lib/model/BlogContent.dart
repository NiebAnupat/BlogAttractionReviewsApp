class BlogContent {
  final String id;
  final int order;
  final int type;
  final dynamic content;

  BlogContent({
    required this.id,
    required this.order,
    required this.type,
    required this.content,
  });

  factory BlogContent.fromJson(Map<String, dynamic> json) {
    return BlogContent(
      id: json['ID'],
      order: json['Order'],
      type: json['Type'],
      content: json['Value'],
    );
  }
}
