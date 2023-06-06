output "rds_endpoint" {
  description = "The connection endpoint"
  value       = aws_db_instance.example.endpoint
}

output "rds_arn" {
  description = "The Amazon Resource Name (ARN) of the RDS instance"
  value       = aws_db_instance.example.arn
}

output "rds_instance_class" {
  description = "The instance type of the RDS instance"
  value       = aws_db_instance.example.instance_class
}

output "rds_name" {
  description = "The name of the RDS instance"
  value       = aws_db_instance.example.name
}
