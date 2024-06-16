provider "aws" {
 region = "us-west-2"
}

resource "aws_dynamodb_table" "example_table" {
 name           = "heavybee-table"
 billing_mode   = "PAY_PER_REQUEST"
 hash_key       = "fromEmail"

 attribute {
   name = "fromEmail"
   type = "S"
 }

 attribute {
   name = "fromName"
   type = "S"
 }

 attribute {
   name = "fromDomain"
   type = "S"
 }
}
