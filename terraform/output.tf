output "postgres_module_rds_endpoint" {
  description = "The connection endpoint"
  value       = module.postgres-module.rds_endpoint
}

output "postgres_module_rds_arn" {
  description = "The Amazon Resource Name (ARN) of the RDS instance"
  value       = module.postgres-module.rds_arn
}

output "postgres_module_rds_instance_class" {
  description = "The instance type of the RDS instance"
  value       = module.postgres-module.rds_instance_class
}

output "postgres_module_rds_name" {
  description = "The name of the RDS instance"
  value       = module.postgres-module.rds_name
}
