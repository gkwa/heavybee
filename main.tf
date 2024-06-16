provider "aws" {
  region = "us-west-2"
}

resource "aws_dynamodb_table" "heavybee_table" {
  name         = "heavybee-table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "fromEmail"

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

  global_secondary_index {
    name            = "fromName-index"
    hash_key        = "fromName"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "fromDomain-index"
    hash_key        = "fromDomain"
    projection_type = "ALL"
  }
}
