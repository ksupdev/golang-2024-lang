# Use official MySQL image from Docker Hub
FROM mysql:latest

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=mydatabase
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Copy custom configuration file
# COPY my.cnf /etc/mysql/my.cnf

# Expose the MySQL port
EXPOSE 3306

# Start MySQL service
CMD ["mysqld"]
