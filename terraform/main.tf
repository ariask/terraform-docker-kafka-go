# The main entry point of the Terraform configuration
# This file is used to configure the backend and to call the module
# The module is located in the postgres-module folder
# 

terraform {
  backend "s3" {
    bucket = "terraform-state-59"
    key    = "terraform.tfstate"
    region = "eu-north-1"
    dynamodb_table = "my-terraform-lock-table"
    encrypt = true
  }
}

module "telenor_devex_postgres_instance" {
    source = "./postgre-module"
    
    database_name           = "telenor_team_59_db"
    allocated_storage       = 20
    cpu_config              = "db.t2.micro"
    username                = var.username
    password                = var.password
    subnet_ids              = ["subnet-0bb1c79de3EXAMPLE", "subnet-09e3cdefdbEXAMPLE"]
    vpc_security_group_ids  = ["sg-01b2a78901EXAMPLE"]  
}

  
  