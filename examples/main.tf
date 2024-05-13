resource "aws_s3_bucket" "bucket1" {
  bucket = "bucket1"
}

resource "aws_s3_object" "file1" {
  bucket  = aws_s3_bucket.bucket1.bucket
  key     = "file1.txt"
  content = "Hello, World"
}

resource "aws_s3_bucket" "bucket2" {
  bucket = "bucket2"
}

resource "aws_s3_object" "obj2" {
  bucket  = aws_s3_bucket.bucket2.bucket
  key     = "file2.txt"
  content = "Foo, Bar"
}