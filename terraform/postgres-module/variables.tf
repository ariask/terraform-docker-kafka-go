variable "allocated_storage" {
  description = "The allocated storage in gibibytes"
  type        = number
}

variable "cpu_configuration" {
  description = "The instance type of the RDS instance"
  type        = string
}

variable "username" {
  description = "Username for the master DB user"
  type        = string
}

variable "password" {
  description = "Password for the master DB user"
  type        = string
  sensitive   = true
}

variable "subnet_ids" {
  description = "A list of VPC subnet IDs"
  type        = list(string)
}

variable "vpc_security_group_ids" {
  description = "A list of VPC security group IDs"
  type        = list(string)
}

variable "database_name" {
  description = "The name of the database to create when the DB instance is created"
  type        = string
}