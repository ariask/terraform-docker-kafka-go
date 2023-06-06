

resource "aws_db_instance" "example" {
  allocated_storage    = var.allocated_storage
  storage_type         = "gp2"
  engine               = "postgres"
  engine_version       = "13.3"
  instance_class       = var.cpu_configuration
  name                 = var.database_name
  username             = var.username
  password             = var.password
  parameter_group_name = "default.postgres13"
  skip_final_snapshot  = true

  vpc_security_group_ids = var.vpc_security_group_ids
  subnet_group           = aws_db_subnet_group.example.name

  tags = {
    Name = var.database_name
    CreatedBy = "Telenor.DeveloperExperience.PostgresModule"
  }
}

resource "aws_db_subnet_group" "example" {
  name       = "my_database_subnet_group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "My database subnet group"
    CreatedBy = "Telenor.DeveloperExperience.PostgresModule"
  }
}
